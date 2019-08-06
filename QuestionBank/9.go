/*
 * @Time     : 2019/5/10 16:07
 * @Author   : cancan
 * @File     : 9.go
 * @Function : 回文数
 */

/*
 * Question:
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * Example 1:
 * 输入: 121
 * 输出: true
 *
 * Example 2:
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 * Example 3:
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 * Follow up:
 * 你能不将整数转为字符串来解决这个问题吗？
 */

package QuestionBank

import "strconv"

func isPalindrome9(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	end := len(s) - 1
	start := 0

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
