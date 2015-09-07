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
        Options   []map[string]string
        CmdBin    string
        CmdArgs   []string
        Mode      string
        Help      bool
        Debug     bool
    }
)

func (config *Config) SetCommand(cmd string) {
    runCmd := strings.TrimFunc(cmd, func(r rune) bool {
        return r == '"' || r == '\''
    })
    tokens := strings.Fields(runCmd)
    config.CmdBin = tokens[0]
    config.CmdArgs = tokens[1:len(tokens)]
    switch config.CmdBin {
    case "top":
        config.Mode = "top"
        break
    default:
        config.Mode = "regex"
        break
    }
}

func LoadConfig() *Config {
    config := &Config{Help: false, Key: os.Getenv("SYNCOPATE_KEY"), Group: os.Getenv("SYNCOPATE_GROUP")}

    configFile := flag.String("c", "syncopate.yaml", "Syncopate YAML config")
    runCmd  := flag.String("r", "", "Command to run")
    key     := flag.String("k", "", "API key")
    group   := flag.String("g", "", "Group name")
    mode    := flag.String("m", "", "Mode: (regex, csv, ...)")
    help    := flag.Bool("help", false, "Show mode usage")
    debug   := flag.Bool("debug", false, "Debug output")
    flag.Parse()

    // 2nd Priority: YAML Config
    source, err := ioutil.ReadFile(*configFile)
    configNotFound := false
    if err != nil {
        log.Printf("Could not locate %s. Running without config...",*configFile)
        configNotFound = true
    }
    err = yaml.Unmarshal(source, config)
    if err != nil {
        log.Fatal(err)
    }

    // 1st Priority: Command Line
    // Command to run takes precedence over mode
    if *runCmd != "" {
        config.SetCommand(*runCmd)
    } else if *mode != "" {
        config.Mode = strings.ToLower(*mode)
    }

    if *key != "" {
        config.Key = *key
    }
    if *group != "" {
        config.Group = *group
    }
    if *help {
        config.Help = true
    }
    if *debug {
        config.Debug = true
    }

    // Try embedded static configs
    if configNotFound {
        switch config.Mode {
        case "top":
            src, _ := Asset("configs/top.yaml")
            err = yaml.Unmarshal(src, config)
            if err != nil {
                log.Fatal(err)
            }
        }
    }

    // Check if config is legal
    if !config.ok() {
        log.Fatalf("Illegal Config: %+v", config)
    }

    return config
}

func (c *Config) ok() bool {
    return c.Key != "" && c.Group != "" && c.Mode != "";
}

