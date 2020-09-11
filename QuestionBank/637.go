/*
 * @Time     : 2020/9/12 00:27
 * @Author   : cancan
 * @File     : 637.go
 * @Function : 二叉树的层平均值
 */

/*
 * 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
 *
 * 示例 1：
 * 输入：
 *     3
 *    / \
 *   9  20
 *     /  \
 *    15   7
 * 输出：[3, 14.5, 11]
 * 解释：
 * 第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。
 *
 * 提示：
 * - 节点值的范围在32位有符号整数范围内。
 */

package QuestionBank

func averageOfLevels(root *TreeNode) []float64 {
	ret := []float64{}
	if root == nil {
		return ret
	}

	tmp := []*TreeNode{root}

	for len(tmp) > 0 {
		newTmp := []*TreeNode{}
		s := 0
		for _, node := range tmp {
			s += node.Val
			if node.Left != nil {
				newTmp = append(newTmp, node.Left)
			}
			if node.Right != nil {
				newTmp = append(newTmp, node.Right)
			}
		}
		ret = append(ret, float64(s)/float64(len(tmp)))
		tmp = newTmp
	}

	return ret
}
