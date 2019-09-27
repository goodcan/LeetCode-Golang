/*
 * @Time     : 2019/9/27 10:14
 * @Author   : cancan
 * @File     : 617.go
 * @Function : 合并二叉树
 */

/**
 * Question:
 * 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
 * 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，
 * 否则不为 NULL 的节点将直接作为新二叉树的节点。
 *
 * Example 1:
 * 输入:
 * 	Tree 1                     Tree 2
 *           1                         2
 *          / \                       / \
 *         3   2                     1   3
 *        /                           \   \
 *       5                             4   7
 * 输出:
 * 合并后的树:
 * 	     3
 * 	    / \
 * 	   4   5
 * 	  / \   \
 * 	 5   4   7
 *
 * Note: 合并必须从两个树的根节点开始。
 */

package QuestionBank

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	switch {
	case t1 == nil:
		return t2
	case t2 == nil:
		return t1
	}

	mergeTreesDfs(t1, t2)

	return t1
}

func mergeTreesDfs(node1, node2 *TreeNode) {
	node1.Val += node2.Val

	if node2.Left != nil {
		if node1.Left == nil {
			node1.Left = node2.Left
		} else {
			mergeTreesDfs(node1.Left, node2.Left)
		}
	}

	if node2.Right != nil {
		if node1.Right == nil {
			node1.Right = node2.Right
		} else {
			mergeTreesDfs(node1.Right, node2.Right)
		}
	}
}
