/*
 * @Time     : 2020/5/7 10:34
 * @Author   : cancan
 * @File     : 572_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_572(t *testing.T) {
	tests := []struct {
		s   *TreeNode
		t   *TreeNode
		ans bool
	}{
		{
			InitTree([]interface{}{3, 4, 5, 1, 2}),
			InitTree([]interface{}{4, 1, 2}),
			true,
		},
		{
			InitTree([]interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}),
			InitTree([]interface{}{4, 1, 2}),
			false,
		},
	}

	for _, test := range tests {
		if isSubtree(test.s, test.t) != test.ans {
			t.Errorf("failure s %v t %v ans %v", test.s, test.t, test.ans)
		}
	}
}
