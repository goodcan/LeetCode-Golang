/*
 * @Time     : 2020/5/10 11:30
 * @Author   : cancan
 * @File     : 236_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_236(t *testing.T) {
	tests := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		ans  *TreeNode
	}{
		{
			InitTree([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			InitTree([]interface{}{5}),
			InitTree([]interface{}{1}),
			InitTree([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
		},
		{
			InitTree([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			InitTree([]interface{}{5}),
			InitTree([]interface{}{4}),
			InitTree([]interface{}{5, 6, 2, nil, nil, 7, 4}),
		},
	}

	for _, test := range tests {
		if !TreeEqual(lowestCommonAncestor(test.root, test.p, test.q), test.ans) {
			t.Errorf("failure root %v p %v q %v ans %v", test.root, test.p, test.q, test.ans)
		}
	}
}
