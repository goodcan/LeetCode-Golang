/*
 * @Time     : 2020/9/16 11:01
 * @Author   : cancan
 * @File     : 226.go
 * @Function : 翻转二叉树
 */

/*
 * 翻转一棵二叉树。
 *
 * 示例：
 * 输入：
 *      4
 *    /   \
 *   2     7
 *  / \   / \
 * 1   3 6   9
 * 输出：
 *      4
 *    /   \
 *   7     2
 *  / \   / \
 * 9   6 3   1
 *
 * 备注:
 * - 这个问题是受到 Max Howell 的 原问题 启发的 ：
 * - 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
 */

package QuestionBank

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := []*TreeNode{root}
	for len(tmp) > 0 {
		newTmp := []*TreeNode{}
		for _, node := range tmp {
			if node.Left != nil {
				newTmp = append(newTmp, node.Left)
			}
			if node.Right != nil {
				newTmp = append(newTmp, node.Right)
			}
			node.Left, node.Right = node.Right, node.Left
			tmp = newTmp
		}
	}
	return root
}
