/*
 * @Time     : 2020/7/17 21:21
 * @Author   : cancan
 * @File     : 796_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_796(t *testing.T) {
	tests := []struct {
		A   string
		B   string
		ans bool
	}{
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},
	}

	for _, test := range tests {
		if rotateString(test.A, test.B) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
