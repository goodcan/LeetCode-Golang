// @Time     : 2019/4/26 10:53
// @Author   : cancan
// @File     : 3_test.go
// @Function : 无重复字符的最长子串 test

package QuestionBank

import "testing"

func Test_3(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"au", 2},
		{"dvdf", 3},
	}

	for _, test := range tests {
		if lengthOfLongestSubstring(test.s) != test.ans {
			t.Errorf("failure s %s ans %d", test.s, test.ans)
		}
	}
}
