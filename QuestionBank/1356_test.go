// @Time     : 2020/11/6 16:20
// @Author   : cancan
// @File     : 1356_test.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_1356(t *testing.T) {
	tests := []struct {
		arr []int
		ans []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
		{[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}},
		{[]int{10000, 10000}, []int{10000, 10000}},
		{[]int{2, 3, 5, 7, 11, 13, 17, 19}, []int{2, 3, 5, 17, 7, 11, 13, 19}},
		{[]int{10, 100, 1000, 10000}, []int{10, 100, 10000, 1000}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(sortByBits(test.arr), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
