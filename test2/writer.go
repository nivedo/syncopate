package main

import (
    "os"
    "fmt"
    "time"
    "math/rand"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    f, err := os.OpenFile("test1000.log", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }

    defer f.Close()

    lastTemp := 300
    lastLoad := 300

    for {
        if rand.Intn(2) == 1 {
            lastTemp = rand.Intn(500)
            if _, err = f.WriteString(fmt.Sprintf("TEMP: %d\n",lastTemp)); err != nil {
                panic(err)
            }
        } else {
            lastLoad = rand.Intn(500)
            if _, err = f.WriteString(fmt.Sprintf("LOAD: %d\n",lastLoad)); err != nil {
                panic(err)
            }
        }
        time.Sleep(time.Duration(rand.Intn(1e9)))
    }
}
