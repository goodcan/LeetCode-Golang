/*
 * @Time     : 2020/5/8 15:49
 * @Author   : cancan
 * @File     : 221.go
 * @Function : 最大正方形
 */

/*
 * Question:
 * 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
 *
 * Example:
 * 输入:
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 * 输出: 4
 */

package QuestionBank

func maximalSquare(matrix [][]byte) int {
	h := len(matrix)
	if h == 0 {
		return 0
	}
	w := len(matrix[0])
	ret := 0

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			maxLen := h - i
			if w-j < maxLen {
				maxLen = w - j
			}
			size := 1
			for k := 1; k < maxLen; k++ {
				if matrix[i+k][j+k] == '0' {
					break
				}

				flag := true

				for ii := i; ii < i+k; ii++ {
					if matrix[ii][j+k] == '0' {
						flag = false
						break
					}
				}

				for jj := j; jj < j+k; jj++ {
					if matrix[i+k][jj] == '0' {
						flag = false
						break
					}
				}

				if flag {
					size++
				} else {
					break
				}
			}

			if size > ret {
				ret = size
			}
		}
	}

	return ret * ret
}
