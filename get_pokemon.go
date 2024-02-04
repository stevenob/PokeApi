package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getPokemon(url string) (Pokemon, error) {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	res, err := http.Get(baseUrl + url)
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
	// pokecache.cache.Add(url, body)
	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	return pokemon, err
}
