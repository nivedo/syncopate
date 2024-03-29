package main

import (
    "bufio"
    "bytes"
    "io"
    "log"
    "os"
    "os/exec"
    "time"
    "strings"
    "runtime"
)

const (
    CMD_TOP = iota
    CMD_DF
    CMD_DU
    CMD_TAIL
    CMD_CUSTOM
    MIN_WATCH = 200 * time.Millisecond
    DATA_EOF = "!!SYNCOPATE_EOF"
)

type (
    CommandInfo struct {
        WatchSec time.Duration
        Bin      string
        Args     []string
        Type     int
        Mode     int
    }
    CommandHub struct {

    }
)

func NewCommandInfo(cmd string, watchSec time.Duration) *CommandInfo {
    info := &CommandInfo{}

    if cmd == "" {
        return nil    
    }

    tokens := strings.Fields(cmd)
    info.Bin = tokens[0]
    if len(tokens) > 1 {
        info.Args = tokens[1:len(tokens)]
    }

    switch info.Bin {
    case "top":
        info.Type = CMD_TOP
        // Make sure batch mode, if not, add batch mode argument
        switch runtime.GOOS {
        case "darwin":
            info.RequireArgs("-l", []string{"-l","0"})
        case "linux":
            info.RequireArgs("-b", []string{"-b"})
        default:
        }
    case "df":
        info.Type = CMD_DF
        info.SetWatchSec(watchSec, 2.0 * time.Second)
    case "du":
        info.Type = CMD_DU
        info.SetWatchSec(watchSec, 2.0 * time.Second)
    case "tail":
        info.Type = CMD_TAIL
        info.RequireArgs("-f", []string{"-f"})
    default:
        info.Type = CMD_CUSTOM
        info.SetWatchSec(watchSec, -1.0)
    }

    return info
}

func (info *CommandInfo) RequireArgs(requiredToken string, requiredArgs []string) {
    hasRequired := false
    for _, a := range info.Args {
        if a == requiredToken {
            hasRequired = true
        }
    }
    if !hasRequired {
        info.Args = append(requiredArgs, info.Args...)
    }
}

func (info *CommandInfo) SetWatchSec(t time.Duration, def time.Duration) {
    if t > 0 {
        // Minimum watch cycle time is 0.2 seconds
        if t < MIN_WATCH {
            t = MIN_WATCH
        }
        info.WatchSec = t
    } else {
        info.WatchSec = def
    }
}

func StartCommands(info *CommandInfo, data chan string) {
    if info != nil && len(info.Bin) > 0 {
        if info.WatchSec > 0 {
            // Run on watch timer
            go WatchCommand(info, data)
        } else {
            // Run once
            go RunCommand(info, data)
        }
    } else {
        // Pipe stdin to reader
        go SendData(os.Stdin, data)
    }
}

func WatchCommand(info *CommandInfo, data chan string) {
    for {
        cmd := exec.Command(info.Bin, info.Args...)
        if output, err := cmd.Output(); err == nil {
            buf := bytes.NewBuffer(output)
            SendData(buf, data)
        } else {
            log.Fatal(err)
        }
        
        time.Sleep(info.WatchSec)
    }
}

func RunCommand(info *CommandInfo, data chan string) {
    // Run command and pipe stdout to reader
    cmd := exec.Command(info.Bin, info.Args...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    SendData(stdout, data)
    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }
}

func SendData(reader io.Reader, data chan string) {
    s := bufio.NewScanner(reader)

    for s.Scan() {
        data <- s.Text()
    }

    data <- DATA_EOF
}
