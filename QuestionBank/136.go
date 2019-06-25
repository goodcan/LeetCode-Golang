// @Time     : 2019/4/29 16:27
// @Author   : cancan
// @File     : 136.go
// @Function : 只出现一次的数字

/*
 * Question:
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *
 * Note：
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * Example 1:
 * 输入: [2,2,1]
 * 输出: 1
 *
 * Example 2:
 * 输入: [4,1,2,1,2]
 * 输出: 4
 */

package QuestionBank

func singleNumber(nums []int) int {
	ans := 0

	for _, v := range nums {
		ans ^= v
	}

	return ans
}
