package _8

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if com(haystack[i:], needle) {
			return i
		}
	}
	return -1
}

func com(s1, s2 string) bool {
	if len(s1) < len(s2) {
		return false
	}

	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
