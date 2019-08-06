/*
 * @Time     : 2019/8/6 14:12
 * @Author   : cancan
 * @File     : 125.go
 * @Function : 验证回文串
 */

/*
 * Question:
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * Note：
 * 本题中，我们将空字符串定义为有效的回文串。
 *
 * Example 1:
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 * Example 2:
 * 输入: "race a car"
 * 输出: false
 */

package QuestionBank

import "unicode"

func isPalindrome(s string) bool {
	t := []rune{}
	for _, v := range []rune(s) {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			t = append(t, v)
		}
	}

	l := len(t)

	for i, j := 0, l-1; i < l-1; i, j = i+1, j-1 {
		if unicode.ToLower(t[i]) != unicode.ToLower(t[j]) {
			return false
		}
	}

	return true
}
