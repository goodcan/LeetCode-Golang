// @Time     : 2019/5/7 18:01
// @Author   : cancan
// @File     : 58_test.go
// @Function : 最后一个单词的长度 test

package QuestionBank

import "testing"

func Test_58(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"H ", 1},
		{"Hello World", 5},
	}

	for _, test := range tests {
		if lengthOfLastWord(test.s) != test.ans {
			t.Errorf("failure s %s ans %d", test.s, test.ans)
		}
	}
}
