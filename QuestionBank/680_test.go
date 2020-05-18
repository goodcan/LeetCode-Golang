/*
 * @Time     : 2020/5/19 00:25
 * @Author   : cancan
 * @File     : 680_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_680(t *testing.T) {
	tests := []struct {
		s   string
		ans bool
	}{
		{"aba", true},
		{"abca", true},
	}

	for _, test := range tests {
		if validPalindrome(test.s) != test.ans {
			t.Errorf("failure s %v ans %v", test.s, test.ans)
		}
	}
}
