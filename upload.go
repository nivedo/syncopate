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
    SERVER_URL = "http://api.blub.io:32794"
)

func upload(cluster *Cluster, key string) {
    url := fmt.Sprintf("%s/clusters/%s", SERVER_URL, key)
    log.Println("[UPLOAD] URL:>", url)

    cjson, err := json.Marshal(cluster)
    if err != nil {
        log.Fatal(err)
    }

    log.Println(string(cjson))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(cjson))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    for i, _ := range cluster.Series {
        cluster.Series[i].Events = nil
    }
}

func hashID(text string) string {
    h := md5.New()
    h.Write([]byte(SALT_ID + text))
    return hex.EncodeToString(h.Sum(nil)[0:12])
}

func hashKey(text string) string {
    h := md5.New()
    h.Write([]byte(SALT_KEY + text))
    return hex.EncodeToString(h.Sum(nil)[0:6])
}

func makeSeriesID(token string, group string, desc string) string {
    strID := fmt.Sprintf("%s.%s.%s", token, group, desc)
    hid := hashID(strID)
    return hid
}

func uploadHelper(events chan WatchEvent) {
    for {
        time.Sleep(time.Second * 1)
        events <- WatchEvent{SeriesIndex: -1}
    }
}