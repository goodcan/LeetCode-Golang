/*
 * @Time     : 2020/5/7 10:25
 * @Author   : cancan
 * @File     : 572.go
 * @Function : 另一个树的子树
 */

/*
 * Question:
 * 给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。
 *
 * Example 1:
 * 给定的树 s:
 *      3
 *     / \
 *    4   5
 *   / \
 *  1   2
 * 给定的树 t：
 *    4
 *   / \
 *  1   2
 * 返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。
 *
 * Example 2:
 * 给定的树 s：
 *      3
 *     / \
 *    4   5
 *   / \
 *  1   2
 *     /
 *    0
 * 给定的树 t：
 *    4
 *   / \
 *  1   2
 * 返回 false。
 */

package QuestionBank

func isSubtree(s *TreeNode, t *TreeNode) bool {
	tmp := []*TreeNode{s}
	for len(tmp) != 0 {
		newTmp := []*TreeNode{}
		for _, i := range tmp {
			if i.Left != nil {
				newTmp = append(newTmp, i.Left)
			}
			if i.Right != nil {
				newTmp = append(newTmp, i.Right)
			}
			if isSubtreeCheck(i, t) {
				return true
			}
		}
		tmp = newTmp
	}
	return false
}

func isSubtreeCheck(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}

	return isSubtreeCheck(n1.Left, n2.Left) && isSubtreeCheck(n1.Right, n2.Right)
}
