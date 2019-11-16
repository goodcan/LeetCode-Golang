/*
 * @Time     : 2019-07-30 23:43
 * @Author   : cancan
 * @File     : 268_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_268(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, test := range tests {
		if missingNumber(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
