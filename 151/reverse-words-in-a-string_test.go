package _51

import (
	"testing"
)

// TestReverseWords 测试 reverseWords 函数
func TestReverseWords(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
	}

	for _, test := range tests {
		actual := reverseWords(test.input)
		if actual != test.expect {
			t.Errorf("reverseWords(%q) = %q; expect %q", test.input, actual, test.expect)
		}
	}
}
