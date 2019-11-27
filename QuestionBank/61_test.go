/*
 * @Time     : 2019/11/27 11:39
 * @Author   : cancan
 * @File     : 61_test.go
 * @Function :
 */

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_61(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		ans  *ListNode
	}{
		{
			InitSinglyLinkList([]int{1, 2, 3, 4, 5}),
			2,
			InitSinglyLinkList([]int{4, 5, 1, 2, 3}),
		},
		{
			InitSinglyLinkList([]int{0, 1, 2}),
			4,
			InitSinglyLinkList([]int{2, 0, 1}),
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(SinglyLinkList2Slice(rotateRight(test.head, test.k)), SinglyLinkList2Slice(test.ans)) {
			t.Errorf("failure head %v k %d ans %v", test.head, test.k, test.ans)
		}
	}
}
