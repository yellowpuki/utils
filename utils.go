package utils

func Contains(str []string, x string) bool {
	for _, s := range str {
		if x == s {
			return true
		}
	}

	return false
}

func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}

	return false
}
