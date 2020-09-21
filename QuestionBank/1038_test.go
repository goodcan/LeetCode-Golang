/*
 * @Time     : 2020/9/21 16:04
 * @Author   : cancan
 * @File     : 1038_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1038(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  *TreeNode
	}{
		{
			InitTree([]interface{}{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}),
			InitTree([]interface{}{30, 36, 21, 36, 35, 26, 15, nil, nil, nil, 33, nil, nil, nil, 8}),
		},
	}

	for _, test := range tests {
		if !TreeEqual(bstToGst(test.root), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
