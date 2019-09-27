/*
 * @Time     : 2019/6/27 10:34
 * @Author   : cancan
 * @File     : 344_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_344(t *testing.T) {
	tests := []struct {
		s   []byte
		ans []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, test := range tests {
		reverseString(test.s)
		if !utils.SliceEqual(test.s, test.ans) {
			t.Errorf("failure s %v ans %v", test.s, test.ans)
		}
	}
}
