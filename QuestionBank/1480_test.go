/*
 * @Time     : 2020/7/14 21:05
 * @Author   : cancan
 * @File     : 1480_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_1480(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(runningSum(test.nums), test.ans) {
			t.Errorf("failure nums %v ans %v", test.nums, test.ans)
		}
	}
}
