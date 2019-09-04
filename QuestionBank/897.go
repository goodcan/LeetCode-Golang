/*
 * @Time     : 2019/9/4 10:10
 * @Author   : cancan
 * @File     : 897.go
 * @Function : 递增顺序查找树
 */

/*
 * Question:
 * 给定一个树，按中序遍历重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。
 *
 * Example:
 * 输入：[5,3,6,2,4,null,8,1,null,null,null,7,9]
 *
 *        5
 *       / \
 *     3    6
 *    / \    \
 *   2   4    8
 *  /        / \
 * 1        7   9
 *
 * 输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
 *
 *  1
 *   \
 *    2
 *     \
 *      3
 *       \
 *        4
 *         \
 *          5
 *           \
 *            6
 *             \
 *              7
 *               \
 *                8
 *                 \
 *                  9
 *
 * Note:
 * 给定树中的结点数介于 1 和 100 之间。
 * 每个结 * 点都有一个从 0 到 1000 范围内的唯一整数值。
 */

package QuestionBank

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	l := traversal897(root)

	ret := &TreeNode{Val: l[0]}
	tmp := ret

	for _, v := range l[1:] {
		tmp.Right = &TreeNode{Val: v}
		tmp = tmp.Right
	}

	return ret
}

func traversal897(root *TreeNode) []int {
	var r []int

	if root == nil {
		return r
	}

	r = append(r, traversal897(root.Left)...)
	r = append(r, root.Val)
	r = append(r, traversal897(root.Right)...)

	return r
}
