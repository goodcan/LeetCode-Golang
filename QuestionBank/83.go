// @Time     : 2019/4/28 14:36
// @Author   : cancan
// @File     : 83.go
// @Function : 删除排序链表中的重复元素

/*
 * Question:
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 *
 * Example 1:
 * 输入: 1->1->2
 * 输出: 1->2
 *
 * Example 2:
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
 */

package QuestionBank

func deleteDuplicates(head *ListNode) *ListNode {
	temp := make(map[int]*ListNode)
	prev := new(ListNode)
	node := head

	for node != nil {
		n, ok := temp[node.Val]
		if ok {
			node = node.Next
			n.Next = node
			prev = n
		} else {
			prev = node
			temp[node.Val] = prev
			node = node.Next
		}
	}

	return head
}
