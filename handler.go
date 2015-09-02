package main

import (
    "log"
)

type (
    Handler interface {
        Parse(data string)
        Run()
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
    default:
        log.Fatal("No mode available.")
    }

    return nil
}
