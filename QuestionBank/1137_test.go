/*
 * @Time     : 2019/8/27 17:05
 * @Author   : cancan
 * @File     : 1137_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1137(t *testing.T) {
	tests := []struct {
		n   int
		ans int
	}{
		{4, 4},
		{25, 1389537},
	}

	for _, test := range tests {
		if tribonacci(test.n) != test.ans {
			t.Errorf("failure n %d ans %d", test.n, test.ans)
		}
	}
}
