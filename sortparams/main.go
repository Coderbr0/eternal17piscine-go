package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// convert into args using os
	args := os.Args[1:]

	i := 0
	for i < len(args)-1 {
		if args[i] > args[i+1] {
			args[i], args[i+1] = args[i+1], args[i]
			i = 0
		} else {
			i++
		}
	}
	for i := range args {
		for _, ch := range args[i] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
