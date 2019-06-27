/*
 * @Time     : 2019/6/27 9:40
 * @Author   : cancan
 * @File     : 232_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_232(t *testing.T) {
	tests := []struct {
		Id      int
		methods []string
		params  [][]interface{}
		ans     []interface{}
	}{
		{
			0,
			[]string{"MyQueue", "Push", "Push", "Peek", "Pop", "Empty"},
			[][]interface{}{{}, {1}, {2}, {}, {}, {}},
			[]interface{}{nil, nil, nil, 1, 1, false},
		},
	}

	s2m := utils.NewString2Method()
	var MyQueue MyQueue

	for _, test := range tests {
		for i, method := range test.methods {
			if i == 0 {
				MyQueue = Constructor232()
			} else {
				if !s2m.Check(s2m.Run(&MyQueue, method, test.params[i]...), test.ans[i]) {
					t.Errorf("failure id %d", test.Id)
				}
			}
		}
	}
}
