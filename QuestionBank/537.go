/*
 * @Time     : 2019/9/3 18:24
 * @Author   : cancan
 * @File     : 537.go
 * @Function : 复数乘法
 */

/*
 * Question:
 * 给定两个表示复数的字符串。
 * 返回表示它们乘积的字符串。注意，根据定义 i2 = -1 。
 *
 * Example 1:
 * 输入: "1+1i", "1+1i"
 * 输出: "0+2i"
 * 解释: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
 *
 * Example 2:
 * 输入: "1+-1i", "1+-1i"
 * 输出: "0+-2i"
 * 解释: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。
 *
 * Note:
 * 输入字符串不包含额外的空格。
 * 输入字符串将以 a+bi 的形式给出，其中整数 a 和 b 的范围均在 [-100, 100] 之间。输出也应当符合这种形式。
 */

package QuestionBank

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	aSplit := strings.Split(a, "+")
	bSplit := strings.Split(b, "+")

	aX, _ := strconv.Atoi(aSplit[0])
	aY, _ := strconv.Atoi(aSplit[1][:len(aSplit[1])-1])

	bX, _ := strconv.Atoi(bSplit[0])
	bY, _ := strconv.Atoi(bSplit[1][:len(bSplit[1])-1])

	ansX := strconv.Itoa(aX*bX - aY*bY)
	ansY := strconv.Itoa(aX*bY + aY*bX)

	return ansX + "+" + ansY + "i"
}
