/*
 * @Time     : 2019/9/25 14:25
 * @Author   : cancan
 * @File     : 1185_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1185(t *testing.T) {
	tests := []struct {
		day   int
		month int
		year  int
		ans   string
	}{
		{31, 8, 2019, "Saturday"},
		{18, 7, 1999, "Sunday"},
		{15, 8, 1993, "Sunday"},
	}

	for _, test := range tests {
		if dayOfTheWeek(test.day, test.month, test.year) != test.ans {
			t.Errorf("failure day %d month %d year %d ans %s",
				test.day, test.month, test.year, test.ans)
		}
	}
}
