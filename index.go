package piscine

func Index(s string, toFind string) int {
	if s == "" || toFind == "" {
		return 0
	}
	firstIndex := -1
	for i := range s {
		if []rune(s)[i] == []rune(toFind)[0] {
			firstIndex = i
			tempString := s[i : i+len([]rune(toFind))]

			if tempString == toFind {
				return firstIndex
			}
		}
	}
	return -1
}
