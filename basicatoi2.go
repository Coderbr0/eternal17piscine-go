package piscine

func BasicAtoi2(s string) int {
	var result int
	// if int of num is between the ascii values of num, covert, else return 0
	// if string contains anything else
	for _, num := range s {
		if int(num) >= 48 && int(num) <= 57 {
			conversion := int(num) - 48
			result = (result * 10) + conversion
		} else {
			return 0
		}
	}
	return result
}
