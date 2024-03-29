/*
 * @Time     : 2022/2/6 17:45
 * @Author   : cancan
 * @File     : 520.go
 * @Function : 检测大写字母
 */

/*
 * Question:
 * 给定一个单词，你需要判断单词的大写使用是否正确。
 * 我们定义，在以下情况时，单词的大写用法是正确的：
 *  1.全部字母都是大写，比如"USA"。
 *  2.单词中所有字母都不是大写，比如"leetcode"。
 *  3.如果单词不只含有一个字母，只有首字母大写， 比如 "Google"。
 * 否则，我们定义这个单词没有正确使用大写字母。
 *
 * Example 1:
 * 输入: "USA"
 * 输出: True
 *
 * Example 2:
 * 输入: "FlaG"
 * 输出: False
 *
 * Note:
 * 输入是由大写和小写拉丁字母组成的非空单词。
 */

package QuestionBank

func detectCapitalUse(word string) bool {
	upper := 0
	for _, v := range word {
		if v >= 'A' && v <= 'Z' {
			upper++
		}
	}
	if upper == 0 || (upper == 1 && word[0] >= 'A' && word[0] <= 'Z') || upper == len(word) {
		return true
	}

	return false
}
