/*
 * @Time     : 2020/10/12 15:59
 * @Author   : cancan
 * @File     : 530_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_530(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  int
	}{
		{InitTree([]interface{}{1, nil, 3, 2}), 1},
	}

	for _, test := range tests {
		if getMinimumDifference(test.root) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
