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
		SeriesIndex int
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
		SeriesIDs []string
		Series    []Series
	}
)

func startCluster(config *Config, events chan SyncEvent) *Cluster {
	cluster := &Cluster{
		Key:   config.Key,
		ID:    hashID(config.Key),
		Group: config.Group,
		Token: hashKey(config.Key)}
	cluster.Series = make([]Series, len(config.Variables))
	cluster.SeriesIDs = make([]string, len(config.Variables))

	log.Printf(">>> STARTING CLUSTER (API: %s) (GROUP: %s) (TOKEN: %s)\n\n", cluster.Key, cluster.Group, cluster.Token)

	for i, wv := range config.Variables {
		seriesID := makeSeriesID(cluster.Token, cluster.Group, wv.Description)
		cluster.SeriesIDs[i] = seriesID
		cluster.Series[i] = Series{ID: seriesID}
		cluster.Series[i].Events = []Event{}
	}

    data := make(chan string)
	go read(config, data)

    handler := newHandler(cluster, config, data, events)
    go handler.run()

	return cluster
}

func read(cfg *Config, data chan string) {
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
