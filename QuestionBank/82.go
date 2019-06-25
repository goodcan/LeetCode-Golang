// @Time     : 2019/4/28 11:22
// @Author   : cancan
// @File     : 82.go
// @Function : 删除排序链表中的重复元素 II

/*
 *Question:
 *给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 *Example 1:
 *输入: 1->2->3->3->4->4->5
 *输出: 1->2->5
 *
 *Example 2:
 *输入: 1->1->1->2->3
 *输出: 2->3
 */

package QuestionBank

func deleteDuplicatesII(head *ListNode) *ListNode {
	temp := make(map[int]*ListNode)
	prev := new(ListNode)

	node := head

	for node != nil {
		n, ok := temp[node.Val]

		if ok {
			node = node.Next
			if n.Next == nil {
				head = node
			} else {
				n.Next = node
			}
			prev = n
		} else {
			temp[node.Val] = prev
			prev = node
			node = node.Next
		}
	}

	return head
}
