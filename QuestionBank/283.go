/*
 * @Time     : 2019/7/9 9:45
 * @Author   : cancan
 * @File     : 283.go
 * @Function : 移动零
 */

/*
 * Question:
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * Example:
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * Note:
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 */

package QuestionBank

func moveZeroes(nums []int) {
	l := len(nums)
	i := 0
	c := 0
	for i < (l - c) {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			c++
		} else {
			i++
		}
	}
}
