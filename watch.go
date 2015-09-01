package main

import (
	"bufio"
	"github.com/ActiveState/tail"
	"io"
	"log"
	"os"
	"regexp"
	"time"
)

type (
	WatchVar struct {
		Pattern     string
		Description string
		Min         int
		Max         int
	}
	WatchFile struct {
		Filename  string
		Variables []WatchVar
	}
	WatchEvent struct {
		Time        int64
		SeriesID    string
		SeriesIndex int
		Key         string
		Value       string
	}
	WatchConfig struct {
		Filename     string
		Patterns     []string
		Descriptions []string
		Index        []int
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

func startCluster(config Config, events chan WatchEvent) Cluster {
	cluster := Cluster{
		Key:   config.Key,
		ID:    hashID(config.Key),
		Group: config.Group,
		Token: hashKey(config.Key)}
	cluster.Series = make([]Series, config.NumSeries)
	cluster.SeriesIDs = make([]string, config.NumSeries)

	log.Printf(">>> STARTING CLUSTER (API: %s) (GROUP: %s) (TOKEN: %s)\n\n", cluster.Key, cluster.Group, cluster.Token)

	seriesIndex := 0

	for _, wf := range config.Data {
		watchCfg := WatchConfig{Filename: wf.Filename}
		for _, wv := range wf.Variables {
			seriesID := makeSeriesID(cluster.Token, cluster.Group, wv.Description)
			cluster.SeriesIDs[seriesIndex] = seriesID
			cluster.Series[seriesIndex] = Series{ID: seriesID}
			cluster.Series[seriesIndex].Events = []Event{}
			watchCfg.Patterns = append(watchCfg.Patterns, wv.Pattern)
			watchCfg.Descriptions = append(watchCfg.Descriptions, wv.Description)
			watchCfg.Index = append(watchCfg.Index, seriesIndex)
			seriesIndex += 1
		}
		go watch(watchCfg, &cluster, events)
	}

	return cluster
}

// Use this function for each mode
func watchHelper(content string, wc WatchConfig, cluster *Cluster, events chan WatchEvent) {
	for i, _ := range wc.Patterns {
		match, _ := regexp.MatchString(wc.Patterns[i], content)
		if match {
			r, _ := regexp.Compile(wc.Patterns[i])
			seriesID := makeSeriesID(cluster.Token, cluster.Group, wc.Descriptions[i])
			allMatch := r.FindAllStringSubmatch(content, -1)
			for _, matchVal := range allMatch {
				events <- WatchEvent{
					SeriesID:    seriesID,
					SeriesIndex: wc.Index[i],
					Key:         wc.Descriptions[i],
					Value:       matchVal[1],
					Time:        time.Now().UTC().UnixNano() / int64(time.Microsecond)}
			}
		}
	}
}

func watch(wc WatchConfig, cluster *Cluster, events chan WatchEvent) {

	watchStart := time.Now().UTC().UnixNano()/int64(time.Microsecond)
	log.Printf("[TRACKING] %s -- Variables: %s\n", wc.Filename, wc.Descriptions)

	if wc.Filename == "_pipe" {
		r := bufio.NewReader(os.Stdin)
		buf := make([]byte, 0, 4*1024)
		for {
			n, err := r.Read(buf[:cap(buf)])
			buf = buf[:n]
			if n == 0 {
				if err == nil {
					continue
				}
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			content := string(buf)
			watchHelper(content, wc, cluster, events)
			// process buf
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
		}
	} else {
		t, err := tail.TailFile(wc.Filename, tail.Config{Follow: true})
		for line := range t.Lines {
			curTime := time.Now().UTC().UnixNano() / int64(time.Microsecond)
			if curTime-watchStart > 1e6 {
				watchHelper(line.Text, wc, cluster, events)
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
