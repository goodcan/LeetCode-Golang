/*
 * @Time     : 2020/7/14 21:25
 * @Author   : cancan
 * @File     : 1470_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_1470(t *testing.T) {
	tests := []struct {
		nums []int
		n    int
		ans  []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(shuffle(test.nums, test.n), test.ans) {
			t.Errorf("failure nums %v n %d ans %v", test.nums, test.n, test.ans)
		}
	}
}
