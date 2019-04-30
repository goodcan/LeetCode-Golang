// @Time     : 2019/4/30 14:16
// @Author   : cancan
// @File     : 706_test.go
// @Function : 设计哈希映射 test

package QuestionBank

import "testing"

func Test_706(t *testing.T) {
	hashMap := Constructor706()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	if hashMap.Get(1) != 1 {
		t.Error("Get method error")
	}
	if hashMap.Get(3) != -1 {
		t.Error("Get method error")
	}
	hashMap.Put(2, 1)
	if hashMap.Get(2) != 1 {
		t.Error("Put method error")
	}
	hashMap.Remove(2)
	if hashMap.Get(2) != -1 {
		t.Error("Remove method error")
	}
}
