/*
 * @Time     : 2019-06-26 21:54
 * @Author   : cancan
 * @File     : 876_test.py
 * @Function :
 */

package QuestionBank

import "testing"

func Test_876(t *testing.T) {
	tests := []struct {
		head *ListNode
		ans  *ListNode
	}{
		{
			InitSinglyLinkList([]int{1, 2, 3, 4, 5}),
			InitSinglyLinkList([]int{3, 4, 5}),
		},
		{
			InitSinglyLinkList([]int{1, 2, 3, 4, 5, 6}),
			InitSinglyLinkList([]int{4, 5, 6}),
		},
	}

	for _, test := range tests {
		if !SinglyLinkListEqual(middleNode(test.head), test.ans) {
			t.Errorf("failure head %v ans %v",
				SinglyLinkList2Slice(test.head),
				SinglyLinkList2Slice(test.ans),
			)
		}
	}
}
