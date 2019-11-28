/*
 * @Time     : 2019/11/28 15:23
 * @Author   : cancan
 * @File     : 704.go
 * @Function : 二分查找
 */

/*
 * Question:
 * 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
 * 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
 *
 * Example 1:
 * 输入: nums = [-1,0,3,5,9,12], target = 9
 * 输出: 4
 * 解释: 9 出现在 nums 中并且下标为 4
 *
 * Example 2:
 * 输入: nums = [-1,0,3,5,9,12], target = 2
 * 输出: -1
 * 解释: 2 不存在 nums 中因此返回 -1
 *
 * NOte:
 * 1.你可以假设 nums 中的所有元素是不重复的。
 * 2.n 将在 [1, 10000]之间。
 * 3.nums 的每个元素都将在 [-9999, 9999]之间。
 */

package QuestionBank

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			right = mid - 1
		case nums[mid] < target:
			left = mid + 1
		}
	}
	return -1
}
