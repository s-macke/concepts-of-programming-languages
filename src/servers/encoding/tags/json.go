package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Bird struct {
	Species     string `json:"birdType"`
	Description string
}

func main() {
	//var Species string `json:"birdType"`

	birdJson := `{"Species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird

	fmt.Println(reflect.TypeOf(bird).Field(0).Tag)
	//fmt.Println(reflect.Indirect(reflect.ValueOf(bird)).Type().Field(0))
	//fmt.Println(reflect.Indirect(reflect.ValueOf(bird)).Type().Field(0))

	err := json.Unmarshal([]byte(birdJson), &bird)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Species: %s\nDescription: %s\n", bird.Species, bird.Description)
	s, _ := json.Marshal(bird)
	fmt.Println(string(s))
}
