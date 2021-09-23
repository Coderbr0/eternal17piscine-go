package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	// count := 0
	for _, arg := range args {
		if arg == "01" || arg == "galaxy" || arg == "01 galaxy" {
			// count++
			fmt.Println("Alert!!!")
		}
	}
	// for i := 1; i <= count; i++ {
	// 	fmt.Println("Alert!!!")
	// }
}
