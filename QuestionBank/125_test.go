/*
 * @Time     : 2019/8/6 14:17
 * @Author   : cancan
 * @File     : 125_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_125(t *testing.T) {
	tests := []struct {
		s   string
		ans bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, test := range tests {
		if isPalindrome(test.s) != test.ans {
			t.Errorf("failure s %s ans %v", test.s, test.ans)
		}
	}
}
