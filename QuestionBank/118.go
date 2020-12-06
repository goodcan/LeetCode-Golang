/*
 * @Time     : 2020/12/6 15:56
 * @Author   : cancan
 * @File     : 118.go
 * @Function : 杨辉三角
 */

/*
 * Question:
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * Example:
 * 输入: 5
 * 输出:
 * [
 *      [1],
 *     [1,1],
 *    [1,2,1],
 *   [1,3,3,1],
 *  [1,4,6,4,1]
 * ]
 */

package QuestionBank

func generate(numRows int) [][]int {
	ret := make([][]int, 0, numRows)
	if numRows == 0 {
		return ret
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	ret = [][]int{{1}}
	r := []int{1, 1}
	for i := 1; i < numRows; i++ {
		ret = append(ret, r)
		t := []int{1}
		for j := 0; j < len(r)-1; j++ {
			t = append(t, r[j]+r[j+1])
		}
		t = append(t, 1)
		r = t
	}
	return ret
}
