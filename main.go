package main

import (
	"errors"
	"fmt"
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

		err, result := handleCEP(cep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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

func handleCEP(cep string) (error, string) {
	return nil, cep
}
