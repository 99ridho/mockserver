package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/99ridho/mockserver/models"
)

func readRulesFromFile(path string) ([]models.Rule, error) {
	data, err := readFile(path)

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
	port := flag.String("port", "8181", "Server port")
	rulesPath := flag.String("rule", "./rules.json", "Rule path file")
	flag.Parse()

	rules, err := readRulesFromFile(*rulesPath)

	if err != nil {
		log.Fatalf("Can't read rule, error thrown: %v\n", err)
	}

	for _, rule := range rules {
		http.HandleFunc(rule.Path, makeHandler(rule))
	}

	log.Printf("Server started at port %s", *port)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
