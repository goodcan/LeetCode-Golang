/*
 * @Time     : 2020/7/14 20:56
 * @Author   : cancan
 * @File     : 1480.go
 * @Function : 一维数组的动态和
 */

/*
 * Question:
 * 给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
 * 请返回 nums 的动态和。
 *
 * Example 1：
 * 输入：nums = [1,2,3,4]
 * 输出：[1,3,6,10]
 * 解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。
 *
 * Example 2：
 * 输入：nums = [1,1,1,1,1]
 * 输出：[1,2,3,4,5]
 * 解释：动态和计算过程为 [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1] 。
 *
 * Example 3：
 * 输入：nums = [3,1,2,10,1]
 * 输出：[3,4,6,16,17]
 *
 * Note：
 * 1 <= nums.length <= 1000
 * -10^6 <= nums[i] <= 10^6
 */

package QuestionBank

func runningSum(nums []int) []int {
	tmp := 0
	ans := make([]int, len(nums))
	for idx, v := range nums {
		tmp += v
		ans[idx] = tmp
	}
	return ans
}
