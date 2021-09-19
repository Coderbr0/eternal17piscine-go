package piscine

func SplitWhiteSpaces(s string) []string {
	wordCount := 0

	// this word will reset to empty after encountering a seperator
	word := ""

	for i, ch := range s {
		if ch == ' ' && s[i+1] != ' ' {
			wordCount++
		}
	}

	finalString := make([]string, wordCount+1)
	stringIndex := 0

	// adding words to final string after every seperator, and resetting word to empty.
	for _, ch := range s {
		if WhiteSpace(ch) {
			if word != "" {
				finalString[stringIndex] = word
				stringIndex++
				word = ""
			}
			// add the rest of the string to word
		} else {
			word += string(ch)
		}
	}
	// for last word as there may be no seperators after word
	x := len(finalString)
	if word != "" {
		finalString[x-1] = word
	}
	return finalString
}

//  checks if the char or rune is a seperator
func WhiteSpace(a rune) bool {
	return a == '\n' || a == ' ' || a == '\t'
}
