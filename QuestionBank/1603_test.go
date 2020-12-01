// @Time     : 2020/12/1 15:37
// @Author   : cancan
// @File     : test_1603.go
// @Function :

package QuestionBank

import (
	"LeetCode-Golang/utils"
	"testing"
)

func Test_1603(t *testing.T) {
	tests := []struct {
		methods []string
		params  [][]interface{}
		ans     []interface{}
	}{
		{
			[]string{"ParkingSystem", "AddCar", "AddCar", "AddCar", "AddCar"},
			[][]interface{}{{1, 1, 0}, {1}, {2}, {3}, {1}},
			[]interface{}{nil, true, true, false, false},
		},
	}

	s2m := utils.NewString2Method()
	var ParkingSystem ParkingSystem

	for index, test := range tests {
		for i, method := range test.methods {
			if i == 0 {
				ParkingSystem = Constructor1603(test.params[i][0].(int), test.params[i][1].(int), test.params[i][2].(int))
			} else {
				if !s2m.Check(s2m.Run(&ParkingSystem, method, test.params[i]...), test.ans[i]) {
					t.Errorf("failure id %d", index)
				}
			}
		}
	}
}
