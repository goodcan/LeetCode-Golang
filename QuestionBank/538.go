/*
 * @Time     : 2020/9/21 15:39
 * @Author   : cancan
 * @File     : 538.go
 * @Function : 把二叉搜索树转换为累加树
 */

/*
 * 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
 *
 * 例如：
 * 输入: 原始二叉搜索树:
 *           5
 *         /   \
 *        2     13
 *
 * 输出: 转换为累加树:
 *          18
 *         /   \
 *       20     13
 */

package QuestionBank

func convertBST(root *TreeNode) *TreeNode {
	l := convertBSTDfs(root)
	if len(l) <= 1 {
		return root
	}

	s := l[0].Val
	for idx := 1; idx < len(l); idx++ {
		s += l[idx].Val
		l[idx].Val = s
	}

	return root
}

func convertBSTDfs(root *TreeNode) []*TreeNode {
	ret := []*TreeNode{}
	if root == nil {
		return ret
	}
	ret = append(ret, convertBSTDfs(root.Right)...)
	ret = append(ret, root)
	ret = append(ret, convertBSTDfs(root.Left)...)
	return ret
}
