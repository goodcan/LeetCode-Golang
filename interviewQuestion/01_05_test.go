/*
 * @Time     : 2022/1/3 00:20
 * @Author   : cancan
 * @File     : 01.05_test.go
 * @Function :
 */

package interviewQuestion

import "testing"

func Test_01_05(t *testing.T) {
	tests := []struct {
		first  string
		second string
		ans    bool
	}{
		{"pale", "ple", true},
		{"pales", "pal", false},
	}

	for _, test := range tests {
		if oneEditAway(test.first, test.second) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
