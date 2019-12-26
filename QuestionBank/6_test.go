/*
 * @Time     : 2019/12/26 19:37
 * @Author   : cancan
 * @File     : 6_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_6(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		ans     string
	}{
		{"LEETCODEISHIRING", 3, "LCIRETOESIIGEDHN"},
		{"LEETCODEISHIRING", 4, "LDREOEIIECIHNTSG"},
	}

	for _, test := range tests {
		if convert(test.s, test.numRows) != test.ans {
			t.Errorf("failure s %s numRows %d ans %s",
				test.s, test.numRows, test.ans)
		}
	}
}
