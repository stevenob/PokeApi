package main

import "fmt"

func commandPokedex(args ...string) error {
	fmt.Println("Your PokeDex: ")
	for poke := range pokedex {
		fmt.Println("  -", pokedex[poke].Name)
	}
	return nil
}
