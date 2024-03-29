package main

import (
    "bufio"
    "fmt"
    //"io"
    //"log"
    "os"
    "regexp"
    "strings"
    "strconv"
)

type WatchEventMap map[string]string

type ParseState struct {
    InTable   bool
    LineCount int
    Headers   []string
}

func UploadWatchEventMap(watchEventMap WatchEventMap) {
    fmt.Println(watchEventMap)

}

func ParseTable(text string) {
}

func ResetParseState(state *ParseState) {
    state.InTable = false
    state.LineCount = 0
}

func ConvertToValidSeriesKey(rawId string) string {
    // Convert #, %
    newId := strings.Replace(rawId, "#", "n", -1)
    newId = strings.Replace(newId, "%", "p", -1)

    // Conver space to _
    newId = strings.Replace(newId, " ", "_", -1)

    return newId
}

func InitTableHeaders(state *ParseState, line string) {
    tokens := strings.Fields(line)
    state.Headers = make([]string, len(tokens))

    for i, t := range tokens {
        state.Headers[i] = ConvertToValidSeriesKey(t)
    }
}

func ParseTableHeaders(state *ParseState, line string, lc int) {
    lineTrim := strings.Replace(line, "N/A", "", -1)
    lineTrim = strings.Trim(lineTrim, " ")
    hasAlpha, _ := regexp.MatchString("[A-Z]+", lineTrim)
    if len(lineTrim) > 0 && line == strings.ToUpper(lineTrim) && hasAlpha {
        state.InTable = true
        fmt.Println("in table at", state.LineCount+lc, "len=", len(line))

        InitTableHeaders(state, line)
        // fmt.Println(line)
        fmt.Println(strings.Join(state.Headers, ", "), len(state.Headers))
    }
}

func InsertWatchEventMap(watchEventMap WatchEventMap, key string, value string) bool {
    valid := len(key) > 0 && len(value) > 0
    if valid {
        watchEventMap[key] = value
    }
    return valid
}

func ParseTopHeaders(state *ParseState, line string, watchEventMap WatchEventMap) bool{
    valid := true
    if !state.InTable {
        verbose := false
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
                            valid = InsertWatchEventMap(watchEventMap, seriesKey, v)
                        } else if len(ptokens) == 1 {
                            // Load avg do not have extra fields
                            seriesKey := ConvertToValidSeriesKey(seriesIdPrefix + "_" + strconv.Itoa(i+1))
                            valid = InsertWatchEventMap(watchEventMap, seriesKey, ptokens[0])
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

func ParseTopMacOSX(state *ParseState, text string, watchEventMap WatchEventMap) {
    lines := strings.Split(text, "\n")

    if len(lines) > 0 {
        if strings.Contains(lines[0], "Processes") {
            UploadWatchEventMap(watchEventMap)
            ResetParseState(state)
            fmt.Println(lines[0])
            
            // NOTE: Cannot assign a new map to watchEventMap because that doesn't
            //       change the reference in other function calls
            for k := range watchEventMap {
                delete(watchEventMap, k)
            }
        }

        state.LineCount += len(lines)
        for i, line := range lines {
            ParseTopHeaders(state, line, watchEventMap)
            ParseTableHeaders(state, line, i)

            if !state.InTable {
                fmt.Println(line)
            }
        }
    }
    verbose := false
    if verbose {
        fmt.Println(">>>>>>>>>>>>>>>")
        fmt.Println("parse")
        fmt.Println(len(lines))
        fmt.Println(lines[0])
        fmt.Println()
        ParseTable(text)
    }
}

func ReadStdin() {
    //nBytes, nChunks := int64(0), int64(0)
    parseState := &ParseState{InTable: false, LineCount: 0}
    r := bufio.NewReader(os.Stdin)
    //buf := make([]byte, 0, 4*1024)
    watchEventMap := make(WatchEventMap)
    for {
        // n is the number of bytes read
        /*
        n, err := r.Read(buf[:cap(buf)])
        buf = buf[:n]
        if n == 0 {
            if err == nil {
                continue
            }
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }
        nChunks++
        nBytes += int64(len(buf))
        line := string(buf)
        */
        line,_ := r.ReadString('\n')
        // fmt.Println(line)
        ParseTopMacOSX(parseState, line, watchEventMap)
    }
}

func main() {
    fmt.Println("running top helper ...")
    ReadStdin()
}

