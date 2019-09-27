/*
 * @Time     : 2019/7/9 10:04
 * @Author   : cancan
 * @File     : 283_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_283(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		//{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0}, []int{0, 0}},
	}

	for _, test := range tests {
		moveZeroes(test.nums)
		if !utils.SliceEqual(test.nums, test.ans) {
			t.Errorf("failure nums %v ans %v", test.nums, test.ans)
		}
	}
}
