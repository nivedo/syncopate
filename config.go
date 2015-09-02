package main

import (
    "log"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type (
    Config struct {
        Key       string
        Group     string
        Variables []Variable
        Mode      string
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

    return config
}