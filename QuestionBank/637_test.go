/*
 * @Time     : 2020/9/12 00:32
 * @Author   : cancan
 * @File     : 637_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_637(t *testing.T) {
	tests := []struct {
		root *TreeNode
		ans  []float64
	}{
		{InitTree([]interface{}{3, 9, 20, 15, 7}), []float64{3, 14.5, 11}},
	}

	for _, test := range tests {
		if !utils.SliceEqual(averageOfLevels(test.root), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
