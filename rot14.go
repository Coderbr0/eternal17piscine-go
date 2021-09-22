package piscine

func Rot14(s string) string {
	sRune := []rune(s)
	var changed []rune

	for i, char := range sRune {
		if char >= 'A' && char <= 'L' || char >= 'a' && char <= 'l' {
			char = char + 14
		} else if char >= 'M' && char <= 'Z' || char >= 'm' && char <= 'z' {
			char = char - 12
		} else {
			i++
		}
		changed = append(changed, char)
	}
	return string(changed)
}
