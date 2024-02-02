package main

import (
	"fmt"
)

func commandMapb() error {
	if page == 0 {
		println("Page already 1")
		return nil
	} else {
		defer println("page ", page-1)
		page--
		loc, _ := getLocations()
		for _, l := range loc.Results {
			fmt.Println(l.Name)
		}
		return nil
	}

}
