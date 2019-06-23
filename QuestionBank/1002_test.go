// @Time     : 2019-06-23 14:12
// @Author   : cancan
// @File     : 1002_test.py
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_commonChars(t *testing.T) {
	tests := []struct {
		A   []string
		ans []string
	}{
		{[]string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{[]string{"cool", "lock", "cook"}, []string{"c", "o"}},
	}

	for _, test := range tests {
		if utils.StringSliceEqual(commonChars(test.A), test.ans) {
			t.Error("failure A %v ans %v", test.A, test.ans)
		}
	}

}
