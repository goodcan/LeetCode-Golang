/*
 * @Time     : 2019/9/4 9:58
 * @Author   : cancan
 * @File     : 144.go
 * @Function : 二叉树的前序遍历
 */

/*
 * Question:
 * 给定一个二叉树，返回它的 前序 遍历。
 *
 * Example:
 * 输入: [1,null,2,3]
 *    1
 *     \
 *      2
 *     /
 *    3
 * 输出: [1,2,3]
 *
 * Follow up:
 * 递归算法很简单，你可以通过迭代算法完成吗？
 */

package QuestionBank

func preorderTraversal(root *TreeNode) []int {
	var r []int

	if root == nil {
		return r
	}

	r = append(r, root.Val)
	r = append(r, preorderTraversal(root.Left)...)
	r = append(r, preorderTraversal(root.Right)...)

	return r
}
