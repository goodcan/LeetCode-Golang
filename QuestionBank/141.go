/*
 * @Time     : 2019/11/22 10:51
 * @Author   : cancan
 * @File     : 141.go
 * @Function : 环形链表
 */

/*
 * Question:
 * 给定一个链表，判断链表中是否有环。
 * 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
 *
 * Example 1：
 * 输入：head = [3,2,0,-4], pos = 1
 * 输出：true
 * 解释：链表中有一个环，其尾部连接到第二个节点。
 *
 * Example 2：
 * 输入：head = [1,2], pos = 0
 * 输出：true
 * 解释：链表中有一个环，其尾部连接到第一个节点。
 *
 * Example 3：
 * 输入：head = [1], pos = -1
 * 输出：false
 * 解释：链表中没有环。
 *
 * Follow up：
 * 你能用 O(1)（即，常量）内存解决此问题吗？
 */

package QuestionBank

func hasCycle(head *ListNode) bool {
	tmp := make(map[*ListNode]*ListNode)
	for head != nil {
		if _, ok := tmp[head]; ok {
			return true
		} else {
			tmp[head] = head
			head = head.Next
		}
	}
	return false
}
