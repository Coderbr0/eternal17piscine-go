package piscine

func Map(f func(int) bool, a []int) []bool {
	// declaring an empty slice of type bool
	var arrOfBool []bool

	// loop through given array and add a function of each a into our array of booleans
	for i := range a {
		arrOfBool = append(arrOfBool, f(a[i]))
	}
	return arrOfBool
}
