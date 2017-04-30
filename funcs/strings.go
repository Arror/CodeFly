package funcs

// Filter high order function for string[]
func Filter(strs []string, condition func(string) bool) []string {

	if strs == nil || len(strs) == 0 {
		return strs
	}

	new := []string{}

	for _, str := range strs {

		if !condition(str) {
			new = append(new, str)
		}
	}

	return new
}
