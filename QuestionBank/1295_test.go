/*
 * @Time     : 2020/7/18 18:50
 * @Author   : cancan
 * @File     : 1295_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1295(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}

	for _, test := range tests {
		if findNumbers(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
