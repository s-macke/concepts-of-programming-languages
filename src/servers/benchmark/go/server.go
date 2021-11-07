package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func StartWithPort(port int) {
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(1000 * time.Millisecond)
		if n, err := fmt.Fprintf(w, "Hello!\n"); err != nil {
			log.Fatalf("Error after writing %d bytes: %v", n, err.Error())
		}
	})

	fmt.Printf("Starting server at port %d\n", port)
	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(port), "cert.crt", "cert.key", m))
}

func main() {
	StartWithPort(8080)
}
