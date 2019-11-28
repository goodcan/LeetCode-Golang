/*
 * @Time     : 2019/11/28 16:09
 * @Author   : cancan
 * @File     : 33_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_33(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		ans    int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	}

	for _, test := range tests {
		if search33(test.nums, test.target) != test.ans {
			t.Errorf("failure nums %v target %d ans %d", test.nums, test.target, test.ans)
		}
	}
}
