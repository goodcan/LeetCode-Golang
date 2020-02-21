/*
 * @Time     : 2020/2/21 11:49
 * @Author   : cancan
 * @File     : 01.01_test.go
 * @Function :
 */

package interviewQuestion

import "testing"

func Test_01_01(t *testing.T) {
	tests := []struct {
		astr string
		ans  bool
	}{
		{"leetcode", false},
		{"abc", true},
	}

	for _, test := range tests {
		if isUnique(test.astr) != test.ans {
			t.Errorf("failure astr %s ans %v", test.astr, test.ans)
		}
	}
}
