// @Time     : 2021/2/4 11:24
// @Author   : cancan
// @File     : 643_test.go
// @Function :

package QuestionBank

import (
	"math"
	"testing"
)

func Test_643(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		ans  float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{7, 4, 5, 8, 8, 3, 9, 8, 7, 6}, 7, 7},
	}

	m := 0.000001
	for _, test := range tests {
		if math.Dim(findMaxAverage(test.nums, test.k), test.ans) > m {
			t.Errorf("failure %+v", test)
		}
	}
}
