package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := flag.String("port", "11000", "Listening port")
	flag.Parse()
	http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("TEST. VERSION: 1.0.1. Timestamp: %v", time.Now().UnixNano())))
	})
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Printf("Failed to start web-server. Error: %v", err)
	}
	log.Fatal(err)
	return
}
