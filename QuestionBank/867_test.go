// @Time     : 2021/2/25 10:04
// @Author   : cancan
// @File     : 867_test.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_867(t *testing.T) {
	tests := []struct {
		matrix [][]int
		ans    [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(transpose(test.matrix), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
