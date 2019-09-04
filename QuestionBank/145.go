/*
 * @Time     : 2019/9/4 10:25
 * @Author   : cancan
 * @File     : 145.go
 * @Function : 二叉树的后序遍历
 */

/*
 * Question:
 * 给定一个二叉树，返回它的 后序 遍历。
 *
 * Example:
 * 输入: [1,null,2,3]
 *    1
 *     \
 *      2
 *     /
 *    3
 * 输出: [3,2,1]
 *
 * Follow up:
 * 递归算法很简单，你可以通过迭代算法完成吗？
 */

package QuestionBank

func postorderTraversal(root *TreeNode) []int {
	var r []int

	if root == nil {
		return r
	}

	r = append(r, postorderTraversal(root.Left)...)
	r = append(r, postorderTraversal(root.Right)...)
	r = append(r, root.Val)

	return r
}
