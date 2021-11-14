package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

type kvstore struct {
	store map[string]string
}

func (kv *kvstore) AddApiRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("key") && r.URL.Query().Has("value") {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		kv.store[key] = value
		//fmt.Printf("Added key: %s, value: %s\m", key, value)
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Printf("Missing key or value")
		http.Error(w, "Missing key or value URL parameter", http.StatusBadRequest)
	}
}

func (kv *kvstore) AddEntriesApiRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method must be POST", http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Unable to decode the json structure", http.StatusBadRequest)
	}
	for key, value := range data {
		kv.store[key] = value
	}
	w.WriteHeader(http.StatusOK)
}

func (kv *kvstore) ListApiRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(kv.store)
	w.Write(data)
}

func (kv *kvstore) ClearApiRequest(w http.ResponseWriter, r *http.Request) {
	kv.store = make(map[string]string)
	w.WriteHeader(http.StatusOK)
}

func NewKvStore() kvstore {
	var kv kvstore
	kv.store = make(map[string]string)
	http.HandleFunc("/api/add", kv.AddApiRequest)
	http.HandleFunc("/api/clear", kv.ClearApiRequest)
	http.HandleFunc("/api/list", kv.ListApiRequest)
	http.HandleFunc("/api/addAll", kv.AddEntriesApiRequest)
	return kv
}

// content is our static web server content.
//go:embed static
var staticFS embed.FS

func AddFilesystem() {
	serverRoot, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(serverRoot)))
}

func main() {
	kv := NewKvStore()
	kv.store["abcd"] = "efgh"
	kv.store["efgh"] = "ijkl"
	kv.store["ijkl"] = "mnop"
	AddFilesystem()

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
