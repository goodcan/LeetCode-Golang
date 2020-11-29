/*
 * @Time     : 2020/11/29 14:17
 * @Author   : cancan
 * @File     : 976_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_976(t *testing.T) {
	tests := []struct {
		A   []int
		ans int
	}{
		{[]int{2, 1, 2}, 5},
		{[]int{1, 2, 1}, 0},
		{[]int{3, 2, 3, 4}, 10},
		{[]int{3, 6, 2, 3}, 8},
	}

	for _, test := range tests {
		if largestPerimeter(test.A) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
