package piscine

import "github.com/01-edu/z01"

func PointOne(n *int) {
	n = new(int)
	*n = 1
	z01.PrintRune(rune(*n))
}
