/*
 * @Time     : 2020/9/16 11:03
 * @Author   : cancan
 * @File     : 226_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_226(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  *TreeNode
	}{
		{
			InitTree([]interface{}{4, 2, 7, 1, 3, 6, 9}),
			InitTree([]interface{}{4, 7, 2, 9, 6, 3, 1}),
		},
	}

	for _, test := range tests {
		if !TreeEqual(invertTree(test.root), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
