/*
 * @Time     : 2020/9/21 15:55
 * @Author   : cancan
 * @File     : 538_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_538(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  *TreeNode
	}{
		{
			InitTree([]interface{}{5, 2, 13}),
			InitTree([]interface{}{18, 20, 13}),
		},
	}

	for _, test := range tests {
		if !TreeEqual(convertBST(test.root), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
