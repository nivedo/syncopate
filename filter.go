package main

import (
    "log"
)

type (
    Filter interface {
        // Match returns true if the filter has been completed.
        // This may take multiple calls as filters are "stateful"
        // and can handle multiple lines.
        Match(data string) bool
        // GetVars() returns a KVList of the last known values
        // to be successfully matched.
        GetVars() KVList
        // Returns number of variables for the filter,
        // which should be a fixed number in most cases.
        NumVars() int
    }
)

func GetFilter(opt Option_t) Filter {
    if IsFilterRegex(opt) {
        return NewFilterRegex(opt)
    }

    log.Fatalf("Filter doesn't exist for option %+v", opt)
    return nil
}