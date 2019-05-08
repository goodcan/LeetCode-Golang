// @Time     : 2019/5/8 11:20
// @Author   : cancan
// @File     : 26_test.go
// @Function : 删除排序数组中的重复项 test

package QuestionBank

import "testing"

func Test_26(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, test := range tests {
		if removeDuplicates(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
