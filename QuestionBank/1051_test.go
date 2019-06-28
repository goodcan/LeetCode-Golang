/*
 * @Time     : 2019/6/28 11:24
 * @Author   : cancan
 * @File     : 1051_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1051(t *testing.T) {
	tests := []struct {
		heights []int
		ans     int
	}{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
	}

	for _, test := range tests {
		if heightChecker(test.heights) != test.ans {
			t.Errorf("failure heights %v ans %d", test.heights, test.ans)
		}
	}
}
