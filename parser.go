package main

import (
    "log"
)

type (
    Parser interface {
        Parse(data string)
        Run()
    }
    ParserInfo struct {
        Config      *Config
        Uploader    *Uploader
        Data        chan string
    }
    AnyParser struct {
        Info         *ParserInfo
        Filters      []Filter
    }
    OrderedParser struct {
        Info         *ParserInfo
        Filters      []Filter
        FilterIndex  int
    }
)

func GetParser(info *ParserInfo) Parser {
    switch info.Config.Mode {
    case "any":
        return NewAnyParser(info)
    case "ordered":
        return NewOrderedParser(info)
    default:
        log.Fatalf("Unknown parser for mode %s.", info.Config.Mode)
    }

    return nil
}

///////////////////////////////////////////////////////////////////
// AnyParser
// ----------
// Distributes each line of data to each filter, uploads
// on any match.
///////////////////////////////////////////////////////////////////

func NewAnyParser(info *ParserInfo) *AnyParser {
    parser := &AnyParser{Info: info}
    for _,opt := range parser.Info.Config.Options {
        parser.Filters = append(parser.Filters, GetFilter(opt))
    }
    log.Printf("[AnyParser] Initialized with %d filters.", len(parser.Filters))
    return parser
}

func (p *AnyParser) Parse(data string) {
    uploader := p.Info.Uploader
    for _,f := range p.Filters {
        if f.Match(data) {
            uploader.UploadKV(f.GetVars())
        }
    }
}

func (p *AnyParser) Run() {
    for {
        data := <-p.Info.Data
        p.Parse(data)
    }
}

///////////////////////////////////////////////////////////////////
// OrderedParser
// ----------
// Filters must succeed in the order they were listed, and 
// results are then uploaded in a batch.
///////////////////////////////////////////////////////////////////

func NewOrderedParser(info *ParserInfo) *OrderedParser {
    parser := &OrderedParser{Info: info}
    for _,opt := range parser.Info.Config.Options {
        parser.Filters = append(parser.Filters, GetFilter(opt))
    }
    log.Printf("[OrderedParser] Initialized with %d filters.", len(parser.Filters))
    return parser
}

func (p *OrderedParser) Parse(data string) {
    match := p.Filters[p.FilterIndex].Match(data) 
    for match {
        p.FilterIndex++
        p.CheckUpload()
        match = p.Filters[p.FilterIndex].Match(data)
    }
}

func (p *OrderedParser) CheckUpload() {
    if p.FilterIndex == len(p.Filters) {
        uploader := p.Info.Uploader
        for _,f := range p.Filters {
            uploader.UploadKV(f.GetVars())
        }
        p.FilterIndex = 0
    }
}

func (p *OrderedParser) Run() {
    for {
        data := <-p.Info.Data
        p.Parse(data)
    }
}