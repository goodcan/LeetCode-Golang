/*
 * @Time     : 2022/1/1 11:04
 * @Author   : cancan
 * @File     : 387.go
 * @Function : 字符串中的第一个唯一字符
 */

/*
 * Question:
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 *
 * Example:
 * s = "leetcode"
 * 返回 0.
 * s = "loveleetcode",
 * 返回 2.
 *
 * Note:
 * 您可以假定该字符串只包含小写字母。
 */

package QuestionBank

import "strings"

func firstUniqChar(s string) int {
	t := map[string]int{}

	for index, temp := range s {
		i := string(temp)
		_, ok := t[i]
		if !ok {
			t[i] = 0
			if strings.Count(s, i) == 1 {
				return index
			}
		} else {
			continue
		}
	}

	return -1
}
