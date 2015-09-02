package main

func main() {
    config := LoadConfig()

    events := make(chan SyncEvent, 1)
    cluster := StartCluster(config, events)

    up := make(chan bool, 1)
    go UploadHelper(up)

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
                Upload(cluster, cluster.ID)
                changed = false
            }
        }
    }
}
