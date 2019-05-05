// @Time     : 2019/5/5 18:46
// @Author   : cancan
// @File     : 771.go
// @Function : 宝石与石头

/*
Question:
给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。
S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

Example 1:
输入: J = "aA", S = "aAAbbbb"
输出: 3

Example 2:
输入: J = "z", S = "ZZ"
输出: 0

Note:
S 和 J 最多含有50个字母。
J 中的字符不重复。
*/

package QuestionBank

func numJewelsInStones(J string, S string) int {
	tmp := make(map[string]int)
	ans := 0

	for _, v := range J {
		tmp[string(v)] = 1
	}

	for _, v := range S {
		_, ok := tmp[string(v)]

		if ok {
			ans++
		}
	}

	return ans
}
