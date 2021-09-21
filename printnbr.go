package piscine

import "github.com/01-edu/z01"

func Convert(nb int) {
	runeNum := '0'

	// accounting for int 0
	if nb == 0 {
		z01.PrintRune(runeNum)
	}

	for i := 1; i <= nb%10; i++ {
		runeNum++
	}

	for i := -1; i >= nb%10; i-- {
		runeNum++
	}

	if nb/10 != 0 {
		Convert(nb / 10)
	}
	z01.PrintRune(runeNum)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	Convert(n)
}
