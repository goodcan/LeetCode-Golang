/*
 * @Time     : 2019/11/26 20:40
 * @Author   : cancan
 * @File     : 707_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_707(t *testing.T) {
	tests := []struct {
		methods []string
		params  [][]interface{}
		ans     []interface{}
	}{
		{
			[]string{"MyLinkedList", "AddAtHead", "AddAtTail", "AddAtIndex", "Get", "DeleteAtIndex", "Get"},
			[][]interface{}{{}, {1}, {3}, {1, 2}, {1}, {1}, {1}},
			[]interface{}{nil, nil, nil, nil, 2, nil, 3},
		},
	}

	s2m := utils.NewString2Method()
	var hashMap MyLinkedList

	for index, test := range tests {
		for i, method := range test.methods {
			if i == 0 {
				hashMap = Constructor707()
			} else {
				if !s2m.Check(s2m.Run(&hashMap, method, test.params[i]...), test.ans[i]) {
					t.Errorf("failure id %d", index)
				}
			}
		}
	}

}
