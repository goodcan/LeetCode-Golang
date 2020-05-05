/*
 * @Time     : 2020/5/5 15:02
 * @Author   : cancan
 * @File     : 1365_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_1365(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{
			[]int{8, 1, 2, 2, 3},
			[]int{4, 0, 1, 1, 3},
		},
		{
			[]int{6, 5, 4, 8},
			[]int{2, 1, 0, 3},
		},
		{
			[]int{7, 7, 7, 7},
			[]int{0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(smallerNumbersThanCurrent(test.nums), test.ans) {
			t.Errorf("failure %v ans %v", test.nums, test.ans)
		}
	}
}
