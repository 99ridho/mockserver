package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/99ridho/mockserver/models"
)

func readRulesFromFile() ([]models.Rule, error) {
	data, err := readFile("./rules.json")

	if err != nil {
		return []models.Rule{}, err
	}

	rules := make([]models.Rule, 0)
	err = json.Unmarshal(data, &rules)

	if err != nil {
		return []models.Rule{}, err
	}

	return rules, nil
}

func readFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

func makeHandler(rule models.Rule) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := ""

		if rule.Response.Type == "JSON" {
			w.Header().Add("Content-Type", "application/json")

			fileContent, _ := readFile(rule.Response.Value)
			body = string(fileContent)
		} else if rule.Response.Type == "TEXT" {
			body = rule.Response.Value
		}

		if rule.Response.StatusCode != 0 && rule.Response.StatusCode >= 100 && rule.Response.StatusCode <= 599 {
			w.WriteHeader(rule.Response.StatusCode)
		}

		fmt.Fprintln(w, body)
	}
}

func main() {
	rules, err := readRulesFromFile()

	if err != nil {
		panic("Can't read rule")
	}

	for _, rule := range rules {
		http.HandleFunc(rule.Path, makeHandler(rule))
	}

	log.Println("Server started")
	http.ListenAndServe(":8181", nil)
}
