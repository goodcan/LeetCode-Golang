/*
 * @Time     : 2019/5/5 17:58
 * @Author   : cancan
 * @File     : 442..go
 * @Function : 数组中重复的数据
 */

/*
 * Question:
 * 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
 * 找到所有出现两次的元素。
 * 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
 *
 * Example:
 * 输入: [4,3,2,7,8,2,3,1]
 * 输出: [2,3]
 */

package QuestionBank

func findDuplicates(nums []int) []int {
	tmp := make(map[int]int)
	ret := []int{}

	for _, v := range nums {
		n, ok := tmp[v]
		if ok {
			if n == 1 {
				ret = append(ret, v)
				tmp[v]++
			}
		} else {
			tmp[v] = 1
		}
	}

	return ret
}
