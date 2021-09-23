package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for _, arg := range args {
		if arg == "01" || arg == "galaxy" || arg == "01 galaxy" {
			fmt.Println("Alert!!!")
		}
	}
}
