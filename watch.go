package main

import (
	"bufio"
	"fmt"
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
		ID          string
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
		numVar := len(wf.Variables)
		patternList := make([]string, numVar)
		descList := make([]string, numVar)
		indexList := make([]int, numVar)
		for i, wv := range wf.Variables {
			seriesStr := fmt.Sprintf("%s.%s.%s", cluster.Token, cluster.Group, wv.Description)
			seriesID := hashID(seriesStr)
			cluster.SeriesIDs[seriesIndex] = seriesID
			cluster.Series[seriesIndex] = Series{ID: seriesID}
			cluster.Series[seriesIndex].Events = []Event{}
			patternList[i] = wv.Pattern
			descList[i] = wv.Description
			indexList[i] = seriesIndex
			seriesIndex += 1
		}
		go watch(
			wf.Filename,
			cluster.Group,
			cluster.Token,
			patternList,
			descList,
			indexList,
			events,
			time.Now().UTC().UnixNano()/int64(time.Microsecond))
	}

	return cluster
}

// TODO: refactor function argument into struct
func watch(
	filename string,
	group string,
	token string,
	patterns []string,
	descriptions []string,
	seriesIndices []int,
	events chan WatchEvent,
	watchStart int64) {
	log.Printf("[TRACKING] %s -- Variables: %s\n", filename, descriptions)
	lc := 0
	if filename == "PIPE" {
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
			line := string(buf)
			curTime := time.Now().UTC().UnixNano() / int64(time.Microsecond)
			if curTime-watchStart > 1e6 {
				for i, _ := range patterns {
					match, _ := regexp.MatchString(patterns[i], line)
					if match {
						r, _ := regexp.Compile(patterns[i])
						seriesStr := fmt.Sprintf("%s.%s.%s", token, group, descriptions[i])
						seriesID := hashID(seriesStr)
						allMatch := r.FindAllStringSubmatch(line, -1)
						for j, matchVal := range allMatch {
							eventStr := fmt.Sprintf("%s-%s-%d-%d-%d-%s", filename, patterns[i], lc, j, curTime, matchVal[0])
							eventID := hashID(eventStr)
							events <- WatchEvent{
								ID:          eventID,
								SeriesID:    seriesID,
								SeriesIndex: seriesIndices[i],
								Key:         descriptions[i],
								Value:       matchVal[1],
								Time:        curTime}
						}
					}
				}
			}
			lc = lc + 1
			// process buf
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
		}
	} else {
		t, err := tail.TailFile(filename, tail.Config{Follow: true})
		for line := range t.Lines {
			curTime := time.Now().UTC().UnixNano() / int64(time.Microsecond)
			if curTime-watchStart > 1e6 {
				for i, _ := range patterns {
					match, _ := regexp.MatchString(patterns[i], line.Text)
					if match {
						r, _ := regexp.Compile(patterns[i])
						seriesStr := fmt.Sprintf("%s.%s.%s", token, group, descriptions[i])
						seriesID := hashID(seriesStr)
						allMatch := r.FindAllStringSubmatch(line.Text, -1)
						for j, matchVal := range allMatch {
							eventStr := fmt.Sprintf("%s-%s-%d-%d-%d-%s", filename, patterns[i], lc, j, curTime, matchVal[0])
							eventID := hashID(eventStr)
							events <- WatchEvent{
								ID:          eventID,
								SeriesID:    seriesID,
								SeriesIndex: seriesIndices[i],
								Key:         descriptions[i],
								Value:       matchVal[1],
								Time:        curTime}
						}
					}
				}
			}
			lc = lc + 1
		}
		if err != nil {
			panic(err)
		}
	}
}
