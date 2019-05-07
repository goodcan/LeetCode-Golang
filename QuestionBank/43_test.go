// @Time     : 2019/5/7 16:28
// @Author   : cancan
// @File     : 43_test.go
// @Function : 字符串相乘 test

package QuestionBank

import "testing"

func Test_43(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		ans  string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"498828660196", "840477629533", "419254329864656431168468"},
	}

	for _, test := range tests {
		if multiply(test.num1, test.num2) != test.ans {
			t.Errorf("failure num1 %s num2 %s ans %s", test.num1, test.num2, test.ans)
		}
	}
}
