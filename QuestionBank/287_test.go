/*
 * @Time     : 2020/5/26 14:32
 * @Author   : cancan
 * @File     : 287——test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_287(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	}

	for _, test := range tests {
		if findDuplicate(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %v", test.nums, test.ans)
		}
	}
}
