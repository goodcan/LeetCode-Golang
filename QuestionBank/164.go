// @Time     : 2020/11/26 11:28
// @Author   : cancan
// @File     : 164.go
// @Function : 最大间距

/*
 * Question:
 * 给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
 * 如果数组元素个数小于 2，则返回 0。
 *
 * Example 1:
 * 输入: [3,6,9,1]
 * 输出: 3
 * 解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
 *
 * Example 2:
 * 输入: [10]
 * 输出: 0
 * 解释: 数组元素个数小于 2，因此返回 0。
 *
 * Note:
 * 1.你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
 * 2.请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
 */

package QuestionBank

import "sort"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	ret := 0
	for i := 1; i < len(nums); i++ {
		t := nums[i] - nums[i-1]
		if t > ret {
			ret = t
		}
	}
	return ret
}
