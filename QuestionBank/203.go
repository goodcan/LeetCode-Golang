/*
 * @Time     : 2019-11-16 15:08
 * @Author   : cancan
 * @File     : 203.go
 * @Function : 移除链表元素
 */

/*
 * Question:
 * 删除链表中等于给定值 val 的所有节点。
 *
 * Example:
 * 输入: 1->2->6->3->4->5->6, val = 6
 * 输出: 1->2->3->4->5
 */

package QuestionBank

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	tmp := head
	var prev *ListNode
	for tmp != nil {
		if tmp.Val == val {
			if prev != nil {
				prev.Next = tmp.Next
				tmp = tmp.Next
			} else {
				head = tmp.Next
				tmp.Next = nil
				tmp = head
			}
		} else {
			prev = tmp
			tmp = tmp.Next
		}
	}

	return head
}
