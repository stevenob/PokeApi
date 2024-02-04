package main

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Attempt to catch pokemon",
			callback:    commandCatch,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Display caught pokemon stats and info",
			callback:    commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokeman in input area",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Displays loactions",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous loactions",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display caught pokemon",
			callback:    commandPokedex,
		},
	}
}
