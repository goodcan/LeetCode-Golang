// @Time     : 2019/4/25 15:44
// @Author   : cancan
// @File     : 1_test.go
// @Function : 两数之和 test

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_1(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		ans    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 7, 11, 15}, 7, []int{}},
	}

	for _, test := range tests {
		if !utils.IntSliceEqual(twoSum1(test.nums, test.target), test.ans) {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
