package main

import (
    "log"
    "flag"
    "os"
    "time"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type (
    Option_t map[string]string
    Config struct {
        Key         string
        Group       string
        Options     []Option_t
        CmdInfo     *CommandInfo
        Mode        string
        Help        bool
        Debug       bool
    }
)

func LoadConfig() *Config {
    config := &Config{Help: false, Key: os.Getenv("SYNCOPATE_KEY"), Group: os.Getenv("SYNCOPATE_GROUP")}

    configFile  := flag.String("c", "syncopate.yaml", "Syncopate YAML config")
    runCmd      := flag.String("r", "", "Command to run")
    watchSec    := flag.Float64("w", -1.0, "Watch cycle time (in seconds)")
    key         := flag.String("k", "", "API key")
    group       := flag.String("g", "", "Group name")
    mode        := flag.String("m", "", "Mode: (csv, top, ...)")
    help        := flag.Bool("help", false, "Show mode usage")
    debug       := flag.Bool("debug", false, "Debug output")
    flag.Parse()

    // 2nd Priority: YAML Config
    source, err := ioutil.ReadFile(*configFile)
    configNotFound := false
    if err != nil {
        log.Printf("Could not locate %s. Running without config...",*configFile)
        configNotFound = true
    } else {
        err = yaml.Unmarshal(source, config)
        if err != nil {
            log.Fatal(err)
        }
    }

    // Command to run will set default mode, which can be overridden
    // as command line argument
    config.CmdInfo = NewCommandInfo(*runCmd, time.Duration(*watchSec) * time.Second)

    // Try embedded static configs
    if configNotFound && config.CmdInfo != nil {
        switch config.CmdInfo.Type {
        case CMD_TOP:
            src, _ := Asset("configs/top.yaml")
            err = yaml.Unmarshal(src, config)
            if err != nil {
                log.Fatal(err)
            }
        case CMD_DF:
            src, _ := Asset("configs/df.yaml")
            err = yaml.Unmarshal(src, config)
            if err != nil {
                log.Fatal(err)
            }
        }
    }

    // 1st Priority: Command Line
    if *key != "" {
        config.Key = *key
    }
    if *group != "" {
        config.Group = *group
    }
    if *mode != "" {
        config.Mode = *mode
    }
    if *help {
        config.Help = true
    }
    if *debug {
        config.Debug = true
    }

    // Check if config is legal
    if !config.IsValid() {
        log.Fatalf("Illegal Config: %+v", config)
    }

    return config
}

func (c *Config) IsValid() bool {
    if c.Key == "" {
        log.Fatal("Missing key.")
        return false
    }
    if c.Group == "" {
        log.Fatal("Missing group.")
        return false
    }
    if c.Mode == "" {
        log.Fatal("Missing mode.")
        return false
    }
    if len(c.Options) == 0 {
        log.Fatal("No options specified.")
        return false
    }
    return true
}

