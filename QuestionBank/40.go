/*
 * @Time     : 2020/9/10 10:07
 * @Author   : cancan
 * @File     : 40.go
 * @Function : 组合总和 II
 */

/*
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。
 *
 * 示例 1:
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 *   [1, 7],
 *   [1, 2, 5],
 *   [2, 6],
 *   [1, 1, 6]
 * ]
 *
 * 示例 2:
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 *   [1,2,2],
 *   [5]
 * ]
 */

package QuestionBank

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(candidates)
	combinationSum2Dfs(target, 0, len(candidates), 0, candidates, []int{}, &ret)
	return ret
}

func combinationSum2Dfs(target, start, candidatesLen, s int, candidates, l []int, ret *[][]int) {
	if s > target {
		return
	}

	if s == target {
		newL := make([]int, len(l))
		copy(newL, l)
		*ret = append(*ret, newL)
		return
	}

	for i := start; i < candidatesLen; i++ {
		if i-1 >= start && candidates[i-1] == candidates[i] {
			continue
		}
		v := candidates[i]
		combinationSum2Dfs(target, i+1, candidatesLen, s+v, candidates, append(l, v), ret)
	}
}
