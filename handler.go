package main

import (
    "log"
    "strings"
    "fmt"
    "time"
)

type (
    KVPair struct {
        K string
        V string
    }
    KVList []KVPair
    Handler interface {
        Load()
        Run()
        Parse(data string)
        Help()
    }
    HandlerInfo struct {
        Cluster     *Cluster
        Config      *Config
        Data        chan string
        Events      chan SyncEvent
    }
)

func GetHandler(info *HandlerInfo) Handler {
    switch info.Config.Mode {
    case "regex":
        return NewRegexHandler(info)
    case "top":
        return NewTopHandler(info)
    case "csv":
        return NewCsvHandler(info)
    default:
        log.Fatal("ERROR: No mode specified.")
    }

    return nil
}

func ConvertToValidSeriesKey(rawId string) string {
    // Convert #, %
    newId := strings.Replace(rawId, "#", "num_", -1)
    newId = strings.Replace(newId, "%", "pct_", -1)

    // Conver space to _
    newId = strings.Replace(newId, " ", "_", -1)
    newId = strings.ToLower(newId)

    return newId
}

func (m KVList) Print() {
    fmt.Printf("{ size: %d\n", len(m))
    for _,v := range m {
        fmt.Printf("%20s: %s\n",v.K,v.V)
    }
    fmt.Println("}")
}

func (info *HandlerInfo) Upload(m KVList) {
    now := time.Now().UTC().UnixNano() / int64(time.Microsecond)
    for _, v := range m {
        seriesID := MakeSeriesID(info.Cluster.Token, info.Cluster.Group, v.K)
        info.Events <- SyncEvent{
            SeriesID:    seriesID,
            Key:         v.K,
            Value:       v.V,
            Time:        now}
    }
}


