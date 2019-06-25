// @Time     : 2019/4/29 16:38
// @Author   : cancan
// @File     : 137.go
// @Function : 只出现一次的数字 II

/*
 * Question:
 * 给定一个非空整数数组，除了某个元素只出现一次以外，
 * 其余每个元素均出现了三次。找出那个只出现了一次的元素。
 *
 * Note
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * Example 1:
 * 输入: [2,2,3,2]
 * 输出: 3
 *
 * Example 2:
 * 输入: [0,1,0,1,0,1,99]
 * 输出: 99
 */

package QuestionBank

func singleNumberII(nums []int) int {
	tmp := make(map[int]int)

	for _, v := range nums {
		_, ok := tmp[v]
		if ok {
			tmp[v]++
		} else {
			tmp[v] = 1
		}
	}

	for k, v := range tmp {
		if v == 1 {
			return k
		}
	}

	return -1
}
