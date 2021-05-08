// @Time     : 2021/5/8 10:22
// @Author   : cancan
// @File     : 17_04_test.go
// @Function :

package interviewQuestion

import "testing"

func Test_17_04(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{0}, 1},
		{[]int{0, 2}, 1},
		{[]int{2, 0, 1}, 3},
		{[]int{3, 0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, test := range tests {
		if missingNumber(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
