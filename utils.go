package utils

func Contains(str []string, x string) bool {
	for _, s := range str {
		if x == s {
			return true
		}
	}

	return false
}
