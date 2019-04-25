package QuestionBank

import "testing"

func Test_674(t *testing.T) {
	tests := []struct {
		nums []int
		ans  int
	}{
		{[]int{}, 0},
		{[]int{2}, 1},
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{2, 2, 2, 2, 2}, 1},
		{[]int{1, 2, 3, 4}, 4},
	}

	for _, test := range tests {
		if findLengthOfLCIS(test.nums) != test.ans {
			t.Errorf("failure nums %v ans %d", test.nums, test.ans)
		}
	}
}
