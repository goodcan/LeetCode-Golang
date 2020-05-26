/*
 * @Time     : 2020/5/26 14:31
 * @Author   : cancan
 * @File     : 287.go
 * @Function : 寻找重复数
 */

/*
 * Question:
 * 给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），
 * 可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
 *
 * Example 1:
 * 输入: [1,3,4,2,2]
 * 输出: 2
 *
 * Example 2:
 * 输入: [3,1,3,4,2]
 * 输出: 3
 *
 * Note：
 * 1.不能更改原数组（假设数组是只读的）。
 * 2.只能使用额外的 O(1) 的空间。
 * 3.时间复杂度小于 O(n2) 。
 * 4.数组中只有一个重复的数字，但它可能不止重复出现一次。
 */

package QuestionBank

func findDuplicate(nums []int) int {
	t := make(map[int]int)
	for _, v := range nums {
		if _, ok := t[v]; ok {
			return v
		} else {
			t[v] = v
		}
	}
	return -1
}
