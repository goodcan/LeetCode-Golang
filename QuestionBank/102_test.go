/*
 * @Time     : 2020/5/13 23:03
 * @Author   : cancan
 * @File     : 102_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_102(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  [][]int
	}{
		{
			InitTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			[][]int{{3}, {9, 20}, {15, 7}},
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(levelOrder(test.root), test.ans) {
			t.Errorf("failure root %v ans %v", test.root, test.ans)
		}
	}
}
