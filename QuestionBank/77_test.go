/*
 * @Time     : 2020/9/9 18:52
 * @Author   : cancan
 * @File     : 77_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_77(t *testing.T) {
	tests := []struct {
		n   int
		k   int
		ans [][]int
	}{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(combine(test.n, test.k), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
