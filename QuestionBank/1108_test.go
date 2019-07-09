/*
 * @Time     : 2019/7/8 17:44
 * @Author   : cancan
 * @File     : 1108_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1108(t *testing.T) {
	tests := []struct {
		address string
		ans     string
	}{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
	}

	for _, test := range tests {
		if defangIPaddr(test.address) != test.ans {
			t.Errorf("failure address %s ans %s", test.address, test.ans)
		}
	}
}
