package main

import (
    "log"
    "strings"
    "fmt"
    "time"
)

type (
    KVPair struct {
        K     string
        V     string
        Force bool
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
        KVMap       map[string]string
    }
)

func GetHandler(info *HandlerInfo) Handler {
    switch info.Config.Mode {
    case "match", "csv":
        return NewMatchHandler(info, false)
    case "batch", "top", "df":
        return NewMatchHandler(info, true)
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

func CreateKVPair(key string, value string) KVPair {
    return KVPair{K: ConvertToValidSeriesKey(key), V: strings.TrimSpace(value)}
}

func (list *KVList) Print() {
    fmt.Printf("{ size: %d\n", len(*list))
    for _,v := range *list {
        fmt.Printf("%20s: %s\n",v.K,v.V)
    }
    fmt.Println("}")
}

func UploadKV(list KVList, info *HandlerInfo) {
    now := time.Now().UTC().UnixNano() / int64(time.Microsecond)
    for _, v := range list {
        if v.Force || info.KVMap[v.K] != v.V {
            if info.Config.Debug {
                log.Printf("[UploadKV] Uploading KV: %s, %s", v.K, v.V)
            }
            seriesID := MakeSeriesID(info.Cluster.Token, info.Cluster.Group, v.K)
            info.KVMap[v.K] = v.V
            info.Events <- SyncEvent{
                SeriesID:    seriesID,
                ID:          HashSeriesID(info.Cluster.Token, info.Cluster.Group, v.K),
                Key:         v.K,
                Value:       v.V,
                Time:        now}
        }
    }
}

