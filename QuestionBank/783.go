/*
 * @Time     : 2020/10/12 16:05
 * @Author   : cancan
 * @File     : 783.go
 * @Function : 二叉搜索树节点最小距离
 */

/*
给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。

示例：
输入: root = [4,2,6,1,3,null,null]
输出: 1
解释:
注意，root是树节点对象(TreeNode object)，而不是数组。

给定的树 [4,2,6,1,3,null,null] 可表示为下图:

          4
        /   \
      2      6
     / \
    1   3

最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。

注意：
1.二叉树的大小范围在 2 到 100。
2.二叉树总是有效的，每个节点的值都是整数，且不重复。
*/

package QuestionBank

func minDiffInBST(root *TreeNode) int {
	ans := -1
	pre := -1
	minDiffInBSTDfs(root, &ans, &pre)
	return ans
}

func minDiffInBSTDfs(node *TreeNode, ans *int, pre *int) {
	if node == nil {
		return
	}
	minDiffInBSTDfs(node.Left, ans, pre)
	if *pre != -1 && *ans == -1 {
		*ans = node.Val - *pre
	} else if node.Val-*pre < *ans {
		*ans = node.Val - *pre
	}
	*pre = node.Val
	minDiffInBSTDfs(node.Right, ans, pre)
}
