package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type (
	Variable struct {
		Pattern     string
		Description string
		Min         int
		Max         int
	}
	SyncEvent struct {
		Time        int64
		SeriesID    string
		Key         string
		Value       string
	}
	Event struct {
		Time int64
		Data map[string]interface{}
	}
	Series struct {
		ID     string
		Events []Event
	}
	Cluster struct {
		Key       string
		ID        string
		Group     string
		Token     string
		Series    map[string]*Series
	}
)

func StartCluster(config *Config, events chan SyncEvent) *Cluster {
	cluster := &Cluster{
		Key:   config.Key,
		ID:    HashID(config.Key),
		Group: config.Group,
		Token: HashKey(config.Key),
        Series: make(map[string]*Series),
    }

	log.Printf(">>> STARTING CLUSTER (API: %s) (GROUP: %s) (TOKEN: %s)\n\n", cluster.Key, cluster.Group, cluster.Token)

    data := make(chan string)
	go Read(config, data)

    handlerInfo := &HandlerInfo{Cluster: cluster, Config: config, Data: data, Events: events}
    handler := GetHandler(handlerInfo)
    go handler.Run()

	return cluster
}

func Read(cfg *Config, data chan string) {
    for _,v := range cfg.Variables {
	   log.Printf("[TRACKING] Variables: %s\n", v.Description)
    }

	r := bufio.NewReader(os.Stdin)

	for {
		line,err := r.ReadString('\n')
		data <- line
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
}
