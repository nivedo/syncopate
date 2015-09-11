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
        Mode        string
        Help        bool
        Debug       bool
        Options     []Option_t
        CmdInfo     *CommandInfo
    }
    Args struct {
        Config      string
        Command     string
        WatchSec    float64
        Key         string
        Group       string
        Mode        string
        Help        bool
        Debug       bool
    }
)

func ParseFlags() *Args {
    f := &Args{}

    flag.StringVar(&f.Config, "c", "syncopate.yaml", "Syncopate YAML config")
    flag.StringVar(&f.Command, "r", "", "Command to run")
    flag.Float64Var(&f.WatchSec, "w", -1.0, "Watch cycle time (in seconds)")
    flag.StringVar(&f.Key, "k", "", "API key")
    flag.StringVar(&f.Group, "g", "", "Group name")
    flag.StringVar(&f.Mode, "m", "", "Mode: (csv, top, ...)")
    flag.BoolVar(&f.Help, "help", false, "Show mode usage")
    flag.BoolVar(&f.Debug, "debug", false, "Debug output")

    flag.Parse()
    return f
}

func LoadConfig() *Config {
    config := &Config{Help: false, Key: os.Getenv("SYNCOPATE_KEY"), Group: os.Getenv("SYNCOPATE_GROUP")}
    args := ParseFlags()

    // 2nd Priority: YAML Config
    source, err := ioutil.ReadFile(args.Config)
    configNotFound := false
    if err != nil {
        log.Printf("Could not locate %s. Running without config...",args.Config)
        configNotFound = true
    } else {
        err = yaml.Unmarshal(source, config)
        if err != nil {
            log.Fatal(err)
        }
    }

    // Command to run will set default mode, which can be overridden
    // as command line argument
    config.CmdInfo = NewCommandInfo(args.Command, time.Duration(args.WatchSec) * time.Second)

    // Try embedded static configs
    if configNotFound && config.CmdInfo != nil {
        config.SetDefaults()
    }

    // 1st Priority: Command Line
    config.Override(args)

    // Check if config is legal
    if !config.IsValid() {
        log.Fatalf("Illegal Config: %+v", config)
    }

    return config
}

func (c *Config) SetDefaults() {
    switch c.CmdInfo.Type {
    case CMD_TOP:
        src, _ := Asset("configs/top.yaml")
        err := yaml.Unmarshal(src, c)
        if err != nil {
            log.Fatal(err)
        }
    case CMD_DF:
        src, _ := Asset("configs/df.yaml")
        err := yaml.Unmarshal(src, c)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func (c *Config) Override(args *Args) {
    if args.Key != "" {
        c.Key = args.Key
    }
    if args.Group != "" {
        c.Group = args.Group
    }
    if args.Mode != "" {
        c.Mode = args.Mode
    }
    if args.Help {
        c.Help = true
    }
    if args.Debug {
        c.Debug = true
    }
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

