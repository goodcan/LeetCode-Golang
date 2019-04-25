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

	for _, tt := range tests {
		if findLengthOfLCIS(tt.nums) != tt.ans {
			t.Errorf("failure nums %v ans %d", tt.nums, tt.ans)
		}
	}
}
