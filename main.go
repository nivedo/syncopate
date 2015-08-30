package main

import (
    "github.com/ActiveState/tail"
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
    "os"
    "io"
    "bufio"
    "bytes"
    "time"
    "regexp"
    "net/http"
    "gopkg.in/yaml.v2"
)

type (
    WatchVar struct {
        Pattern     string
        Description string
        Min         int
        Max         int
    }
    WatchFile struct {
        Filename    string
        Variables   []WatchVar
    }
    WatchEvent struct {
        Time        int64
        ID          string
        SeriesID    string
        SeriesIndex int
        Key         string
        Value       string
    }
    Config struct {
        Key     string
        Data    []WatchFile
    }
    Pair struct {
        Key    string
        Value  string
    }
    Event struct {
        Time   int64
        Data   map[string]interface{}
    }
    Series struct {
        ID     string
        Events []Event
    }
    Cluster struct {
        Key       string
        ID        string
        SeriesIDs []string
        Series    []Series
    }
)

const (
    SERVER_URL = "http://api.blub.io:32794"
)

func main() {
    var config Config
    source, err := ioutil.ReadFile("syncopate.yaml")
    if err != nil {
        fmt.Println("Could not locate syncopate.yaml. Terminating...")
        return
    }
    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }

    fmt.Println("Syncopate is initializing...")

    numSeries := 0

    for _,wf := range config.Data {
        numSeries += len(wf.Variables)
    }

    events := make(chan WatchEvent, 1)
    cluster := Cluster{Key: config.Key, ID: hash(config.Key)}
    cluster.Series = make([]Series, numSeries)
    cluster.SeriesIDs = make([]string, numSeries)
    seriesIndex := 0

    for _,wf := range config.Data {
        /*
        filename, err := filepath.Abs(wf.Filename)
        if err != nil {
            panic(err)
        }
        */
        numVar := len(wf.Variables)
        patternList := make([]string, numVar)
        descList    := make([]string, numVar)
        indexList   := make([]int, numVar)
        for i,wv := range wf.Variables {
            //seriesStr := fmt.Sprintf("%s_%s",filename,wv.Pattern)
            seriesStr := fmt.Sprintf("%s.%s",cluster.Key,wv.Description)
            seriesID := hash(seriesStr)
            cluster.SeriesIDs[seriesIndex] = seriesID
            cluster.Series[seriesIndex] = Series{ID: seriesID}
            cluster.Series[seriesIndex].Events = []Event{}
            patternList[i] = wv.Pattern
            descList[i] = wv.Description
            indexList[i] = seriesIndex
            seriesIndex += 1
        }
        go watch(wf.Filename, cluster.Key, patternList, descList, indexList, events, time.Now().UTC().UnixNano() / int64(time.Microsecond))
    }

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
            upload(&cluster, config.Key)
            changed = false
        }
    }
}

func upload(cluster *Cluster, key string) {
	url := fmt.Sprintf("%s/clusters/%s", SERVER_URL, key)
	fmt.Println("[UPLOAD] URL:>", url)

	cjson, err := json.Marshal(cluster)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cjson))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(cjson))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	for i, _ := range cluster.Series {
		cluster.Series[i].Events = nil
	}
}

func hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil)[0:12])
}

func uploadHelper(events chan WatchEvent) {
	for {
		time.Sleep(time.Second * 1)
		events <- WatchEvent{SeriesIndex: -1, ID: "upload"}
	}
}

func watch(filename string, clusterKey string, patterns []string, descriptions []string, seriesIndices []int, events chan WatchEvent, watchStart int64) {
    fmt.Printf("[TRACKING] %s -- Variables: %s\n", filename, descriptions)
    lc := 0
    if filename != "PIPE" {
        t, err := tail.TailFile(filename, tail.Config{Follow: true})
        for line := range t.Lines {
            curTime := time.Now().UTC().UnixNano() / int64(time.Microsecond)
            if curTime - watchStart > 1e6 {
                for i,_ := range patterns {
                    match,_ := regexp.MatchString(patterns[i], line.Text)
                    if match {
                        r,_ := regexp.Compile(patterns[i])
                        //seriesStr := fmt.Sprintf("%s_%s",filename,patterns[i])
                        seriesStr := fmt.Sprintf("%s.%s",clusterKey,descriptions[i])
                        seriesID := hash(seriesStr)
                        allMatch := r.FindAllStringSubmatch(line.Text, -1)
                        for j,matchVal := range allMatch {
                            eventStr := fmt.Sprintf("%s-%s-%d-%d-%d-%s",filename,patterns[i],lc,j,curTime,matchVal[0])
                            eventID := hash(eventStr)
                            events <- WatchEvent{ID: eventID, SeriesID: seriesID, SeriesIndex: seriesIndices[i],
                                Key: descriptions[i], Value: matchVal[1], Time: curTime}
                        }
                    }
                }
            }
            lc = lc+1
        }
        if err != nil {
            panic(err)
        }
    } else {
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
            if curTime - watchStart > 1e6 {
                for i,_ := range patterns {
                    match,_ := regexp.MatchString(patterns[i], line)
                    if match {
                        r,_ := regexp.Compile(patterns[i])
                        seriesStr := fmt.Sprintf("%s.%s",clusterKey,descriptions[i])
                        seriesID := hash(seriesStr)
                        allMatch := r.FindAllStringSubmatch(line, -1)
                        for j,matchVal := range allMatch {
                            eventStr := fmt.Sprintf("%s-%s-%d-%d-%d-%s",filename,patterns[i],lc,j,curTime,matchVal[0])
                            eventID := hash(eventStr)
                            events <- WatchEvent{ID: eventID, SeriesID: seriesID, SeriesIndex: seriesIndices[i],
                                Key: descriptions[i], Value: matchVal[1], Time: curTime}
                        }
                    }
                }
            }
            lc = lc+1
            // process buf
            if err != nil && err != io.EOF {
                log.Fatal(err)
            }
        }
    }
}
