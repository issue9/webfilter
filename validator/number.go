// SPDX-License-Identifier: MIT

package validator

import "github.com/issue9/web/filter"

type Number interface {
	float32 | float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Between 判断数值区间 (min, max)
func Between[T Number](min, max T) filter.ValidatorFuncOf[T] {
	if max < min {
		panic("max 必须大于等于 min")
	}

	return func(val T) bool { return val > min && val < max }
}

// BetweenEqual 判断数值区间 [min, max]
func BetweenEqual[T Number](min, max T) filter.ValidatorFuncOf[T] {
	if max < min {
		panic("max 必须大于等于 min")
	}

	return func(val T) bool { return val >= min && val <= max }
}

func Less[T Number](num T) filter.ValidatorFuncOf[T] {
	return func(t T) bool { return t < num }
}

func LessEqual[T Number](num T) filter.ValidatorFuncOf[T] {
	return func(t T) bool { return t <= num }
}

func Great[T Number](num T) filter.ValidatorFuncOf[T] {
	return func(t T) bool { return t > num }
}

func GreatEqual[T Number](num T) filter.ValidatorFuncOf[T] {
	return func(t T) bool { return t >= num }
}
