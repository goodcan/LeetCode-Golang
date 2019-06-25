// @Time     : 2019/6/25 15:44
// @Author   : cancan
// @File     : 739_test.go
// @Function :

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_T39(t *testing.T) {
	tests := []struct {
		T   []int
		ans []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}

	for _, test := range tests {
		if !utils.IntSliceEqual(dailyTemperatures(test.T), test.ans) {
			t.Errorf("failure T %v ans %v \n", test.T, test.ans)
		}
	}

}
