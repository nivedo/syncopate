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
        Variables []Variable
        Matches   map[string]bool
        Mode      string
        Help      bool
    }
)

func LoadConfig() *Config {
    config := &Config{Help: false, Key: os.Getenv("SYNCOPATE_KEY"), Group: os.Getenv("SYNCOPATE_GROUP")}

    configFile := flag.String("c", "syncopate.yaml", "Syncopate YAML config")
    key := flag.String("k", "", "API Key")
    group := flag.String("g", "", "Group Name")
    mode := flag.String("m", "", "Mode: (Regex, CSV, ...)")
    help := flag.Bool("help", false, "Show Mode Usage")
    flag.Parse()

    // 2nd Priority: YAML Config
    source, err := ioutil.ReadFile(*configFile)
    hasConfig := true
    if err != nil {
        log.Printf("Could not locate %s. Running without config...",*configFile)
        hasConfig = false
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

    if !hasConfig {
        LoadDefault(config)
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

func LoadDefault(config *Config) {
    switch config.Mode {
    case "top":
        LoadDefaultTop(config)
    default:
    }
}

func LoadDefaultTop(config *Config) {
    defaults := []string{
        "cpu_usage_user",
        "cpu_usage_sys",
        "cpu_usage_idle"}
    config.Matches = make(map[string]bool)
    for _, s := range defaults {
        config.Matches[s] = true
    }
}


