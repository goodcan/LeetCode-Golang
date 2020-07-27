/*
 * @Time     : 2020/7/28 00:11
 * @Author   : cancan
 * @File     : 104.go
 * @Function : 二叉树的最大深度
 */

/*
 * Question:
 * 给定一个二叉树，找出其最大深度。
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 *
 * Note:
 * 叶子节点是指没有子节点的节点。
 *
 * Example:
 * 给定二叉树 [3,9,20,null,null,15,7]，
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 返回它的最大深度 3 。
 */

package QuestionBank

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1
	tmp := []*TreeNode{}
	if root.Left != nil {
		tmp = append(tmp, root.Left)
	}
	if root.Right != nil {
		tmp = append(tmp, root.Right)
	}

	for len(tmp) > 0 {
		ans++
		newTmp := []*TreeNode{}
		for _, item := range tmp {
			if item.Left != nil {
				newTmp = append(newTmp, item.Left)
			}
			if item.Right != nil {
				newTmp = append(newTmp, item.Right)
			}
			tmp = newTmp
		}
	}

	return ans
}
