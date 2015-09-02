package main

import (
    "fmt"
    "regexp"
    "strings"
    "strconv"
)

type (
    WatchEventMap map[string]string
    ParseState struct {
        InTable   bool
        LineCount int
        Headers   []string
    }
    TopHandler struct {
        Info    *HandlerInfo
        Map     WatchEventMap
        State   ParseState
    }
)

func NewTopHandler(info *HandlerInfo) *TopHandler {
    h := &TopHandler{Info: info}
    h.State = ParseState{InTable: false, LineCount: 0}
    h.Map = make(WatchEventMap)
    return h
}

func (h *TopHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
        h.Print()
    }
}

func (h *TopHandler) Parse(data string) {
    h.ParseTopMacOSX(data)
}

func (h *TopHandler) Help() {
    // TODO: list all variables
}

func ConvertToValidSeriesKey(rawId string) string {
    // Convert #, %
    newId := strings.Replace(rawId, "#", "num_", -1)
    newId = strings.Replace(newId, "%", "pct_", -1)

    // Conver space to _
    newId = strings.Replace(newId, " ", "_", -1)
    newId = strings.ToLower(newId)

    return newId
}

func (h *TopHandler) Upload() {
}

func (h *TopHandler) ParseTable(text string) {
}

func (h *TopHandler) Reset() {
    h.State.InTable = false
    h.State.LineCount = 0

    // NOTE: Cannot assign a new map to watchEventMap because that doesn't
    //       change the reference in other function calls
    for k := range h.Map {
        delete(h.Map, k)
    }
}

func (h *TopHandler) Print() {
    for k,v := range h.Map {
        fmt.Printf("%20s: %s\n",k,v)
    }
}

func (h *TopHandler) InitTableHeaders(line string) {
    tokens := strings.Fields(line)
    h.State.Headers = make([]string, len(tokens))

    for i, t := range tokens {
        h.State.Headers[i] = ConvertToValidSeriesKey(t)
    }
}

func (h *TopHandler) ParseTableHeaders(line string) {
    //lineTrim := strings.Replace(line, "N/A", "", -1)
    lineTrim := strings.Trim(line, " ")
    hasAlpha, _ := regexp.MatchString("[A-Z]+", lineTrim)
    if len(lineTrim) > 0 && line == strings.ToUpper(lineTrim) && hasAlpha {
        h.State.InTable = true
        h.InitTableHeaders(line)
        fmt.Println(strings.Join(h.State.Headers, ", "), len(h.State.Headers))
    }
}

func (h *TopHandler) AddEvent(key string, value string) bool {
    valid := len(key) > 0 && len(value) > 0
    if valid {
        h.Map[key] = value
    }
    return valid
}

func (h *TopHandler) ParseTopHeaders(line string) bool {
    valid := true
    if !h.State.InTable {
        verbose := false
        line = strings.TrimSpace(line)
        line = strings.TrimRight(line, ".")
        tokens := strings.Split(line, ":")

        numTokens := len(tokens)
        if numTokens > 0 {
            hasAlpha, _ := regexp.MatchString("[a-zA-Z]+", tokens[0])
            if hasAlpha && numTokens == 2 {
                seriesIdPrefix := tokens[0]

                // Get rid of (number)'s, confuses the split on ,()
                reg, _ := regexp.Compile("\\([0-9]+\\)")
                str := reg.ReplaceAllString(tokens[1], "")

                // Split on ,()
                statPairs := strings.FieldsFunc(str, func(r rune) bool {
                    return r == ',' || r == '(' || r == ')'
                })
                if verbose {
                    fmt.Println(statPairs)
                }
                // StatPairs separated by commas
                for i, p := range statPairs {
                    pair := strings.TrimLeft(p, " ")
                    pair = strings.TrimRight(pair, " ")
                    // Check empty string
                    if len(pair) > 0 {
                        ptokens := strings.Split(pair, " ")
                        if len(ptokens) >= 2 {
                            v := ptokens[0]
                            k := strings.Join(ptokens[1:], "_")
                            seriesKey := ConvertToValidSeriesKey(seriesIdPrefix + "_" + k)
                            if verbose {
                                fmt.Print(seriesKey, "=", v, ",")
                            }
                            valid = h.AddEvent(seriesKey, v)
                        } else if len(ptokens) == 1 {
                            // Load avg do not have extra fields
                            seriesKey := ConvertToValidSeriesKey(seriesIdPrefix + "_" + strconv.Itoa(i+1))
                            valid = h.AddEvent(seriesKey, ptokens[0])
                        }
                    }
                }
                if verbose {
                    fmt.Println("")
                }
            }
        }
    }
    return valid
}

func (h *TopHandler) ParseTopMacOSX(data string) {
    if strings.Contains(data, "Processes") {
        h.Upload()
        h.Reset()
        fmt.Println(data)
    }

    h.State.LineCount++
    h.ParseTopHeaders(data)
    h.ParseTableHeaders(data)

    if !h.State.InTable {
        fmt.Println(data)
    }

    verbose := false
    if verbose {
        fmt.Println(">>>>>>>>>>>>>>>")
        fmt.Println("parse")
        fmt.Println(data)
        fmt.Println()
        h.ParseTable(data)
    }
}

