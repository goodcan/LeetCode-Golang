/*
 * @Time     : 2021/11/6 14:26
 * @Author   : cancan
 * @File     : 237.go
 * @Function : 删除链表的结点
 */

/*
 * Question:
 * 请编写一个函数，使其可以删除某个链表中给定的（非末尾的）节点，您将只被给予要求被删除的节点。
 * 比如：假设该链表为 1 -> 2 -> 3 -> 4  ，给定您的为该链表中值为 3 的第三个节点，
 * 那么在调用了您的函数之后，该链表则应变成 1 -> 2 -> 4 。
 */

package QuestionBank

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
