/*
 * @Time     : 2020/9/27 20:31
 * @Author   : cancan
 * @File     : 235_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_235(t *testing.T) {
	tests := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		ans  *TreeNode
	}{
		{
			InitTree([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			InitTree([]interface{}{2}),
			InitTree([]interface{}{8}),
			InitTree([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
		},
		{
			InitTree([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			InitTree([]interface{}{2}),
			InitTree([]interface{}{4}),
			InitTree([]interface{}{2, 0, 4, nil, nil, 3, 5}),
		},
	}

	for _, test := range tests {
		if !TreeEqual(lowestCommonAncestor235(test.root, test.p, test.q), test.ans) {
			t.Errorf("failure root %v p %v q %v ans %v", test.root, test.p, test.q, test.ans)
		}
	}
}
