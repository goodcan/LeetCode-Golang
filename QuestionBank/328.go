/*
 * @Time     : 2019/6/27 11:09
 * @Author   : cancan
 * @File     : 328.go
 * @Function : 奇偶链表
 */

/*
 *Question:
 *给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，
 *这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
 *请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
 *
 *Example 1:
 *输入: 1->2->3->4->5->NULL
 *输出: 1->3->5->2->4->NULL
 *
 *Example 2:
 *输入: 2->1->3->5->6->4->7->NULL
 *输出: 2->3->6->7->1->5->4->NULL
 *
 *Note:
 *应当保持奇数节点和偶数节点的相对顺序。
 *链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
 */

package QuestionBank

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	flag := 1

	prev1 := head
	var prev2 *ListNode
	var tmp *ListNode

	ret := head

	for {
		if flag%2 != 0 {
			if head.Next == nil {
				head.Next = tmp
				if prev2 != nil {
					prev2.Next = nil
				}
				break
			}
			prev1 = head
		} else {
			if tmp == nil {
				tmp = head
			}
			if head.Next == nil {
				if prev2 != nil {
					prev2.Next = head
				}
				prev1.Next = tmp
				break
			}
			prev1.Next = head.Next
			if prev2 != nil {
				prev2.Next = head
			}
			prev2 = head
		}

		head = head.Next
		flag++
	}

	return ret
}
