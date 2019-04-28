// @Time     : 2019/4/28 15:50
// @Author   : cancan
// @File     : 98_test.go
// @Function : 验证二叉搜索树 test

package QuestionBank

import "testing"

func Test_98(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  bool
	}{
		{InitTree([]interface{}{1, 1}), false},
		{InitTree([]interface{}{2, 1, 3}), true},
		{InitTree([]interface{}{5, 1, 4, nil, nil, 3, 6}), false},
	}

	for _, test := range tests {
		if isValidBST(test.root) != test.ans {
			t.Errorf("failure root %v ans %v", test.root, test.ans)
		}
	}
}
