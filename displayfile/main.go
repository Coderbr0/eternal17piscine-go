package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]

	// more than one argument

	if len(arguments) > 1 {
		fmt.Println("Too many arguments")
	} else if len(arguments) < 1 {
		fmt.Println("File name missing")
	} else {
		content, _ := ioutil.ReadFile("quest8.txt")
		fmt.Printf(string(content))

	}
}
