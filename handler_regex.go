package main

import (
    "regexp"
    "time"
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