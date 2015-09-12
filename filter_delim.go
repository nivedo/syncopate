package main

import (
    "log"
    "regexp"
    "strconv"
    "fmt"
    "strings"
)

type (
    FilterDelim struct {
        Desc        string      // {{ colname: $1 }}
        Delims      string      // ",|"
        Indices     []int       // [ 1 ]
        MaxIndex    int
        Vars        KVList
    }
)

func NewFilterDelim(opt Option_t, headerMap map[string]int) *FilterDelim {
    if !IsFilterDelim(opt) {
        log.Fatalf("Illegal FilterDelim options: %+v", opt)
    }

    f := &FilterDelim{
        Desc:    opt["columns"],
        Delims:  opt["delims"],
    }

    fvars     := GetFilterVars(f.Desc)
    f.Indices  = make([]int, len(fvars))
    f.Vars     = make(KVList, len(fvars))

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
            f.Indices[i] = int(num)
            if len(label) == 0 {
                label = fmt.Sprintf("col%d", num)
            }
        } else {
            // (2) Matches ${header} pattern
            rule = strings.ToLower(rule[1:])
            if num, ok := headerMap[rule]; ok {
                f.Indices[i] = num
            } else {
                log.Fatalf("Could not find column %s.", rule)
            }
            if len(label) == 0 {
                label = ConvertToValidSeriesKey(rule)
            }
        }

        if f.Indices[i] > f.MaxIndex {
            f.MaxIndex = f.Indices[i]
        }

        f.Vars[i].K = label

        if v.Type != 0 {
            f.Vars[i].Type = v.Type
        } else {
            f.Vars[i].Type = S_CHAR
        }
    }

    return f
}

func (f *FilterDelim) GetVars() KVList {
    return f.Vars
}

func (f *FilterDelim) Match(data string) bool {
    l := strings.TrimSpace(data)
    tokens := strings.FieldsFunc(l, func(r rune) bool {
        return strings.ContainsRune(f.Delims, r)
    })

    numTokens := len(tokens)
    if numTokens <= f.MaxIndex {
        return false
    }
    for i, fi := range f.Indices {
        f.Vars[i].V = strings.TrimSpace(tokens[fi])
    }

    return true
}

func IsFilterDelim(opt Option_t) bool {
    if _, ok := opt["delims"]; !ok {
        return false 
    }
    if _, ok := opt["columns"]; !ok {
        return false 
    }
    return true
}