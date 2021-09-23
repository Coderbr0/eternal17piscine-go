package piscine

func Index(s string, toFind string) int {
	var firstIndex int

	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			firstIndex = i
			temp := s[i : i+len(toFind)]
			if temp == toFind {
				return firstIndex
			}
		}
	}

	return -1
}
