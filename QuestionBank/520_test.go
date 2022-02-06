/*
 * @Time     : 2022/2/6 17:51
 * @Author   : cancan
 * @File     : 520_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_520(t *testing.T) {
	tests := []struct {
		word string
		ans  bool
	}{
		{"USA", true},
		{"FlaG", false},
	}

	for _, test := range tests {
		if detectCapitalUse(test.word) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
