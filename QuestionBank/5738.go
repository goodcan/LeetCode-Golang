// @Time     : 2021/4/25 17:04
// @Author   : cancan
// @File     : 5738.go
// @Function : K 进制表示下的各位数字总和

/*
 * Question:
 * 给你一个整数 n（10 进制）和一个基数 k ，请你将 n 从 10 进制表示转换为 k 进制表示，计算并返回转换后各位数字的 总和 。
 * 转换后，各位数字应当视作是 10 进制数字，且它们的总和也应当按 10 进制表示返回。
 *
 * 示例 1：
 * 输入：n = 34, k = 6
 * 输出：9
 * 解释：34 (10 进制) 在 6 进制下表示为 54 。5 + 4 = 9 。
 *
 * 示例 2：
 * 输入：n = 10, k = 10
 * 输出：1
 * 解释：n 本身就是 10 进制。 1 + 0 = 1 。
 *
 * 提示：
 *     1 <= n <= 100
 *     2 <= k <= 10
 */

package QuestionBank

func sumBase(n int, k int) int {
	ans := 0
	for n != 0 {
		a := n / k
		b := n % k
		for b > k {
			n = a
			a = n / k
			b = n % k
		}
		ans += b
		n = a
	}
	return ans
}
