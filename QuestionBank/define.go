/*
 * @Time     : 2019/4/25 19:12
 * @Author   : cancan
 * @File     : define.go
 * @Function : 定义类型
 */

package QuestionBank

import "LeetCode-Golang/utils"

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func InitSinglyLinkList(nums []int) *ListNode {
	prev := new(ListNode)

	ret := prev

	for _, v := range nums {
		temp := new(ListNode)
		temp.Val = v
		prev.Next = temp
		prev = temp
	}

	return ret.Next
}

func SinglyLinkList2Slice(l *ListNode) []int {
	ret := []int{}

	for l != nil {
		ret = append(ret, l.Val)
		l = l.Next
	}

	return ret
}

func SinglyLinkListEqual(l1, l2 *ListNode) bool {
	return utils.SliceEqual(SinglyLinkList2Slice(l1), SinglyLinkList2Slice(l2))
}

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InitTree(nums []interface{}) *TreeNode {
	l := len(nums)

	if l == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0].(int)}
	nodeQueue := []*TreeNode{root}
	front := 0
	index := 1

	for index < l {
		node := nodeQueue[front]
		front++

		item := nums[index]
		index++

		if item != nil {
			node.Left = &TreeNode{Val: item.(int)}
			nodeQueue = append(nodeQueue, node.Left)
		}

		if index >= l {
			break
		}

		item = nums[index]
		index++
		if item != nil {
			node.Right = &TreeNode{Val: item.(int)}
			nodeQueue = append(nodeQueue, node.Right)
		}
	}

	return root
}

func Tree2Slice(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)
	ans = append(ans, Tree2Slice(root.Left)...)
	ans = append(ans, Tree2Slice(root.Right)...)

	return ans
}

func TreeEqual(t1, t2 *TreeNode) bool {
	return utils.SliceEqual(Tree2Slice(t1), Tree2Slice(t2))
}
