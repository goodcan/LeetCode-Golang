/*
 * @Time     : 2019/6/27 11:28
 * @Author   : cancan
 * @File     : 328_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"
)

func Test_328(t *testing.T) {
	tests := []struct {
		head *ListNode
		ans  *ListNode
	}{
		{
			InitSinglyLinkList([]int{}),
			InitSinglyLinkList([]int{}),
		},
		{
			InitSinglyLinkList([]int{1}),
			InitSinglyLinkList([]int{1}),
		},
		{
			InitSinglyLinkList([]int{1, 1}),
			InitSinglyLinkList([]int{1, 1}),
		},
		{
			InitSinglyLinkList([]int{1, 2, 3, 4, 5}),
			InitSinglyLinkList([]int{1, 3, 5, 2, 4}),
		},
		{
			InitSinglyLinkList([]int{2, 1, 3, 5, 6, 4, 7}),
			InitSinglyLinkList([]int{2, 3, 6, 7, 1, 5, 4}),
		},
		{
			InitSinglyLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8}),
			InitSinglyLinkList([]int{1, 3, 5, 7, 2, 4, 6, 8}),
		},
	}

	for _, test := range tests {
		if !SinglyLinkListEqual(oddEvenList(test.head), test.ans) {
			t.Errorf("failure head %v ans %v",
				SinglyLinkList2Slice(test.head),
				SinglyLinkList2Slice(test.ans),
			)
		}
	}
}
