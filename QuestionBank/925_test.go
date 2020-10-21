/*
 * @Time     : 2020/10/21 22:39
 * @Author   : cancan
 * @File     : 925_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_925(t *testing.T) {
	tests := []struct {
		name  string
		typed string
		ans   bool
	}{
		{"alex", "aaleex", true},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"laiden", "laiden", true},
	}

	for _, test := range tests {
		if isLongPressedName(test.name, test.typed) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
