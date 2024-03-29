package main

import (
    "log"
    "strings"
    "regexp"
    "fmt"
    "time"
    "bytes"
    "strconv"
    "net/http"
    "net/url"
)

const (
    S_INT    = 1 << iota
    S_FLOAT  = 1 << iota
    S_CHAR   = 1 << iota
    S_UNKNOWN = 0
    SYNC_URL = "http://localhost:8000/cluster-sync/"
)

type (
    KVPair struct {
        K     string
        V     string
        Force bool
        Type  uint8
    }
    KVList []KVPair
    UploadEvent struct {
        Time        int64
        ID          uint64
        Key         string
        Value       string
        Type        uint8
    }
    Uploader struct {
        Config      *Config
        Token       uint32
        Events      chan UploadEvent
        LastVal     map[string]string
    }
)

func GetType(v string) uint8 {
    regInt,_ := regexp.Compile("^\\d+$")
    if regInt.MatchString(v) {
        return S_INT
    }
    regFloat,_ := regexp.Compile("^(?:\\d+\\.?\\d*)$|^(?:\\.\\d+)$")
    if regFloat.MatchString(v) {
        return S_FLOAT
    }
    return S_CHAR
}

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
    // Trim whitespace
    newId := strings.TrimSpace(rawId)

    // Convert #, %
    newId = strings.Replace(rawId, "#", "num_", -1)
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

func (u *Uploader) SyncKV(list KVList) bool {
    data := url.Values{}
    data.Set("group",u.Config.Group)
    for _,v := range list {
        data.Add("topic",v.K)
    }

    client := &http.Client{}
    r, err := http.NewRequest("POST", SYNC_URL, bytes.NewBufferString(data.Encode()))
    if err != nil {
        log.Printf("API Key: %s request failed.", u.Config.Key)
        return false
    }
    r.Header.Add("Authorization", "OAuth " + u.Config.Key)
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    resp, _ := client.Do(r)
    if resp.StatusCode != http.StatusNoContent {
        log.Printf("API Key: %s synchronization failure.", u.Config.Key)
        return false
    }

    log.Printf("API Key: %s synchronized with %+v",u.Config.Key, data)
    return true
}

func (u *Uploader) UploadKV(list KVList) {
    now := time.Now().UTC().UnixNano() / int64(time.Microsecond)
    if u.Config.Debug {
        log.Printf("[UploadKV] Processing %+v", list)
    }
    for _, v := range list {
        if v.Type == S_UNKNOWN {
            log.Fatalf("Unknown type for var %+v", v)
        }
        if v.Force || u.LastVal[v.K] != v.V {
            if u.Config.Debug {
                log.Printf("[UploadKV] Uploading KV: %s, %s, %d", v.K, v.V, v.Type)
            }
            u.LastVal[v.K] = v.V
            u.Events <- UploadEvent{
                ID:          HashSeriesID(u.Token, u.Config.Group, v.K),
                Key:         v.K,
                Value:       v.V,
                Time:        now,
                Type:        v.Type,
            }
        }
    }
}
