// @Time     : 2019/4/28 10:10
// @Author   : cancan
// @File     : 50.go
// @Function : Pow(x, n)

/*
 * Question:
 * 实现 pow(x, n) ，即计算 x 的 n 次幂函数。
 *
 * Example 1:
 * 输入: 2.00000, 10
 * 输出: 1024.00000
 *
 * Example 2:
 * 输入: 2.10000, 3
 * 输出: 9.26100
 *
 * Example 3:
 * 输入: 2.00000, -2
 * 输出: 0.25000
 * 解释: 2-2 = 1/22 = 1/4 = 0.25
 *
 * Note:
 * -100.0 < x < 100.0
 * n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
 */

package QuestionBank

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	ans := 1.0

	for n > 0 {
		if n%2 == 1 {
			ans *= x
		}

		x *= x
		n /= 2
	}

	return ans
}
