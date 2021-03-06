/*
 * @Time     : 2020/5/27 11:11
 * @Author   : cancan
 * @File     : 974.go
 * @Function : 和可被 K 整除的子数组
 */

/*
 * Question:
 * 给定一个整数数组 A，返回其中元素之和可被 K 整除的（连续、非空）子数组的数目。
 *
 * Example：
 * 输入：A = [4,5,0,-2,-3,1], K = 5
 * 输出：7
 * 解释：
 * 有 7 个子数组满足其元素之和可被 K = 5 整除：
 * [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
 *
 * Note：
 * 1 <= A.length <= 30000
 * -10000 <= A[i] <= 10000
 * 2 <= K <= 10000
 */

package QuestionBank

func subarraysDivByK(A []int, K int) int {
	tmp := map[int]int{0: 1}
	total, ans := 0, 0
	for _, v := range A {
		total += v
		m := (total%K + K) % K
		ans += tmp[m]
		tmp[m]++
	}
	return ans
}
