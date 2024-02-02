package main

import (
	"fmt"
)

func commandMap() error {
	loc, _ := getLocations()
	for _, l := range loc.Results {
		fmt.Println(l.Name)
	}
	println("page: ", page)
	page++
	return nil
}
