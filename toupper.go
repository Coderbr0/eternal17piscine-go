package piscine

func ToUpper(s string) string {
	a := []byte(s)
	var cap []byte
	for i, number := range a {
		if 97 <= number && number <= 122 {
			number = number - 32
			cap = append(cap, number)
		} else {
			cap = append(cap, number)
			i++
		}
	}
	return string(cap)
}
