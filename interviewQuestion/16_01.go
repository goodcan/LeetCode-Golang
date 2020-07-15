/*
 * @Time     : 2020/7/15 21:33
 * @Author   : cancan
 * @File     : 16_01.go
 * @Function : 交换数字
 */

/*
 * Question:
 * 编写一个函数，不用临时变量，直接交换numbers = [a, b]中a与b的值。
 *
 * Example：
 * 输入: numbers = [1,2]
 * 输出: [2,1]
 *
 * Note：
 *     numbers.length == 2
 */

package interviewQuestion

func swapNumbers(numbers []int) []int {
	numbers[0], numbers[1] = numbers[1], numbers[0]
	return numbers
}
