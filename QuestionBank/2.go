/*
 * @Time     : 2019/4/25 19:10
 * @Author   : cancan
 * @File     : 2.go
 * @Function : 两数相加
 */

/*
 * Question:
 * 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。
 * 将两数相加返回一个新的链表。你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 * Example：
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 */

package QuestionBank

import "strconv"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := link2Str(l1)
	s2 := link2Str(l2)

	s1l := len(s1)
	s2l := len(s2)

	var l int

	if s1l >= s2l {
		l = s1l
	} else {
		l = s2l
	}

	prev := new(ListNode)
	ret := prev
	flag := 0

	for i := 0; i < l; i++ {
		n1 := 0
		n2 := 0
		if i < s1l {
			n1, _ = strconv.Atoi(string(s1[i]))
		}

		if i < s2l {
			n2, _ = strconv.Atoi(string(s2[i]))
		}

		n := n1 + n2 + flag

		prev = addNode(prev, n)

		flag = n / 10
	}

	if flag == 1 {
		addNode(prev, flag)
	}

	return ret.Next
}

func link2Str(head *ListNode) string {
	ret := ""

	for head != nil {
		ret += strconv.Itoa(head.Val)
		head = head.Next
	}

	return ret
}

func addNode(prev *ListNode, n int) *ListNode {
	temp := new(ListNode)
	temp.Val = n % 10
	prev.Next = temp
	prev = temp

	return prev
}
