package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for i := range a {
		runes := []rune(a[i])
		for i := range runes {
			z01.PrintRune(runes[i])
		}
		z01.PrintRune('\n')
	}
}
