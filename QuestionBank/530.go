/*
 * @Time     : 2020/10/12 15:54
 * @Author   : cancan
 * @File     : 530.go
 * @Function : 二叉搜索树的最小绝对差
 */

/*
 * 给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。
 *
 * 示例：
 * 输入：
 *
 *    1
 *     \
 *      3
 *     /
 *    2
 * 输出：
 * 1
 *
 * 解释：
 * 最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
 *
 * 提示：
 * 树中至少有 2 个节点。
 */

package QuestionBank

func getMinimumDifference(root *TreeNode) int {
	ans := -1
	pre := -1
	getMinimumDifferenceDfs(root, &ans, &pre)
	return ans
}

func getMinimumDifferenceDfs(node *TreeNode, ans *int, pre *int) {
	if node == nil {
		return
	}
	getMinimumDifferenceDfs(node.Left, ans, pre)
	if *pre != -1 && *ans == -1 {
		*ans = node.Val - *pre
	} else if node.Val-*pre < *ans {
		*ans = node.Val - *pre
	}
	*pre = node.Val
	getMinimumDifferenceDfs(node.Right, ans, pre)
}
