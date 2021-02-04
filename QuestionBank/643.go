// @Time     : 2021/2/4 11:20
// @Author   : cancan
// @File     : 643.go
// @Function : 子数组最大平均数 I

/*
 * 给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
 *
 * 示例：
 * 输入：[1,12,-5,-6,50,3], k = 4
 * 输出：12.75
 * 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
 *
 * 提示：
 * - 1 <= k <= n <= 30,000。
 * - 所给数据范围 [-10,000，10,000]。
 */

package QuestionBank

func findMaxAverage(nums []int, k int) float64 {
	var tmp, ans int
	for i := 0; i < k; i++ {
		tmp += nums[i]
		ans += nums[i]
	}

	for i := k; i < len(nums); i++ {
		tmp = tmp - nums[i-k] + nums[i]
		if tmp > ans {
			ans = tmp
		}
	}

	return float64(ans) / float64(k)
}
