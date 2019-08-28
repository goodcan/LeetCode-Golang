/*
 * @Time     : 2019/8/28 10:57
 * @Author   : cancan
 * @File     : 1161_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1161(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  int
	}{
		{
			InitTree([]interface{}{1, 7, 0, 7, -8, nil, nil}),
			2,
		},
	}

	for _, test := range tests {
		if maxLevelSum(test.root) != test.ans {
			t.Errorf("failure root %+v ans %d", test.root, test.ans)
		}
	}
}
