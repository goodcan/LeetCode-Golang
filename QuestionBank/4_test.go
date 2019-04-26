// @Time     : 2019/4/26 11:52
// @Author   : cancan
// @File     : 4_test.go
// @Function : 寻找两个有序数组的中位数 test

package QuestionBank

import "testing"

func Test_4(t *testing.T) {
	tests := []struct {
		num1 []int
		num2 []int
		ans  float64
	}{
		{[]int{}, []int{}, 0.0},
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, test := range tests {
		if findMedianSortedArrays(test.num1, test.num2) != test.ans {
			t.Errorf("failure num1 %v num2 %v ans %f", test.num1, test.num2, test.ans)
		}
	}
}
