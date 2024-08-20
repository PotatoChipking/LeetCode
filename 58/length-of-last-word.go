package _8

func lengthOfLastWord(s string) int {
	if s == " " {
		return 0
	}

	ans := 0
	var l = false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && l == true {
			break
		}
		if s[i] != ' ' {
			ans++
		}
		if s[i] != ' ' && !l {
			l = true
		}

	}
	return ans

}
