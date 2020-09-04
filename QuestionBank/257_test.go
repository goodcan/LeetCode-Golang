/*
 * @Time     : 2020/9/4 15:13
 * @Author   : cancan
 * @File     : 257_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_257(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  []string
	}{
		{nil, []string{}},
		{InitTree([]interface{}{1, 2, 3, nil, 5}), []string{"1->2->5", "1->3"}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(binaryTreePaths(test.root), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
