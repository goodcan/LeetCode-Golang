/*
 * @Time     : 2019/8/29 15:26
 * @Author   : cancan
 * @File     : 1154_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1154(t *testing.T) {
	tests := []struct {
		date string
		ans  int
	}{
		{"2019-01-09", 9},
		{"2019-02-10", 41},
		{"2003-03-01", 60},
		{"2004-03-01", 61},
	}

	for _, test := range tests {
		if dayOfYear(test.date) != test.ans {
			t.Errorf("failure date %s ans %d", test.date, test.ans)
		}
	}
}
