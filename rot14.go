package piscine

func Rot14(s string) string {
	ar := []rune(s)

	for i, ch := range ar {
		if 'a' <= ch && ch <= 'l' || 'A' <= ch && ch <= 'L' {
			ar[i] += 14
		} else if 'm' <= ch && ch <= 'z' || 'M' <= ch && ch <= 'Z' {
			ar[i] -= 12
		}
	}
	return string(ar)
}
