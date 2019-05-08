// @Time     : 2019/5/8 14:42
// @Author   : cancan
// @File     : 27_test.go
// @Function : 移除元素 test

package QuestionBank

import "testing"

func Test_27(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		ans  int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, test := range tests {
		if removeElement(test.nums, test.val) != test.ans {
			t.Errorf("failure nums %v val %d ans %d", test.nums, test.val, test.ans)
		}
	}
}
