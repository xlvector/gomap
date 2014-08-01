package main

import (
	"github.com/xlvector/gomap"
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/nearby", gomap.NewNearByService())
	http.Handle("/place", gomap.NewPlaceService())
	http.Handle("/direction", gomap.NewDirectionService())
	s := &http.Server{
		Addr:           ":8903",
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
