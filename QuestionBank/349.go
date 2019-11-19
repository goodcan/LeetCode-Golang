/*
 * @Time     : 2019/11/19 16:08
 * @Author   : cancan
 * @File     : 349.go
 * @Function : 两个数组的交集
 */

/*
 * Question:
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * Example 1:
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2]
 *
 * Example 2:
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [9,4]
 *
 * Note:
 * 1.输出结果中的每个元素一定是唯一的。
 * 2.我们可以不考虑输出结果的顺序。
 */

package QuestionBank

func intersection(nums1 []int, nums2 []int) []int {
	ans := []int{}
	tmp := make(map[int]int)
	for _, v := range nums1 {
		tmp[v]++
	}
	hasAdd := make(map[int]int)
	for _, v := range nums2 {
		if _, ok := tmp[v]; ok {
			if _, ok := hasAdd[v]; !ok {
				ans = append(ans, v)
				hasAdd[v]++
			}
		}
	}
	return ans
}
