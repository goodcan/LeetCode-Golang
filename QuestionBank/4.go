// @Time     : 2019/4/26 11:49
// @Author   : cancan
// @File     : 4.go
// @Function : 寻找两个有序数组的中位数

/*
 * Question:
 * 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
 * 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
 * 你可以假设 nums1 和 nums2 不会同时为空。
 *
 * Example 1:
 * nums1 = [1, 3]
 * nums2 = [2]
 * 则中位数是 2.0
 *
 * Example 2:
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * 则中位数是 (2 + 3)/2 = 2.5
 */

package QuestionBank

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)

	sort.Ints(nums)

	l := len(nums)

	if l == 0 {
		return 0
	}

	if l%2 == 0 {
		return float64(nums[l/2-1]+nums[l/2]) / 2
	} else {
		return float64(nums[l/2])
	}
}
