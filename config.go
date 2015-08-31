package main

import (
    "log"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type (
    Config struct {
        Key       string
        Data      []WatchFile
        NumSeries int
    }
)

func loadConfig() Config {
    var config Config
    source, err := ioutil.ReadFile("syncopate.yaml")
    if err != nil {
        log.Fatal("Could not locate syncopate.yaml. Terminating...")
    }
    err = yaml.Unmarshal(source, &config)
    if err != nil {
        log.Fatal(err)
    }
    for _,wf := range config.Data {
        config.NumSeries += len(wf.Variables)
    }

    return config
}