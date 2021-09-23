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
			count++
			arr[count] = ch
		}
	}
	*ptr = arr
	return length
}
