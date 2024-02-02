package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocations() (locations, error) {
	url := "https://pokeapi.co/api/v2/location/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	location := locations{}
	err = json.Unmarshal(body, &location)
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return location, nil
}
