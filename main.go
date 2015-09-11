package main

import (
    "log"
)

func main() {
    config := LoadConfig()
    log.Printf(">>> STARTING SYNCOPATE (API: %s) (GROUP: %s) (TOKEN: %d)\n\n", config.Key, config.Group, Hash32(config.Key))

    // Start Syncopate Commands
    data := make(chan string, 256)
    go StartCommands(config.CmdInfo, data)

    // Start Uploader
    uploader := NewUploader(config)
    go uploader.Start()

    // Start Handler
    // handlerInfo := &HandlerInfo{Config: config, Uploader: uploader, Data: data}
    // handler := GetHandler(handlerInfo)
    // handler.Run()

    parserInfo := &ParserInfo{Config: config, Uploader: uploader, Data: data}
    p := NewAnyParser(parserInfo)
    p.Run()
}
