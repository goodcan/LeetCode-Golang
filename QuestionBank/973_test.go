// @Time     : 2020/11/9 11:41
// @Author   : cancan
// @File     : 973_test.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_973(t *testing.T) {
	tests := []struct {
		points [][]int
		K      int
		ans    [][]int
	}{
		{[][]int{{1, 2}, {1, 1}}, 1, [][]int{{1, 1}}},
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(kClosest(test.points, test.K), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
