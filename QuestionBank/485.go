/*
 * @Time     : 2021/2/15 13:24
 * @Author   : cancan
 * @File     : 485.go
 * @Function : 最大连续1的个数
 */

/*
 * 给定一个二进制数组， 计算其中最大连续1的个数。
 *
 * 示例 1:
 * 输入: [1,1,0,1,1,1]
 * 输出: 3
 * 解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
 *
 * 注意：
 * 输入的数组只包含 0 和1。
 * 输入数组的长度是正整数，且不超过 10,000。
 */

package QuestionBank

func findMaxConsecutiveOnes(nums []int) int {
	t := map[int]int{1: 0}
	r := t[1]
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v == 1 {
			t[v] += 1
		} else {
			if t[1] > r {
				r = t[1]
			}
			t[1] = 0
		}
	}
	if t[1] > r {
		return t[1]
	} else {
		return r
	}
}
