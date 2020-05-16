/*
 * @Time     : 2020/5/16 23:08
 * @Author   : cancan
 * @File     : 25_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_25(t *testing.T) {
	tests := []struct {
		head *ListNode
		k    int
		ans  *ListNode
	}{
		{
			InitSinglyLinkList([]int{1, 2, 3, 4, 5}),
			1,
			InitSinglyLinkList([]int{1, 2, 3, 4, 5}),
		},
		{
			InitSinglyLinkList([]int{1, 2, 3, 4, 5}),
			2,
			InitSinglyLinkList([]int{2, 1, 4, 3, 5}),
		},
		{
			InitSinglyLinkList([]int{1, 2}),
			2,
			InitSinglyLinkList([]int{2, 1}),
		},
	}

	for _, test := range tests {
		if !SinglyLinkListEqual(reverseKGroup(test.head, test.k), test.ans) {
			t.Errorf("failure head %v k %d ans %v", test.head, test.k, test.ans)
		}
	}
}
