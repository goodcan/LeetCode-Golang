/*
 * @Time     : 2019/8/27 17:03
 * @Author   : cancan
 * @File     : 1137.go
 * @Function : 第 N 个泰波那契数
 */

/*
 * Question:
 * 泰波那契序列 Tn 定义如下：
 * T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2
 * 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。
 *
 * Example 1：
 * 输入：n = 4
 * 输出：4
 * 解释：
 * T_3 = 0 + 1 + 1 = 2
 * T_4 = 1 + 1 + 2 = 4
 *
 * Example 2：
 * 输入：n = 25
 * 输出：1389537
 *
 * Note：
 * 0 <= n <= 37
 * 答案保证是一个 32 位整数，即 answer <= 2^31 - 1。
 */

package QuestionBank

func tribonacci(n int) int {
	tmp := [3]int{0, 1, 1}

	if n <= 2 {
		return tmp[n]
	}

	n -= 2

	for n > 0 {
		tmp[0], tmp[1], tmp[2] = tmp[1], tmp[2], tmp[0]+tmp[1]+tmp[2]
		n--
	}

	return tmp[2]
}
