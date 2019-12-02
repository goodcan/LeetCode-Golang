/*
 * @Time     : 2019/12/2 17:58
 * @Author   : cancan
 * @File     : 658_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_658(t *testing.T) {
	tests := []struct {
		arr []int
		k   int
		x   int
		ans []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, 1, []int{1, 2, 3, 4}},
		{[]int{1, 2, 2, 2, 5, 5, 5, 8, 9, 9}, 4, 0, []int{1, 2, 2, 2}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(findClosestElements(test.arr, test.k, test.x), test.ans) {
			t.Errorf("failure arr %v k %d x %d ans %v", test.arr, test.k, test.x, test.ans)
		}
	}
}
