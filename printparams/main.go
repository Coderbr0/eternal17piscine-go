package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	// loop over the arguments
	for i := 1; i <= len(arguments)-1; i++ {
		// loop over each arguments and print the word, char by char
		for _, word := range arguments[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
