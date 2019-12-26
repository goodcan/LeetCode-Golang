/*
 * @Time     : 2019/12/26 19:28
 * @Author   : cancan
 * @File     : 6.go
 * @Function : Z 字形变换
 */

/*
 * Question:
 * 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
 * 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
 *     L   C   I   R
 *     E T O E S I I G
 *     E   D   H   N
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
 * 请你实现这个将字符串进行指定行数变换的函数：
 * string convert(string s, int numRows);
 *
 * Example 1:
 * 输入: s = "LEETCODEISHIRING", numRows = 3
 * 输出: "LCIRETOESIIGEDHN"
 *
 * Example 2:
 * 输入: s = "LEETCODEISHIRING", numRows = 4
 * 输出: "LDREOEIIECIHNTSG"
 * 解释:
 *     L     D     R
 *     E   O E   I I
 *     E C   I H   N
 *     T     S     G
 */

package QuestionBank

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	tmp := [][]string{}
	for i := 0; i < numRows; i++ {
		tmp = append(tmp, []string{})
	}

	var idx int
	flag1 := 0
	flag2 := 0

	for _, v := range s {
		if (flag2 == 0 && flag1 == numRows) || (flag2 == 1 && flag1 == numRows-2) {
			flag1 = 0
			if numRows > 2 {
				if flag2 == 1 {
					flag2 = 0
				} else {
					flag2 = 1
				}
			}
		}

		if flag2 == 0 {
			idx = flag1
		} else {
			idx = numRows - flag1 - 2
		}
		tmp[idx] = append(tmp[idx], string(v))

		flag1++
	}

	ret := ""
	for _, v := range tmp {
		ret += strings.Join(v, "")
	}

	return ret
}
