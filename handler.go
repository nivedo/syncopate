package main

import (
    "log"
)

type (
    Handler interface {
        parse(data string)
        run()
        help()
    }
    HandlerInfo struct {
        Cluster     *Cluster
        Config      *Config
        Data        chan string
        Events      chan SyncEvent
    }
)

func getHandler(info *HandlerInfo) Handler {
    switch info.Config.Mode {
    case "regex":
        return NewRegexHandler(info)
    case "top":
        return NewTopHandler(info)
    default:
        log.Fatal("No mode available.")
    }

    return nil
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
