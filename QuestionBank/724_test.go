// @Time     : 2021/1/28 14:23
// @Author   : cancan
// @File     : 724_test.go
// @Function :

package QuestionBank

import "testing"

func Test_724(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
	}

	for _, test := range tests {
		if pivotIndex(test.nums) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
