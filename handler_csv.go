package main

import (
    //"log"
    "fmt"
    "regexp"
    "strings"
    "strconv"
    //"time"
)

type (
    ColNameMap map[int]string
    CsvHandler struct {
        Info        *HandlerInfo
        Map         KVMap
        ColNames    ColNameMap
        ColRequests []string
        LineCount   int
    }
)

func NewCsvHandler(info *HandlerInfo) *CsvHandler {
    h := &CsvHandler{Info :info}
    h.Load()
    return h
}

func (h *CsvHandler) Load() {
    defaultNumCols := 5
    cFields := h.Info.Config.Fields

    if len(cFields) == 0 {
        for i := 0; i < defaultNumCols; i++ {
            h.ColNames[i] = fmt.Sprintf("col%d", i)
        }
    } else {
        for _, v := range cFields {
            reg, _ := regexp.Compile("^\\$[0-9]+$")
            if reg.Match([]byte(v["pattern"])) {
                // Matches $column_index pattern
                num, err := strconv.ParseInt(v["pattern"][1:], 10, 64)
                if err == nil {
                    h.ColNames[int(num)] = v["desc"]
                }
            } else {
                // Column name specified
                h.ColRequests = append(h.ColRequests, v["desc"]) 
            }
        }
    }
}

func (h *CsvHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
    }
}

func (h *CsvHandler) ParseLine(data string) {
    if len(h.ColNames) > 0 {
        tokens := strings.Split(data, ",")
        numCols := len(tokens)
        for i, k := range h.ColNames {
            if i >= 0 && i < numCols {
                h.Map[k] = tokens[i]
            }
        }
    }
}

func (h *CsvHandler) Parse(data string) {
    // First line
    if h.LineCount == 0 {

    }

    h.LineCount++
}

func (h *CsvHandler) Upload() {
    /*
    now := time.Now().UTC().UnixNano() / int64(time.Microsecond)
    for k, v := range h.Map {
        seriesID := MakeSeriesID(h.Info.Cluster.Token, h.Info.Cluster.Group, k)
        h.Info.Events <- SyncEvent{
            SeriesID:    seriesID,
            Key:         k,
            Value:       ConvertToUnit(v),
            Time:        now}
    }
    */
}


