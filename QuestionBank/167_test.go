/*
 * @Time     : 2019/8/20 15:27
 * @Author   : cancan
 * @File     : 167_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_167(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		ans     []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
	}

	for _, test := range tests {
		if !utils.IntSliceEqual(twoSum167(test.numbers, test.target), test.ans) {
			t.Errorf("failure numbers %v target %d ans %v",
				test.numbers, test.target, test.ans)
		}
	}
}
