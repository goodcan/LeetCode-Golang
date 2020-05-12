/*
 * @Time     : 2020/5/12 9:51
 * @Author   : cancan
 * @File     : 155_test.go
 * @Function : 最小栈
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_155(t *testing.T) {
	tests := []struct {
		methods []string
		params  [][]interface{}
		ans     []interface{}
	}{
		{
			[]string{"MinStack", "Push", "Push", "Push", "GetMin", "Pop", "Top", "GetMin"},
			[][]interface{}{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			[]interface{}{nil, nil, nil, nil, -3, nil, 0, -2},
		},
		{
			[]string{"MinStack", "Push", "Push", "Push", "GetMin", "Top", "Pop", "GetMin"},
			[][]interface{}{{}, {-2}, {0}, {-1}, {}, {}, {}, {}},
			[]interface{}{nil, nil, nil, nil, -2, -1, nil, -2},
		},
	}

	s2m := utils.NewString2Method()
	var MinStack MinStack

	for index, test := range tests {
		for i, method := range test.methods {
			if i == 0 {
				MinStack = Constructor155()
			} else {
				if !s2m.Check(s2m.Run(&MinStack, method, test.params[i]...), test.ans[i]) {
					t.Errorf("failure id %d", index)
				}
			}
		}
	}
}
