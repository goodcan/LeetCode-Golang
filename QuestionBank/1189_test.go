/*
 * @Time     : 2019/9/25 11:51
 * @Author   : cancan
 * @File     : 1189_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1189(t *testing.T) {
	tests := []struct {
		text string
		ans  int
	}{
		{"nlaebolko", 1},
		{"loonbalxballpoon", 2},
		{"leetcode", 0},
	}

	for _, test := range tests {
		if maxNumberOfBalloons(test.text) != test.ans {
			t.Errorf("failure text %s ans %d", test.text, test.ans)
		}
	}

}
