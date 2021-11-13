package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Nothing to see here\n")
}

func EchoParameters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(r.URL.Query())
	w.Write(data)
}

func main() {
	http.HandleFunc("/", HandleNotFound)
	http.HandleFunc("/api/echoParams", EchoParameters)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
