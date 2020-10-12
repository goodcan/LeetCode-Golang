/*
 * @Time     : 2020/10/12 16:14
 * @Author   : cancan
 * @File     : 783_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_783(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  int
	}{
		{InitTree([]interface{}{1, nil, 3, 2}), 1},
	}

	for _, test := range tests {
		if minDiffInBST(test.root) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
