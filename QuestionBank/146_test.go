/*
 * @Time     : 2020/5/25 14:14
 * @Author   : cancan
 * @File     : 146_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_146(t *testing.T) {
	tests := []struct {
		methods []string
		params  [][]interface{}
		ans     []interface{}
	}{
		{

			[]string{"LRUCache", "Put", "Put", "Get", "Put", "Get", "Put", "Get", "Get", "Get"},
			[][]interface{}{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			[]interface{}{nil, nil, nil, 1, nil, -1, nil, -1, 3, 4},
		},
	}

	s2m := utils.NewString2Method()
	var LRUCache LRUCache

	for index, test := range tests {
		for i, method := range test.methods {
			if i == 0 {
				LRUCache = Constructor146(test.params[0][0].(int))
			} else {
				if !s2m.Check(s2m.Run(&LRUCache, method, test.params[i]...), test.ans[i]) {
					t.Errorf("failure id %d", index)
				}
			}
		}
	}
}
