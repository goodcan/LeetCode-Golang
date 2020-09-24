/*
 * @Time     : 2020/9/24 10:27
 * @Author   : cancan
 * @File     : 501.go
 * @Function : 二叉搜索树中的众数
 */

/*
 * 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
 *
 * 假定 BST 有如下定义：
 * - 结点左子树中所含结点的值小于等于当前结点的值
 * - 结点右子树中所含结点的值大于等于当前结点的值
 * - 左子树和右子树都是二叉搜索树
 *
 * 例如：
 * 给定 BST [1,null,2,2],
 *    1
 *     \
 *      2
 *     /
 *    2
 * 返回[2].
 *
 * 提示：如果众数超过1个，不需考虑输出顺序
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
 */

package QuestionBank

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	data := map[int]int{
		root.Val: 1,
	}
	countMax := 1
	tmp := []*TreeNode{root}

	for len(tmp) > 0 {
		newTmp := []*TreeNode{}
		for _, node := range tmp {
			if node.Left != nil {
				newTmp = append(newTmp, node.Left)
				data[node.Left.Val]++
				if data[node.Left.Val] > countMax {
					countMax = data[node.Left.Val]
				}
			}

			if node.Right != nil {
				newTmp = append(newTmp, node.Right)
				data[node.Right.Val]++
				if data[node.Right.Val] > countMax {
					countMax = data[node.Right.Val]
				}
			}
		}
		tmp = newTmp
	}

	ret := []int{}
	for k, v := range data {
		if v == countMax {
			ret = append(ret, k)
		}
	}

	return ret
}
