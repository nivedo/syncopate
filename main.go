package main

func main() {
    config := loadConfig()

    events := make(chan SyncEvent, 1)
    cluster := startCluster(config, events)

    up := make(chan bool, 1)
    go uploadHelper(up)

    changed := false
    for {
        select {
        case wEvent := <-events:
            changed = true
            newEvent := Event{Time: wEvent.Time}
            newEvent.Data = make(map[string]interface{})
            newEvent.Data[wEvent.Key] = wEvent.Value
            cluster.Series[wEvent.SeriesIndex].Events = append(cluster.Series[wEvent.SeriesIndex].Events, newEvent)
        case <-up:
            if changed {
                upload(cluster, cluster.ID)
                changed = false
            }
        }
    }
}
