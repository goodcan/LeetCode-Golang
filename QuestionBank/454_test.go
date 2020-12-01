// @Time     : 2020/11/27 10:25
// @Author   : cancan
// @File     : 454_test.go
// @Function :

package QuestionBank

import "testing"

func Test_454(t *testing.T) {
	tests := []struct {
		A   []int
		B   []int
		C   []int
		D   []int
		ans int
	}{
		{
			[]int{1, 2},
			[]int{-2, -1},
			[]int{-1, 2},
			[]int{0, 2},
			2,
		},
	}

	for _, test := range tests {
		if fourSumCount(test.A, test.B, test.C, test.D) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
