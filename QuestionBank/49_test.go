// @Time     : 2020/12/14 9:39
// @Author   : cancan
// @File     : 49_test.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_49(t *testing.T) {
	tests := []struct {
		strs []string
		ans  [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(groupAnagrams(test.strs), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
