package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	programName := arg[0]
	nRune := []rune(programName)

	for i := range nRune {
		z01.PrintRune(nRune[i])
	}
	z01.PrintRune('\n')
}
