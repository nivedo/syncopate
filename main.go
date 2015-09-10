package main

func main() {
    config := LoadConfig()

    events := make(chan SyncEvent, 1)
    StartCluster(config, events)

    d := NewDispatcher(config.Key, events)
    d.Run()
}
