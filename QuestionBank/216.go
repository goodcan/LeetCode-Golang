/*
 * @Time     : 2020/9/11 0:56
 * @Author   : cancan
 * @File     : 216.go
 * @Function : 组合总和 III
 */

/*
 * 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
 *
 * 说明：
 * - 所有数字都是正整数。
 * - 解集不能包含重复的组合。
 *
 * 示例 1:
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 *
 * 示例 2:
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 */

package QuestionBank

func combinationSum3(k int, n int) [][]int {
	ret := [][]int{}
	combinationSum3Dfs(1, 0, k, n, []int{}, &ret)
	return ret
}

func combinationSum3Dfs(start, s, k, n int, l []int, ret *[][]int) {
	if len(l) > k {
		return
	}
	if s > n {
		return
	}
	if len(l) == k && s == n {
		newL := make([]int, len(l))
		copy(newL, l)
		*ret = append(*ret, newL)
		return
	}

	for i := start; i < 10; i++ {
		combinationSum3Dfs(i+1, s+i, k, n, append(l, i), ret)
	}
}
