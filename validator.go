// SPDX-License-Identifier: MIT

// Package validator 符合 [web.validation] 的验证器
//
// [web.validation]: https://pkg.go.dev/github.com/issue9/web/validation
package validator

import (
	"reflect"

	"github.com/issue9/web/validation"
)

func And[T any](v ...validation.ValidatorOf[T]) validation.ValidatorOf[T] {
	return validation.And(v...)
}

func AndFunc[T any](v ...func(T) bool) validation.ValidatorOf[T] {
	return validation.AndFunc(v...)
}

func Or[T any](v ...validation.ValidatorOf[T]) validation.ValidatorOf[T] {
	return validation.Or(v...)
}

func OrFunc[T any](v ...func(T) bool) validation.ValidatorOf[T] {
	return validation.OrFunc(v...)
}

func Not[T any](v validation.ValidatorOf[T]) validation.ValidatorOf[T] {
	return validation.Not(v)
}

func NotFunc[T any](v func(T) bool) validation.ValidatorOf[T] {
	return validation.NotFunc(v)
}

func Zero[T any](v T) bool { return reflect.ValueOf(v).IsZero() }
