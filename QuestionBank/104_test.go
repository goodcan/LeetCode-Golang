/*
 * @Time     : 2020/7/28 00:15
 * @Author   : cancan
 * @File     : 104_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_104(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  int
	}{
		{InitTree([]interface{}{3, 9, 20, nil, nil, 15, 7}), 3},
	}

	for _, test := range tests {
		if maxDepth104(test.root) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
