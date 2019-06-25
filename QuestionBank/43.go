/*
 * @Time     : 2019/5/7 16:16
 * @Author   : cancan
 * @File     : 43.go
 * @Function : 字符串相乘
 */

/*
 * Question:
 * 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
 *
 * Example 1:
 * 输入: num1 = "2", num2 = "3"
 * 输出: "6"
 *
 * Example 2:
 * 输入: num1 = "123", num2 = "456"
 * 输出: "56088"
 *
 * Note：
 * 1.num1 和 num2 的长度小于110。
 * 2.num1 和 num2 只包含数字 0-9。
 * 3.num1 和 num2 均不以零开头，除非是数字 0 本身。
 * 4.不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 */

package QuestionBank

import "bytes"

func multiply(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	r := make([]int16, l1+l2)

	for i1 := l1 - 1; i1 >= 0; i1-- {
		idx := l2 + i1
		for i2 := l2 - 1; i2 >= 0; i2-- {
			r[idx] += int16(num1[i1]-'0') * int16(num2[i2]-'0')
			idx--
		}
	}

	for idx := len(r) - 1; idx >= 0; idx-- {
		if r[idx] > 9 {
			r[idx-1] += (r[idx] / 10)
			r[idx] %= 10
		}
	}

	var buf bytes.Buffer
	for _, v := range r {
		if buf.Len() > 0 || v != 0 {
			buf.WriteByte(byte(v) + '0')
		}
	}

	if buf.Len() == 0 {
		buf.WriteByte('0')
	}

	return buf.String()
}
