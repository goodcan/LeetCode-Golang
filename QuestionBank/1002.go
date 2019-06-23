// @Time     : 2019-06-23 14:06
// @Author   : cancan
// @File     : 1002.py
// @Function : 查找常用字符

/*
Question:
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）
组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。
你可以按任意顺序返回答案。

示例 1：
输入：["bella","label","roller"]
输出：["e","l","l"]

示例 2：
输入：["cool","lock","cook"]
输出：["c","o"]

提示：
1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] 是小写字母
*/

package QuestionBank

func commonChars(A []string) []string {
	tmp1 := make(map[string]int)
	tmp2 := make(map[string]int)
	tmp3 := make(map[string]int)

	pop := len(A) - 1

	for _, v := range A[pop] {
		if _, ok := tmp1[string(v)]; ok {
			tmp1[string(v)]++
		} else {
			tmp1[string(v)] = 1
		}
	}

	A = A[:pop]

	for len(A) != 0 {
		pop = len(A) - 1
		a := A[pop]
		A = A[:pop]

		tmp2 = map[string]int{}
		for _, v := range a {
			if _, ok := tmp2[string(v)]; ok {
				tmp2[string(v)]++
			} else {
				tmp2[string(v)] = 1
			}
		}

		tmp3 = map[string]int{}
		for k, v := range tmp2 {
			if _, ok := tmp1[k]; ok {
				tmp3[k] = tmp1[k]
				if v < tmp1[k] {
					tmp3[k] = v
				}
			}
		}

		tmp1 = tmp3
	}

	ret := []string{}

	for k, v := range tmp1 {
		for i := 0; i < v; i++ {
			ret = append(ret, k)
		}
	}

	return ret
}
