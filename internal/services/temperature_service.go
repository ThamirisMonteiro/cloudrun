package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

var WeatherAPIKey = "f4325b77228346ef861132533240112"

func GetTemperature(location string) (string, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", WeatherAPIKey, location)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response WeatherAPIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	tempF := response.Current.TempC*1.8 + 32
	tempK := response.Current.TempC + 273.15

	tempC := math.Round(response.Current.TempC*10) / 10
	tempF = math.Round(tempF*10) / 10
	tempK = math.Round(tempK*10) / 10

	temperature := map[string]float64{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	temperatureJSON, err := json.Marshal(temperature)
	if err != nil {
		return "", err
	}

	return string(temperatureJSON), nil
}
