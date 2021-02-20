// @Time     : 2021/2/20 12:00
// @Author   : cancan
// @File     : 6_test.go
// @Function :

package LCP

import "testing"

func Test_6(t *testing.T) {
	tests := []struct {
		coins []int
		ans   int
	}{
		{[]int{4, 2, 1}, 4},
		{[]int{2, 3, 10}, 8},
	}

	for _, test := range tests {
		if minCount(test.coins) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
