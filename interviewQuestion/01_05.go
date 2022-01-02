/*
 * @Time     : 2022/1/3 00:11
 * @Author   : cancan
 * @File     : 01.05.go
 * @Function : 一次编辑
 */

/*
 * Question:
 * 字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
 *
 * Example 1:
 * 输入:
 * first = "pale"
 * second = "ple"
 * 输出: True
 *
 * Example 2:
 * 输入:
 * first = "pales"
 * second = "pal"
 * 输出: False
 */

package interviewQuestion

import "math"

func oneEditAway(first string, second string) bool {
	l1 := len(first)
	l2 := len(second)

	if math.Abs(float64(l1-l2)) > 1 {
		return false
	}

	if l1 == l2 {
		fn := 0
		for idx := 0; idx < l1; idx++ {
			if fn > 1 {
				return false
			}
			if first[idx] != second[idx] {
				fn++
			}
		}
		return fn <= 1
	} else if l1 > l2 {
		for idx := 0; idx < l2; idx++ {
			if first[idx] != second[idx] {
				return first[idx+1:] == second[idx:]
			}
		}
		return true
	} else if l1 < l2 {
		for idx := 0; idx < l1; idx++ {
			if first[idx] != second[idx] {
				return first[idx:] == second[idx+1:]
			}
		}
		return true
	}

	return true
}
