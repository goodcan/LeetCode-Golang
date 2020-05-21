/*
 * @Time     : 2020/5/21 11:04
 * @Author   : cancan
 * @File     : 5.go
 * @Function : 最长回文子串
 */

/*
 * Question:
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 *
 * Example 1：
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 *
 * Example 2：
 * 输入: "cbbd"
 * 输出: "bb"
 */

package QuestionBank

func longestPalindrome(s string) string {
	l := len(s)
	maxL := 0
	i := 0
	var ret string

	for l-i > maxL/2 {
		left := i
		right := i
		for right+1 < l && s[right+1] == s[i] {
			right++
		}
		i = right
		for left >= 0 && right < l && s[left] == s[right] {
			if right-left+1 > maxL {
				ret = s[left : right+1]
				maxL = right - left + 1
			}
			left--
			right++
		}
		i++
	}
	return ret
}
