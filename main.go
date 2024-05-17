package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var path string
	var port string

	flag.StringVar(&port, "p", "8080", "port to serve on")
	flag.StringVar(&path, "d", ".", "filepath")
	flag.Parse()

	s, err := os.Stat(path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		log.Fatalf("path '%s' not found", path)
	} else if err != nil {
		log.Fatal(err)
	}

	if s.IsDir() {
		http.Handle("/", http.FileServer(http.Dir(path)))
	} else {
		http.HandleFunc("/", serveFile(path))
	}

	log.Printf("Listening on :%s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveFile(path string) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, path)
	}
}
