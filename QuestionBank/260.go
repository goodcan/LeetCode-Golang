/*
 * @Time     : 2019-11-17 22:22
 * @Author   : cancan
 * @File     : 260.go
 * @Function : 只出现一次的数字 III
 */

/*
 # Question:
 # 给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
 #
 # Example:
 # 输入: [1,2,1,3,2,5]
 # 输出: [3,5]
 #
 # Note：
 # 1.结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
 # 2.你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现
*/

package QuestionBank

func singleNumberIII(nums []int) []int {
	tmp := make(map[int]int)
	for _, v := range nums {
		tmp[v]++
	}
	ans := []int{}
	for k, v := range tmp {
		if v > 1 {
			continue
		}
		ans = append(ans, k)
	}
	return ans
}
