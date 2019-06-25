/*
 * @Time     : 2019/4/28 16:04
 * @Author   : cancan
 * @File     : 100.go
 * @Function : 相同的树
 */

/*
 * 给定两个二叉树，编写一个函数来检验它们是否相同。
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 *
 * Example 1:
 * 输入:       1         1
 *           / \       / \
 *          2   3     2   3
 *         [1,2,3],   [1,2,3]
 * 输出: true
 *
 * Example 2:
 * 输入:      1          1
 *           /           \
 *          2             2
 *         [1,2],     [1,null,2]
 * 输出: false
 *
 * Example 3:
 * 输入:       1         1
 *           / \       / \
 *          2   1     1   2
 *         [1,2,1],   [1,1,2]
 * 输出: false
 */

package QuestionBank

func isSameTree(p *TreeNode, q *TreeNode) bool {
	t1 := []*TreeNode{}
	t2 := []*TreeNode{}

	if p != nil && q != nil {
		t1 = append(t1, p)
		t2 = append(t2, q)
	} else if p == nil && q == nil {
		return true
	} else {
		return false
	}

	for len(t1) != 0 {
		tt1 := []*TreeNode{}
		tt2 := []*TreeNode{}

		for i, v1 := range t1 {
			v2 := t2[i]
			if v1.Val != v2.Val {
				return false
			}

			if v1.Left != nil && v2.Left != nil {
				tt1 = append(tt1, v1.Left)
				tt2 = append(tt2, v2.Left)
			} else if v1.Left == nil && v2.Left == nil {

			} else {
				return false
			}

			if v1.Right != nil && v2.Right != nil {
				tt1 = append(tt1, v1.Right)
				tt2 = append(tt2, v2.Right)
			} else if v1.Right == nil && v2.Right == nil {

			} else {
				return false
			}
		}

		t1 = tt1
		t2 = tt2
	}

	return true
}
