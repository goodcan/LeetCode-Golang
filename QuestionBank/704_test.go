/*
 * @Time     : 2019/11/28 15:29
 * @Author   : cancan
 * @File     : 704_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_704(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		ans    int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for _, test := range tests {
		if search(test.nums, test.target) != test.ans {
			t.Errorf("failure nums %v target %d ans %d", test.nums, test.target, test.ans)
		}
	}
}
