package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	fileServer := http.FileServer(http.Dir(*directory))
	http.Handle("/", fileServer)
	fmt.Printf("Starting server at port " + *port + "\n")
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
