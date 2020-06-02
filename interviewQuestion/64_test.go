/*
 * @Time     : 2020/6/2 11:47
 * @Author   : cancan
 * @File     : 64_test.go
 * @Function :
 */

package interviewQuestion

import "testing"

func Test_64(t *testing.T) {
	tests := []struct {
		n   int
		ans int
	}{
		{3, 6},
		{9, 45},
	}

	for _, test := range tests {
		if sumNums(test.n) != test.ans {
			t.Errorf("failure n %d ans %d", test.n, test.ans)
		}
	}
}
