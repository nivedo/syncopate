package main

import (
    "log"
)

type (
    Handler interface {
        parse(data string)
        run()
        help()
    }
    HandlerInfo struct {
        Cluster     *Cluster
        Config      *Config
        Data        chan string
        Events      chan SyncEvent
    }
)

func getHandler(info *HandlerInfo) Handler {
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
