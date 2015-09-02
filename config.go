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

    // 2nd Priority: YAML Config
    configFile := flag.String("c", "syncopate.yaml", "config")
    source, err := ioutil.ReadFile(*configFile)
    if err != nil {
        log.Printf("Could not locate %s. Running without config...",*configFile)
    }
    err = yaml.Unmarshal(source, config)
    if err != nil {
        log.Fatal(err)
    }

    // 1st Priority: Command Line
    

    if !config.ok() {
        log.Fatalf("Illegal Config: %+v", config)
    }

    return config
}

func (c *Config) ok() bool {
    return c.Key != "" && c.Group != "" && c.Mode != "";
}