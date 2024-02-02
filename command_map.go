package main

import (
	"fmt"
)

func commandMap() error {
	location, _ := getLocations()
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
