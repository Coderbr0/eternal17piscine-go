package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for i := 3; i <= nb; i++ {
		if nb%i != 0 {
			return nb
		}
	}
	return 0
}
