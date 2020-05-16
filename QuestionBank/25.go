/*
 * @Time     : 2020/5/16 22:50
 * @Author   : cancan
 * @File     : 25.go
 * @Function : K 个一组翻转链表
 */

/*
 * Question:
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 * k 是一个正整数，它的值小于或等于链表的长度。
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 *
 * Example：
 * 给你这个链表：1->2->3->4->5
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 *
 * Note：
 * 1.你的算法只能使用常数的额外空间。
 * 2.你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 */

package QuestionBank

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	count := 0
	tmp := head
	ret := head
	var start, end *ListNode
	for head != nil {
		count++
		head = head.Next
		if count%k == 0 {
			lastEnd := end
			start, end = reverseSomeListNoe(tmp, k)
			if count/k == 1 {
				ret = start
			} else {
				lastEnd.Next = start
			}
			tmp = head
			end.Next = head
		}
	}
	return ret
}

func reverseSomeListNoe(head *ListNode, k int) (*ListNode, *ListNode) {
	pre := head
	next := head.Next
	head.Next = nil
	n := 1
	for next != nil {
		tmp := next.Next
		next.Next = pre
		pre = next
		next = tmp
		n++
		if n == k {
			return pre, head
		}
	}
	return nil, nil
}
