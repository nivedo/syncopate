package main

import (
    "regexp"
    "time"
    "fmt"
)

type (
    RegexHandler struct {
        Info *HandlerInfo
    }
)

func NewRegexHandler(info *HandlerInfo) *RegexHandler {
    return &RegexHandler{Info :info}
}

func (h *RegexHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
    }
}

func (h *RegexHandler) Parse(data string) {
    vars := h.Info.Config.Variables
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