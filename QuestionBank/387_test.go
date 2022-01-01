/*
 * @Time     : 2022/1/1 11:05
 * @Author   : cancan
 * @File     : 387_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"
)

func Test_387(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
	}

	for _, test := range tests {
		if firstUniqChar(test.s) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
