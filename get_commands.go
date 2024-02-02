package main

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
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
	}
}
