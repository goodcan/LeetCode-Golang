/*
 * @Time     : 2019/8/28 10:51
 * @Author   : cancan
 * @File     : 1161.go
 * @Function : 最大层内元素和
 */

/*
 * Question:
 * 给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
 * 请你找出层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中 最小 的那个。
 *
 * Example
 *     1
 *    / \
 *   7   0
 *  / \
 * 7  -8
 * 输入：[1,7,0,7,-8,null,null]
 * 输出：2
 * 解释：
 * 第 1 层各元素之和为 1，
 * 第 2 层各元素之和为 7 + 0 = 7，
 * 第 3 层各元素之和为 7 + -8 = -1，
 * 所以我们返回第 2 层的层号，它的层内元素之和最大。
 *
 * Note：
 * 树中的节点数介于 1 和 10^4 之间
 * -10^5 <= node.val <= 10^5
 */

package QuestionBank

func maxLevelSum(root *TreeNode) int {
	node := root

	if node == nil {
		return 0
	}

	maxLevel := 1
	maxSum := node.Val

	level := 1
	nodes := []*TreeNode{node}

	for len(nodes) != 0 {
		var tmp []*TreeNode
		level++
		s := 0

		for _, v := range nodes {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
				s += v.Left.Val
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
				s += v.Right.Val
			}
		}

		if s > maxSum {
			maxSum = s
			maxLevel = level
		}

		nodes = tmp
	}

	return maxLevel
}
