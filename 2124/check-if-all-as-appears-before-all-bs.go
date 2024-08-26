package _124

func checkString(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			for j := i + 1; j < len(s); j++ {
				if s[j] == 'a' {
					return false
				}
			}
		}
	}
	return true
}
