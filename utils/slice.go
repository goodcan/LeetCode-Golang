/*
 * @Time     : 2019/4/25 15:53
 * @Author   : cancan
 * @File     : slice.go
 * @Function : 切片相关工具
 */

package utils

import (
	"reflect"
)

func SliceEqual(a, b interface{}) bool {
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)

	if v1.Kind() != reflect.Slice || v2.Kind() != reflect.Slice {
		return false
	}

	if v1.IsNil() != v2.IsNil() {
		return false
	}

	if v1.Len() != v2.Len() {
		return false
	}

	for i := 0; i < v1.Len(); i++ {
		e1 := v1.Index(i)
		e2 := v2.Index(i)
		isSlice1 := false
		isSlice2 := false
		switch e1.Interface().(type) {
		case []int:
			isSlice1 = true
		}
		switch e2.Interface().(type) {
		case []int:
			isSlice2 = true
		}
		if isSlice1 && isSlice2 {
			if !SliceEqual(e1.Interface(), e2.Interface()) {
				return false
			}
		} else {
			if e1.Interface() != e2.Interface() {
				return false
			}
		}
	}

	return true
}
