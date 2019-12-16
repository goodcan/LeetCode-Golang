/*
 * @Time     : 2019/12/16 14:36
 * @Author   : cancan
 * @File     : 230.go
 * @Function : 二叉搜索树中第K小的元素
 */

/*
Question:
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

Note：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

Example 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1

Example 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3

Follow up：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
*/

package QuestionBank

import "sort"

func kthSmallest(root *TreeNode, k int) int {
	l := []int{}
	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		tmp := []*TreeNode{}
		for _, node := range nodes {
			l = append(l, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		nodes = tmp
	}

	sort.Ints(l)
	return l[k-1]
}
