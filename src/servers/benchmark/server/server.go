package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1000 * time.Millisecond)
		if n, err := fmt.Fprintf(w, "Hello!\n"); err != nil {
			log.Fatalf("Error after writing %d bytes: %v", n, err.Error())
		}
	})

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServeTLS(":8080", "cert.crt", "cert.key", nil))
}
