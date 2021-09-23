package piscine

import "github.com/01-edu/z01"

func Convert(nbr int) {
	numRune := '0'

	for i := 1; i <= nbr%10; i++ {
		numRune++
	}

	for i := -1; i >= nbr%10; i-- {
		numRune++
	}

	if nbr/10 != 0 {
		Convert(nbr / 10)
	}
	z01.PrintRune(numRune)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	Convert(n)
}
