// @Time     : 2019/4/29 17:13
// @Author   : cancan
// @File     : 206.go
// @Function : 反转链表

/*
Question:
反转一个单链表。

Example:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL

Follow up:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/

package QuestionBank

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode = nil

	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}

	return pre
}
