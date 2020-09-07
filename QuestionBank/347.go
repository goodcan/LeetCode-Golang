/*
 * @Time     : 2020/9/7 23:25
 * @Author   : cancan
 * @File     : 347.go
 * @Function : 前 K 个高频元素
 */

/*
 * 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
 *
 * 示例 1:
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 *
 * 示例 2:
 * 输入: nums = [1], k = 1
 * 输出: [1]
 *
 * 提示：
 * - 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
 * - 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
 * - 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
 * - 你可以按任意顺序返回答案。
 */

package QuestionBank

import "sort"

type topKFrequentSort [][2]int

func (t topKFrequentSort) Len() int {
	return len(t)
}

func (t topKFrequentSort) Less(i, j int) bool {
	return t[i][1] > t[j][1]
}

func (t topKFrequentSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func topKFrequent(nums []int, k int) []int {
	tmp := make(map[int]int)
	for _, v := range nums {
		tmp[v]++
	}

	sortData := topKFrequentSort{}
	for k, v := range tmp {
		sortData = append(sortData, [2]int{k, v})
	}

	sort.Sort(sortData)

	ret := []int{}
	for _, item := range sortData[:k] {
		ret = append(ret, item[0])
	}

	return ret
}
