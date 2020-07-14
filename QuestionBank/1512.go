/*
 * @Time     : 2020/7/14 20:18
 * @Author   : cancan
 * @File     : 1512.go
 * @Function : 好数对的数目
 */

/*
 * Question:
 * 给你一个整数数组 nums 。
 * 如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。
 * 返回好数对的数目。
 *
 * Example 1：
 * 输入：nums = [1,2,3,1,1,3]
 * 输出：4
 * 解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始
 *
 * Example 2：
 * 输入：nums = [1,1,1,1]
 * 输出：6
 * 解释：数组中的每组数字都是好数对
 *
 * Example 3：
 * 输入：nums = [1,2,3]
 * 输出：0
 *
 * Note
 * 1 <= nums.length <= 100
 * 1 <= nums[i] <= 100
 */

package QuestionBank

func numIdenticalPairs(nums []int) int {
	tmp := make(map[int]int)
	ans := 0
	for _, v := range nums {
		ans += tmp[v]
		tmp[v]++
	}
	return ans
}
