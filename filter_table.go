package main

import (
    "log"
    "strings"
    "regexp"
    "fmt"
    "strconv"
    "unicode/utf8"
)

type (
    ColRange struct {
        Start   int
        End     int
    }
    FilterTable struct {
        Desc            string      // {{ pid: @"PID" }}
        ReqIndices      []int       // [ 0 ] Elements are sorted in this order
        HeaderPattern   string
        EndPattern      string
        RowBuffer       []string
        HeaderBuffer    string
        NumRows         int
        InTable         bool        // In table
        Initialized     bool        // Col mask initialized
        ColMask         []int
        ColRanges       []ColRange
        RowIndex        int
        Vars            KVList
    }
)

func NewFilterTable(opt Option_t) *FilterTable {
    if !IsFilterTable(opt) {
        log.Fatalf("Illegal NewFilterTable options: %+v", opt)
    }
    f := &FilterTable{
        Desc:           opt["table"],
        HeaderPattern:  opt["header"],
        EndPattern:     DATA_EOF,
    }
    if end, ok := opt["end"]; ok {
        f.EndPattern = end
    }
    nrow, err := strconv.ParseInt(opt["rows"], 10 /* base 10 */, 64 /* int64 */)
    if err != nil {
        log.Fatalf("Unable to parse \"rows\" option: %s", opt["rows"])
    }
    f.NumRows = int(nrow)
    f.RowBuffer = make([]string, nrow)

    return f
}

func (f *FilterTable) GetVars() KVList {
    return f.Vars
}

// This should be done immediately after initialization 
// to preallocate Vars (KVList).  HeaderMap should be lower case!!
func (f *FilterTable) Prealloc() {
    headerMap := f.GetHeaderMap()
    fvars     := GetFilterVars(f.Desc)

    baseLabels := make([]string, len(fvars))
    types      := make([]uint8, len(fvars))

    for i,v := range fvars {
        label  := v.Name
        rule   := v.Rule
        reg, _ := regexp.Compile("^\\$[0-9]+$")

        if reg.MatchString(rule) {
            // (1) Matches ${index} pattern
            num, err := strconv.ParseInt(rule[1:], 10, 64)
            if err != nil {
                log.Fatalf("%s not a valid column rule.", rule)
            }
            f.ReqIndices = append(f.ReqIndices, int(num))
            if len(label) == 0 {
                label = fmt.Sprintf("col%d", num)
            }
        } else {
            // (2) Matches ${header} pattern
            rule = strings.ToLower(rule[1:])
            if num, ok := headerMap[rule]; ok {
                f.ReqIndices = append(f.ReqIndices, num)
            } else {
                log.Fatalf("Could not find column %s.", rule)
            }
            if len(label) == 0 {
                label = ConvertToValidSeriesKey(rule)
            }
        }

        baseLabels[i] = label
        if v.Type != 0 {
            types[i] = v.Type
        } else {
            types[i] = S_CHAR
        }
    }

    numBase := len(f.ReqIndices)
    f.Vars = make(KVList, numBase * f.NumRows)
    for i,_ := range baseLabels {
        for j := 0; j < f.NumRows; j++ {
            index := i + j * numBase
            f.Vars[index].K = fmt.Sprintf("%s_%d", baseLabels[i], j)
            f.Vars[index].Type = types[i]
        }
    }
}

func (f *FilterTable) InitColRange() int {
    count := 0
    sep := true
    f.ColRanges = nil
    tokenRange := ColRange{Start: -1, End: -1}
    for i, m := range f.ColMask {
        if (m == 0 || i == len(f.ColMask)-1) && !sep {
            // End
            count++
            tokenRange.End = i
            f.ColRanges = append(f.ColRanges, tokenRange)
            sep = true
        } else if m == 0 && sep {
            // Extend range as far as possible to be safe
            f.ColRanges[count-1].End = i
        } else if m > 0 && sep {
            // Start
            tokenRange = ColRange{Start: i, End: i}
            sep = false
        }
    }
    // Check column ranges valid
    for _, r := range f.ColRanges {
        if !(r.Start >= 0 && r.End >= 0 && r.Start <= r.End) {
            log.Fatalf("Invalid table col range, start: %d, end: %d.", r.Start, r.End)
        }
    }

    return count
}

func (f *FilterTable) GetHeaderMap() map[string]int {
    headerMap := make(map[string]int)
    if len(f.HeaderBuffer) == 0 {
        log.Fatal("Header buffer is empty.")
    }
    for i, r := range f.ColRanges {
        var h string
        if i == len(f.ColRanges)-1 {
            h = strings.TrimSpace(f.HeaderBuffer[r.Start:])
        } else {
            h = strings.TrimSpace(f.HeaderBuffer[r.Start:r.End])
        }
        h = strings.ToLower(h)
        headerMap[h] = i
    }
    return headerMap
}

func (f *FilterTable) Initialize() {
    f.InitColRange()
    f.Prealloc()
    f.Initialized = true
}

func (f *FilterTable) GetValueAtIndex(line string, index int) string {
    var s string
    if index < len(f.ColRanges) {
        r := f.ColRanges[index]
        if index == len(f.ColRanges)-1 {
            s = strings.TrimSpace(line[r.Start:])
        } else {
            s = strings.TrimSpace(line[r.Start:r.End])
        }
    } else {
        log.Fatal("Table col index %d is out of range %d", index, len(f.ColRanges))
    }
    return s
}

func (f *FilterTable) ParseRowForMask(data string) {
    for i, r := range data {
        if i < len(f.ColMask) && r != ' ' {
            f.ColMask[i]++
        }
    }
}

func (f *FilterTable) ParseRow(data string, rowNum int) {
    // Parse requested indices
    offset := rowNum * len(f.ReqIndices)
    for i, index := range f.ReqIndices {
        varIndex := offset + i
        f.Vars[varIndex].V = f.GetValueAtIndex(data, index)
    }
}

func (f *FilterTable) Match(data string) bool {
    matchHeader, _ := regexp.MatchString(f.HeaderPattern, data)
    if matchHeader {
        f.InTable = true
        f.ColMask = make([]int, utf8.RuneCountInString(data))
        f.RowBuffer = make([]string, f.NumRows)
        f.HeaderBuffer = data
        return false
    }
    if f.InTable {
        matchEnd, _ := regexp.MatchString(f.EndPattern, data)
        if matchEnd {
            f.InTable = false
            f.RowIndex = 0
            if !f.Initialized {
                f.Initialize()
            } else {
                f.InitColRange()
            }
            for i, row := range f.RowBuffer {
                f.ParseRow(row, i)
            }
            return true
        } else {
            if strings.TrimSpace(data) == "" {
                return false
            }
            f.ParseRowForMask(data)
            if f.RowIndex < f.NumRows {
                f.RowBuffer[f.RowIndex] = data
                f.RowIndex++
            }
        }
    }
    return false
}

func IsFilterTable(opt Option_t) bool {
    if _, ok := opt["table"]; !ok {
        return false 
    }
    if _, ok := opt["header"]; !ok {
        return false 
    }
    if _, ok := opt["rows"]; !ok {
        return false
    }
    return true
}