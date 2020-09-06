/*
 * @Time     : 2020/9/6 22:51
 * @Author   : cancan
 * @File     : 107.go
 * @Function : 二叉树的层次遍历 II
 */

/*
 * 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 返回其自底向上的层次遍历为：
 * [
 *   [15,7],
 *   [9,20],
 *   [3]
 * ]
 */

package QuestionBank

func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		tmp := []*TreeNode{}
		d := []int{}
		for _, node := range nodes {
			d = append(d, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		ret = append([][]int{d}, ret...)
		nodes = tmp
	}
	return ret
}
