package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	count := 0
	for _, arg := range args {
		if arg == "01" || arg == "galaxy" || arg == "01 galaxy" {
			count++
		}
	}
	if count >= 1 {
		fmt.Println("Alert!!!")
	}
}
