package main

import (
    "log"
)

type (
    Cluster struct {
        Key       string
        ID        uint64
        Group     string
        Token     uint32
    }
)

func StartCluster(config *Config, events chan SyncEvent) *Cluster {
    cluster := &Cluster{
        Key:   config.Key,
        ID:    Hash64(config.Key),
        Group: config.Group,
        Token: Hash32(config.Key),
    }

    log.Printf(">>> STARTING SYNCOPATE (API: %s) (GROUP: %s) (TOKEN: %d)\n\n", cluster.Key, cluster.Group, cluster.Token)

    data := make(chan string, 256)
    StartCommands(config.CmdInfo, data)

    handlerInfo := &HandlerInfo{Cluster: cluster, Config: config, Data: data, Events: events, KVMap: make(map[string]string)}
    handler := GetHandler(handlerInfo)
    go handler.Run()

    return cluster
}
