package main

import (
    "log"
    "fmt"
    "regexp"
    "strings"
    "strconv"
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
    }
    Match interface {
        // Eval returns an array of labels, values, and match success
        Eval(line string) ([]string, []string, bool)
        NumVars() int
    }
    // Implements Match interface
    MatchRegex struct {
        Desc    string      // {{ pcnt:%p }}
        Pattern string      // ((?:\\d+\\.?\\d*)|(?:\\.\\d+))%
        Labels  []string    // pcnt
    }
    MatchColumns struct {
        Desc    string      // {{ colname: $1 }}
        Delims  string      // ",|"
        Indices []int       // [ 1 ]
        Labels  []string    // [ "colname" ]
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
            m := NewMatch(desc, v)
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
            m := NewMatch(desc, v)
            h.AddMatch(m)
            numVars += m.NumVars()
        }
        if h.Batch {
            if fail, ok := v["fail"]; ok {
                h.FailMatch = NewMatch(fail, v)
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
}

func (h *MatchHandler) AddMatch(r Match) {
    h.Matches = append(h.Matches, r)
}

func (h *MatchHandler) AddValues(keys []string, values []string) int {
    for i,_ := range values {
        kv := CreateKVPair(keys[i], values[i])
        if h.Info.Config.Debug {
            log.Println("[MatchHandler] Adding KV:", kv)
        }
        h.Vars[h.VarIndex] = kv
        h.VarIndex++
    }
    return len(values)
}

func (h *MatchHandler) Eval(line string, matchIndex int) int {
    rule := h.Matches[matchIndex]
    keys, vals, _ := rule.Eval(line)

    // Repeated Rules, add suffix
    if h.Batch && h.Repeats[h.MatchIndex] > 0 {
        skeys := make([]string, len(keys))
        suffix := h.Runs[h.MatchIndex]
        for i,_ := range keys {
            skeys[i] = fmt.Sprintf("%s_%d",keys[i],suffix)
        }
        return h.AddValues(skeys, vals)
    }

    return h.AddValues(keys, vals)
}

func (h *MatchHandler) ParseSingle(line string) bool {
    success := false
    for i,_ := range h.Matches {
        n := h.Eval(line,i)
        if n > 0 {
            UploadKV(h.Vars[:(n-1)], h.Info)
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
            UploadKV(h.Vars[:(h.VarIndex-1)], h.Info)
            h.MatchIndex = 0
            h.VarIndex = 0
            h.Runs = make([]int, len(h.Runs))
        }

        n = h.Eval(line, h.MatchIndex)
    }
    return success
}

func (h *MatchHandler) BatchFailed(line string) bool {
    _, _, failed := h.FailMatch.Eval(line)
    return failed
}

func GetMatchType(option Option_t) string {
    if _, ok := option["match"]; ok {
        return "match"
    } else if _, ok := option["fail"]; ok {
        return "match"
    } else if _, ok := option["columns"]; ok {
        return "columns"
    } else {
        log.Fatal("Unable to discern match type, invalid option.")
        return "match"
    }
}

func NewMatch(desc string, option Option_t) Match {
    switch GetMatchType(option) {
    case "match":
        return NewMatchRegex(desc)
    case "columns":
        return NewMatchColumns(desc, option)
    default:
        log.Fatal("Unknown match type.")
        return NewMatchRegex(desc)
    }
}

///////////////////////////////////////////////////////////////////
// MatchRegex
// ----------
// Example -- CPU usage: {{ cpu_usage_user:%p }} user, {{ cpu_usage_sys:%p }} sys
///////////////////////////////////////////////////////////////////

func NewMatchRegex(desc string) *MatchRegex {
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

func (r *MatchRegex) Eval(line string) ([]string, []string, bool) {
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

func NewMatchColumns(desc string, option Option_t) *MatchColumns {
    delimiters := ","
    if delims, ok := option["delimiters"]; ok {
        delimiters = delims
    }
    r, _ := regexp.Compile("\\{\\{\\s*(\\w+):(.+?)\\}\\}")
    tokens := r.FindAllStringSubmatch(desc, -1)

    labels  := make([]string, len(tokens))
    indices := make([]int, len(tokens))

    for i,token := range tokens {
        labels[i] = strings.TrimSpace(token[1])
        rule := strings.TrimSpace(token[2])
        reg, _ := regexp.Compile("^\\$[0-9]+$")
        if reg.Match([]byte(rule)) {
            // Matches $column_index pattern
            num, err := strconv.ParseInt(rule[1:], 10, 64)
            if err == nil {
                indices[i] = int(num)
            } else {
                log.Fatalf("%s not a valid column rule.", rule)
            }
        } else {
            log.Fatalf("%s not a valid column rule.", rule)
        }
    }

    return &MatchColumns{
        Desc:       desc,
        Delims:     delimiters,
        Indices:    indices,
        Labels:     labels}
}


func (c *MatchColumns) Eval(line string) ([]string, []string, bool) {
    tokens := strings.FieldsFunc(line, func(r rune) bool {
        return strings.ContainsRune(c.Delims, r)
    })
    numTokens := len(tokens)
    values := make([]string, numTokens)
    match := true
    for i, cindex := range c.Indices {
        if cindex < numTokens {
            values[i] = tokens[cindex]
        } else {
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

