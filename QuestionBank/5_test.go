/*
 * @Time     : 2020/5/21 11:04
 * @Author   : cancan
 * @File     : 5_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_5(t *testing.T) {
	tests := []struct {
		s   string
		ans string
	}{
		//{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, test := range tests {
		if longestPalindrome(test.s) != test.ans {
			t.Errorf("failure s %v ans %v", test.s, test.ans)
		}
	}
}
