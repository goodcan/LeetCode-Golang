/*
 * @Time     : 2019/11/19 16:18
 * @Author   : cancan
 * @File     : 349_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_349(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		ans   []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}

	for _, test := range tests {
		if utils.SliceEqual(intersection(test.nums1, test.nums2), test.ans) {
			t.Errorf("failure nums1 %v nums2 %v ans %v", test.nums1, test.nums2, test.ans)
		}
	}
}
