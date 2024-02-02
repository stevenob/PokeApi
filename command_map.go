package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	type locations struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous any    `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}
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
	// fmt.Printf("%v", location.Results)
	return nil
}
