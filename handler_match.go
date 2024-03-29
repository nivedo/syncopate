package main

import (
    "log"
    "fmt"
    "regexp"
    "strings"
    "strconv"
    "bufio"
    "os"
    "unicode/utf8"
)

/*
 *  Two types of MatchHandlers:
 *  ==========================
 *
 *  1) BATCH = false -- for each line, evaluate all given rules and upload on ANY match.
 *  2) BATCH = true  -- for each line, evaluate only rule N until rule N passes, then iterate N++.
 *     Do not upload anything until all rules pass.
 */

type (
    MatchHandler struct {
        Info       *HandlerInfo
        Vars       KVList
        Matches    []Match

        // Helpers
        VarIndex    int
        MatchIndex  int

        // Batch Parameters
        Batch      bool
        Runs       []int
        Repeats    []int
        FailMatch  Match
        LineCount  int
    }
    Match interface {
        // Eval returns an array of labels, values, and match success
        EvalAndParse(line string, h *MatchHandler) ([]string, []string, bool)
        NumVars() int
    }
    // Implements Match interface
    MatchRegex struct {
        Desc    string      // {{ pcnt:%p }}
        Pattern string      // ((?:\\d+\\.?\\d*)|(?:\\.\\d+))%
        Labels  []string    // pcnt
    }
    MatchColumns struct {
        Desc        string      // {{ colname: $1 }}
        Delims      string      // ",|"
        Indices     []int       // [ 1 ]
        Labels      []string    // [ "colname" ]
        HasHeader   bool
    }
    Range_t struct {
        Start   int
        End     int
    }
    MatchTable struct {
        Desc            string      // {{ pid: @"PID" }}
        IReqIndices     []int       // [ 0 ]
        IReqLabels      []string    // [ "pid" ]
        HReqHeaders     []string
        HReqLabels      []string
        HeaderPattern   string
        EndPattern      string
        RowBuffer       []string
        HeaderMap       map[string]int
        HeaderBuffer    string
        NumRows         int
        BufferRowIndex  int
        InTable         bool        // In table
        HasMask         bool        // Col mask initialized
        ColMask         []int
        ColRanges       []Range_t
        VarIndex        int
    }
)

func NewMatchHandler(info *HandlerInfo, batch bool) *MatchHandler {
    h := &MatchHandler{Info: info, VarIndex: 0, MatchIndex: 0, Batch: batch}
    h.Load()
    return h
}

func (h *MatchHandler) Load() {
    numVars := 0
    for _,v := range h.Info.Config.Options {
        if desc, ok := v["match"]; ok {
            m := h.NewMatch(desc, v)
            h.AddMatch(m)
            repeats := 1

            if h.Batch {
                log.Printf("[MatchHandler] TRACKING %s\n", desc)
                h.Repeats = append(h.Repeats, 0)

                if rep, ok := v["repeat"]; ok {
                    repeats, _ = strconv.Atoi(rep)
                    if repeats > 1 {
                        h.Repeats[len(h.Repeats)-1] = repeats
                        log.Printf("[MatchHandler] Repeat %s %d times\n", desc, repeats)
                    }
                }
            }

            numVars += repeats * m.NumVars()
        } else if desc, ok := v["columns"]; ok {
            m := h.NewMatch(desc, v)
            h.AddMatch(m)
            numVars += m.NumVars()
            h.Repeats = append(h.Repeats, make([]int, numVars)...)
        } else if desc, ok := v["table"]; ok {
            m := h.NewMatch(desc, v)
            h.AddMatch(m)
            numVars += m.NumVars()
            h.Repeats = append(h.Repeats, make([]int, numVars)...)
        }
        if h.Batch {
            if fail, ok := v["fail"]; ok {
                h.FailMatch = h.NewMatch(fail, v)
                log.Printf("[MatchHandler] Faliure Condition - %s\n", fail)
            }
        }
    }
    log.Printf("[MatchHandler] Initialized with %d variables, %d matches.", numVars, len(h.Matches))
    h.Runs = make([]int, len(h.Matches))
    h.Vars = make(KVList, numVars)
}

func (h *MatchHandler) Help() {

}

func (h *MatchHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
    }
}

func (h *MatchHandler) Parse(data string) {
    if !h.Batch {
        h.ParseSingle(data)
    } else {
        if !h.ParseBatch(data) {
            if h.BatchFailed(data) {
                log.Fatalf("[BatchHandler] FAILED match %+v with line %s", h.Matches[h.MatchIndex], data)
            }
        }
    }
    h.LineCount++
}

func (h *MatchHandler) AddMatch(r Match) {
    h.Matches = append(h.Matches, r)
}

func (h *MatchHandler) AddValues(keys []string, values []string) int {
    for i,_ := range values {
        kv := CreateKVPair(keys[i], values[i])
        if h.Info.Config.Debug {
            log.Printf("[MatchHandler] Adding KV: %s, index: %d", kv, h.VarIndex)
        }
        h.Vars[h.VarIndex] = kv
        h.VarIndex++
    }
    return len(values)
}

func (h *MatchHandler) Eval(line string, matchIndex int) int {
    // rule := h.Matches[matchIndex]
    keys, vals, _ := h.Matches[matchIndex].EvalAndParse(line, h)

    // Repeated Rules, add suffix
    if h.Batch && h.Repeats[h.MatchIndex] > 0 {
        skeys := make([]string, len(keys))
        suffix := h.Runs[h.MatchIndex]
        for i,_ := range keys {
            skeys[i] = fmt.Sprintf("%s_%d",keys[i],suffix)
        }
        return h.AddValues(skeys, vals)
    }
    // log.Print(line)
    return h.AddValues(keys, vals)
}

func (h *MatchHandler) ParseSingle(line string) bool {
    success := false
    for i,_ := range h.Matches {
        n := h.Eval(line,i)
        if n > 0 {
            h.Info.Uploader.UploadKV(h.Vars[:(n-1)])
            h.VarIndex = 0
            success = true
        }
    }
    return success
}

func (h *MatchHandler) ParseBatch(line string) bool {
    success := false
    n := h.Eval(line, h.MatchIndex)
    for n > 0 {
        // Match passes, advance to next rule
        success = true

        // Repeated rules
        if h.Runs[h.MatchIndex] < h.Repeats[h.MatchIndex]-1 {
            h.Runs[h.MatchIndex]++
            break
        }

        h.MatchIndex++
        if h.MatchIndex == len(h.Matches) {
            // All rules pass, upload KVList
            h.Info.Uploader.UploadKV(h.Vars[:h.VarIndex])
            h.MatchIndex = 0
            h.VarIndex = 0
            h.Runs = make([]int, len(h.Runs))
        }

        n = h.Eval(line, h.MatchIndex)
    }
    return success
}

func (h *MatchHandler) BatchFailed(line string) bool {
    _, _, failed := h.FailMatch.EvalAndParse(line, h)
    return failed
}

func GetMatchType(option Option_t) string {
    if _, ok := option["match"]; ok {
        return "match"
    } else if _, ok := option["fail"]; ok {
        return "match"
    } else if _, ok := option["columns"]; ok {
        return "columns"
    } else if _, ok := option["table"]; ok {
        return "table"
    } else {
        log.Fatal("Unable to discern match type, invalid option.")
        return "match"
    }
}

func (h *MatchHandler) NewMatch(desc string, option Option_t) Match {
    switch GetMatchType(option) {
    case "match":
        return h.NewMatchRegex(desc)
    case "columns":
        return h.NewMatchColumns(desc, option)
    case "table":
        return h.NewMatchTable(desc, option)
    default:
        log.Fatal("Unknown match type.")
        return h.NewMatchRegex(desc)
    }
}

///////////////////////////////////////////////////////////////////
// MatchRegex
// ----------
// Example -- CPU usage: {{ cpu_usage_user:%p }} user, {{ cpu_usage_sys:%p }} sys
///////////////////////////////////////////////////////////////////

func (h *MatchHandler) NewMatchRegex(desc string) *MatchRegex {
    r, _ := regexp.Compile("\\{\\{\\s*(\\w+):(.+?)\\}\\}")
    tokens := r.FindAllStringSubmatch(desc, -1)

    result := desc
    var labels []string

    subs := make(map[string]string)

    for i,token := range tokens {
        labels = append(labels, strings.TrimSpace(token[1]))
        rule := strings.TrimSpace(token[2])
        subtoken := fmt.Sprintf("SYNCVAR_%d",i)
        switch rule {
        case "%p":
            subs[subtoken] = "((?:\\d+\\.?\\d*)|(?:\\.\\d+))%"
        case "%f":
            subs[subtoken] = "([+-]?(?:\\d+\\.?\\d*)|(?:\\.\\d+))"
        case "%d":
            subs[subtoken] = "(\\d+)"
        case "%w":
            subs[subtoken] = "(\\w+)"
        case "%mem":
            subs[subtoken] = "(\\d+[BKMG]?)[+-]?"
        default:
            // Use user specified regex
            subs[subtoken] = rule
        }
        result = strings.Replace(result, token[0], subtoken, 1)
    }

    // Escape Special Characters
    result = regexp.QuoteMeta(result)

    // Treat newlines as "match all"
    r2, _ := regexp.Compile("[\\t\\n\\r]+")
    result = r2.ReplaceAllString(result, ".*")

    // Replace Subtokens
    for k,v := range subs {
        result = strings.Replace(result, k, v, 1)
    }

    // Whitespace is arbitrary
    r3, _ := regexp.Compile("\\s+")
    result = r3.ReplaceAllString(result, "\\s*")

    return &MatchRegex{Desc: desc, Pattern: result, Labels: labels}
}

func (r *MatchRegex) EvalAndParse(line string, h *MatchHandler) ([]string, []string, bool) {
    match, _ := regexp.MatchString(r.Pattern, line)
    if match {
        reg, _ := regexp.Compile(r.Pattern)
        allMatch := reg.FindAllStringSubmatch(line, -1)
        return r.Labels, allMatch[0][1:], true
    }
    return nil, nil, false
}

func (r *MatchRegex) NumVars() int {
    return len(r.Labels)
}

///////////////////////////////////////////////////////////////////
// MatchColumns
///////////////////////////////////////////////////////////////////

func (h *MatchHandler) GetColumnIndexMapFromHeaders(delimiters string, indexMap map[string]int) bool {
    // Look for filename from run command
    filename := ""//h.Info.Config.CmdFile
    if len(filename) > 0 {
        file, err := os.Open(filename)
        if err == nil {
            scanner := bufio.NewScanner(file)
            if scanner.Scan() {
                // Successfuly scanned first line in file
                line := scanner.Text()
                log.Printf("%s header: %s", filename, line)
                headers := strings.FieldsFunc(line, func(r rune) bool {
                    return strings.ContainsRune(delimiters, r)
                })
                for i, ht := range headers {
                    h := strings.TrimSpace(ht)
                    indexMap[h] = i
                }
                return true
            }
        }
        file.Close()
    }
    return false
}

func (h *MatchHandler) NewMatchColumns(desc string, option Option_t) *MatchColumns {
    delimiters := ","
    if delims, ok := option["delimiters"]; ok {
        delimiters = delims
    }
    headerIndexMap := make(map[string]int)
    hasHeaderIndex := false
    hasHeader := false
    if headers, ok := option["headers"]; ok && strings.ToLower(headers) == "true"{
        hasHeader = true
        hasHeaderIndex = h.GetColumnIndexMapFromHeaders(delimiters, headerIndexMap)
    }

    r, _ := regexp.Compile("\\{\\{\\s*(\\w*):?\\$?([\\w\\d\\-]+)\\s*\\}\\}")
    tokens := r.FindAllStringSubmatch(desc, -1)
    // log.Print(tokens)

    labels  := make([]string, len(tokens))
    indices := make([]int, len(tokens))

    for i,token := range tokens {
        // log.Print(token)
        label := strings.TrimSpace(token[1])
        rule := strings.TrimSpace(token[2])
        reg, _ := regexp.Compile("^[0-9]+$")
        if reg.Match([]byte(rule)) {
            // (1) Matches ${numeric} pattern
            num, err := strconv.ParseInt(rule, 10 /* base 10 */, 64 /* int64 */)
            if err == nil {
                indices[i] = int(num)
                if len(label) > 0 {
                    labels[i] = label
                } else {
                    labels[i] = "column" + rule
                }
            } else {
                log.Fatalf("%s not a valid column rule.", rule)
            }
        } else if hasHeaderIndex {
            // (2) Matches ${alpha numeric} pattern
            if index, ok := headerIndexMap[rule]; ok {
                indices[i] = index
                if len(label) > 0 {
                    labels[i] = label
                } else {
                    labels[i] = rule
                }
            } else {
                log.Fatalf("%s does not exist in header index map.", rule)
            }
        } else {
            log.Fatalf("%s not a valid column rule.", rule)
        }
    }

    return &MatchColumns{
        Desc:       desc,
        Delims:     delimiters,
        Indices:    indices,
        Labels:     labels,
        HasHeader:  hasHeader}
}


func (c *MatchColumns) EvalAndParse(line string, h *MatchHandler) ([]string, []string, bool) {
    // Ignore headers
    if h.Info.Config.CmdInfo != nil && c.HasHeader && h.LineCount == 0 {
        return nil, nil, false
    }
    l := strings.TrimSpace(line)
    // log.Printf("Match columns eval: %s, delims=%s", l, c.Delims)
    // log.Println(c)
    tokens := strings.FieldsFunc(l, func(r rune) bool {
        return strings.ContainsRune(c.Delims, r)
    })
    numTokens := len(tokens)
    values := make([]string, numTokens)
    match := true
    for i, cindex := range c.Indices {
        if cindex < numTokens {
            values[i] = tokens[cindex]
        } else {
            log.Printf("Index request %d for %s exceeds number of columns %d.", cindex, c.Labels[i], numTokens)
            match = false
            break
        }
    }
    if match {
        return c.Labels, values, true
    } else {
        return nil, nil, false
    }
}

func (c *MatchColumns) NumVars() int {
    return len(c.Labels)
}

///////////////////////////////////////////////////////////////////
// MatchTable
///////////////////////////////////////////////////////////////////

func (h *MatchHandler) NewMatchTable(desc string, option Option_t) *MatchTable {
    var headerPattern, endPattern string
    var ok bool
    numRows := 5    // Default number of rows
    if headerPattern, ok = option["header"]; !ok {
        log.Fatal("Missing header pattern for table.") 
    }
    if endPattern, ok = option["end"]; !ok {
        log.Fatal("Missing end pattern for table.")
    }
    if rows, ok := option["rows"]; ok {
        num, err := strconv.ParseInt(rows, 10 /* base 10 */, 64 /* int64 */)
        if err != nil {
            log.Fatal("Unable to parse \"rows\" option: %s", rows)
        } else {
            numRows = int(num)
        }
    }
    r, _ := regexp.Compile("\\{\\{\\s*(\\w*):?\\$([^}]+)\\s*\\}\\}")
    tokens := r.FindAllStringSubmatch(desc, -1)
    // log.Print(tokens)

    var iReqLabels  []string
    var iReqIndices []int
    var hReqLabels  []string
    var hReqHeaders []string

    for _,token := range tokens {
        // log.Print(token)
        label := strings.TrimSpace(token[1])
        rule := strings.TrimSpace(token[2])
        reg, _ := regexp.Compile("^[0-9]+$")
        if reg.Match([]byte(rule)) {
            // (1) Matches ${index} pattern
            num, err := strconv.ParseInt(rule, 10 /* base 10 */, 64 /* int64 */)
            if err == nil {
                iReqIndices = append(iReqIndices, int(num))
                if len(label) > 0 {
                    iReqLabels = append(iReqLabels, label)
                } else {
                    iReqLabels = append(iReqLabels, "column" + rule)
                }
            } else {
                log.Fatalf("%s not a valid column rule.", rule)
            }
        } else {
            // (2) Matches ${header} pattern
            hReqHeaders = append(hReqHeaders, rule)
            if len(label) > 0 {
                hReqLabels = append(hReqLabels, label)
            } else {
                hReqLabels = append(hReqLabels, ConvertToValidSeriesKey(rule))
            }
        }
    }

    return &MatchTable{
        Desc:           desc,
        IReqIndices:    iReqIndices,
        IReqLabels:     iReqLabels,
        HReqHeaders:    hReqHeaders,
        HReqLabels:     hReqLabels,
        HeaderPattern:  headerPattern,
        HeaderMap:      make(map[string]int),
        EndPattern:     endPattern,
        HasMask:        false,
        InTable:        false,
        NumRows:        numRows,
        RowBuffer:      make([]string, numRows),
        BufferRowIndex: 0,
        VarIndex:       0}
}

func (t *MatchTable) GetValueAtIndex(line string, index int) string {
    var s string
    if index < len(t.ColRanges) {
        r := t.ColRanges[index]
        if index == len(t.ColRanges)-1 {
            s = strings.TrimSpace(line[r.Start:])
        } else {
            s = strings.TrimSpace(line[r.Start:r.End])
        }
    } else {
        log.Fatal("Table col index %d is out of range %d", index, len(t.ColRanges))
    }
    return s
}

func (t *MatchTable) ParseRow(line string, rowIndex int, labels []string, values []string) {
    // Parse requested indices
    for i, index := range t.IReqIndices {
        values[t.VarIndex] = t.GetValueAtIndex(line, index)
        labels[t.VarIndex] = fmt.Sprintf("%s_%d", t.IReqLabels[i], rowIndex)
        t.VarIndex++
    }

    // Parse requested headers
    // fmt.Println(t.HReqHeaders)
    for i, header := range t.HReqHeaders {
        if index, ok := t.HeaderMap[header]; ok {
            values[t.VarIndex] = t.GetValueAtIndex(line, index)
            labels[t.VarIndex] = fmt.Sprintf("%s_%d", t.HReqLabels[i], rowIndex)
        } else {
            log.Fatal("Request header %s for %s does not exist.", header, t.HReqLabels[i])
        }
        t.VarIndex++
    }
}

func (t *MatchTable) InitColRange() {
    fmt.Println(t.ColMask)
    sep := true
    tokenRange := Range_t{Start: -1, End: -1}
    for i, m := range t.ColMask {
        if (m == 0 || i == len(t.ColMask)-1) && !sep {
            // End
            tokenRange.End = i
            t.ColRanges = append(t.ColRanges, tokenRange)
            sep = true
        } else if m > 0 && sep {
            // Start
            tokenRange = Range_t{Start: i, End: i}
            sep = false
        }
    }
    // Check column ranges valid
    for _, r := range t.ColRanges {
        if !(r.Start >= 0 && r.End >= 0 && r.Start <= r.End) {
            log.Fatalf("Invalid table col range, start: %d, end: %d.", r.Start, r.End)
        }
    }
    fmt.Println(t.ColRanges)
    t.HasMask = true
}

func (t *MatchTable) InitHeaderMap() {
    // t.HeaderMap = make(map[string]int)
    if len(t.HeaderBuffer) == 0 {
        log.Fatal("Header buffer is empty.")
    }
    for i, r := range t.ColRanges {
        var h string
        if i == len(t.ColRanges)-1 {
            h = strings.TrimSpace(t.HeaderBuffer[r.Start:])
        } else {
            h = strings.TrimSpace(t.HeaderBuffer[r.Start:r.End])
        }
        t.HeaderMap[h] = i
    }
    fmt.Print(t.HeaderMap)
}

func (t *MatchTable) ParseRowForMask(line string) {
    for i, r := range line {
        if i < len(t.ColMask) && r != ' ' {
            t.ColMask[i]++
        }
    }
}

func (t *MatchTable) EvalAndParse(line string, h *MatchHandler) ([]string, []string, bool) {
    // log.Print(t.ColMask)
    matchHeader, _ := regexp.MatchString(t.HeaderPattern, line)
    if matchHeader {
        t.InTable = true
        t.BufferRowIndex = 0
        t.VarIndex = 0
        if !t.HasMask {
            t.ColMask = make([]int, utf8.RuneCountInString(line))
            t.HeaderBuffer = line
            t.ParseRowForMask(line)
        }
        // t.RowBuffer = make([]string, t.NumRows)
        return nil, nil, false
    }
    labels := make([]string, t.NumVars())
    values := make([]string, t.NumVars())
    if t.InTable {
        matchEnd, _ := regexp.MatchString(t.EndPattern, line)
        if matchEnd {
            t.InTable = false
            if !t.HasMask {
                t.InitColRange()
                t.InitHeaderMap()
            }
            for i, row := range t.RowBuffer {
                t.ParseRow(row, i, labels, values)
            }
            return labels, values, true
        }
        if !t.HasMask {
            t.ParseRowForMask(line)
        }
        if t.BufferRowIndex < t.NumRows {
            t.RowBuffer[t.BufferRowIndex] = line
            t.BufferRowIndex++
        }
    }
    return nil, nil, false
}

func (t *MatchTable) NumVars() int {
    return (len(t.IReqLabels) + len(t.HReqLabels)) * t.NumRows
}

