/*
 * @Time     : 2020/9/9 18:41
 * @Author   : cancan
 * @File     : 77.go
 * @Function : 组合
 */

/*
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 *
 * 示例:
 * 输入: n = 4, k = 2
 * 输出:
 * [
 *   [2,4],
 *   [3,4],
 *   [2,3],
 *   [1,2],
 *   [1,3],
 *   [1,4],
 * ]
 */

package QuestionBank

func combine(n int, k int) [][]int {
	ret := [][]int{}
	combineDfs(1, n, k, []int{}, &ret)
	return ret
}

func combineDfs(start, n, k int, l []int, ret *[][]int) {
	if len(l) > k {
		return
	}

	if len(l) == k {
		newL := make([]int, len(l))
		copy(newL, l)
		*ret = append(*ret, newL)
		return
	}

	for i := start; i <= n; i++ {
		combineDfs(i+1, n, k, append(l, i), ret)
	}
}
