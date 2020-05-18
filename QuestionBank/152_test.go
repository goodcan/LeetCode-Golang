/*
 * @Time     : 2020/5/18 11:07
 * @Author   : cancan
 * @File     : 152_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_152(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	}

	for _, test := range tests {
		if maxProduct(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %v", test.nums, test.ans)
		}
	}
}
