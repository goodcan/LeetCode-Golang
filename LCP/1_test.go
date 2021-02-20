/*
 * @Time     : 2019/9/26 14:43
 * @Author   : cancan
 * @File     : LCP1_test.go
 * @Function :
 */

package LCP

import "testing"

func Test_LCP1(t *testing.T) {
	tests := []struct {
		guess  []int
		answer []int
		ans    int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, 3},
		{[]int{2, 2, 3}, []int{3, 2, 1}, 1},
	}

	for _, test := range tests {
		if game(test.guess, test.answer) != test.ans {
			t.Errorf("failure guess %v answer %v ans %d",
				test.guess, test.answer, test.ans)
		}
	}
}
