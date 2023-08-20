// SPDX-License-Identifier: MIT

package validator

type Number interface {
	float32 | float64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Between 判断数值区间 (min, max)
func Between[T Number](min, max T) func(T) bool {
	if max < min {
		panic("max 必须大于等于 min")
	}

	return func(val T) bool { return val > min && val < max }
}

// BetweenEqual 判断数值区间 [min, max]
func BetweenEqual[T Number](min, max T) func(T) bool {
	if max < min {
		panic("max 必须大于等于 min")
	}

	return func(val T) bool { return val >= min && val <= max }
}

func Less[T Number](num T) func(T) bool {
	return func(t T) bool { return t < num }
}

func LessEqual[T Number](num T) func(T) bool {
	return func(t T) bool { return t <= num }
}

func Great[T Number](num T) func(T) bool {
	return func(t T) bool { return t > num }
}

func GreatEqual[T Number](num T) func(T) bool {
	return func(t T) bool { return t >= num }
}

// HTTPStatus 是否为有效的 HTTP 状态码
func HTTPStatus(s int) bool { return BetweenEqual(100, 599)(s) }

// ZeroOr 判断值为零值或是非零情况下符合 v 的要求
func ZeroOr[T Number](v func(T) bool) func(T) bool {
	return Or(func(v T) bool { return v == 0 }, v)
}
