/*
 * @Time     : 2020/5/19 00:17
 * @Author   : cancan
 * @File     : 680.go
 * @Function : 验证回文字符串 Ⅱ
 */

/*
 * Question:
 * 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
 *
 * Example 1:
 * 输入: "aba"
 * 输出: True
 *
 * Example 2:
 * 输入: "abca"
 * 输出: True
 * 解释: 你可以删除c字符。
 *
 * Note:
 * 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
 */

package QuestionBank

func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			return validPalindromeCheck(s, low, high-1) || validPalindromeCheck(s, low+1, high)
		}
	}
	return true
}

func validPalindromeCheck(s string, low, high int) bool {
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}
