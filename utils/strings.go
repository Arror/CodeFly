package utils

// Filter high order function for string[]
func Filter(strs []string, where func(string) bool) []string {

	if strs == nil || len(strs) == 0 {
		return strs
	}

	new := []string{}

	for _, str := range strs {

		if !where(str) {
			new = append(new, str)
		}
	}

	return new
}
