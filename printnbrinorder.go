package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var numbers []int

	if n == 0 {
		numbers = append(numbers, n)
	}

	for n > 0 {
		numbers = append(numbers, n%10)
		n /= 10
	}

	i := 0

	for i < len(numbers)-1 {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			i = 0
		} else {
			i++
		}
	}
	for i := range numbers {
		z01.PrintRune(rune('0' + numbers[i]))
	}

}
