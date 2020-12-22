// @Time     : 2020/12/22 10:06
// @Author   : cancan
// @File     : 103.go
// @Function : 二叉树的锯齿形层序遍历

/*
 * 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 返回锯齿形层序遍历如下：
 * [
 *   [3],
 *   [20,9],
 *   [15,7]
 * ]
 */

package QuestionBank

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int

	if root == nil {
		return ans
	}

	tmp := []*TreeNode{root}
	flag := 1

	for len(tmp) > 0 {
		newTmp := []*TreeNode{}
		ansTMp := []int{}

		for _, node := range tmp {
			if flag == 1 {
				ansTMp = append(ansTMp, node.Val)
			} else {
				ansTMp = append([]int{node.Val}, ansTMp...)
			}

			if node.Left != nil {
				newTmp = append(newTmp, node.Left)
			}

			if node.Right != nil {
				newTmp = append(newTmp, node.Right)
			}
		}

		if flag == 1 {
			flag = 0
		} else {
			flag = 1
		}

		ans = append(ans, ansTMp)
		tmp = newTmp

	}

	return ans
}
