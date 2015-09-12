package main

import (
    "log"
    "regexp"
    "strings"
)

type (
    Filter interface {
        // Match returns true if the filter has been completed.
        // This may take multiple calls as filters are "stateful"
        // and can handle multiple lines.
        Match(data string) bool
        // GetVars() returns a KVList of the last known values
        // to be successfully matched.
        GetVars() KVList
    }
    FilterVar struct {
        Desc    string
        Type    uint8   // Optional
        Name    string  // Optional
        Rule    string  // Required
    }
)

func GetFilter(opt Option_t) Filter {
    if IsFilterRegex(opt) {
        return NewFilterRegex(opt)
    }
    if IsFilterTable(opt) {
        return NewFilterTable(opt)
    }

    log.Fatalf("Filter doesn't exist for option %+v", opt)
    return nil
}

// Example usage: {{ (int) my_var_name:my_var_rule }}
func GetFilterVars(desc string) []FilterVar {
    var vars []FilterVar
    r := regexp.MustCompile("\\{\\{\\s*(?:\\((?P<type>\\w+)\\))?(?:\\s*(?P<name>\\w+)\\s*:)?\\s*(?P<rule>.+?)\\}\\}")
    matches := r.FindAllStringSubmatch(desc,-1)
    
    for _,match := range matches {
        f := FilterVar{
            Desc: strings.TrimSpace(match[0]),
            Name: ConvertToValidSeriesKey(match[2]), 
            Rule: strings.TrimSpace(match[3]),
        }
        switch match[1] {
        case "int":
            f.Type = S_INT
        case "float":
            f.Type = S_FLOAT
        case "string":
            f.Type = S_CHAR
        }
        vars = append(vars, f)
    }

    return vars
}