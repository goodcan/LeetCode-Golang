/*
 * @Time     : 2020/9/19 14:13
 * @Author   : cancan
 * @File     : 404.go
 * @Function : 左叶子之和
 */

/*
 * 计算给定二叉树的所有左叶子之和。
 *
 * 示例：
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 *
 * 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 */

package QuestionBank

func sumOfLeftLeaves(root *TreeNode) int {
	ret := 0
	if root == nil {
		return ret
	}

	tmp := []*TreeNode{root}
	for len(tmp) > 0 {
		newTmp := []*TreeNode{}
		for _, node := range tmp {
			if node.Left != nil {
				if node.Left.Left == nil && node.Left.Right == nil {
					ret += node.Left.Val
				} else {
					newTmp = append(newTmp, node.Left)
				}
			}
			if node.Right != nil {
				newTmp = append(newTmp, node.Right)
			}
		}
		tmp = newTmp
	}

	return ret
}
