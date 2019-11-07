/*
 * @Time     : 2019/11/7 10:20
 * @Author   : cancan
 * @File     : 1249_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_1249(t *testing.T) {
	tests := []struct {
		s   string
		ans string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
		{"(a(b(c)d)", "a(b(c)d)"},
	}

	for _, test := range tests {
		if minRemoveToMakeValid(test.s) != test.ans {
			t.Errorf("failure s %s ans %s", test.s, test.ans)
		}
	}
}
