package main

import (
    "log"
)

func main() {
    log.Println("Syncopate is initializing...")

    config := loadConfig()
    events := make(chan WatchEvent, 1)

    cluster := startCluster(config, events)
    go uploadHelper(events)

    changed := false
    for {
        wEvent := <-events
        if wEvent.SeriesIndex >= 0 {
            changed = true
            newEvent := Event{Time: wEvent.Time}
            newEvent.Data = make(map[string]interface{})
            newEvent.Data[wEvent.Key] = wEvent.Value
            cluster.Series[wEvent.SeriesIndex].Events = append(cluster.Series[wEvent.SeriesIndex].Events, newEvent)
        }
        if changed && wEvent.SeriesIndex < 0 {
            upload(&cluster, cluster.ID)
            changed = false
        }
    }
}
