/*
 * @Time     : 2021/2/16 13:01
 * @Author   : cancan
 * @File     : 561_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_561(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 4, 3, 2}, 4},
		{[]int{6, 2, 6, 5, 1, 2}, 9},
	}

	for _, test := range tests {
		if arrayPairSum(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
