package piscine

func Capitalize(s string) string {
	// convert string to rune
	a := []rune(s)

	// at first char
	c1 := true

	// check if character is alpha numerical & first char of word
	// if so and lowercase, make capital
	for i, ch := range a {
		if AlphaN(a[i]) == true && c1 {
			if 'a' <= ch && ch <= 'z' {
				a[i] -= 32
			}
			// for other chars, ie not c1, if capital, lowercase them,
			c1 = false
		} else if 'A' <= ch && ch <= 'Z' {
			a[i] += 32
			// if char is not alphanumerical, keep iterating until c1 becomes true
		} else if AlphaN(a[i]) == false {
			c1 = true
		}
	}
	return string(a)
}

// checks if rune is an alphanumerical
func AlphaN(c rune) bool {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		return true
	}
	return false
}
