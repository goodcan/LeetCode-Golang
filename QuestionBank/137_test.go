// @Time     : 2019/4/29 16:39
// @Author   : cancan
// @File     : 137_test.go
// @Function : 只出现一次的数字 II test

package QuestionBank

import "testing"

func Test_137(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{2, 2, 2, 3, 3, 3}, -1},
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}

	for _, test := range tests {
		if singleNumberII(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
