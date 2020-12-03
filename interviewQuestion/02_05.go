// @Time     : 2020/12/3 20:09
// @Author   : cancan
// @File     : 02_05.go
// @Function : 链表求和

/*
 * 给定两个用链表表示的整数，每个节点包含一个数位。
 * 这些数位是反向存放的，也就是个位排在链表首部。
 * 编写函数对这两个整数求和，并用链表形式返回结果。
 *
 * 示例：
 * 输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
 * 输出：2 -> 1 -> 9，即912
 * 进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?
 *
 * 示例：
 * 输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
 * 输出：9 -> 1 -> 2，即912
 */

package interviewQuestion

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	startTmp := &ListNode{}
	pre := startTmp
	add := 0
	for l1 != nil || l2 != nil {
		tmp := add
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		add = tmp / 10

		pre.Next = &ListNode{Val: tmp % 10}
		pre = pre.Next
	}

	if add != 0 {
		pre.Next = &ListNode{Val: add}
	}

	return startTmp.Next
}
