/*
 * @Time     : 2019/10/8 11:12
 * @Author   : cancan
 * @File     : 5205.go
 * @Function : 独一无二的出现次数
 */

/*
 * Question:
 * 给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。
 * 如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
 *
 * Example 1：
 * 输入：arr = [1,2,2,1,1,3]
 * 输出：true
 * 解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。
 *
 * Example 2：
 * 输入：arr = [1,2]
 * 输出：false
 *
 * Example 3：
 * 输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
 * 输出：true
 *
 * Note：
 * 1 <= arr.length <= 1000
 * -1000 <= arr[i] <= 1000
 */

package QuestionBank

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int)
	check := make(map[int]int)

	for _, v := range arr {
		counter[v]++
	}

	for k, v := range counter {
		if _, ok := check[v]; ok {
			return false
		} else {
			check[v] = k
		}
	}

	return true
}
