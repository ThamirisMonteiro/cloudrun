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

type viacepResponse struct {
	Localidade string `json:"localidade"`
	Estado     string `json:"estado"`
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
		_, err = w.Write([]byte(result.Localidade + " - " + result.Estado))
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
