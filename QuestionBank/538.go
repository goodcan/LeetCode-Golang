/*
 * @Time     : 2020/9/21 15:39
 * @Author   : cancan
 * @File     : 538.go
 * @Function :
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
