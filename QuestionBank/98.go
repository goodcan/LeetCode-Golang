/*
 * @Time     : 2019/4/28 15:35
 * @Author   : cancan
 * @File     : 98.go
 * @Function : 验证二叉搜索树
 */

/*
 * Question:
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 假设一个二叉搜索树具有如下特征：
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 * Example 1:
 * 输入:
 *     2
 *    / \
 *   1   3
 * 输出: true
 *
 * Example 2:
 * 输入:
 *     5
 *    / \
 *   1   4
 *      / \
 *     3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 *      根节点的值为 5 ，但是其右子节点值为 4 。
 */

package QuestionBank

func isValidBST(root *TreeNode) bool {
	list := traversal(root)

	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false
		}
	}

	return true
}

func traversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	ans = append(ans, traversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, traversal(root.Right)...)

	return ans
}
