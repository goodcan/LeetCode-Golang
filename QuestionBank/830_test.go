// @Time     : 2021/1/5 9:51
// @Author   : cancan
// @File     : 830_test.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_830(t *testing.T) {
	tests := []struct {
		s   string
		ans [][]int
	}{
		{"abbxxxxzzy", [][]int{{3, 6}}},
		{"abc", [][]int{}},
		{"abcdddeeeeaabbbcd", [][]int{{3, 5}, {6, 9}, {12, 14}}},
		{"aba", [][]int{}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(largeGroupPositions(test.s), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
