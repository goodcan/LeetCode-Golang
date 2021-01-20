// @Time     : 2021/1/20 10:41
// @Author   : cancan
// @File     : 628.go
// @Function : 三个数的最大乘积

/*
 * 给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
 *
 * 示例 1:
 * 输入: [1,2,3]
 * 输出: 6
 *
 * 示例 2:
 * 输入: [1,2,3,4]
 * 输出: 24
 *
 * 注意:
 * - 给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
 * - 输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
 */

package QuestionBank

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	tmp1 := nums[0] * nums[1] * nums[l-1]
	tmp2 := nums[l-1] * nums[l-2] * nums[l-3]
	if tmp1 >= tmp2 {
		return tmp1
	} else {
		return tmp2
	}
}
