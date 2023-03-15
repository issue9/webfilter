// SPDX-License-Identifier: MIT

package validator

import "github.com/issue9/web/validation"

type Number interface {
	float32 | float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Range 声明判断数值大小的验证规则
func Range[T Number](min, max T) validation.ValidatorFuncOf[T] {
	if max < min {
		panic("max 必须大于等于 min")
	}

	return func(val T) bool { return val >= min && val <= max }
}

// Min 声明判断数值不小于 min 的验证规则
func Min[T Number](min T) validation.ValidatorFuncOf[T] {
	return func(v T) bool { return v >= min }
}

// Max 声明判断数值不大于 max 的验证规则
func Max[T Number](max T) validation.ValidatorFuncOf[T] {
	return func(v T) bool { return v <= max }
}
