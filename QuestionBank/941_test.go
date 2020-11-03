// @Time     : 2020/11/3 10:33
// @Author   : cancan
// @File     : 941_test.go
// @Function :

package QuestionBank

import "testing"

func Test_941(t *testing.T) {
	tests := []struct {
		A   []int
		ans bool
	}{
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{6, 5, 3}, false},
		{[]int{0, 3, 2, 1}, true},
		{[]int{0, 3, 2, 1, 3}, false},
	}

	for _, test := range tests {
		if validMountainArray(test.A) != test.ans {
			t.Errorf("failure %+v", test)
		}
	}
}
