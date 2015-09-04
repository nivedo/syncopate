package main

import (
    "log"
    "fmt"
    "regexp"
    "strings"
)

/*
 *  Two types of MatchHandlers:
 *  ==========================
 *
 *  1) BATCH = false -- for each line, evaluate all given rules and upload on ANY match.
 *  2) BATCH = true  -- for each line, evaluate only rule N until rule N passes, then iterate N++.
 *     Do not upload anything until all rules pass.
 */

type (
    MatchHandler struct {
        Info  *HandlerInfo
        Vars  KVList
        Matches []Match
        Start int
        N     int
        Batch bool
    }
    Match interface {
        // Eval inserts elements into vars from vars[start:], returning num elements inserted
        Eval(line string, vars *KVList, start int) int
    }
    MatchRegex struct {
        Pattern string
        Labels  []string
    }
)

func NewMatchHandler(info *HandlerInfo, batch bool) *MatchHandler {
    h := &MatchHandler{Info: info, Start: 0, N: 0, Batch: batch}
    h.Load()
    return h
}

func (h *MatchHandler) Load() {
    for _,v := range h.Info.Config.Options {
        if rule, ok := v["match"]; ok {
            h.AddMatch(NewMatch(rule))
            log.Printf("[TRACKING] Match: %s\n", rule)
        }
    }
    h.Vars = make(KVList, 100)
}

func (h *MatchHandler) Help() {

}

func (h *MatchHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
    }
}

func (h *MatchHandler) Parse(data string) {
    if !h.Batch {
        h.ParseSingle(data)
    } else {
        h.ParseBatch(data)
    }
}

func (h *MatchHandler) AddMatch(r Match) {
    h.Matches = append(h.Matches, r)
}

func (h *MatchHandler) ParseSingle(line string) {
    for _,rule := range h.Matches {
        n := rule.Eval(line, &h.Vars, 0)
        if(n > 0) {
            UploadKV(h.Vars[:(n-1)], h.Info)
        }
    }
}

func (h *MatchHandler) ParseBatch(line string) {
    n := h.Matches[h.N].Eval(line, &h.Vars, h.Start)
    for n > 0 {
        // Match passes, advance to next rule
        h.N++
        h.Start = h.Start + n
        if h.N == len(h.Matches) {
            // All rules pass, upload KVList
            UploadKV(h.Vars[:(h.Start-1)], h.Info)
            h.N = 0
            h.Start = 0
        }
        n = h.Matches[h.N].Eval(line, &h.Vars, h.Start)
    }
}

func NewMatch(pattern string) Match {
    return NewMatchRegex(pattern)
}

/* Match Regex 
 * ==========
 * Example -- CPU usage: {{ cpu_usage_user:%p }} user, {{ cpu_usage_sys:%p }} sys
 */
func NewMatchRegex(pattern string) *MatchRegex {
    r, _ := regexp.Compile("\\{\\{\\s*(\\w+):(.+?)\\}\\}")
    tokens := r.FindAllStringSubmatch(pattern, -1)

    result := pattern
    var labels []string

    subs := make(map[string]string)

    for i,token := range tokens {
        labels = append(labels, strings.TrimSpace(token[1]))
        rule := strings.TrimSpace(token[2])
        subtoken := fmt.Sprintf("SYNCVAR_%d",i)
        switch rule {
        case "%p":
            subs[subtoken] = "((?:\\d+\\.?\\d*)|(?:\\.\\d+))%"
        case "%f":
            subs[subtoken] = "([+-]?(?:\\d+\\.?\\d*)|(?:\\.\\d+))"
        case "%d":
            subs[subtoken] = "(\\d+)"
        case "%w":
            subs[subtoken] = "(\\w+)"
        case "%mem":
            subs[subtoken] = "(\\d+[BKMG]?)[+-]?"
        default:
            // Use user specified regex
            subs[subtoken] = rule
        }
        result = strings.Replace(result, token[0], subtoken, 1)
    }

    // Escape Special Characters
    result = regexp.QuoteMeta(result)

    // Replace Subtokens
    for k,v := range subs {
        result = strings.Replace(result, k, v, 1)
    }

    // Whitespace is arbitrary
    r3, _ := regexp.Compile("\\s+")
    result = r3.ReplaceAllString(result, "\\s*")

    return &MatchRegex{Pattern: result, Labels: labels}
}

func (r *MatchRegex) Eval(line string, vars *KVList, start int) int {
    index := start
    match, _ := regexp.MatchString(r.Pattern, line)
    if match {
        reg, _ := regexp.Compile(r.Pattern)
        allMatch := reg.FindAllStringSubmatch(line, -1)
        for i,v := range allMatch[0][1:] {
            (*vars)[index] = KVPair{K: r.Labels[i], V: strings.TrimSpace(v)}
            log.Println((*vars)[index])
            index++
        }
        return index - start
    }
    return 0
}