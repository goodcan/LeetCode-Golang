// @Time     : 2020/11/9 11:30
// @Author   : cancan
// @File     : 973.go
// @Function : 最接近原点的 K 个点

/*
 * Question:
 * 我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。
 * （这里，平面上两点之间的距离是欧几里德距离。）
 * 你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。
 *
 * Example 1：
 * 输入：points = [[1,3],[-2,2]], K = 1
 * 输出：[[-2,2]]
 *
 * 解释：
 * (1, 3) 和原点之间的距离为 sqrt(10)，
 * (-2, 2) 和原点之间的距离为 sqrt(8)，
 * 由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
 * 我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
 *
 * Example 2：
 * 输入：points = [[3,3],[5,-1],[-2,4]], K = 2
 * 输出：[[3,3],[-2,4]]
 * （答案 [[-2,4],[3,3]] 也会被接受。）
 *
 * Note：
 * 1. 1 <= K <= points.length <= 10000
 * 2. -10000 < points[i][0] < 10000
 * 3. -10000 < points[i][1] < 10000
 */

package QuestionBank

import "sort"

func kClosest(points [][]int, K int) [][]int {
	tmp := &kClosestItem{}
	for idx, v := range points {
		tmp.data = append(tmp.data, [2]int{idx, v[0]*v[0] + v[1]*v[1]})
	}
	sort.Sort(tmp)

	ret := [][]int{}
	for _, v := range tmp.data[:K] {
		ret = append(ret, points[v[0]])
	}

	return ret
}

type kClosestItem struct {
	data [][2]int
}

func (k *kClosestItem) Len() int {
	return len(k.data)
}

func (k *kClosestItem) Less(i, j int) bool {
	return k.data[i][1] < k.data[j][1]
}

func (k *kClosestItem) Swap(i, j int) {
	k.data[i], k.data[j] = k.data[j], k.data[i]
}
