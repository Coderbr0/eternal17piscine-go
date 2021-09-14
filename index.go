package piscine

func Index(s string, toFind string) int {
	sRune := []rune(s)
	tfRune := []rune(toFind)
	firstIndex := -1
	for i := 0; i < len(sRune); i++ {
		if sRune[i] == tfRune[0] {
			firstIndex = i
			tempString := sRune[i : i+len(tfRune)]

			if string(tempString) == string(toFind) {
				return firstIndex
			}
		}
	}
	return -1
}
