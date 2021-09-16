package piscine

func BasicAtoi(s string) int {
	var result int

	for _, num := range s {
		conversion := int(num) - 48
		result = (result * 10) + conversion
	}
	return result
}
