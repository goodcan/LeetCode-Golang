/*
 * @Time     : 2020/9/7 23:33
 * @Author   : cancan
 * @File     : 347_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_347(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		ans  []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{4, 1, -1, 2, -1, 2, 3},
			2,
			[]int{-1, 2},
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(topKFrequent(test.nums, test.k), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
