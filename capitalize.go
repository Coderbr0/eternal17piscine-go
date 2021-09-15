package piscine

func Capitalize(s string) string {
	sRune := []rune(s)

	//used for first char of the word in a string
	c1 := true

	//checking if rune and first char of word, to capitalize
	for i := range sRune {
		if IsRune(sRune[i]) == true && c1 {
			if 'a' <= sRune[i] && sRune[i] <= 'z' {
				sRune[i] -= 32
			}
			// rest of the characters in the each word in the string, if capital, make lowercase.
			c1 = false
		} else if 'A' <= sRune[i] && sRune[i] <= 'Z' {
			sRune[i] += 32
			// finally if not rune i.e if special character, go until first letter
		} else if IsRune(sRune[i]) == false {
			c1 = true
		}
	}
	return string(sRune)
}

//function to check if character is a rune, returns boolean
func IsRune(c rune) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
