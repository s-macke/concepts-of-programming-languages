package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("p", "8081", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	dirFs := os.DirFS(*directory)
	http.Handle("/", http.FileServer(http.FS(dirFs)))
	fmt.Printf("Starting server at port " + *port + "\n")
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
