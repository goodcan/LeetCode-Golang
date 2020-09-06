/*
 * @Time     : 2020/9/6 22:57
 * @Author   : cancan
 * @File     : 107_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_107(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  [][]int
	}{
		{
			InitTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			[][]int{{15, 7}, {9, 20}, {3}},
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(levelOrderBottom(test.root), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
