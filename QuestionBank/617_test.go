/*
 * @Time     : 2019/9/27 10:19
 * @Author   : cancan
 * @File     : 617_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_617(t *testing.T) {
	tests := []struct {
		t1  *TreeNode
		t2  *TreeNode
		ans *TreeNode
	}{
		{
			t1:  InitTree([]interface{}{1, 3, 2, 5}),
			t2:  InitTree([]interface{}{2, 1, 3, nil, 4, nil, 7}),
			ans: InitTree([]interface{}{3, 4, 5, 5, 4, nil, 7}),
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(Tree2Slice(mergeTrees(test.t1, test.t2)), Tree2Slice(test.ans)) {
			t.Errorf("faiulre t1 %v t2 %v ans %v", test.t1, test.t2, test.ans)
		}
	}
}
