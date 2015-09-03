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
    for _,v := range h.Info.Config.Fields {
        h.Fields = append(h.Fields, RegexField{Pattern: v["pattern"], Description: v["desc"]})
        log.Printf("[TRACKING] Field: %s\n", v["desc"])
    }
}

func (h *RegexHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
        /*
        v,l := ExtractValues(data,"CPU usage: {%p:cpu_usage_user} user, {%p:cpu_usage_sys} sys")
        for i,_ := range v {
            fmt.Printf("%s:%s\n",l[i],v[i])
        }
        */
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
    r, _ := regexp.Compile("\\{(%[a-z]):(\\w+)\\}")
    tokens := r.FindAllStringSubmatch(regex, -1)

    result := regex
    var labels []string

    for _,token := range tokens {
        labels = append(labels, token[2])
        switch token[1] {
        case "%p":
            result = strings.Replace(result, token[0], "(\\d*\\.?\\d*)%", 1)
        case "%f":
            result = strings.Replace(result, token[0], "(\\d*\\.?\\d*)", 1)
        case "%d":
            result = strings.Replace(result, token[0], "(\\d+)", 1)
        }
    }
    return result, labels
}

func ExtractValues(data string, regex string) ([]string, []string) {
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
