package _2

func intToRoman(num int) string {
	var ans string
	a := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	b := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	for i := len(a) - 1; i >= 0; i-- {
		for num >= a[i] {
			ans += b[i]
			num -= a[i]
		}
	}
	return ans
}

//m := map[int]string{
//	1:  "I",
//	4:  "IV",
//	5:  "V",
//	9:  "IX",
//	10: "X",
//	40:	 "XL",
//	50: "L",
//	90: "XC",
//	100: "C",
//	400: "CD",
//	500: "D",
//	900: "CM",
//	1000: "M",
//}
