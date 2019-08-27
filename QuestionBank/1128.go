/*
 * @Time     : 2019/8/27 16:09
 * @Author   : cancan
 * @File     : 1128.go
 * @Function : 等价多米诺骨牌对的数量
 */

/*
 * Question:
 * 给你一个由一些多米诺骨牌组成的列表 dominoes。
 * 如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。
 * 形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。
 * 在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i, j) 的数量。
 *
 * Example：
 * 输入：dominoes = [[1,2],[2,1],[3,4],[5,6]]
 * 输出：1
 *
 * Note：
 * 1 <= dominoes.length <= 40000
 * 1 <= dominoes[i][j] <= 9
 */

package QuestionBank

func numEquivDominoPairs(dominoes [][]int) int {
	tmp := make(map[[2]int]int)
	ans := 0

	for _, v := range dominoes {
		if _, ok := tmp[[2]int{v[0], v[1]}]; ok {
			tmp[[2]int{v[0], v[1]}]++
		} else if _, ok := tmp[[2]int{v[1], v[0]}]; ok {
			tmp[[2]int{v[1], v[0]}]++
		} else {
			tmp[[2]int{v[0], v[1]}] = 1
		}
	}

	for _, v := range tmp {
		if v == 1 {
			continue
		}
		ans += (v * (v - 1)) / 2
	}

	return ans
}
