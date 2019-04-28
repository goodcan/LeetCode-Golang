// @Time     : 2019/4/28 10:10
// @Author   : cancan
// @File     : 50_test.go
// @Function : Pow(x, n) test

package QuestionBank

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_50(t *testing.T) {
	tests := []struct {
		x   float64
		n   int
		ans float64
	}{
		{0.00000, 0, 1.0},
		{0.44528, 0, 1.0},
		{2.00000, 10, 1024.00000},
		{2.00000, -2, 0.25000},
		{2.10000, 3, 9.26100},
	}

	for _, test := range tests {
		if val, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", myPow(test.x, test.n)), 64); val != test.ans {
			t.Errorf("failure x %f n %d ans %f", test.x, test.n, test.ans)
		}
	}
}
