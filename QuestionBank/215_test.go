/*
 * @Time     : 2019/7/9 18:05
 * @Author   : cancan
 * @File     : 215_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_215(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		ans  int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for _, test := range tests {
		if findKthLargest(test.nums, test.k) != test.ans {
			t.Errorf("failure num %v k %d ans %d", test.nums, test.k, test.ans)
		}
	}
}
