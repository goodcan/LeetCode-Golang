/*
 * @Time     : 2019/11/29 18:34
 * @Author   : cancan
 * @File     : 34_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_34(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		ans    []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(searchRange(test.nums, test.target), test.ans) {
			t.Errorf("failure nums %v target %d ans %v", test.nums, test.target, test.ans)
		}
	}
}
