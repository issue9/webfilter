// SPDX-License-Identifier: MIT

// Package validator 符合 [web.filter] 的验证器
//
// [web.filter]: https://pkg.go.dev/github.com/issue9/web#Filter
package validator

import "reflect"

// And 以与的形式串联多个验证器函数
func And[T any](v ...func(T) bool) func(T) bool {
	return func(val T) bool {
		for _, validator := range v {
			if !validator(val) {
				return false
			}
		}
		return true
	}
}

// Or 以或的形式并联多个验证器函数
func Or[T any](v ...func(T) bool) func(T) bool {
	return func(val T) bool {
		for _, validator := range v {
			if validator(val) {
				return true
			}
		}
		return false
	}
}

// Not 验证器的取反
func Not[T any](v func(T) bool) func(T) bool {
	return func(val T) bool { return !v(val) }
}

func Zero[T any](v T) bool { return reflect.ValueOf(v).IsZero() }

// Equal 生成判断值是否等于 v 的验证器
func Equal[T comparable](v T) func(T) bool {
	return func(t T) bool { return t == v }
}

func NilOr[T any](v func(T) bool) func(T) bool {
	return Or(func(v T) bool { return reflect.ValueOf(v).IsNil() }, v)
}
