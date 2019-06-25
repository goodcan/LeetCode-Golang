/*
 * @Time     : 2019/4/30 15:21
 * @Author   : cancan
 * @File     : struct.go
 * @Function : 接口类工具
 */

package utils

import (
	"reflect"
)

type String2Method interface {
	Run(obj interface{}, methodName string, args ...interface{}) []reflect.Value
	Check(ret []reflect.Value, ans ...interface{}) bool
}

type string2Method struct{}

func NewString2Method() String2Method {
	return &string2Method{}
}

func (this *string2Method) Run(obj interface{}, methodName string, args ...interface{}) []reflect.Value {
	inputs := []reflect.Value{}
	for _, v := range args {
		inputs = append(inputs, reflect.ValueOf(v))
	}
	return reflect.ValueOf(obj).MethodByName(methodName).Call(inputs)
}

func (this *string2Method) Check(ret []reflect.Value, ans ...interface{}) bool {
	if ret == nil && ans == nil {
		return true
	} else if ret != nil && ans != nil {
		for i, v := range ret {
			if v.Interface() != ans[i] {
				return false
			}
		}
		return true
	} else {
		return true
	}
}
