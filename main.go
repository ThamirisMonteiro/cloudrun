package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var key = "f4325b77228346ef861132533240112"

type viacepResponse struct {
	Localidade string `json:"localidade"`
	Estado     string `json:"estado"`
}

type weatherapiResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cep := strings.TrimPrefix(r.URL.Path, "/")
		err := validateCEP(cep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		err, result, status := handleCEP(cep)
		if err != nil {
			http.Error(w, err.Error(), status)
		}

		temp, err := getTemperature(result.Localidade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		_, err = w.Write([]byte(temp))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func getTemperature(localidade string) (string, error) {
	url := "https://api.weatherapi.com/v1/current.json?key=" + key + "&q=" + localidade
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	response := &weatherapiResponse{}
	err = json.Unmarshal(body, &response)

	tempF := response.Current.TempC*1.8 + 32
	tempK := response.Current.TempC + 273

	temperature := map[string]float64{
		"temp_C": response.Current.TempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	temperatureJSON, err := json.Marshal(temperature)
	if err != nil {
		return "", err
	}

	return string(temperatureJSON), nil
}

func validateCEP(cep string) error {
	re := regexp.MustCompile(`^\d{8}$`)
	if !re.MatchString(cep) {
		return errors.New("invalid zipcode")
	}
	return nil
}

func handleCEP(cep string) (error, *viacepResponse, int) {
	result, err, status := checkViaCEP(cep)
	if err != nil {
		return err, nil, status
	}
	response := &viacepResponse{}
	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		return err, nil, http.StatusInternalServerError
	}
	return nil, response, status
}

func checkViaCEP(cep string) (string, error, int) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err, http.StatusNotFound
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	if erro, exists := response["erro"]; exists && erro == "true" {
		return "", errors.New("can not find zipcode"), http.StatusNotFound
	}

	return string(body), nil, http.StatusOK
}
