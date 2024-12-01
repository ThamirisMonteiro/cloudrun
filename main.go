package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

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
		_, err = w.Write([]byte(result))
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
	fmt.Println(re)
	if !re.MatchString(cep) {
		return errors.New("invalid zipcode")
	}
	return nil
}

func handleCEP(cep string) (error, string, int) {
	result, err, status := checkViaCEP(cep)
	if err != nil {
		return err, "", status
	}
	return nil, result, status
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
