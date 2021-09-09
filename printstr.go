package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for s := 0; s <= 10; s++ {
		z01.PrintRune(rune(s))
	}
}
