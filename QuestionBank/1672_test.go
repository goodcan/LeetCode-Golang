// @Time     : 2020/12/1 15:21
// @Author   : cancan
// @File     : 1672_test.go
// @Function :

package QuestionBank

import "testing"

func Test_1672(t *testing.T) {
	tests := []struct {
		accounts [][]int
		ans      int
	}{
		{[][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{[][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
		{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17},
	}

	for _, test := range tests {
		if maximumWealth(test.accounts) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
