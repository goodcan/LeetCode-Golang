// @Time     : 2020/11/26 11:29
// @Author   : cancan
// @File     : 164_test.go
// @Function :

package QuestionBank

import "testing"

func Test_164(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{3, 6, 9, 1}, 3},
		{[]int{10}, 0},
	}

	for _, test := range tests {
		if maximumGap(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
