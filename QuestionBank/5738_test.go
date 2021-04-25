// @Time     : 2021/4/25 17:07
// @Author   : cancan
// @File     : 5738_test.go
// @Function :

package QuestionBank

import "testing"

func Test_5738(t *testing.T) {
	tests := []struct {
		n   int
		k   int
		ans int
	}{
		{34, 6, 9},
		{10, 10, 1},
	}

	for _, test := range tests {
		if sumBase(test.n, test.k) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
