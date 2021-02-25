// @Time     : 2021/2/24 14:53
// @Author   : cancan
// @File     : 832_test.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_832(t *testing.T) {
	tests := []struct {
		A   [][]int
		ans [][]int
	}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(flipAndInvertImage(test.A), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
