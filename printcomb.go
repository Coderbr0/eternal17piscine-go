package piscine

import "github.com/01-edu/z01"

func main() {
	for digit1 := '0'; digit1 <= '7'; digit1++ {
		for digit2 := '1'; digit2 <= '8'; digit2++ {
			for digit3 := '2'; digit3 <= '9'; digit3++ {
				if digit1 < digit2 && digit2 < digit3 {
					z01.PrintRune(digit1)
					z01.PrintRune(digit2)
					z01.PrintRune(digit3)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				if digit1 == '7' && digit2 == '8' && digit3 == '9' {
					z01.PrintRune(digit1)
					z01.PrintRune(digit2)
					z01.PrintRune(digit3)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
