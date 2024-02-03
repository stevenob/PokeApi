package main

import "testing"

//Test get_commands.getCommands()
func TestGetCommands(t *testing.T) {
	cmd := getCommands()
	if cmd == nil {
		t.Fatal("Failed to read command list")
	}
}

//Test command_help.commandHelp()
func TestCommandHelp(t *testing.T) {
	err := commandHelp()
	if err != nil {
		t.Fatal(err)
	}
}

//Test command_map.commandMap()
func TestCommandMap(t *testing.T) {
	err := commandMap()
	if err != nil {
		t.Fatal(err)
	}
}

//Test command_mapb.commandMapb()
func TestCommandMapb(t *testing.T) {
	err := commandMapb()
	if err != nil {
		t.Fatal(err)
	}
}
