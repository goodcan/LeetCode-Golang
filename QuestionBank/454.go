// @Time     : 2020/11/27 10:22
// @Author   : cancan
// @File     : 454.go
// @Function : 四数相加 II

/*
 * 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
 * 为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。
 *
 * 例如:
 * 输入:
 * A = [ 1, 2]
 * B = [-2,-1]
 * C = [-1, 2]
 * D = [ 0, 2]
 * 输出:
 * 2
 *
 * 解释:
 * 两个元组如下:
 * 1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
 * 2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
 */

package QuestionBank

func fourSumCount(A []int, B []int, C []int, D []int) int {
	tmp := make(map[int]int)
	for _, va := range A {
		for _, vb := range B {
			tmp[0-va-vb]++
		}
	}
	count := 0
	for _, vc := range C {
		for _, vd := range D {
			if num, ok := tmp[vc+vd]; ok {
				count += num
			}
		}
	}

	return count
}
