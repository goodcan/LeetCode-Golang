/*
 * @Time     : 2019/9/4 17:00
 * @Author   : cancan
 * @File     : 1006_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1006(t *testing.T) {
	tests := []struct {
		N   int
		ans int
	}{
		{4, 7},
		{10, 12},
		{1, 1},
	}

	for _, test := range tests {
		if clumsy(test.N) != test.ans {
			t.Errorf("failure N %d ans %d", test.N, test.ans)
		}
	}
}
