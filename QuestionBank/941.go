// @Time     : 2020/11/3 10:28
// @Author   : cancan
// @File     : 941.go
// @Function : 有效的山脉数组

/*
 * 给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。
 * 让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：
 *
 * A.length >= 3
 * 在 0 < i A.length - 1 条件下，存在 i 使得：
 *     A[0] < A[1] < ... A[i-1] < A[i]
 *     A[i] > A[i+1] > ... > A[A.length - 1]
 *
 * 示例 1：
 * 输入：[2,1]
 * 输出：false
 *
 * 示例 2：
 * 输入：[3,5,5]
 * 输出：false
 *
 * 示例 3：
 * 输入：[0,3,2,1]
 * 输出：true
 *
 * 提示：
 * 0 <= A.length <= 10000
 * 0 <= A[i] <= 10000
 */

package QuestionBank

func validMountainArray(A []int) bool {
	l := len(A)
	if l <= 2 {
		return false
	}

	// 1 上升
	// 2 下降
	flag := 1
	count1 := 0 // 上升的数量

	for idx := 1; idx < l; idx++ {
		if flag == 1 && A[idx-1] >= A[idx] {
			flag = 2
		}
		if flag == 1 {
			count1++
		}
		if flag == 2 && A[idx-1] <= A[idx] {
			return false
		}
	}

	if count1 == 0 || flag == 1 {
		return false
	}

	return true
}
