/*
 * @Time     : 2021/2/15 13:31
 * @Author   : cancan
 * @File     : 485_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_485(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
	}

	for _, test := range tests {
		if findMaxConsecutiveOnes(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
