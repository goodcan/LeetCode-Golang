/*
 * @Time     : 2020/7/14 20:22
 * @Author   : cancan
 * @File     : 1512_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1512(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
	}

	for _, test := range tests {
		if numIdenticalPairs(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
