/*
 * @Time     : 2019/9/4 10:26
 * @Author   : cancan
 * @File     : 145_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_145(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  []int
	}{
		{
			InitTree([]interface{}{1, nil, 2, 3}),
			[]int{3, 2, 1},
		},
	}

	for _, test := range tests {
		if !utils.IntSliceEqual(postorderTraversal(test.root), test.ans) {
			t.Errorf("failure root %v ans %v", test.root, test.ans)
		}
	}
}
