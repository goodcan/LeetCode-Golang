/*
 * @Time     : 2020/9/19 14:15
 * @Author   : cancan
 * @File     : 404_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_404(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  int
	}{
		{InitTree([]interface{}{3, 9, 10, nil, nil, 15, 7}), 24},
		{InitTree([]interface{}{1, 2, 3, 4, 5}), 4},
	}

	for _, test := range tests {
		if sumOfLeftLeaves(test.root) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
