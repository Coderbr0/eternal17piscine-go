package piscine

func ConcatParams(args []string) string {
	// needed string that will be added to
	var needed string

	// loop over the words in the slice of strings and add each word to the empty string
	// add a line break until last index, so there is no extra line break on end
	for i, word := range args {
		needed += word
		if i != len(args)-1 {
			needed += "\n"
		}

	}
	return needed
}
