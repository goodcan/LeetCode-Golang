/*
 * @Time     : 2020/2/21 11:49
 * @Author   : cancan
 * @File     : 01.01.go
 * @Function : 判定字符是否唯一
 */

/*
 * Question:
 * 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
 *
 * Example 1：
 * 输入: s = "leetcode"
 * 输出: false
 *
 * Example 2：
 * 输入: s = "abc"
 * 输出: true
 *
 * Note：
 * 1. 0 <= len(s) <= 100
 * 2. 如果你不使用额外的数据结构，会很加分。
 */

package interviewQuestion

func isUnique(astr string) bool {
	tmp := make(map[string]bool)
	for _, v := range astr {
		if _, ok := tmp[string(v)]; ok {
			return false
		} else {
			tmp[string(v)] = true
		}
	}
	return true
}
