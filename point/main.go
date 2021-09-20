// package piscine

// import "fmt"

// func setPoint(x, y *int) {
// 	*x = 42
// 	*y = 21
// }

// func main() {
// 	points := &point{}

// 	setPoint(points)

// 	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
// }

package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
