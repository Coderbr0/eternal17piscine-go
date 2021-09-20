package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:]
	evenS := "I have an even number of arguments"
	oddS := "I have an odd number of arguments"
	if isEven(len(args)) {
		for i := range evenS {
			z01.PrintRune(rune(evenS[i]))
		}
	} else {
		for i := range oddS {
			z01.PrintRune(rune(oddS[i]))
		}
	}
	z01.PrintRune('\n')
}
