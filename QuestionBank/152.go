/*
 * @Time     : 2020/5/18 10:58
 * @Author   : cancan
 * @File     : 152.go
 * @Function : 乘积最大子序列
 */

/*
 * Question：
 * 给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。
 *
 * Example 1:
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 *
 * Example 2:
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 */

package QuestionBank

import "math"

func maxProduct(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	mx, mn, ret := nums[0], nums[0], nums[0]
	for i := 1; i < l; i++ {
		newMx := int(math.Max(float64(nums[i]), math.Max(float64(nums[i]*mx), float64(nums[i]*mn))))
		newMn := int(math.Min(float64(nums[i]), math.Min(float64(nums[i]*mx), float64(nums[i]*mn))))
		ret = int(math.Max(float64(ret), float64(newMx)))
		mx = newMx
		mn = newMn
	}
	return ret
}
