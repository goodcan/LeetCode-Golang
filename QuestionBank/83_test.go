// @Time     : 2019/4/28 14:36
// @Author   : cancan
// @File     : 83_test.go
// @Function : 删除排序链表中的重复元素 test

package QuestionBank

import "testing"

func Test_83(t *testing.T) {
	tests := []struct {
		head *ListNode
		ans  *ListNode
	}{
		{
			InitSinglyLinkList([]int{1, 1, 2}),
			InitSinglyLinkList([]int{1, 2}),
		},
		{
			InitSinglyLinkList([]int{1, 1, 2, 3, 3}),
			InitSinglyLinkList([]int{1, 2, 3}),
		},
	}

	for _, test := range tests {
		if !SinglyLinkListEqual(deleteDuplicates(test.head), test.ans) {
			t.Errorf("failure head %v ans %v",
				SinglyLinkList2Slice(test.head),
				SinglyLinkList2Slice(test.ans),
			)
		}
	}
}
