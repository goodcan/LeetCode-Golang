/*
 * @Time     : 2020/9/24 10:36
 * @Author   : cancan
 * @File     : 501_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_501(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  []int
	}{
		{InitTree([]interface{}{1, nil, 2, 2}), []int{2}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(findMode(test.root), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
