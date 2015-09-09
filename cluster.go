package main

import (
    "bufio"
    "io"
    "log"
    "os"
    "os/exec"
    "time"
)

type (
    SyncEvent struct {
        Time        int64
        ID          uint64
        Token       uint32
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
        ID        uint64
        Group     string
        Token     uint32
        Series    map[string]*Series
    }
)

func StartCluster(config *Config, events chan SyncEvent) *Cluster {
    cluster := &Cluster{
        Key:   config.Key,
        ID:    Hash64(config.Key),
        Group: config.Group,
        Token: Hash32(config.Key),
        Series: make(map[string]*Series),
    }

    log.Printf(">>> STARTING CLUSTER (API: %s) (GROUP: %s) (TOKEN: %d)\n\n", cluster.Key, cluster.Group, cluster.Token)

    data := make(chan string)
    go Read(config, data)

    handlerInfo := &HandlerInfo{Cluster: cluster, Config: config, Data: data, Events: events, KVMap: make(map[string]string)}
    handler := GetHandler(handlerInfo)
    go handler.Run()

    return cluster
}

func Read(cfg *Config, data chan string) {
    if len(cfg.CmdBin) > 0 {
        if cfg.CmdWatchSec > 0 {
            // Run on watch timer
            var sleepTime = time.Duration(cfg.CmdWatchSec * 1e9)
            for {
                RunAndReadToBuffer(cfg.CmdBin, cfg.CmdArgs, data, true)
                time.Sleep(sleepTime)
            }
        } else {
            // Run once
            RunAndReadToBuffer(cfg.CmdBin, cfg.CmdArgs, data, false)
        }
    } else {
        // Pipe stdin to reader
        ReadToBuffer(os.Stdin, data, false)
    }
}

func RunAndReadToBuffer(bin string, args []string, data chan string, stopAtEOF bool) {
    // Run command and pipe stdout to reader
    cmd := exec.Command(bin, args...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    ReadToBuffer(stdout, data, stopAtEOF)
    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }
}

func ReadToBuffer(reader io.Reader, data chan string, stopAtEOF bool) {
    r := bufio.NewReader(reader)

    for {
        line,err := r.ReadString('\n')
        data <- line
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }
        if stopAtEOF && err == io.EOF {
            break
        }
    }
}
