/*
 * @Time     : 2020/7/15 21:21
 * @Author   : cancan
 * @File     : 1431_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_1431(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		ans          []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(kidsWithCandies(test.candies, test.extraCandies), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
