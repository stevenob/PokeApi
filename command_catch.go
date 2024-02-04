package main

import (
	"fmt"
	"math/rand"
)

var pokedex = make(map[string]Pokemon)

func commandCatch(args ...string) error {
	name := args[0]
	fmt.Printf("\nThrowing pokeball at %s...\n", name)
	pokemon, err := getPokemon(name)
	chance := rand.Intn(pokemon.BaseExperience)
	if chance < 40 {
		fmt.Printf("%s was caught!\n", name)
		pokedex[name] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return err
}
