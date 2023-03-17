// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/issue9/sliceutil"
	"github.com/issue9/web/filter"
)

// In 声明枚举类型的验证规则
//
// 要求验证的值必须包含在 element 元素中，如果不存在，则返回 msg 的内容。
func In[T comparable](element ...T) filter.ValidatorFuncOf[T] {
	return func(v T) bool {
		return sliceutil.Exists(element, func(elem T) bool { return elem == v })
	}
}

// NotIn 声明不在枚举中的验证规则
func NotIn[T comparable](element ...T) filter.ValidatorFuncOf[T] {
	return Not(In(element...))
}
