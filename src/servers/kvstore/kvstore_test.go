package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var kv kvstore

func init() {
	kv = NewKvStore()
}

func GetList(t *testing.T) map[string]string {
	req := httptest.NewRequest(http.MethodGet, "/api/list", nil)
	w := httptest.NewRecorder()
	kv.ListApiRequest(w, req)
	response := w.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Expected 200, got ", response.StatusCode)
	}
	var data map[string]string
	err := json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func AddEntry(t *testing.T, key string, value string) {
	req := httptest.NewRequest(http.MethodGet, "/api/add?key="+key+"&value="+value, nil)
	w := httptest.NewRecorder()
	kv.AddApiRequest(w, req)
	response := w.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Expected 202, got ", response.StatusCode)
	}
}

func AddEntries(t *testing.T) {
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	dataAsJson, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/api/addAll", bytes.NewBuffer(dataAsJson))
	w := httptest.NewRecorder()
	kv.AddEntriesApiRequest(w, req)
	response := w.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Expected 202, got ", response.StatusCode)
	}
}

func ClearAll(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/clear", nil)
	w := httptest.NewRecorder()
	kv.ClearApiRequest(w, req)
	response := w.Result()
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatal("Expected 200, got ", response.StatusCode)
	}
}

func TestGetEmptyList(t *testing.T) {
	data := GetList(t)
	if len(data) != 0 {
		t.Fatal("Expected empty map, got ", data)
	}
}

func TestAddEntry(t *testing.T) {
	key := "mykey"
	value := "myvalue"
	AddEntry(t, key, value)
	data := GetList(t)
	if len(data) != 1 {
		t.Fatal("Expected map with one entry, got ", data)
	}

	ClearAll(t)
	data = GetList(t)
	if len(data) != 0 {
		t.Fatal("Expected empty map, got ", data)
	}
}

func TestAddEntries(t *testing.T) {
	ClearAll(t)
	AddEntries(t)
	data := GetList(t)
	if len(data) != 3 {
		t.Fatal("Expected map with one entry, got ", data)
	}
}
