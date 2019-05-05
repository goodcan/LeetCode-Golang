// @Time     : 2019/5/5 18:46
// @Author   : cancan
// @File     : 771_test.go
// @Function : 宝石与石头 test

package QuestionBank

import "testing"

func Test_771(t *testing.T) {
	tests := []struct {
		J   string
		S   string
		ans int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, test := range tests {
		if numJewelsInStones(test.J, test.S) != test.ans {
			t.Errorf("failure J %s S %s ans %d", test.J, test.S, test.ans)
		}
	}
}
