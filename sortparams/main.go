// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	// convert into args using os
// 	args := os.Args[1:]

// 	for i := range args {
// 		ar := []rune(args[i])
// 		j := 0
// 		if '0' <= ar[j] && ar[j] <= '9' {
// 			for _, j := range args[i] {
// 				z01.PrintRune(j)
// 				z01.PrintRune('\n')
// 			}
// 		}

// 	}

// }
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if runes[j] >= '0' && runes[j] <= '9' {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if runes[j] >= 'A' && runes[j] <= 'Z' {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
	for i := range arguments {
		runes := []rune(arguments[i])
		j := 0
		if runes[j] >= 'a' && runes[j] <= 'z' {
			for _, j := range arguments[i] {
				z01.PrintRune(j)
				z01.PrintRune('\n')
			}
		}
	}
}
