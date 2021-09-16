package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	programName := arg[0]
	nRune := []rune(programName)

	for i := 2; i <= len(nRune)-1; i++ {
		z01.PrintRune(nRune[i])
	}
	z01.PrintRune('\n')
}
