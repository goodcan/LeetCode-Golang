// @Time     : 2019-06-23 14:12
// @Author   : cancan
// @File     : 1002_test.go
// @Function :

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_commonChars(t *testing.T) {
	tests := []struct {
		A   []string
		ans []string
	}{
		{[]string{"bella", "label", "roller"}, []string{"l", "l", "e"}},
		{[]string{"cool", "lock", "cook"}, []string{"o", "c"}},
	}

	for _, test := range tests {
		if utils.SliceEqual(commonChars(test.A), test.ans) {
			t.Errorf("failure A %v ans %v", test.A, test.ans)
		}
	}

}
