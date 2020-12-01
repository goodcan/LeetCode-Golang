// @Time     : 2020/12/1 16:47
// @Author   : cancan
// @File     : 1266_test.go
// @Function :

package QuestionBank

import (
	"testing"
)

func Test_1266(t *testing.T) {
	tests := []struct {
		points [][]int
		ans    int
	}{
		{[][]int{{1, 1}, {3, 4}, {-1, 0}}, 7},
		{[][]int{{3, 2}, {-2, 2}}, 5},
	}

	for _, test := range tests {
		if minTimeToVisitAllPoints(test.points) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
