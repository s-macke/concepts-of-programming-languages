package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Weather struct {
	City        string `json:"city"` // The tag defines the exact name in the json.
	Sky         string `json:"sky"`
	Temperature int    `json:"temp"`
}

func main() {
	var weather Weather
	// Tags are accessible via the reflection
	fmt.Println("Tag of first field entry", reflect.TypeOf(weather).Field(0).Tag)

	weatherJson := `{"city": "Rosenheim","sky": "clear", "temp": 15}`
	err := json.Unmarshal([]byte(weatherJson), &weather)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("City: %s\nSky: %s\nTemperature: %d\n", weather.City, weather.Sky, weather.Temperature)
	s, _ := json.Marshal(weather)
	fmt.Println(string(s))
}
