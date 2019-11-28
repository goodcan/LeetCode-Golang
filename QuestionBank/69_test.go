/*
 * @Time     : 2019/11/28 15:50
 * @Author   : cancan
 * @File     : 69_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_69(t *testing.T) {
	tests := []struct {
		x   int
		ans int
	}{
		{4, 2},
		{8, 2},
	}

	for _, test := range tests {
		if mySqrt(test.x) != test.ans {
			t.Errorf("failure x %d ans %d", test.x, test.ans)
		}
	}
}
