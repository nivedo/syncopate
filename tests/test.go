package main

import (
    "log"
    "regexp"
    "strings"
)

const (
    S_INT    = 1 << iota
    S_FLOAT  = 1 << iota
    S_CHAR   = 1 << iota
)

type (
    FilterVar struct {
        Type    int     // Optional
        Name    string  // Optional
        Rule    string  // Required
    }
)

// Example usage: {{ (int) my_var_name:my_var_rule }}
func GetFilterVars(desc string) []FilterVar {
    var vars []FilterVar
    r := regexp.MustCompile("\\{\\{\\s*(?:\\((?P<type>\\w+)\\))?(?:\\s*(?P<name>\\w+)\\s*:)?\\s*(?P<rule>.+?)\\}\\}")
    matches := r.FindAllStringSubmatch(desc,-1)
    
    for _,match := range matches {
        f := FilterVar{
            Name: strings.ToLower(strings.TrimSpace(match[2])), 
            Rule: strings.TrimSpace(match[3]),
        }
        switch match[1] {
        case "int":
            f.Type = S_INT
        case "float":
            f.Type = S_FLOAT
        default:
            f.Type = S_CHAR
        }
        vars = append(vars, f)
    }

    return vars
}

func main() {
    log.Println(GetFilterVars("{{ $1 }} {{ hello_world:$abcd_efgh }} {{ (int) $abcd_efgh }}"))
}