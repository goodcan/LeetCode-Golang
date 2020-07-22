/*
 * @Time     : 2020/7/22 22:00
 * @Author   : cancan
 * @File     : 1313_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_1313(t *testing.T) {
	tests := []struct {
		nums []int
		ans  []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4, 4, 4}},
		{[]int{1, 1, 2, 3}, []int{1, 3, 3}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(decompressRLElist(test.nums), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
