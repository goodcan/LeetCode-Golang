/*
 * @Time     : 2019/11/28 17:11
 * @Author   : cancan
 * @File     : 162_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_162(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
	}

	for _, test := range tests {
		if findPeakElement(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
