package _2

func intToRoman(num int) string {
	var ans string
	for num > 0 {
		cur := num % 10
		num /= 10

	}

	return ans
}

func getRoman(num int) string {
	cur := num % 10
	switch cur {
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
}
