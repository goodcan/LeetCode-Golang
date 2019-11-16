/*
 * @Time     : 2019-11-16 15:21
 * @Author   : cancan
 * @File     : 203_test.go
 * @Function :
 */

package QuestionBank

import (
	"testing"

	"LeetCode-Golang/utils"
)

func Test_203(t *testing.T) {
	tests := []struct {
		head *ListNode
		val  int
		ans  *ListNode
	}{
		{
			InitSinglyLinkList([]int{1, 2, 6, 3, 4, 5, 6}),
			6,
			InitSinglyLinkList([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		if !utils.SliceEqual(SinglyLinkList2Slice(removeElements(test.head, test.val)), SinglyLinkList2Slice(test.ans)) {
			t.Errorf("failure %v %d %v", test.head, test.val, test.ans)
		}
	}

}
