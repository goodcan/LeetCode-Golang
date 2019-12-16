/*
 * @Time     : 2019/12/16 14:49
 * @Author   : cancan
 * @File     : 230_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_230(t *testing.T) {
	tests := []struct {
		root *TreeNode
		k    int
		ans  int
	}{
		{
			InitTree([]interface{}{3, 1, 4, nil, 2}),
			1,
			1,
		},
		{
			InitTree([]interface{}{5, 3, 6, 2, 4, nil, nil, 1}),
			3,
			3,
		},
	}

	for _, test := range tests {
		if kthSmallest(test.root, test.k) != test.ans {
			t.Errorf("failure root %v k %d ans %d", test.root, test.k, test.ans)
		}
	}
}
