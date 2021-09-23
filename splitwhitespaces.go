package piscine

func SplitWhiteSpaces(s string) []string {
	sr := []rune(s)

	var strA []string

	word := ""

	for i := range sr {

		if sr[i] == '\n' || sr[i] == '\t' || sr[i] == ' ' {
			if word != "" {
				strA = append(strA, word)
				word = ""
			}
		} else {
			word = word + string(sr[i])
		}
		if i == len(sr)-1 {
			strA = append(strA, word)
		}
	}
	return strA
}
