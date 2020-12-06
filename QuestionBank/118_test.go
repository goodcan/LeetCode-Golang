/*
 * @Time     : 2020/12/6 16:02
 * @Author   : cancan
 * @File     : 118_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_118(t *testing.T) {
	tests := []struct {
		numRows int
		ans     [][]int
	}{
		{5, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(generate(test.numRows), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
