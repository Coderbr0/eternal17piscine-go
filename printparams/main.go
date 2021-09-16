package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	// loop over the arguments, convert each one into a rune so we can use printRune
	for i := 1; i <= len(arguments)-1; i++ {
		runes := []rune(arguments[i])
		// print each rune followed by a line break
		for i := range runes {
			z01.PrintRune(runes[i])
		}
		z01.PrintRune('\n')
	}
}
