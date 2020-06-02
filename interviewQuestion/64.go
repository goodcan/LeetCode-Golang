/*
 * @Time     : 2020/6/2 11:47
 * @Author   : cancan
 * @File     : 64.go
 * @Function : 求1+2+…+n
 */

/*
Question:
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

Example 1：
输入: n = 3
输出: 6

Example 2：
输入: n = 9
输出: 45

Note：
1 <= n <= 10000
*/

package interviewQuestion

func sumNums(n int) int {
	return (1 + n) * n / 2
}
