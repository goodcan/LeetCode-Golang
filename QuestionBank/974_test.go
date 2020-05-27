/*
 * @Time     : 2020/5/27 11:14
 * @Author   : cancan
 * @File     : 974_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_974(t *testing.T) {
	tests := []struct {
		A   []int
		K   int
		ans int
	}{
		{[]int{4, 5, 0, -2, -3, 1}, 5, 7},
		{[]int{-1, 2, 9}, 2, 2},
		{[]int{2, -2, 2, -4}, 6, 2},
	}

	for _, test := range tests {
		if subarraysDivByK(test.A, test.K) != test.ans {
			t.Errorf("failure A %v K %d ans %d", test.A, test.K, test.ans)
		}
	}
}
