// @Time     : 2020/12/30 11:06
// @Author   : cancan
// @File     : 1046_test.go
// @Function :

package QuestionBank

import "testing"

func Test_1046(t *testing.T) {
	tests := []struct {
		stones []int
		ans    int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1, 3}, 2},
		{[]int{2, 2}, 0},
		{[]int{316, 157, 73, 106, 771, 828, 46, 212, 926, 604, 600, 992, 71, 51, 477, 869, 425, 405, 859, 924, 45, 187, 283, 590, 303, 66, 508, 982, 464, 398}, 0},
	}

	for _, test := range tests {
		if lastStoneWeight(test.stones) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
