package piscine

func StrRev(s string) string {
	// converting to runes
	runes := []rune(s)

	// iterating from 0 forwards and last character of string backwards
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// swapping around the characters
		runes[i], runes[j] = runes[j], runes[i]
	}
	// converting the runes back to string as our function needs to return string
	return string(runes)
}
