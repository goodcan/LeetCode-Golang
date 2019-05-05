// @Time     : 2019/5/5 18:11
// @Author   : cancan
// @File     : 509_test.go
// @Function : 斐波那契数 test

package QuestionBank

import "testing"

func Test_509(t *testing.T) {
	tests := []struct {
		N   int
		ans int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}

	for _, test := range tests {
		if fib(test.N) != test.ans {
			t.Errorf("failure N %d ans %d", test.N, test.ans)
		}
	}

}
