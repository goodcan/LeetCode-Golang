/*
 * @Time     : 2021/2/17 11:37
 * @Author   : cancan
 * @File     : 566_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_566(t *testing.T) {
	tests := []struct {
		nums [][]int
		r    int
		c    int
		ans  [][]int
	}{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(matrixReshape(test.nums, test.r, test.c), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
