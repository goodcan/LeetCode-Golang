/*
 * @Time     : 2019/10/8 11:48
 * @Author   : cancan
 * @File     : 5079_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_5079(t *testing.T) {
	tests := []struct {
		arr1 []int
		arr2 []int
		arr3 []int
		ans  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}, []int{1, 5}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(arraysIntersection(test.arr1, test.arr2, test.arr3), test.ans) {
			t.Errorf("failure arr1 %v arr2 %v arr3 %v ans %v",
				test.arr1, test.arr2, test.arr3, test.ans)
		}
	}
}
