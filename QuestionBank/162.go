/*
 * @Time     : 2019/11/28 17:03
 * @Author   : cancan
 * @File     : 162.go
 * @Function : 寻找峰值
 */

/*
 * Question:
 * 峰值元素是指其值大于左右相邻值的元素。
 * 给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。
 * 数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
 * 你可以假设 nums[-1] = nums[n] = -∞。
 *
 * Example 1:
 * 输入: nums = [1,2,3,1]
 * 输出: 2
 * 解释: 3 是峰值元素，你的函数应该返回其索引 2。
 *
 * Example 2:
 * 输入: nums = [1,2,1,3,5,6,4]
 * 输出: 1 或 5
 * 解释: 你的函数可以返回索引 1，其峰值元素为 2；
 *      或者返回索引 5， 其峰值元素为 6。
 *
 * Note:
 * 你的解法应该是 O(logN) 时间复杂度的。
 */

package QuestionBank

func findPeakElement(nums []int) int {
	max := []int{nums[0], 0}
	for i, v := range nums {
		if v > max[0] {
			max[0] = v
			max[1] = i
		}
	}
	return max[1]
}
