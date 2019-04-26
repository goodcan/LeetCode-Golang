// @Time     : 2019/4/25 19:12
// @Author   : cancan
// @File     : define.go
// @Function : 定义类型

package QuestionBank

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
