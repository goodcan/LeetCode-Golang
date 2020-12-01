// @Time     : 2020/12/1 17:12
// @Author   : cancan
// @File     : 1372_test.go
// @Function :

package QuestionBank

import "testing"

func Test_1572(t *testing.T) {
	tests := []struct {
		mat [][]int
		ans int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 25},
		{[][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}, 8},
		{[][]int{{5}}, 5},
	}

	for _, test := range tests {
		if diagonalSum(test.mat) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
