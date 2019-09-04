/*
 * @Time     : 2019/9/4 10:03
 * @Author   : cancan
 * @File     : 144_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_144(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  []int
	}{
		{
			InitTree([]interface{}{1, nil, 2, 3}),
			[]int{1, 2, 3},
		},
	}

	for _, test := range tests {
		if !utils.IntSliceEqual(preorderTraversal(test.root), test.ans) {
			t.Errorf("failure root %v ans %v", test.root, test.ans)
		}
	}
}
