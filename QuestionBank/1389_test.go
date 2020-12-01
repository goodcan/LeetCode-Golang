// @Time     : 2020/12/1 16:27
// @Author   : cancan
// @File     : 1389_test.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_1389(t *testing.T) {
	tests := []struct {
		nums  []int
		index []int
		ans   []int
	}{
		{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}, []int{0, 4, 1, 3, 2}},
		{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}, []int{0, 1, 2, 3, 4}},
		{[]int{1}, []int{0}, []int{1}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(createTargetArray(test.nums, test.index), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
