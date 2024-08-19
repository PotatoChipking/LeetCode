package _3

func romanToInt(s string) int {
	// 每两个数进行计算，sum+=
	ans := 0
	i := 0
	for i < len(s) {
		l := getNum(rune(s[i]))
		if i+1 < len(s) {
			r := getNum(rune(s[i+1]))
			if l < r {
				ans += r - l
				i += 2
				continue
			}
		}
		ans += l
		i++
	}
	return ans
}

func getNum(c rune) int {
	cur := 0
	switch c {
	case 'I':
		cur = 1
	case 'V':
		cur = 5
	case 'X':
		cur = 10
	case 'L':
		cur = 50
	case 'C':
		cur = 100
	case 'D':
		cur = 500
	case 'M':
		cur = 1000
	default:
		cur = 0
	}
	return cur
}
