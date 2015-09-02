package main

import (
    "log"
    "flag"
    "os"
    "strings"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type (
    Config struct {
        Key       string
        Group     string
        Fields    []map[string]string
        Mode      string
        Help      bool
        Debug     bool
    }
)

func LoadConfig() *Config {
    config := &Config{Help: false, Key: os.Getenv("SYNCOPATE_KEY"), Group: os.Getenv("SYNCOPATE_GROUP")}

    configFile := flag.String("c", "syncopate.yaml", "Syncopate YAML config")
    key     := flag.String("k", "", "API key")
    group   := flag.String("g", "", "Group name")
    mode    := flag.String("m", "", "Mode: (regex, csv, ...)")
    help    := flag.Bool("help", false, "Show mode usage")
    debug   := flag.Bool("debug", false, "Debug output")
    flag.Parse()

    // 2nd Priority: YAML Config
    source, err := ioutil.ReadFile(*configFile)
    if err != nil {
        log.Printf("Could not locate %s. Running without config...",*configFile)
    }
    err = yaml.Unmarshal(source, config)
    if err != nil {
        log.Fatal(err)
    }

    // 1st Priority: Command Line
    if *key != "" {
        config.Key = *key
    }
    if *group != "" {
        config.Group = *group
    }
    if *mode != "" {
        config.Mode = strings.ToLower(*mode)
    }
    if *help {
        config.Help = true
    }
    if *debug {
        config.Debug = true
    }

    // Check if config is legal
    if !config.ok() {
        log.Fatalf("Illegal Config: %+v", config)
    } else {
        log.Printf("Running Syncopate with config: %+v", config)
    }

    return config
}

func (c *Config) ok() bool {
    return c.Key != "" && c.Group != "" && c.Mode != "";
}

