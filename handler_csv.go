package main

import (
    "log"
    "fmt"
    "regexp"
    "strings"
    "strconv"
    "time"
)

type (
    IKMap map[int]string
    CsvHandler struct {
        Info        *HandlerInfo
        Map         KVMap
        IndexMap    IKMap
        LineCount   int
    }
)

func NewCsvHandler(info *HandlerInfo) *CsvHandler {
    h := &CsvHandler{Info :info}
    h.Map = make(KVMap)
    h.IndexMap = make(IKMap)
    h.Load()
    return h
}

func (h *CsvHandler) Help() {
    // List all variables
    fmt.Println("---- Csv Help ----\n")
    for i, k := range h.IndexMap {
        fmt.Printf("%5s: %s", strconv.Itoa(i), k)
    }
}

func (h *CsvHandler) Load() {
    defaultNumCols := 5
    cFields := h.Info.Config.Fields

    if len(cFields) == 0 {
        for i := 0; i < defaultNumCols; i++ {
            h.IndexMap[i] = fmt.Sprintf("col%d", i)
        }
    } else {
        for _, v := range cFields {
            reg, _ := regexp.Compile("^\\$[0-9]+$")
            if reg.Match([]byte(v["pattern"])) {
                // Matches $column_index pattern
                num, err := strconv.ParseInt(v["pattern"][1:], 10, 64)
                if err == nil {
                    h.IndexMap[int(num)] = v["desc"]
                }
            }
        }
    }
    for i, k := range h.IndexMap {
        log.Printf("[TRACKING] Index: %d, Name: %s\n", i, k) 
    }
}

func (h *CsvHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
    }
}

func (h *CsvHandler) ParseLine(data string) {
    if len(h.IndexMap) > 0 {
        tokens := strings.Split(data, ",")
        numCols := len(tokens)
        for i, k := range h.IndexMap {
            if i >= 0 && i < numCols {
                h.Map[k] = strings.TrimSpace(tokens[i])
            }
        }
        h.Map.Print()
    }
}

func (h *CsvHandler) Parse(data string) {
    h.ParseLine(data)
    h.Upload()
    h.LineCount++
}

func (h *CsvHandler) Upload() {
    now := time.Now().UTC().UnixNano() / int64(time.Microsecond)
    for k, v := range h.Map {
        seriesID := MakeSeriesID(h.Info.Cluster.Token, h.Info.Cluster.Group, k)
        h.Info.Events <- SyncEvent{
            SeriesID:    seriesID,
            Key:         k,
            Value:       v,
            Time:        now}
    }
}


