package piscine

func Compact(ptr *[]string) int {
	length := 0

	for _, ch := range *ptr {
		if ch != "" {
			length++
		}
	}
	count := 0
	arr := make([]string, length)

	for _, ch := range *ptr {
		if ch != "" {
			arr[count] = ch
			count++
		}
	}
	*ptr = arr
	return length
}
