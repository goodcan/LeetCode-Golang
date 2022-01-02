// @Time     : 2021/5/8 10:59
// @Author   : cancan
// @File     : 17_14_test.go
// @Function :

package interviewQuestion

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_17_14(t *testing.T) {
	tests := []struct {
		arr []int
		k   int
		ans []int
	}{
		{[]int{1, 3, 5, 7, 2, 4, 6, 8}, 4, []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(smallestK(test.arr, test.k), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
