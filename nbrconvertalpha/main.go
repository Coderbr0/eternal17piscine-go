package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	caps := false

	for i := range args {
		if args[i] == "--upper" {
			caps = true
		}
	}
	for _, arg := range args {
		num := 0
		for _, n := range arg {
			num = num*10 + int(n-48)
		}
		if 1 <= num && num <= 26 {
			if caps == false {
				z01.PrintRune(rune(num + 96))
			} else {
				z01.PrintRune(rune(num + 64))
			}
		}
		if num > 26 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
