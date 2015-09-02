package main

import (
    "log"
    "flag"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type (
    Config struct {
        Key       string
        Group     string
        Variables []Variable
        Mode      string
        Help      bool
    }
)

func LoadConfig() *Config {
    config := &Config{Help: false}

    configFile := flag.String("c", "syncopate.yaml", "config")
    source, err := ioutil.ReadFile(*configFile)
    if err != nil {
        log.Fatal("Could not locate syncopate.yaml. Terminating...")
    }
    err = yaml.Unmarshal(source, config)
    if err != nil {
        log.Fatal(err)
    }

    return config
}