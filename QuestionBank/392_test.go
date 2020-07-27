/*
 * @Time     : 2020/7/28 00:04
 * @Author   : cancan
 * @File     : 392_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_392(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		ans bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, test := range tests {
		if isSubsequence(test.s, test.t) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
