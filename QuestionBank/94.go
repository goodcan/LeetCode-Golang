/*
 * @Time     : 2019/4/28 14:47
 * @Author   : cancan
 * @File     : 94.go
 * @Function : 二叉树的中序遍历
 */

/*
 * Question:
 * 给定一个二叉树，返回它的中序 遍历。
 *
 * Example:
 * 输入: [1,null,2,3]
 *  1
 *   \
 *    2
 *   /
 *  3
 *
 * 输出: [1,3,2]
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */

package QuestionBank

func inOrderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	ans = append(ans, inOrderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inOrderTraversal(root.Right)...)

	return ans
}
