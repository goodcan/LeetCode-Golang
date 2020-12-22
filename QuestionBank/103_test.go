// @Time     : 2020/12/22 10:18
// @Author   : cancan
// @File     : 103_test.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_103(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  [][]int
	}{
		{InitTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			[][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}

	for idx, test := range tests {
		if !utils.SliceEqual(zigzagLevelOrder(test.root), test.ans) {
			t.Errorf("failure idx %d %+v", idx, test)
		}
	}
}
