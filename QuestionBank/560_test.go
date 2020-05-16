/*
 * @Time     : 2020/5/16 21:00
 * @Author   : cancan
 * @File     : 560_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_560(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		ans  int
	}{
		{
			[]int{1, 1, 1},
			2,
			2,
		},
	}

	for _, test := range tests {
		if subarraySum(test.nums, test.k) != test.ans {
			t.Errorf("failure nums %v k %v ans %v", test.ans, test.k, test.ans)
		}
	}
}
