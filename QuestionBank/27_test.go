// @Time     : 2019/5/8 14:42
// @Author   : cancan
// @File     : 27_test.go
// @Function : 移除元素 test

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_27(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		ansNum   int
		ansSlice []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
	}

	for _, test := range tests {
		if removeElement(test.nums, test.val) != test.ansNum || utils.IntSliceEqual(test.nums, test.ansSlice) {
			t.Errorf("failure nums %v val %d ansNum %d ansSlice %v",
				test.nums, test.val, test.ansNum, test.ansSlice,
			)
		}
	}
}
