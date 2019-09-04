/*
 * @Time     : 2019/9/4 10:14
 * @Author   : cancan
 * @File     : 897_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_897(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  *TreeNode
	}{
		{
			InitTree([]interface{}{5, 3, 6, 2, 4, nil, 8, 1, nil, nil, nil, 7, 9}),
			InitTree([]interface{}{1, nil, 2, nil, 3, nil, 4, nil, 5, nil, 6, nil, 7, nil, 8, nil, 9}),
		},
	}

	for _, test := range tests {
		if !utils.IntSliceEqual(Tree2Slice(increasingBST(test.root)), Tree2Slice(test.ans)) {
			t.Errorf("failure root %v ans %v", test.root, test.ans)
		}
	}
}
