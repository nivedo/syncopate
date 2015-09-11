package main

import (
    "log"
)

type (
    Handler interface {
        Load()
        Run()
        Parse(data string)
        Help()
    }
    HandlerInfo struct {
        Config      *Config
        Uploader    *Uploader
        Data        chan string
    }
)

func GetHandler(info *HandlerInfo) Handler {
    switch info.Config.Mode {
    case "match":
        return NewMatchHandler(info, false)
    case "batch":
        return NewMatchHandler(info, true)
    default:
        log.Fatal("ERROR: No mode specified.")
    }

    return nil
}
