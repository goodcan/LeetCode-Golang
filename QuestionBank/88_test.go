/*
 * @Time     : 2019/7/9 18:19
 * @Author   : cancan
 * @File     : 88_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_88(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		ans   []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		if !utils.SliceEqual(test.nums1, test.ans) {
			t.Errorf("failure nums1 %v m %d nums2 %v n %d ans %v",
				test.nums1, test.m, test.nums2, test.n, test.ans)
		}
	}
}
