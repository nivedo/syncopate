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
        case sEvent := <-events:
            changed = true
            newEvent := Event{Time: sEvent.Time}
            newEvent.Data = make(map[string]interface{})
            newEvent.Data[sEvent.Key] = sEvent.Value
            if _, ok := cluster.Series[sEvent.SeriesID]; !ok {
                cluster.Series[sEvent.SeriesID] = &Series{ID: sEvent.SeriesID, Events: []Event{}}
            }
            cluster.Series[sEvent.SeriesID].Events = append(cluster.Series[sEvent.SeriesID].Events, newEvent)
        case <-up:
            if changed {
                Upload(cluster, config)
                changed = false
            }
        }
    }
}
