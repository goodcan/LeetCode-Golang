// @Time     : 2021/5/8 10:40
// @Author   : cancan
// @File     : 17_10_test.go
// @Function :

package interviewQuestion

import "testing"

func Test_17_10(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 2, 5, 9, 5, 9, 5, 5, 5}, 5},
		{[]int{3, 2}, -1},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, test := range tests {
		if majorityElement(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
