package main

import "github.com/01-edu/z01"

func main() {
// Declare a variable string
	var digits string = "123456789"
// Loop through string until length -1 as length is returned as bytes
	for i := 0; i <= len(digits) -1; i++ {
// use printrune and convert the string digits into a rune
	z01.PrintRune(rune(digits[i]))
	}
	//end the line
	z01.PrintRune('\n')
}

