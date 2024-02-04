package main

import (
	"fmt"
)

func commandInspect(args ...string) error {
	name := args[0]
	pokemon, err := getPokemon(name)
	fmt.Printf("\nName: %s\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	statLst := []string{"-hp:", "-attack:", "-defense:", "-special-attack:", "-special-defense:", "-speed:"}
	fmt.Print("Stats:")
	for s := range pokemon.Stats {
		fmt.Printf("\n   %s %v", statLst[s], pokemon.Stats[s].BaseStat)
	}
	fmt.Print("\nTypes:")
	for t := range pokemon.Types {
		fmt.Print("\n   -", pokemon.Types[t].Type.Name)
	}
	println("")

	return err
}
