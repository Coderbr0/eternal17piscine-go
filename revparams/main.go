package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	// looping through the arguments backwards
	for i := len(arguments) - 1; i > 0; i-- {
		// print each word followed by a line break
		for _, word := range arguments[i] {
			z01.PrintRune(word)
		}
		z01.PrintRune('\n')
	}
}
