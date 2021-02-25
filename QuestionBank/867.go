// @Time     : 2021/2/25 10:03
// @Author   : cancan
// @File     : 867.go
// @Function : 转置矩阵

/*
 * Question:
 * 给定一个矩阵 A， 返回 A 的转置矩阵。
 * 矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
 *
 * Example 1：
 * 输入：[[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[[1,4,7],[2,5,8],[3,6,9]]
 *
 * Example 2：
 * 输入：[[1,2,3],[4,5,6]]
 * 输出：[[1,4],[2,5],[3,6]]
 *
 * Note
 *     1 <= A.length <= 1000
 *     1 <= A[0].length <= 1000
 */

package QuestionBank

func transpose(matrix [][]int) [][]int {
	h := len(matrix)
	w := len(matrix[0])
	ans := make([][]int, 0, w)
	for i := 0; i < w; i++ {
		tmp := make([]int, 0, h)
		for j := 0; j < h; j++ {
			tmp = append(tmp, matrix[j][i])
		}
		ans = append(ans, tmp)
	}
	return ans
}
