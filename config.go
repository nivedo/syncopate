package main

import (
    "log"
    "flag"
    "os"
    "strings"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "runtime"
    "math"
)

type (
    Config struct {
        Key       string
        Group     string
        Options   []map[string]string
        WatchSec  float64
        CmdBin    string
        CmdArgs   []string
        Mode      string
        Help      bool
        Debug     bool
    }
)

func (config *Config) SetCommandMode(cmd string, mode string) {
    if cmd != "" {
        runCmd := strings.TrimFunc(cmd, func(r rune) bool {
            return r == '"' || r == '\''
        })
        tokens := strings.Fields(runCmd)
        config.CmdBin = tokens[0]
        if len(tokens) > 1 {
            config.CmdArgs = tokens[1:len(tokens)]
        }
        switch config.CmdBin {
        case "top":
            config.Mode = "top"
            // Make sure batch mode, if not, add batch mode argument
            switch runtime.GOOS {
            case "darwin":
                config.SetRequiredArgument("-l", []string{"-l","0"})
                break
            case "linux":
                config.SetRequiredArgument("-b", []string{"-b"})
                break
            default:
                break
            }
            break
        case "df":
            config.SetRequiredWatchSec(2.0)
            break
        case "du":
            config.SetRequiredWatchSec(2.0)
            break
        default:
            if len(tokens) == 1 {
                // Only one token in command
                fname := tokens[0]
                if _, err := os.Stat(fname); err == nil {
                    // Token is file
                    ftokens := strings.Split(fname, ".")
                    fsuffix := ftokens[len(ftokens)-1]
                    if fsuffix == "csv" {
                        config.Mode = "csv"
                        break
                    }
                }
            }
            config.Mode = "regex"
            break
        }
    } else if mode != "" {
        config.Mode = strings.ToLower(mode)
    }
}

func (config *Config) SetRequiredArgument(requiredToken string, requiredArgs []string) {
    hasRequired := false
    for _, a := range config.CmdArgs {
        if a == requiredToken {
            hasRequired = true
        }
    }
    if !hasRequired {
        config.CmdArgs = append(config.CmdArgs, requiredArgs...)
    }
}

func (config *Config) SetWatchSec(watchSec float64) {
    if watchSec > 0 {
        // Minimum watch cycle time is 0.2 seconds
        config.WatchSec = math.Max(watchSec, 0.2)
    }
}

func (config *Config) SetRequiredWatchSec(watchSec float64) {
    if config.WatchSec <= 0 {
        config.SetWatchSec(watchSec)
    }
}

func LoadConfig() *Config {
    config := &Config{Help: false, Key: os.Getenv("SYNCOPATE_KEY"), Group: os.Getenv("SYNCOPATE_GROUP")}

    configFile  := flag.String("c", "syncopate.yaml", "Syncopate YAML config")
    runCmd      := flag.String("r", "", "Command to run")
    watchSec    := flag.Float64("w", -1.0, "Watch cycle time (in seconds)")
    key         := flag.String("k", "", "API key")
    group       := flag.String("g", "", "Group name")
    mode        := flag.String("m", "", "Mode: (regex, csv, ...)")
    help        := flag.Bool("help", false, "Show mode usage")
    debug       := flag.Bool("debug", false, "Debug output")
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
    config.SetWatchSec(*watchSec)
    config.SetCommandMode(*runCmd, *mode)

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

