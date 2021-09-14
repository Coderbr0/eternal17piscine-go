package piscine

func Index(s string, toFind string) int {
	if s == "" || toFind == "" {
		return 0
	}
	firstIndex := -1
	for i := range s {
		if s[i] == toFind[0] {
			firstIndex = i
			tempString := s[i : i+len(toFind)]

			if tempString == toFind {
				return firstIndex
			}
		}
	}
	return -1
}
