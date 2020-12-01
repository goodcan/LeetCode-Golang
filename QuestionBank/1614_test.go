// @Time     : 2020/12/1 11:47
// @Author   : cancan
// @File     : 1614_test.go
// @Function :

package QuestionBank

import "testing"

func Test_1614(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
		{"1+(2*3)/(2-1)", 1},
		{"1", 0},
	}

	for _, test := range tests {
		if maxDepth1614(test.s) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
