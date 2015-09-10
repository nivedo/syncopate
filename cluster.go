package main

import (
    "log"
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

    data := make(chan string, 256)
    StartCommands(config.CmdInfo, data)

    handlerInfo := &HandlerInfo{Cluster: cluster, Config: config, Data: data, Events: events, KVMap: make(map[string]string)}
    handler := GetHandler(handlerInfo)
    go handler.Run()

    return cluster
}
