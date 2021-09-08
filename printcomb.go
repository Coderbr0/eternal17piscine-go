package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i < 56; i++ {
		for j := i + 1; j < 57; j++ {
			for k := j + 1; k < 58; k++ {
				if i == 55 && j == 56 && k == 57 {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(j))
					z01.PrintRune(rune(k))
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
		}
	}
}
