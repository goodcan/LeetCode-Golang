// @Time     : 2021/5/8 10:38
// @Author   : cancan
// @File     : 17_10.go
// @Function : 主要元素

/*
 * 数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。
 *
 * 示例 1：
 * 输入：[1,2,5,9,5,9,5,5,5]
 * 输出：5
 *
 * 示例 2：
 * 输入：[3,2]
 * 输出：-1
 *
 * 示例 3：
 * 输入：[2,2,1,1,1,2,2]
 * 输出：2
 */

package interviewQuestion

func majorityElement(nums []int) int {
	n := len(nums) / 2
	tmp := make(map[int]int)
	for _, v := range nums {
		tmp[v]++
		if tmp[v] > n {
			return v
		}
	}
	return -1
}
