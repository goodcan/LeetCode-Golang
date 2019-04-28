// @Time     : 2019/4/28 14:47
// @Author   : cancan
// @File     : 94_test.go
// @Function : 二叉树的中序遍历 test

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_94(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  []int
	}{
		{InitTree([]interface{}{1, nil, 2, 3}), []int{1, 3, 2}},
	}

	for _, test := range tests {
		if !utils.IntSliceEqual(inOrderTraversal(test.root), test.ans) {
			t.Errorf("failure root %v ans %v", test.root, test.ans)
		}
	}

}
