// @Time     : 2020/10/15 15:45
// @Author   : cancan
// @File     : 116.go
// @Function : 填充同一层的兄弟节点

/*
 * Question:
 * 给定一个二叉树
 * struct TreeLinkNode {
 *   TreeLinkNode *left;
 *   TreeLinkNode *right;
 *   TreeLinkNode *next;
 * }
 * 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，
 * 则将 next 指针设置为 NULL。
 * 初始状态下，所有 next 指针都被设置为 NULL。
 *
 * Note:
 * 1.你只能使用额外常数空间。
 * 2.使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
 * 3.你可以假设它是一个完美二叉树（即所有叶子节点都在同一层，每个父节点都有两个子节点）。
 *
 * Example:
 * 给定完美二叉树，
 *      1
 *    /  \
 *   2    3
 *  / \  / \
 * 4  5  6  7
 * 调用你的函数后，该完美二叉树变为：
 *      1 -> NULL
 *    /  \
 *   2 -> 3 -> NULL
 *  / \  / \
 * 4->5->6->7 -> NULL
 */

package QuestionBank

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	tmp := []*Node{root}
	for len(tmp) > 0 {
		newTmp := []*Node{}
		l := len(tmp)
		for i := 0; i < l-1; i++ {
			node := tmp[i]
			node.Next = tmp[i+1]
			if node.Left != nil {
				newTmp = append(newTmp, node.Left)
			}
			if node.Right != nil {
				newTmp = append(newTmp, node.Right)
			}
		}

		node := tmp[l-1]
		if node.Left != nil {
			newTmp = append(newTmp, node.Left)
		}
		if node.Right != nil {
			newTmp = append(newTmp, node.Right)
		}

		tmp = newTmp
	}

	return root
}
