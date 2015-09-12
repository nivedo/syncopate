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

    // Start Parser
    parserInfo := &ParserInfo{Config: config, Uploader: uploader, Data: data}
    p := GetParser(parserInfo)
    p.Run()
}
