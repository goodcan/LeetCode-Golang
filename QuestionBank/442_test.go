// @Time     : 2019/5/5 18:02
// @Author   : cancan
// @File     : 442_test.go
// @Function : 数组中重复的数据 test

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_442(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(findDuplicates(test.nums), test.ans) {
			t.Errorf("failure nums %v ans %v", test.nums, test.ans)
		}
	}
}
