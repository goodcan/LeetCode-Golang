/*
 * @Time     : 2019/8/27 16:16
 * @Author   : cancan
 * @File     : 1128_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"
)

func Test_1128(t *testing.T) {
	tests := []struct {
		dominoes [][]int
		ans      int
	}{
		// [[1,2],[2,1],[3,4],[5,6]]
		{
			[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}},
			1,
		},
	}

	for _, test := range tests {
		if numEquivDominoPairs(test.dominoes) != test.ans {
			t.Errorf("failure dominoes %v ans %d", test.dominoes, test.ans)
		}
	}
}
