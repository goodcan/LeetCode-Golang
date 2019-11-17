/*
 * @Time     : 2019-11-17 22:30
 * @Author   : cancan
 * @File     : 260_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_260(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{
			[]int{1, 2, 1, 3, 2, 5},
			[]int{3, 5},
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(singleNumberIII(test.nums), test.ans) {
			t.Errorf("failure nums %v ans %v", test.nums, test.ans)
		}
	}
}
