/*
 * @Time     : 2019/7/9 18:02
 * @Author   : cancan
 * @File     : 215.go
 * @Function : 数组中的第K个最大元素
 */

/*
 * Question:
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * Example 1:
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 * Example 2:
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 * Note:
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 */

package QuestionBank

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
