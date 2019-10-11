/*
 * @Time     : 2019/10/8 11:19
 * @Author   : cancan
 * @File     : 5205_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_5205(t *testing.T) {
	tests := []struct {
		arr []int
		ans bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, test := range tests {
		if uniqueOccurrences(test.arr) != test.ans {
			t.Errorf("failure arr %v ans %v", test.arr, test.ans)
		}
	}
}
