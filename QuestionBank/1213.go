/*
 * @Time     : 2019/10/8 11:43
 * @Author   : cancan
 * @File     : 5079.go
 * @Function : 三个有序数组的交集
 */

/*
 * Question:
 * 给出三个均为 严格递增排列 的整数数组 arr1，arr2 和 arr3。
 * 返回一个由 仅 在这三个数组中 同时出现 的整数所构成的有序数组。
 *
 * Example：
 * 输入: arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9], arr3 = [1,3,4,5,8]
 * 输出: [1,5]
 * 解释: 只有 1 和 5 同时在这三个数组中出现.
 *
 * Note：
 * 1 <= arr1.length, arr2.length, arr3.length <= 1000
 * 1 <= arr1[i], arr2[i], arr3[i] <= 2000
 */

package QuestionBank

import "sort"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	m1 := mergeSlice(arr1, arr2)
	m2 := mergeSlice(m1, arr3)
	sort.Ints(m2)
	return m2
}

func mergeSlice(a, b []int) []int {
	ret := []int{}
	check := map[int]int{}

	for _, v := range a {
		check[v]++
	}

	for _, v := range b {
		check[v]++
	}

	for k, v := range check {
		if v > 1 {
			ret = append(ret, k)
		}
	}

	return ret
}
