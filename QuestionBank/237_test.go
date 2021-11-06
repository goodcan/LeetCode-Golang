/*
 * @Time     : 2021/11/6 14:28
 * @Author   : cancan
 * @File     : 237_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"
)

func Test_237(t *testing.T) {
	tests := []struct {
		node   *ListNode
		delVal int
		ans    *ListNode
	}{
		{InitSinglyLinkList([]int{1, 2, 3, 4, 5}), 4, InitSinglyLinkList([]int{1, 2, 3, 5})},
	}

	for _, test := range tests {
		delNode := test.node
		for delNode.Val != test.delVal {
			delNode = delNode.Next
		}
		deleteNode(delNode)
		if !SinglyLinkListEqual(test.node, test.ans) {
			t.Errorf("failure %+v", test)
		}
	}
}
