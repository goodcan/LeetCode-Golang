/*
 * @Time     : 2020/7/15 20:51
 * @Author   : cancan
 * @File     : 1486_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1386(t *testing.T) {
	tests := []struct {
		n     int
		start int
		ans   int
	}{
		{5, 0, 8},
		{4, 3, 8},
		{1, 7, 7},
		{10, 5, 2},
	}

	for _, test := range tests {
		if xorOperation(test.n, test.start) != test.ans {
			t.Errorf("failure n %d start %d ans %d", test.n, test.start, test.ans)
		}
	}
}
