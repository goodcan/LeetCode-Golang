// @Time     : 2020/12/18 9:54
// @Author   : cancan
// @File     : 389_test.go
// @Function :

package QuestionBank

import "testing"

func Test_389(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		ans string
	}{
		{"abcd", "abcde", "e"},
		{"", "y", "y"},
		{"a", "aa", "a"},
		{"ae", "aea", "a"},
	}

	for _, test := range tests {
		if string(findTheDifference(test.s, test.t)) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
