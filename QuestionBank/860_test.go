// @Time     : 2020/12/10 10:51
// @Author   : cancan
// @File     : 860_test.go
// @Function :

package QuestionBank

import "testing"

func Test_860(t *testing.T) {
	tests := []struct {
		bills []int
		ans   bool
	}{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
		{[]int{5, 5, 10, 10, 20}, false},
	}

	for _, test := range tests {
		if lemonadeChange(test.bills) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
