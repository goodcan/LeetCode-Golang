// @Time     : 2019/4/28 11:12
// @Author   : cancan
// @File     : 75_test.go
// @Function : 颜色分类 test

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_75(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, test := range tests {
		sortColors(test.nums)
		if !utils.IntSliceEqual(test.nums, test.ans) {
			t.Errorf("failure nums %v ans %v", test.nums, test.ans)
		}
	}
}
