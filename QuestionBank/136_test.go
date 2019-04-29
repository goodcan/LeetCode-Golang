// @Time     : 2019/4/29 16:27
// @Author   : cancan
// @File     : 136_test.go
// @Function : 只出现一次的数字 test

package QuestionBank

import "testing"

func Test_136(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, test := range tests {
		if singleNumber(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
