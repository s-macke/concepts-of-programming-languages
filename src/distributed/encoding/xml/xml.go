package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Weather struct {
	XMLName xml.Name `xml:"weather"`
	Text    string   `xml:",chardata"`
	City    string   `xml:"city"`
	Sky     string   `xml:"sky"`
	Temp    struct {
		Text string `xml:",chardata"`
		Unit string `xml:"unit,attr"`
	} `xml:"temp"`
}

func main() {
	var weather Weather
	weatherXml := `<weather><city>Rosenheim</city><sky>clear</sky><temp unit="°C">15</temp></weather>`
	err := xml.Unmarshal([]byte(weatherXml), &weather)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("City: %s\nSky: %s\nTemperature: %s%s\n", weather.City, weather.Sky, weather.Temp.Text, weather.Temp.Unit)
	s, _ := xml.Marshal(weather)
	fmt.Println(string(s))
}
