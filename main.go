package main

import (
	"fmt"
	"io/ioutil"
)

func test() {
	files, err := ioutil.ReadDir("testdata")
	if err != nil {
		panic(err)
	}
	for i := range files {
		fmt.Println("Reading: " + files[i].Name())
		file, err := ioutil.ReadFile("testdata/" + files[i].Name())
		if err != nil {
			panic(err)
		}
		parse(file)
	}
}

func main() {
	test()
}
