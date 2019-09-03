/*
 * @Time     : 2019/9/3 18:38
 * @Author   : cancan
 * @File     : 537_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_537(t *testing.T) {
	tests := []struct {
		a   string
		b   string
		ans string
	}{
		{"1+1i", "1+1i", "0+2i"},
		{"1+-1i", "1+-1i", "0+-2i"},
	}

	for _, test := range tests {
		if complexNumberMultiply(test.a, test.b) != test.ans {
			t.Errorf("failure a %s b %s ans %s", test.a, test.b, test.ans)
		}
	}
}
