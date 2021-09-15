package piscine

func ToLowerCase(s string) string {
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

func Capitalize(s string) string {
	lower := ToLowerCase(s)
	lb := []byte(lower)
	var capString []byte
	for _, c := range lb {
		// if first char is non alpha, change next alpha to upper
		if 97 <= c && c <= 122 {
			c = c - 32
			capString = append(capString, c)
		} else {
			capString = append(capString, c)
		}
	}
	return string(capString)
}

// if (20 <= n && n <= 64) || (91 <= n && n <= 96) || (123 <= n && n <= 127)
