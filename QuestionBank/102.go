/*
 * @Time     : 2020/5/13 22:21
 * @Author   : cancan
 * @File     : 102.go
 * @Function : 二叉树的层序遍历
 */

/*
 * Question:
 * 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 * Example:
 * 给定二叉树: [3,9,20,null,null,15,7],
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 返回其层次遍历结果：
 * [
 *   [3],
 *   [9,20],
 *   [15,7]
 * ]
 */

package QuestionBank

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	tmp := []*TreeNode{root}
	for len(tmp) > 0 {
		newTmp := []*TreeNode{}
		layer := []int{}
		for _, item := range tmp {
			layer = append(layer, item.Val)
			if item.Left != nil {
				newTmp = append(newTmp, item.Left)
			}
			if item.Right != nil {
				newTmp = append(newTmp, item.Right)
			}
		}
		ret = append(ret, layer)
		tmp = newTmp
	}
	return ret
}
