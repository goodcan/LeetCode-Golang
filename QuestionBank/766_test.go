// @Time     : 2021/2/22 14:22
// @Author   : cancan
// @File     : 766_test.go
// @Function :

package QuestionBank

import "testing"

func Test_766(t *testing.T) {
	tests := []struct {
		matrix [][]int
		ans    bool
	}{
		{[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true},
		{[][]int{{1, 2}, {2, 2}}, false},
	}

	for _, test := range tests {
		if isToeplitzMatrix(test.matrix) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
