/*
 * @Time     : 2020/7/15 21:38
 * @Author   : cancan
 * @File     : 16_01_test.go
 * @Function :
 */

package interviewQuestion

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_16_01(t *testing.T) {
	tests := []struct {
		numbers []int
		ans     []int
	}{
		{[]int{1, 2}, []int{2, 1}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(swapNumbers(test.numbers), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
