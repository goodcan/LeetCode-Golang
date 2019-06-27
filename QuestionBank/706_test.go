// @Time     : 2019/4/30 14:16
// @Author   : cancan
// @File     : 706_test.go
// @Function : 设计哈希映射 test

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_706(t *testing.T) {
	tests := []struct {
		methods []string
		params  [][]interface{}
		ans     []interface{}
	}{
		{
			[]string{"MyHashMap", "Put", "Put", "Get", "Get", "Put", "Get", "Remove", "Get"},
			[][]interface{}{{}, {1, 1}, {2, 2}, {1}, {3}, {2, 1}, {2}, {2}, {2}},
			[]interface{}{nil, nil, nil, 1, -1, nil, 1, nil, -1},
		},
	}

	s2m := utils.NewString2Method()
	var hashMap MyHashMap

	for index, test := range tests {
		for i, method := range test.methods {
			if i == 0 {
				hashMap = Constructor706()
			} else {
				if !s2m.Check(s2m.Run(&hashMap, method, test.params[i]...), test.ans[i]) {
					t.Errorf("failure id %d", index)
				}
			}
		}
	}

}
