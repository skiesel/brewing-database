package main

import (
	"fmt"
	"github.com/skiesel/brewing-database/BrewingElements"
	"io/ioutil"
)

func test() {
	files, err := ioutil.ReadDir("TestData")
	if err != nil {
		panic(err)
	}
	for i := range files {
		fmt.Println("Reading: " + files[i].Name())
		file, err := ioutil.ReadFile("TestData/" + files[i].Name())
		if err != nil {
			panic(err)
		}
		elements.Parse(file)
	}
}

func main() {
	test()
}
