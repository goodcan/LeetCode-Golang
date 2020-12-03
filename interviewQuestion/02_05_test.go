// @Time     : 2020/12/3 20:26
// @Author   : cancan
// @File     : 02_05_test.go
// @Function :

package interviewQuestion

import "testing"

func Test_02_05(t *testing.T) {
	tests := []struct {
		l1  *ListNode
		l2  *ListNode
		ans *ListNode
	}{
		{
			InitSinglyLinkList([]int{7, 1, 6}),
			InitSinglyLinkList([]int{5, 9, 2}),
			InitSinglyLinkList([]int{2, 1, 9}),
		},
		{
			InitSinglyLinkList([]int{5}),
			InitSinglyLinkList([]int{5}),
			InitSinglyLinkList([]int{0, 1}),
		},
	}

	for idx, test := range tests {
		if !SinglyLinkListEqual(addTwoNumbers(test.l1, test.l2), test.ans) {
			t.Errorf("failure %+v", idx)
		}
	}
}
