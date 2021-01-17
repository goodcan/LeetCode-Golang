/*
 * @Time     : 2021/1/17 20:31
 * @Author   : cancan
 * @File     : 1232_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1232(t *testing.T) {
	tests := []struct {
		coordinates [][]int
		ans         bool
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}, true},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, false},
		{[][]int{{-1, 1}, {0, 0}, {1, -1}}, true},
	}

	for _, test := range tests {
		if checkStraightLine(test.coordinates) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
