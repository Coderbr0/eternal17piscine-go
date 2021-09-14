package piscine

func Index(s string, toFind string) int {
	if s == "" || toFind == "" {
		return 0
	}
	for i := range s {
		if s[i] == toFind[0] {
			tempString := s[i : i+len(toFind)]

			if tempString == toFind {
				return i
			}
		}
	}
	return -1
}
