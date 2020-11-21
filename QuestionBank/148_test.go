/*
 * @Time     : 2020/11/21 11:17
 * @Author   : cancan
 * @File     : 148_test.go
 * @Function :
 */

package QuestionBank

import "testing"

func Test_148(t *testing.T) {
	tests := []struct {
		head *ListNode
		ans  *ListNode
	}{
		{
			InitSinglyLinkList([]int{4, 2, 1, 3}),
			InitSinglyLinkList([]int{1, 2, 3, 4}),
		},
		{
			InitSinglyLinkList([]int{4, 2, 1, 3}),
			InitSinglyLinkList([]int{1, 2, 3, 4}),
		},
	}

	for _, test := range tests {
		if !SinglyLinkListEqual(sortList(test.head), test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
