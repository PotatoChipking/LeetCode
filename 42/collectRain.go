package _2

func trap(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0

	leftMax := 0
	rightMax := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] > leftMax {
				leftMax = height[l]
			} else {
				ans += leftMax - height[l]
			}

			l++
		} else {
			if height[r] > rightMax {
				rightMax = height[r]
			} else {
				ans += rightMax - height[r]
			}
			r--
		}
	}
	return ans
}
