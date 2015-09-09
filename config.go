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
    Option_t map[string]string
    Config struct {
        Key         string
        Group       string
        Options     []Option_t
        CmdWatchSec float64
        CmdBin      string
        CmdArgs     []string
        CmdFile     string
        Mode        string
        Help        bool
        Debug       bool
    }
)

func (config *Config) InitCommand(
    cmd string,
    mode string,
    watchSec float64) {
    config.SetWatchSec(watchSec)
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
            case "linux":
                config.SetRequiredArgument("-b", []string{"-b"})
            default:
            }
        case "df":
            config.Mode = "df"
            config.SetRequiredWatchSec(2.0)
        case "du":
            config.Mode = "du"
            config.SetRequiredWatchSec(2.0)
        default:
            if len(tokens) == 1 {
                // Only one token in command
                fname := tokens[0]
                // TODO: Make sure fname is a file and NOT a binary
                if _, err := os.Stat(fname); err == nil {
                    // Token is file, tail file
                    config.CmdBin = "tail"
                    config.CmdArgs = []string{"-f",fname}
                    config.CmdWatchSec = -1
                    config.CmdFile = fname

                    // Check file format
                    ftokens := strings.Split(fname, ".")
                    fsuffix := ftokens[len(ftokens)-1]
                    if fsuffix == "csv" {
                        config.Mode = "csv"
                        break
                    }
                }
            }
            config.Mode = "match"
        }
    }
    // Command-line argument mode takes final precedence
    if mode != "" {
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
        config.CmdWatchSec = math.Max(watchSec, 0.2)
    }
}

func (config *Config) SetRequiredWatchSec(watchSec float64) {
    if config.CmdWatchSec <= 0 {
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
    }
    err = yaml.Unmarshal(source, config)
    if err != nil {
        log.Fatal(err)
    }

    // Command to run will set default mode, which can be overridden
    // as command line argument
    config.InitCommand(*runCmd, *mode, *watchSec)

    // Try embedded static configs
    if configNotFound {
        switch config.Mode {
        case "top":
            src, _ := Asset("configs/top.yaml")
            err = yaml.Unmarshal(src, config)
            if err != nil {
                log.Fatal(err)
            }
        case "df":
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

