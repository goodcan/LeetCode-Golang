/*
 * @Time     : 2020/9/4 15:12
 * @Author   : cancan
 * @File     : 257.go
 * @Function : 二叉树的所有路径
 */

/*
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 * 输入:
 *
 *    1
 *  /   \
 * 2     3
 *  \
 *   5
 *
 * 输出: ["1->2->5", "1->3"]
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
 */

package QuestionBank

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	return dfs(root, []string{})
}

func dfs(node *TreeNode, l []string) []string {
	if node == nil {
		return []string{}
	}
	l = append(l, strconv.Itoa(node.Val))
	if node.Right == nil && node.Left == nil {
		return []string{strings.Join(l, "->")}
	}

	ret := []string{}
	if node.Left != nil {
		ret = append(ret, dfs(node.Left, l)...)
	}

	if node.Right != nil {
		ret = append(ret, dfs(node.Right, l)...)
	}

	return ret
}
