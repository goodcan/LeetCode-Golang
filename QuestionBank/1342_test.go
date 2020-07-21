/*
 * @Time     : 2020/7/21 21:46
 * @Author   : cancan
 * @File     : 1342_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1342(t *testing.T) {
	tests := []struct {
		num int
		ans int
	}{
		{14, 6},
		{8, 4},
		{123, 12},
	}

	for _, test := range tests {
		if numberOfSteps(test.num) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
