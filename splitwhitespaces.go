package piscine

func SplitWhiteSpaces(s string) []string {
	// accounts for an empty string
	wordCount := 0
	word := ""

	// only count words after a seperator, if next is not seperator
	for i, ch := range s {
		if ch == ' ' && s[i+1] != ' ' {
			wordCount++
		}
	}

	finalString := make([]string, wordCount+1)
	stringIndex := 0

	// adding chars to final string
	for _, ch := range s {
		if WhiteSpace(ch) {
			if word != "" {
				finalString[stringIndex] = word
				stringIndex++
				word = ""
			}
		} else {
			// final chars after last seperator
			word += string(ch)
		}
	}
	if word != "" {
		finalString[len(finalString)-1] = word
	}
	return finalString
}

//  checks if the char or rune is a seperator
func WhiteSpace(a rune) bool {
	return a == '\n' || a == ' ' || a == '\t'
}
