package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherForecast struct {
	CurrentCondition []struct {
		FeelsLikeC       string `json:"FeelsLikeC"`
		FeelsLikeF       string `json:"FeelsLikeF"`
		Cloudcover       string `json:"cloudcover"`
		Humidity         string `json:"humidity"`
		LocalObsDateTime string `json:"localObsDateTime"`
		ObservationTime  string `json:"observation_time"`
		PrecipInches     string `json:"precipInches"`
		PrecipMM         string `json:"precipMM"`
		Pressure         string `json:"pressure"`
		PressureInches   string `json:"pressureInches"`
		TempC            string `json:"temp_C"`
		TempF            string `json:"temp_F"`
		UvIndex          string `json:"uvIndex"`
		Visibility       string `json:"visibility"`
		VisibilityMiles  string `json:"visibilityMiles"`
		WeatherCode      string `json:"weatherCode"`
		WeatherDesc      []struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WeatherIconURL []struct {
			Value string `json:"value"`
		} `json:"weatherIconUrl"`
		Winddir16Point string `json:"winddir16Point"`
		WinddirDegree  string `json:"winddirDegree"`
		WindspeedKmph  string `json:"windspeedKmph"`
		WindspeedMiles string `json:"windspeedMiles"`
	} `json:"current_condition"`
	NearestArea []struct {
		AreaName []struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Country []struct {
			Value string `json:"value"`
		} `json:"country"`
		Latitude   string `json:"latitude"`
		Longitude  string `json:"longitude"`
		Population string `json:"population"`
		Region     []struct {
			Value string `json:"value"`
		} `json:"region"`
		WeatherURL []struct {
			Value string `json:"value"`
		} `json:"weatherUrl"`
	} `json:"nearest_area"`
	Request []struct {
		Query string `json:"query"`
		Type  string `json:"type"`
	} `json:"request"`
	Weather []struct {
		Astronomy []struct {
			MoonIllumination string `json:"moon_illumination"`
			MoonPhase        string `json:"moon_phase"`
			Moonrise         string `json:"moonrise"`
			Moonset          string `json:"moonset"`
			Sunrise          string `json:"sunrise"`
			Sunset           string `json:"sunset"`
		} `json:"astronomy"`
		AvgtempC string `json:"avgtempC"`
		AvgtempF string `json:"avgtempF"`
		Date     string `json:"date"`
		Hourly   []struct {
			DewPointC        string `json:"DewPointC"`
			DewPointF        string `json:"DewPointF"`
			FeelsLikeC       string `json:"FeelsLikeC"`
			FeelsLikeF       string `json:"FeelsLikeF"`
			HeatIndexC       string `json:"HeatIndexC"`
			HeatIndexF       string `json:"HeatIndexF"`
			WindChillC       string `json:"WindChillC"`
			WindChillF       string `json:"WindChillF"`
			WindGustKmph     string `json:"WindGustKmph"`
			WindGustMiles    string `json:"WindGustMiles"`
			Chanceoffog      string `json:"chanceoffog"`
			Chanceoffrost    string `json:"chanceoffrost"`
			Chanceofhightemp string `json:"chanceofhightemp"`
			Chanceofovercast string `json:"chanceofovercast"`
			Chanceofrain     string `json:"chanceofrain"`
			Chanceofremdry   string `json:"chanceofremdry"`
			Chanceofsnow     string `json:"chanceofsnow"`
			Chanceofsunshine string `json:"chanceofsunshine"`
			Chanceofthunder  string `json:"chanceofthunder"`
			Chanceofwindy    string `json:"chanceofwindy"`
			Cloudcover       string `json:"cloudcover"`
			Humidity         string `json:"humidity"`
			PrecipInches     string `json:"precipInches"`
			PrecipMM         string `json:"precipMM"`
			Pressure         string `json:"pressure"`
			PressureInches   string `json:"pressureInches"`
			TempC            string `json:"tempC"`
			TempF            string `json:"tempF"`
			Time             string `json:"time"`
			UvIndex          string `json:"uvIndex"`
			Visibility       string `json:"visibility"`
			VisibilityMiles  string `json:"visibilityMiles"`
			WeatherCode      string `json:"weatherCode"`
			WeatherDesc      []struct {
				Value string `json:"value"`
			} `json:"weatherDesc"`
			WeatherIconURL []struct {
				Value string `json:"value"`
			} `json:"weatherIconUrl"`
			Winddir16Point string `json:"winddir16Point"`
			WinddirDegree  string `json:"winddirDegree"`
			WindspeedKmph  string `json:"windspeedKmph"`
			WindspeedMiles string `json:"windspeedMiles"`
		} `json:"hourly"`
		MaxtempC    string `json:"maxtempC"`
		MaxtempF    string `json:"maxtempF"`
		MintempC    string `json:"mintempC"`
		MintempF    string `json:"mintempF"`
		SunHour     string `json:"sunHour"`
		TotalSnowCm string `json:"totalSnow_cm"`
		UvIndex     string `json:"uvIndex"`
	} `json:"weather"`
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://wttr.in/Rosenheim?format=j1", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var forecast WeatherForecast
	err = json.Unmarshal(bodyBytes, &forecast)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Temperature in Rosenheim: %s°C\n", forecast.CurrentCondition[0].TempC)
}
