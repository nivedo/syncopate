package main

import (
    "log"
    "strings"
    "fmt"
    "time"
)

type (
    KVMap map[string]string
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

func (m KVMap) Print() {
    fmt.Printf("{ size: %d\n", len(m))
    for k,v := range m {
        fmt.Printf("%20s: %s\n",k,v)
    }
    fmt.Println("}")
}

func (info *HandlerInfo) Upload(m KVMap) {
    now := time.Now().UTC().UnixNano() / int64(time.Microsecond)
    for k, v := range m {
        seriesID := MakeSeriesID(info.Cluster.Token, info.Cluster.Group, k)
        info.Events <- SyncEvent{
            SeriesID:    seriesID,
            Key:         k,
            Value:       v,
            Time:        now}
    }
}


