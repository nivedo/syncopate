package main

import (
    "fmt"
    "log"
    "time"
    "bytes"
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "net/http"
)

const (
    SALT_ID = "1V3S#F"
    SALT_KEY = "AB#*FP"
    //SERVER_URL = "http://localhost:8080"
    //SERVER_URL = "http://api.blub.io:8080"
    SERVER_URL = "http://52.8.222.214:8080"
)

/* Profiling utility */
func Profile(start time.Time, desc string) {
    t := time.Since(start)
    log.Printf("[Profile] %s took %s time", desc, t)
}

func Upload(cluster *Cluster, config *Config) {
    url := fmt.Sprintf("%s/clusters/%s", SERVER_URL, cluster.ID)
    if config.Debug {
        log.Println("[UPLOAD] URL:>", url)
    }

    cjson, err := json.Marshal(cluster)
    if err != nil {
        log.Fatal(err)
    }
    
    if config.Debug {
        log.Println(string(cjson))
    }

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(cjson))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json; charset=UTF-8")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    for k := range cluster.Series {
        delete(cluster.Series, k)
    }
}

func HashID(text string) string {
    h := md5.New()
    h.Write([]byte(SALT_ID + text))
    return hex.EncodeToString(h.Sum(nil)[0:12])
}

func HashKey(text string) string {
    h := md5.New()
    h.Write([]byte(SALT_KEY + text))
    return hex.EncodeToString(h.Sum(nil)[0:6])
}

func MakeSeriesID(token string, group string, desc string) string {
    strID := fmt.Sprintf("%s.%s.%s", token, group, desc)
    hid := HashID(strID)
    return hid
}

func UploadHelper(up chan bool) {
    for {
        time.Sleep(time.Second * 1)
        up <- true
    }
}
