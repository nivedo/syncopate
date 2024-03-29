package main

import (
    "log"
    "fmt"
    "regexp"
    "strings"
    "strconv"
)

type (
    ParseState struct {
        InTable         bool
        TableRowCount   int
        LineCount       int
        Headers         []string
    }
    TopHandler struct {
        Info      *HandlerInfo
        Vars      KVList
        State     ParseState
        Matches   map[string]int
    }
)

func NewTopHandler(info *HandlerInfo) *TopHandler {
    h := &TopHandler{Info: info}
    h.State = ParseState{
        InTable: false,
        TableRowCount: 0,
        LineCount: 0}
    h.Load()
    return h
}

func (h *TopHandler) Load() {
    defaults := []string{
        "processes_total",
        "processes_running",
        "processes_stuck",
        "processes_sleeping",
        "processes_threads",
        "cpu_usage_user",
        "cpu_usage_sys",
        "cpu_usage_idle",
        "physmem_used",
        "physmem_unused",
        "physmem_wired"}
    h.Matches = make(map[string]int)
    cFields := h.Info.Config.Options
    numVar := 0

    if len(cFields) == 0 {
        for _, s := range defaults {
            h.Matches[s] = numVar
            numVar++
        }
    } else {
        for _, v := range cFields {
            h.Matches[v["desc"]] = numVar
            numVar++
        }
    }

    for k,_ := range h.Matches {
        log.Printf("[TRACKING] Field: %s\n", k)
    }

    h.Vars = make(KVList, numVar)
}

func (h *TopHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
    }
}

func (h *TopHandler) Help() {
    // List all variables
    keys := make([]string,0)
    for k := range h.Matches {
        keys = append(keys, k)
    }

    fmt.Println("---- Top Help ----\n")
    fmt.Printf("%20s: %s", "variables", strings.Join(keys, ", "))
    /*
    keys := make([]string, 0, len(h.Map))
    for k := range h.Map {
        keys = append(keys, k)
    }
    fmt.Println("top handler help --")
    fmt.Printf("%20s: %s", "variables", strings.Join(keys, ", "))
    */
}

func TrimSuffix(s, suffix string) string {
    if strings.HasSuffix(s, suffix) {
        s = s[:len(s)-len(suffix)]
    }
    return s
}

func ConvertToUnit(value string) string{
    // Get rid of %
    newValue := TrimSuffix(value, "%")
   
    // Convert unit
    reg, _ := regexp.Compile("^[0-9]+([,.][0-9]+)?(B|K|M|G|T)$")
    if reg.Match([]byte(newValue)) {
        num, err := strconv.ParseFloat(newValue[:len(newValue)-1], 64)
        if err == nil {
            var factor float64 = 1
            switch newValue[len(newValue)-1:] {
            case "B":
                break
            case "K":
                factor = 1000
                break
            case "M":
                factor = 1000000
                break
            case "G":
                factor = 1000000000
                break
            case "T":
                factor = 1000000000000
                break
            default:
                break
            }
            newValue = strconv.Itoa(int(num * factor))
        }
    }
    return newValue
}

func (h *TopHandler) Reset() {
    h.State.InTable = false
    h.State.LineCount = 0
    h.State.TableRowCount = 0

    // NOTE: Cannot assign a new map to watchEventMap because that doesn't
    //       change the reference in other function calls
    // h.Vars = h.Vars[:0]
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
        // fmt.Println(strings.Join(h.State.Headers, ", "), len(h.State.Headers))
    }
}

func (h *TopHandler) ParseTable(line string) {
    if h.State.InTable {
        tokens := strings.Fields(line)
        if len(tokens) == len(h.State.Headers) {
            for i, v := range tokens {
                k := fmt.Sprintf("table_r%d_c%d", h.State.TableRowCount, i)
                // fmt.Printf("%20s: %s\n", k, v)
                h.AddEvent(k, v)
            }
            h.State.TableRowCount++
        } else {
            fmt.Printf("ERROR: Number of columns mismatch, #headers: %d, #columns: %d\n",
                len(h.State.Headers), len(tokens))
            fmt.Println(line)
        }
    }
}

func (h *TopHandler) AddEvent(key string, value string) bool {
    valid := len(key) > 0 && len(value) > 0
    if valid {
        if index, ok := h.Matches[key]; ok {
            h.Vars[index] = KVPair{K: key, V: ConvertToUnit(value)}
        }
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

func (h *TopHandler) Parse(data string) {
    if strings.Contains(data, "Processes") {
        UploadKV(h.Vars, h.Info)
        h.Reset()
    }

    h.State.LineCount++
    h.ParseTopHeaders(data)
    h.ParseTableHeaders(data)
}

