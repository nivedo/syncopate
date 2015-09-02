package main

import (
    "regexp"
    "time"
    "log"
)

type (
    Handler struct {
        Cluster     *Cluster
        Config      *Config
        Data        chan string
        Events      chan SyncEvent
    }
)

func newHandler(
    c *Cluster,
    cfg *Config,
    data chan string,
    events chan SyncEvent) *Handler {
    return &Handler{
        Cluster: c,
        Config: cfg,
        Data: data,
        Events: events}
}

func (h *Handler) run() {
    var parse func(data string)

    switch h.Config.Mode {
    case "regex":
        parse = h.regex
    default:
        // TODO: Detect mode if nothing specified.
        log.Fatal("No mode available.")
    }

    for {
        data := <-h.Data
        parse(data)
    }
}

// TODO: move this handler to another package.
func (h *Handler) regex(data string) {
    vars := h.Config.Variables
    for i, _ := range vars {
        match, _ := regexp.MatchString(vars[i].Pattern, data)
        if match {
            r, _ := regexp.Compile(vars[i].Pattern)
            seriesID := makeSeriesID(h.Cluster.Token, h.Cluster.Group, vars[i].Description)
            allMatch := r.FindAllStringSubmatch(data, -1)
            for _, matchVal := range allMatch {
                h.Events <- SyncEvent{
                    SeriesID:    seriesID,
                    SeriesIndex: i,
                    Key:         vars[i].Description,
                    Value:       matchVal[1],
                    Time:        time.Now().UTC().UnixNano() / int64(time.Microsecond)}
            }
        }
    }
}
