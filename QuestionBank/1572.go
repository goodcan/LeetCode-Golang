// @Time     : 2020/12/1 17:09
// @Author   : cancan
// @File     : 1572.go
// @Function : 矩阵对角线元素的和

/*
 * 给你一个正方形矩阵 mat，请你返回矩阵对角线元素的和。
 * 请你返回在矩阵主对角线上的元素和副对角线上且不在主对角线上元素的和。
 *
 * 示例 1：
 * 输入：mat = [[1,2,3],
 *            [4,5,6],
 *            [7,8,9]]
 * 输出：25
 * 解释：对角线的和为：1 + 5 + 9 + 3 + 7 = 25
 * 请注意，元素 mat[1][1] = 5 只会被计算一次。
 *
 * 示例 2：
 * 输入：mat = [[1,1,1,1],
 *            [1,1,1,1],
 *            [1,1,1,1],
 *            [1,1,1,1]]
 * 输出：8
 *
 * 示例 3：
 * 输入：mat = [[5]]
 * 输出：5
 *
 * 提示：
 * - n == mat.length == mat[i].length
 * - 1 <= n <= 100
 * - 1 <= mat[i][j] <= 100
 */

package QuestionBank

func diagonalSum(mat [][]int) int {
	ans := 0
	l := len(mat)
	r := l / 2
	if l%2 != 0 {
		r++
	}

	for i := 0; i < r; i++ {
		if i == l-i-1 && l%2 != 0 {
			ans += mat[i][i]
		} else {
			ans += mat[i][i] + mat[i][l-i-1] + mat[l-i-1][i] + mat[l-i-1][l-i-1]
		}
	}

	return ans
}
