package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func main() {
	targetPath := flag.String("path", ".", "-path /path/to")
	addr := flag.String("addr", "127.0.0.1:8080", "-addr 127.0.0.1:8080")
	flag.Parse()

	fs := http.FileServer(http.Dir(*targetPath))
	http.StripPrefix("/", fs)
	server := http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         *addr,
		Handler:      http.StripPrefix("/", fs),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
