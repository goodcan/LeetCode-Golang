/*
 * @Time     : 2020/11/21 10:54
 * @Author   : cancan
 * @File     : 148.go
 * @Function : 排序链表
 */

/*
 * Question:
 * 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。
 *
 * Example 1:
 * 输入: 4->2->1->3
 * 输出: 1->2->3->4
 *
 * Example 2:
 * 输入: -1->5->3->4->0
 * 输出: -1->0->3->4->5
 */

package QuestionBank

import "sort"

func sortList(head *ListNode) *ListNode {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	sort.Ints(vals)

	var ret *ListNode
	if len(vals) > 0 {
		ret = &ListNode{Val: vals[0]}
	}
	tmp := ret
	for i := 1; i < len(vals); i++ {
		tmp.Next = &ListNode{Val: vals[i]}
		tmp = tmp.Next
	}

	return ret
}
