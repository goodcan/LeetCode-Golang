// @Time     : 2021/1/20 10:48
// @Author   : cancan
// @File     : 628_test.go
// @Function :

package QuestionBank

import "testing"

func Test_628(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
	}

	for _, test := range tests {
		if maximumProduct(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
