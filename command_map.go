package main

import (
	"fmt"
)

var offset int = page * 20
var areaUrl string = fmt.Sprintf("?offset=%v&limit=20", offset)

func commandMap(args ...string) error {
	loc, err := getLocations(areaUrl)
	for _, l := range loc.Results {
		fmt.Println(l.Name)
	}
	println("page: ", page)
	page++
	return err
}

func commandMapb(args ...string) error {
	if page == 0 {
		println("Page already 1")
		return nil
	} else {
		defer println("page ", page-1)
		page--
		loc, _ := getLocations(areaUrl)
		for _, l := range loc.Results {
			fmt.Println(l.Name)
		}
		return nil
	}

}
