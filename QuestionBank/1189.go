/*
 * @Time     : 2019/9/25 11:47
 * @Author   : cancan
 * @File     : 1189.go
 * @Function : “气球” 的最大数量
 */

/*
 * Question:
 * 给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。
 * 字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。
 *
 * Example 1：
 * 输入：text = "nlaebolko"
 * 输出：1
 *
 * Example 2：
 * 输入：text = "loonbalxballpoon"
 * 输出：2
 *
 * Example 3：
 * 输入：text = "leetcode"
 * 输出：0
 *
 * Note：
 * 1 <= text.length <= 10^4
 * text 全部由小写英文字母组成
 */

package QuestionBank

func maxNumberOfBalloons(text string) int {
	tmp := map[string]int{
		"b": 0, "a": 0, "l": 0, "o": 0, "n": 0,
	}

	for _, v := range text {
		s := string(v)
		if _, ok := tmp[s]; ok {
			tmp[s]++
		}
	}

	tmp["l"] = tmp["l"] / 2
	tmp["o"] = tmp["o"] / 2

	ans := tmp["b"]

	for _, v := range tmp {
		if v < ans {
			ans = v
		}
	}

	return ans
}
