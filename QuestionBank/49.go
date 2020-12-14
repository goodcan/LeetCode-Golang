// @Time     : 2020/12/14 9:29
// @Author   : cancan
// @File     : 49.go
// @Function : 字母异位词分组

/*
 * Question:
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * Example:
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * 输出:
 * [
 *   ["ate","eat","tea"],
 *   ["nat","tan"],
 *   ["bat"]
 * ]
 *
 * Note
 * 1.所有输入均为小写字母。
 * 2.不考虑答案输出的顺序。
 */

package QuestionBank

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	data := make(map[string][]string)
	for _, item := range strs {
		tmp := []string{}
		for _, s := range item {
			tmp = append(tmp, string(s))
		}
		sort.Strings(tmp)
		newStr := strings.Join(tmp, "")
		data[newStr] = append(data[newStr], item)
	}
	ret := make([][]string, 0, len(strs))
	for _, ss := range data {
		ret = append(ret, ss)
	}
	return ret
}
