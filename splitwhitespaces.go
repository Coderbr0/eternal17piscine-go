package piscine

func SplitWhiteSpaces(s string) []string {
	// accounts for an empty string
	wordCount := 1
	word := ""

	for _, ch := range s {
		if WhiteSpace(ch) {
			wordCount++
		}
	}

	finalString := make([]string, wordCount)
	stringIndex := 0

	// adding chars to final string
	for i, ch := range s {
		if i+1 == len(s) {
			finalString[stringIndex] = word + string(s[i])
		}
		// if whitespace add all chars before whitespace and reset word
		if WhiteSpace(ch) && stringIndex <= wordCount {
			finalString[stringIndex] = word
			stringIndex++
			word = ""
		} else {
			// final chars after last seperator
			word += string(ch)
		}
	}
	return finalString
}

//  checks if the char or rune is a seperator
func WhiteSpace(a rune) bool {
	return a == '\n' || a == ' ' || a == '\t'
}
