/*
 * @Time     : 2020/9/13 01:22
 * @Author   : cancan
 * @File     : 79_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_79(t *testing.T) {
	tests := []struct {
		board [][]byte
		word  string
		ans   bool
	}{
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "SEE", true},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCB", false},
	}

	for _, test := range tests {
		if exist(test.board, test.word) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
