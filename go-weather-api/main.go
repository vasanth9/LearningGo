package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct{
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
       TempC float64 `json:"temp_c"`
	   Condition struct {
           Text string `json:"text"`
	   } `json:"condition"`
	}`json:"current"`
	ForeCast struct {
		ForeCastDay []struct{
			Hour []struct{
              TimeEpoch int64 `json:"time_epoch"`
			  TempC float64 `json:"temp_c"`
			  Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
			   ChanceOfRain float64 `json:"chance_of_rain"`
				}`json:"hour"`
		}`json:"forecastday"`
	} `json:"forecast"`
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }
  

func main() {
	q := "Vizag"
	if len(os.Args) >= 2{
      q = os.Args[1]
	}
    apiKey := goDotEnvVariable("WEATHER_API_KEY")
	response, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key="+apiKey+"&q="+q+"&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		panic("Weather API is not available")
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var weather Weather
	err = json.Unmarshal(body,&weather)
	if err !=nil {
		panic(err)
	}
	location, current, hours := weather.Location, weather.Current, weather.ForeCast.ForeCastDay[0].Hour

	fmt.Printf(
		"%s, %s: %.0fC %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)
	for _, hour := range(hours){
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}
		message := fmt.Sprintf(
			"%s - %.0fC, %.0f%%, %s\n",
            date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		if hour.ChanceOfRain < 20 {
			color.Green(message)
		} else if hour.ChanceOfRain < 50 {
			color.Yellow(message)
		} else {
			color.Red(message)
		}
	}

}
