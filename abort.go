package piscine

func Abort(a, b, c, d, e int) int {
	var arr []int
	arr = append(arr, a, b, c, d, e)
	i := 0
	for i < len(arr)-1 {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			i = 0
		} else {
			i++
		}
	}
	return arr[2]
}
