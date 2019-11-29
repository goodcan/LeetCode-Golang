/*
 * @Time     : 2019/11/29 18:24
 * @Author   : cancan
 * @File     : 153_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_153(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	}

	for _, test := range tests {
		if findMin(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
