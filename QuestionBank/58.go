// @Time     : 2019/5/7 17:58
// @Author   : cancan
// @File     : 58.go
// @Function : 最后一个单词的长度

/*
Question:
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
如果不存在最后一个单词，请返回 0 。
说明：一个单词是指由字母组成，但不包含任何空格的字符串。

Example:
输入: "Hello World"
输出: 5
*/

package QuestionBank

func lengthOfLastWord(s string) int {
	l := len(s)
	ans := 0

	for i := l - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			ans += 1
		} else if ans != 0 {
			return ans
		}
	}

	return ans
}
