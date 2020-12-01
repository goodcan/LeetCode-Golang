// @Time     : 2020/12/1 16:04
// @Author   : cancan
// @File     : 1662_test.go
// @Function :

package QuestionBank

import "testing"

func Test_1662(t *testing.T) {
	tests := []struct {
		word1 []string
		word2 []string
		ans   bool
	}{
		{
			[]string{"ab", "c"},
			[]string{"a", "bc"},
			true,
		},
		{
			[]string{"a", "cb"},
			[]string{"ab", "c"},
			false,
		},
		{
			[]string{"abc", "d", "defg"},
			[]string{"abcddefg"},
			true,
		},
	}

	for _, test := range tests {
		if arrayStringsAreEqual(test.word1, test.word2) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
