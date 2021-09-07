package main

import "github.com/01-edu/z01"


func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

// func reverseString(s string) string {
// 	chars := []rune(s)
// 	for i, j := 0, len(chars)-1; i < j;   
// }

