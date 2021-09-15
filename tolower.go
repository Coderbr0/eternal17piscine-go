package piscine

func ToLower(s string) string {
	a := []byte(s)
	var low []byte
	for _, number := range a {
		if 65 <= number && number <= 90 {
			number = number + 32
			low = append(low, number)
		} else {
			low = append(low, number)
		}
	}
	return string(low)
}
