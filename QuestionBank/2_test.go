// @Time     : 2019/4/25 19:31
// @Author   : cancan
// @File     : 2_test.go
// @Function : 两数相加 test

package QuestionBank

import (
	"testing"
)

func Test_2(t *testing.T) {
	tests := []struct {
		l1  *ListNode
		l2  *ListNode
		ans *ListNode
	}{
		{
			InitSinglyLinkList([]int{1}),
			InitSinglyLinkList([]int{9, 9}),
			InitSinglyLinkList([]int{0, 0, 1}),
		},
		{
			InitSinglyLinkList([]int{5}),
			InitSinglyLinkList([]int{5}),
			InitSinglyLinkList([]int{0, 1}),
		},
		{
			InitSinglyLinkList([]int{2, 4, 3}),
			InitSinglyLinkList([]int{5, 6, 4}),
			InitSinglyLinkList([]int{7, 0, 8}),
		},
		{
			InitSinglyLinkList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			InitSinglyLinkList([]int{5, 6, 4}),
			InitSinglyLinkList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}

	for _, test := range tests {
		if !SinglyLinkListEqual(addTwoNumbers(test.l1, test.l2), test.ans) {
			t.Errorf("failure l1 %v l2 %v asn %v",
				SinglyLinkList2Slice(test.l1),
				SinglyLinkList2Slice(test.l2),
				SinglyLinkList2Slice(test.ans),
			)
		}
	}
}
