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

func (h *RegexHandler) run() {
    for {
        data := <-h.Info.Data
        h.parse(data)
    }
}

func (h *RegexHandler) parse(data string) {
    vars := h.Info.Config.Variables
    for i, _ := range vars {
        match, _ := regexp.MatchString(vars[i].Pattern, data)
        if match {
            r, _ := regexp.Compile(vars[i].Pattern)
            seriesID := makeSeriesID(h.Info.Cluster.Token, h.Info.Cluster.Group, vars[i].Description)
            allMatch := r.FindAllStringSubmatch(data, -1)
            for _, matchVal := range allMatch {
                h.Info.Events <- SyncEvent{
                    SeriesID:    seriesID,
                    SeriesIndex: i,
                    Key:         vars[i].Description,
                    Value:       matchVal[1],
                    Time:        time.Now().UTC().UnixNano() / int64(time.Microsecond)}
            }
        }
    }
}

func (h *RegexHandler) help() {
    fmt.Println("---- Regex Help ----\n")
    fmt.Println("Decimals: (\\d*\\.?\\d*)")
    fmt.Println("Integers: (\\d+)")
}