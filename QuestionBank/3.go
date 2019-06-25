// @Time     : 2019/4/26 10:44
// @Author   : cancan
// @File     : 3.go
// @Function : 无重复字符的最长子串

/*
 * Question:
 * 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 * Example 1:
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 * Example 2:
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 * Example 3:
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 *      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 */

package QuestionBank

func lengthOfLongestSubstring(s string) int {
	left := 0
	ans := 0
	temp := make(map[string]int)

	for index, value := range s {
		d, ok := temp[string(value)]

		if ok && d >= left {
			left = d + 1
		}

		temp[string(value)] = index

		if index-left+1 > ans {
			ans = index - left + 1
		}
	}

	return ans
}
