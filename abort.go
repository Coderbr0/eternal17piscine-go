package piscine

func Abort(a, b, c, d, e int) int {
	var arr []int
	arr = append(arr, a, b, c, d, e)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			i = 0
		}
	}
	return arr[2]
}
