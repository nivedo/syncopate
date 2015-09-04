package main

import (
    "log"
    "fmt"
    "regexp"
    "strings"
)

/*
 *  Two types of RuleHandlers:
 *  ==========================
 *
 *  1) GROUP = false -- for each line, evaluate all given rules and upload on ANY match.
 *  2) GROUP = true  -- for each line, evaluate only rule N until rule N passes, then iterate N++.
 *     Do not upload anything until all rules pass.
 */

type (
    RuleHandler struct {
        Info  *HandlerInfo
        Vars  KVList
        Rules []Rule
        Start int
        N     int
        Batch bool
    }
    Rule interface {
        // Eval inserts elements into vars from vars[start:], returning num elements inserted
        Eval(line string, vars *KVList, start int) int
    }
    RuleRegex struct {
        Pattern string
        Labels  []string
    }
)

func NewRuleHandler(info *HandlerInfo, batch bool) *RuleHandler {
    h := &RuleHandler{Info: info, Start: 0, N: 0, Batch: batch}
    h.Load()
    return h
}

func (h *RuleHandler) Load() {
    for _,v := range h.Info.Config.Options {
        if rule, ok := v["match"]; ok {
            h.AddRule(NewRule(rule))
            log.Printf("[TRACKING] Rule: %s\n", rule)
        }
    }
    h.Vars = make(KVList, 50)
}

func (h *RuleHandler) Help() {

}

func (h *RuleHandler) Run() {
    for {
        data := <-h.Info.Data
        h.Parse(data)
    }
}

func (h *RuleHandler) Parse(data string) {
    if !h.Batch {
        h.ParseSingle(data)
    } else {
        h.ParseBatch(data)
    }
}

func (h *RuleHandler) AddRule(r Rule) {
    h.Rules = append(h.Rules, r)
}

func (h *RuleHandler) ParseSingle(line string) {
    for _,rule := range h.Rules {
        n := rule.Eval(line, &h.Vars, 0)
        if(n > 0) {
            UploadKV(h.Vars[:(n-1)], h.Info)
        }
    }
}

func (h *RuleHandler) ParseBatch(line string) {
    n := h.Rules[h.N].Eval(line, &h.Vars, h.Start)
    for n > 0 {
        // Rule passes, advance to next rule
        h.N++
        h.Start = h.Start + n
        if h.N == len(h.Rules) {
            // All rules pass, upload KVList
            UploadKV(h.Vars[:(h.Start-1)], h.Info)
            h.N = 0
            h.Start = 0
        }
        n = h.Rules[h.N].Eval(line, &h.Vars, h.Start)
    }
}

func NewRule(pattern string) Rule {
    return NewRuleRegex(pattern)
}

/* Rule Regex 
 * ==========
 * Example -- CPU usage: {{ cpu_usage_user:%p }} user, {{ cpu_usage_sys:%p }} sys
 */
func NewRuleRegex(pattern string) *RuleRegex {
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
            subs[subtoken] = "(\\d*\\.?\\d*)%"
        case "%f":
            subs[subtoken] = "(\\d*\\.?\\d*)"
        case "%d":
            subs[subtoken] = "(\\d+)"
        case "%mem":
            subs[subtoken] = "(\\d+[BKMG]?)"
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

    return &RuleRegex{Pattern: result, Labels: labels}
}

func (r *RuleRegex) Eval(line string, vars *KVList, start int) int {
    index := start
    match, _ := regexp.MatchString(r.Pattern, line)
    if match {
        reg, _ := regexp.Compile(r.Pattern)
        allMatch := reg.FindAllStringSubmatch(line, -1)
        for i,v := range allMatch[0][1:] {
            (*vars)[index] = KVPair{K: r.Labels[i], V: v}
            log.Println((*vars)[index])
            index++
        }
        return index - start
    }
    return 0
}