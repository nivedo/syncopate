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
    UploadEvent struct {
        Time        int64
        ID          uint64
        Key         string
        Value       string
    }
    Uploader struct {
        Config      *Config
        Token       uint32
        Events      chan UploadEvent
        LastVal     map[string]string
    }
)

func NewUploader(config *Config) *Uploader {
    return &Uploader{
        Config: config,
        Token: Hash32(config.Key),
        Events: make(chan UploadEvent), 
        LastVal: make(map[string]string),
    }
}

func (kv *KVPair) String() string {
    if kv.Force {
        return fmt.Sprintf("{ %s: %s F }", kv.K, kv.V)
    } else {
        return fmt.Sprintf("{ %s: %s }", kv.K, kv.V)
    }
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

func (u *Uploader) Start() {
    d := NewTCPDispatcher(u.Config.Key, u.Events)
    d.Run()
}

func (u *Uploader) UploadKV(list KVList) {
    now := time.Now().UTC().UnixNano() / int64(time.Microsecond)
    for _, v := range list {
        if v.Force || u.LastVal[v.K] != v.V {
            if u.Config.Debug {
                log.Printf("[UploadKV] Uploading KV: %s, %s", v.K, v.V)
            }
            u.LastVal[v.K] = v.V
            u.Events <- UploadEvent{
                ID:          HashSeriesID(u.Token, u.Config.Group, v.K),
                Key:         v.K,
                Value:       v.V,
                Time:        now}
        }
    }
}
