package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for d1 := '0'; d1 <= '9'; d1++ {
		for d2 := '0'; d2 <= '9'; d2++ {
			d4 := d2 + 1
			for d3 := '0'; d3 <= '9'; d3++ {
				for ; d4 <= '9'; d4++ {
					z01.PrintRune(d1)
					z01.PrintRune(d2)
					z01.PrintRune(' ')
					z01.PrintRune(d3)
					z01.PrintRune(d4)
					if d1 < '9' || d2 < '8' || d3 < '9' || d4 < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				d4 = '0'
			}
		}

	}
	z01.PrintRune('\n')
}
