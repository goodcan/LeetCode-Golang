// @Time     : 2019/5/10 16:10
// @Author   : cancan
// @File     : 9_test.go
// @Function : 回文数 test

package QuestionBank

import "testing"

func Test_9(t *testing.T) {
	tests := []struct {
		x   int
		ans bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, test := range tests {
		if isPalindrome9(test.x) != test.ans {
			t.Errorf("failure x %d ans %v", test.x, test.ans)
		}
	}
}
