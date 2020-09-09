/*
 * @Time     : 2020/9/9 14:11
 * @Author   : cancan
 * @File     : 39_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_39(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		ans        [][]int
	}{
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{{2, 2, 3}, {7}},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(combinationSum(test.candidates, test.target), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
