/*
 * @Time     : 2020/9/10 10:09
 * @Author   : cancan
 * @File     : 40_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_40(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		ans        [][]int
	}{
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{{1, 2, 2}, {5}},
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(combinationSum2(test.candidates, test.target), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
