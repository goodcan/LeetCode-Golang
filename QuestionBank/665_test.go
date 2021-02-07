// @Time     : 2021/2/7 10:41
// @Author   : cancan
// @File     : 665_test.go
// @Function :

package QuestionBank

import "testing"

func Test_665(t *testing.T) {
	tests := []struct {
		nums []int
		ans  bool
	}{
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},
		{[]int{3, 4, 2, 3}, false},
	}

	for _, test := range tests {
		if checkPossibility(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
