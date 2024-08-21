package _51

func reverseWords(s string) string {
	var ans string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		cur := i
		for cur >= 0 && s[cur] != ' ' {
			cur--
		}
		// 左闭右开
		ans += s[cur+1:i+1] + " "
		i = cur
	}

	return ans[:len(ans)-1]
}
