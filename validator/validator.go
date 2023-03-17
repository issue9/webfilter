// SPDX-License-Identifier: MIT

// Package validator 符合 [web.filter] 的验证器
//
// [web.filter]: https://pkg.go.dev/github.com/issue9/web/filter
package validator

import (
	"reflect"

	"github.com/issue9/web/filter"
)

func And[T any](v ...func(T) bool) filter.ValidatorFuncOf[T] {
	return filter.And(v...)
}

func Or[T any](v ...func(T) bool) filter.ValidatorFuncOf[T] {
	return filter.Or(v...)
}

func Not[T any](v func(T) bool) filter.ValidatorFuncOf[T] {
	return filter.Not(v)
}

func Zero[T any](v T) bool { return reflect.ValueOf(v).IsZero() }

// Equal 生成判断值是否等于 v 的验证器
func Equal[T comparable](v T) filter.ValidatorFuncOf[T] {
	return func(t T) bool { return t == v }
}
