/*
 * @Time     : 2019/7/9 18:12
 * @Author   : cancan
 * @File     : 88.go
 * @Function : 合并两个有序数组
 */

/*
 * Question:
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 *
 * Note:
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 * Example:
 * 输入:
 *     nums1 = [1,2,3,0,0,0], m = 3
 *     nums2 = [2,5,6],       n = 3
 * 输出:
 *     [1,2,2,3,5,6]
 */

package QuestionBank

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}
