package main

import (
    "bufio"
    "io"
    "log"
    "os"
    "os/exec"
)

type (
    SyncEvent struct {
        Time        int64
        SeriesID    string
        Key         string
        Value       string
    }
    Event struct {
        Time int64
        Data map[string]interface{}
    }
    Series struct {
        ID     string
        Events []Event
    }
    Cluster struct {
        Key       string
        ID        string
        Group     string
        Token     string
        Series    map[string]*Series
    }
)

func StartCluster(config *Config, events chan SyncEvent) *Cluster {
    cluster := &Cluster{
        Key:   config.Key,
        ID:    HashID(config.Key),
        Group: config.Group,
        Token: HashKey(config.Key),
        Series: make(map[string]*Series),
    }

    log.Printf(">>> STARTING CLUSTER (API: %s) (GROUP: %s) (TOKEN: %s)\n\n", cluster.Key, cluster.Group, cluster.Token)

    data := make(chan string)
    go Read(config, data)

    handlerInfo := &HandlerInfo{Cluster: cluster, Config: config, Data: data, Events: events}
    handler := GetHandler(handlerInfo)
    go handler.Run()

    return cluster
}

func Read(cfg *Config, data chan string) {
    if len(cfg.Cmd) > 0 {
        // Run command and pipe stdout to reader
        cmd := exec.Command("ls", "-l")
        stdout, err := cmd.StdoutPipe()
        if err != nil {
            log.Fatal(err)
        }
        if err := cmd.Start(); err != nil {
            log.Fatal(err)
        }
        ReadToBuffer(stdout, data)
        if err := cmd.Wait(); err != nil {
            log.Fatal(err)
        }
    } else {
        // Pipe stdin to reader
        ReadToBuffer(os.Stdin, data)
    }
}

func ReadToBuffer(reader io.Reader, data chan string) {
    r := bufio.NewReader(reader)

    for {
        line,err := r.ReadString('\n')
        data <- line
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
    }
}
