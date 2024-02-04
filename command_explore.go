package main

import (
	"fmt"
)

func commandExplore(args ...string) error {
	a := args[0]
	fmt.Printf("Exploring %s...\n", a)
	pokemon, _ := getLocalPokemon(a)
	fmt.Println("Found Pokemon: ")
	for p := range pokemon.PokemonEncounters {
		fmt.Println("-", pokemon.PokemonEncounters[p].Pokemon.Name)
	}
	// fmt.Print(pokemon.PokemonEncounters[1].Pokemon.Name)
	return nil

}
