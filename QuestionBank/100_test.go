// @Time     : 2019/4/28 16:16
// @Author   : cancan
// @File     : 100_test.go
// @Function : 相同的树 test

package QuestionBank

import "testing"

func Test_100(t *testing.T) {
	tests := []struct {
		p   *TreeNode
		q   *TreeNode
		ans bool
	}{
		{
			InitTree([]interface{}{1, 2, 3}),
			InitTree([]interface{}{1, 2, 3}),
			true,
		},
		{
			InitTree([]interface{}{1, 2}),
			InitTree([]interface{}{1, nil, 2}),
			false,
		},
		{
			InitTree([]interface{}{1, 2, 1}),
			InitTree([]interface{}{1, 1, 2}),
			false,
		},
	}

	for _, test := range tests {
		if isSameTree(test.p, test.q) != test.ans {
			t.Errorf("failure p %v q %v ans %v",
				test.p, test.q, test.ans)
		}
	}
}
