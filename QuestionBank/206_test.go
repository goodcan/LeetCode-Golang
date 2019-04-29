// @Time     : 2019/4/29 17:13
// @Author   : cancan
// @File     : 206_test.go
// @Function : 反转链表 test

package QuestionBank

import "testing"

func Test_206(t *testing.T) {
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
			InitSinglyLinkList([]int{1, 2, 3, 4, 5}),
			InitSinglyLinkList([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, test := range tests {
		if !SinglyLinkListEqual(reverseList(test.head), test.ans) {
			t.Errorf("failure head %v ans %v",
				SinglyLinkList2Slice(test.head),
				SinglyLinkList2Slice(test.ans),
			)
		}
	}
}
