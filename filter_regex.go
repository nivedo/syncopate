package main

import (
    "regexp"
    "strings"
    "strconv"
    "log"
    "fmt"
)

type (
    FilterRegex struct {
        Desc    string      // {{ pcnt:%p }}
        Pattern string      // ((?:\\d+\\.?\\d*)|(?:\\.\\d+))%
        Repeat  int
        NumPass int
        NumVars int
        Vars    KVList
    }
)

////////////////////////////////////////////////////////////////////////
// FilterRegex
// ----------
// RegexFilter parses each blob of text independently by applying
// a regular expression with multiple fields.
// 
// Examples: 
// (1) Built in parameters 
//     e.g. MyFloat: {{ my_float:%f }} MyInt: {{ my_int:%d }}
// (2) Custom Regex (requires ONE capturing parentheses, default S_CHAR)
//     e.g. MyRegex: {{ (int) my_value:(\d+)[abc]+ }}
////////////////////////////////////////////////////////////////////////

func NewFilterRegex(opt Option_t) *FilterRegex {
    if !IsFilterRegex(opt) {
        log.Fatalf("Illegal NewFilterRegex options: %+v", opt)
    }

    desc := opt["match"]
    repeat := 0
    if rep, ok := opt["repeat"]; ok {
        repeat, _ = strconv.Atoi(rep)
    }

    f := &FilterRegex{Desc: desc, Repeat: repeat}
    f.Init()

    return f
}

func IsFilterRegex(opt Option_t) bool {
    _, ok := opt["match"]
    return ok
}

func (f *FilterRegex) GetVars() KVList {
    return f.Vars
}

func (f *FilterRegex) Match(data string) bool {
    reg, _ := regexp.Compile(f.Pattern)
    allMatch := reg.FindAllStringSubmatch(data, -1)
    if len(allMatch) > 0 {
        offset := f.NumPass * f.NumVars
        for i,val := range allMatch[0][1:] {
            f.Vars[offset + i].V = val
        }
        f.NumPass++
        if f.NumPass >= f.Repeat {
            f.NumPass = 0
            return true
        }
    }
    return false
}

func (f *FilterRegex) Init() {
    fvars   := GetFilterVars(f.Desc)
    pattern := f.Desc
    labels  := make([]string, len(fvars))
    types   := make([]uint8, len(fvars))
    subs    := make(map[string]string)

    for i,v := range fvars {
        labels[i] = v.Name
        types[i]  = S_CHAR // Default to [16]char
        subtoken := fmt.Sprintf("SYNCVAR_%d",i)

        // Substitute regex shortcuts
        switch v.Rule {
        case "%p":
            subs[subtoken] = "((?:\\d+\\.?\\d*)|(?:\\.\\d+))%"
            types[i] = S_FLOAT
        case "%f":
            subs[subtoken] = "([+-]?(?:\\d+\\.?\\d*)|(?:\\.\\d+))"
            types[i] = S_FLOAT
        case "%d":
            subs[subtoken] = "(\\d+)"
            types[i] = S_INT
        case "%w":
            subs[subtoken] = "(\\w+)"
        case "%mem":
            subs[subtoken] = "(\\d+[BKMG]?)[+-]?"
        default:
            // Use user specified regex
            subs[subtoken] = v.Rule
        }

        // Cast type if available
        if v.Type != 0 {
            types[i] = v.Type
        }

        pattern = strings.Replace(pattern, v.Desc, subtoken, 1)
    }

    // Escape Special Characters
    pattern = regexp.QuoteMeta(pattern)

    // Treat newlines as "match all"
    r2, _ := regexp.Compile("[\\t\\n\\r]+")
    pattern = r2.ReplaceAllString(pattern, ".*")

    // Replace Subtokens
    for k,v := range subs {
        pattern = strings.Replace(pattern, k, v, 1)
    }

    // Whitespace is arbitrary
    r3, _ := regexp.Compile("\\s+")
    pattern = r3.ReplaceAllString(pattern, "\\s*")

    f.Pattern = pattern
    f.NumVars = len(labels)

    // Initialize Vars
    if f.Repeat > 1 {
        // TODO: Handle this later
        f.Vars = make(KVList, f.Repeat * f.NumVars)
        for i,label := range labels {
            for j := 0; j < f.Repeat; j++ {
                index := i + j * len(labels)
                f.Vars[index].K = fmt.Sprintf("%s_%d", label, j)
                f.Vars[index].Type = types[i]
            }
        }
    } else {
        f.Vars = make(KVList, f.NumVars)
        for i,label := range labels {
            f.Vars[i].K = label
            f.Vars[i].Type = types[i]
        }
    }
}