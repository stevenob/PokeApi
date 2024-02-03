package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var page int = 0

func getLocations() (Locations, error) {
	offset := page * 20
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=20", offset)
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
	location := Locations{}
	err = json.Unmarshal(body, &location)
	return location, err
}
