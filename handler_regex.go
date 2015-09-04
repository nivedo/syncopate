package main

import (
    "regexp"
    "time"
    "strings"
    "fmt"
    "log"
)

type (
    RegexHandler struct {
        Info     *HandlerInfo
        Fields   []RegexField
    }
    RegexField struct {
        Pattern     string
        Description string
    }
)

func NewRegexHandler(info *HandlerInfo) *RegexHandler {
    h := &RegexHandler{Info :info}
    h.Load()
    return h
}

func (h *RegexHandler) Load() {
    for _,v := range h.Info.Config.Options {
        h.Fields = append(h.Fields, RegexField{Pattern: v["pattern"], Description: v["desc"]})
        log.Printf("[TRACKING] Field: %s\n", v["desc"])
    }
}

func (h *RegexHandler) Run() {
    var list KVList
    rule := NewRuleRegex("CPU usage: {{ cpu_usage_user:%p }} user, {{ cpu_usage_sys:%p }} sys")
    rule2 := NewRuleRegex("Load Avg: {{ load_1:%f }}, {{ load_2:%f }}, {{ load_3:%f }}")
    list = make(KVList,10)
    for {
        data := <-h.Info.Data
        //h.Parse(data)
        /*
        v,l := ExtractRegex(data,"CPU usage: {{ cpu_usage_user:%p }} user, {{ cpu_usage_sys:%p }} sys")
        //v,l := ExtractCSV(data,"{{ $1,3,5,6,2:c1,c3 }}")
        for i,_ := range v {
            fmt.Printf("%s:%s\n",l[i],v[i])
        }
        */
        if(rule.Eval(data, &list, 0) > 0) {
            fmt.Println(list)
        }
        if(rule2.Eval(data, &list, 0) > 0) {
            fmt.Println(list)
        }
    }
}

func (h *RegexHandler) Parse(data string) {
    vars := h.Fields
    for i, _ := range vars {
        match, _ := regexp.MatchString(vars[i].Pattern, data)
        if match {
            r, _ := regexp.Compile(vars[i].Pattern)
            seriesID := MakeSeriesID(h.Info.Cluster.Token, h.Info.Cluster.Group, vars[i].Description)
            allMatch := r.FindAllStringSubmatch(data, -1)
            for _, matchVal := range allMatch {
                h.Info.Events <- SyncEvent{
                    SeriesID:    seriesID,
                    Key:         vars[i].Description,
                    Value:       matchVal[1],
                    Time:        time.Now().UTC().UnixNano() / int64(time.Microsecond)}
            }
        }
    }
}

func (h *RegexHandler) Help() {
    fmt.Println("---- Regex Help ----\n")
    fmt.Println("Decimals: (\\d*\\.?\\d*)")
    fmt.Println("Integers: (\\d+)")
}

func FormatRegex(regex string) (string, []string) {
    //r, _ := regexp.Compile("\\{(%[a-z]):(\\w+)\\}")
    r, _ := regexp.Compile("\\{\\{\\s*(\\w+):(.+?)\\}\\}")
    tokens := r.FindAllStringSubmatch(regex, -1)

    result := regex
    var labels []string

    for _,token := range tokens {
        labels = append(labels, strings.TrimSpace(token[1]))
        rule := strings.TrimSpace(token[2])
        switch rule {
        case "%p":
            result = strings.Replace(result, token[0], "(\\d*\\.?\\d*)%", 1)
        case "%f":
            result = strings.Replace(result, token[0], "(\\d*\\.?\\d*)", 1)
        case "%d":
            result = strings.Replace(result, token[0], "(\\d+)", 1)
        default:
            // Use user specified regex
            result = strings.Replace(result, token[0], rule, 1)
        }
    }

    // Whitespace is arbitrary
    r2, _ := regexp.Compile("\\s+")
    result = r2.ReplaceAllString(result, "\\s+")

    return result, labels
}

func ExtractRegex(data string, regex string) ([]string, []string) {
    fr, labels := FormatRegex(regex)
    match, _ := regexp.MatchString(fr, data)
    if match {
        r, _ := regexp.Compile(fr)
        allMatch := r.FindAllStringSubmatch(data, -1)
        return allMatch[0][1:], labels
    } else {
        return nil, nil
    }
}
