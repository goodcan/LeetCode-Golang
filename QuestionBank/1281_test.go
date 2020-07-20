/*
 * @Time     : 2020/7/20 21:36
 * @Author   : cancan
 * @File     : 1281_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1281(t *testing.T) {
	tests := []struct {
		n   int
		ans int
	}{
		{234, 15},
		{4421, 21},
	}

	for _, test := range tests {
		if subtractProductAndSum(test.n) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
