package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byts, err := ioutil.ReadFile("main.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", byts)
}
