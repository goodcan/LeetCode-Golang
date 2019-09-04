/*
 * @Time     : 2019/9/4 16:52
 * @Author   : cancan
 * @File     : 1006.go
 * @Function : 笨阶乘
 */

/*
 * Question:
 * 通常，正整数 n 的阶乘是所有小于或等于 n 的正整数的乘积。
 * 例如，factorial(10) = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1。
 * 相反，我们设计了一个笨阶乘 clumsy：在整数的递减序列中，
 * 我们以一个固定顺序的操作符序列来依次替换原有的乘法操作符：乘法(*)，除法(/)，加法(+)和减法(-)。
 * 例如，clumsy(10) = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1。
 * 然而，这些运算仍然使用通常的算术运算顺序：我们在任何加、减步骤之前执行所有的乘法和除法步骤，并且按从左到右处理乘法和除法步骤。
 * 另外，我们使用的除法是地板除法（floor division），所以 10 * 9 / 8 等于 11。这保证结果是一个整数。
 * 实现上面定义的笨函数：给定一个整数 N，它返回 N 的笨阶乘。
 *
 * Example 1：
 * 输入：4
 * 输出：7
 * 解释：7 = 4 * 3 / 2 + 1
 *
 * Example 2：
 * 输入：10
 * 输出：12
 * 解释：12 = 10 * 9 / 8 + 7 - 6 * 5 / 4 + 3 - 2 * 1
 *
 * Note：
 * 1 <= N <= 10000
 * -2^31 <= answer <= 2^31 - 1  （答案保证符合 32 位整数。）
 */

package QuestionBank

func clumsy(N int) int {
	l := createSlice(N)

	d := N % 4
	ans := 0

	for i := 0; i < N-d; i += 4 {
		tmp := int(l[i] * l[i+1] / l[i+2])
		if ans == 0 {
			ans += tmp + l[i+3]
		} else {
			ans -= tmp - l[i+3]
		}
	}

	if d != 0 {
		tmp := l[N-d]
		d--
		if d != 0 {
			tmp *= l[N-d]
			d--
			if d != 0 {
				tmp = int(tmp / l[N-d])
			}
		}

		if ans == 0 {
			ans += tmp
		} else {
			ans -= tmp
		}
	}

	return ans
}

func createSlice(N int) []int {
	var ret []int
	for N != 0 {
		ret = append(ret, N)
		N--
	}
	return ret
}
