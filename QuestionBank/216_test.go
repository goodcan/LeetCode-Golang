/*
 * @Time     : 2020/9/11 1:02
 * @Author   : cancan
 * @File     : 216_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_216(t *testing.T) {
	tests := []struct {
		k   int
		n   int
		ans [][]int
	}{
		{3, 7, [][]int{{1, 2, 4}}},
		{3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(combinationSum3(test.k, test.n), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
