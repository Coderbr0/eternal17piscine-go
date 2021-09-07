package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	for i := 0; i <= 9; i++ {
		fmt.Println(i)
		z01.PrintRune('\n')
	}
}
