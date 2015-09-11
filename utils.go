package main

import (
    "time"
    "log"
)

type (
  empty struct{}
)

/* Profiling utility */
func Profile(start time.Time, desc string) {
    t := time.Since(start)
    log.Printf("[Profile] %s took %s time", desc, t)
}

func Now() time.Time {
    return time.Now()
}
