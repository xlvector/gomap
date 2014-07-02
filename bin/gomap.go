package main

import (
    "net/http"
    "github.com/xlvector/gomap"
    "time"
    "log"
)

func main() {
    http.Handle("/nearby", gomap.NewNearByService())
    s := &http.Server{
        Addr:           ":8903",
        ReadTimeout:    30 * time.Second,
        WriteTimeout:   30 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())
}