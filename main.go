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

func main() {
	handleApiData()
}

func handleApiData() {
	userApiKey := getApiKey()

	apiData := getApiData(userApiKey)
	displayApiData(apiData)
}

func getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}
	apiKey := os.Getenv("API_KEY")

	return apiKey
}

func getApiData(apiKey string) []byte {
	cityFlag := cityFlag()
	userApiUrl := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=2&aqi=no&alerts=no", apiKey, cityFlag)

	res, err := http.Get(userApiUrl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("API not available right now")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	return body
}



func displayApiData(dataBody []byte) {
	var weather Weather
	err := json.Unmarshal(dataBody, &weather)
	if err != nil {
		panic(err)
	}

	var (
		location = weather.Location
		hours = weather.Forecast.Forecastday[0].ByHour
		dateNow = weather.Forecast.Forecastday[0].Date
	)

	parsedWeather := fmt.Sprintf("%s, %s | %s |", location.Name, location.Country, dateNow)
	fmt.Println(parsedWeather)
	

	tempFlag := tempFlag() // returns "c" or "f"

	var weatherByHours string
	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if tempFlag == "c" {
			weatherByHours = fmt.Sprintf("%s | %.0fC, %d%%, %s", date.Format("15:04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)
		}
		if tempFlag == "f" {
			weatherByHours = fmt.Sprintf("%s | %.0fF, %d%%, %s", date.Format("15:04"), hour.TempF, hour.ChanceOfRain, hour.Condition.Text)
		}

		if date.Hour() == time.Now().Hour() {
			color.Green(weatherByHours)
		} else {
			fmt.Println(weatherByHours)
		}

	}
}


func cityFlag() string {
	cityFlag := "London"
	if len(os.Args) >= 2 {
		cityFlag = os.Args[1]
	}
	return cityFlag
}

func tempFlag() string {
	tempFlag := "c"
	if len(os.Args) >= 3 { 
		tempFlag = os.Args[2]
	}
	return tempFlag
}
