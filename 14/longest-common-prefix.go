package _4

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	cur := strs[0]
	ans := ""
	for i := 1; i < len(strs); i++ {
		ans1 := ""
		l := min(len(cur), len(strs[i]))
		for j := 0; j < l; j++ {
			if j < len(strs[i]) && cur[j] == strs[i][j] {
				ans1 += string(cur[j])
			} else {
				break
			}
		}
		if ans1 == "" {
			return ans1
		}
		// 赋值，否则判断不执行
		if len(ans) == 0 {
			ans = ans1
		}
		// 若未赋值，这里不会触发，便不会更新ans的值，return时便会返回旧值
		if len(ans1) < len(ans) {
			ans = ans1
		}
	}
	return ans
}

func longestCommonPrefix_FenZhi(strs []string) string {
	return ""
}

func reverse(l int, r int, s []string) string {
	if l == r {
		return ""
	}
	if l == r-1 {
		getTwo(s[l], s[r])
	}
	r = (r + l) / 2
	reverse(l, r, s)
	reverse(r+1, r, s)
	return getTwo(s[l], s[r])

}

func getTwo(string1 string, string2 string) string {
	return ""
}
