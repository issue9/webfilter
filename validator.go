// SPDX-License-Identifier: MIT

// Package validator 符合 [web/validation] 的验证方法
//
// [web/validation]: https://pkg.go.dev/github.com/issue9/web/validation
package validator

import "github.com/issue9/web/validation"

type (
	Validator = validation.Validator

	// Func 是 [Validator] 的函数形式
	Func = validation.ValidateFunc
)

func And(v ...Validator) Validator { return validation.And(v...) }

func Or(v ...Validator) Validator { return validation.Or(v...) }

func AndFunc(f ...func(any) bool) Validator { return validation.AndF(f...) }

func OrFunc(f ...func(any) bool) Validator { return validation.OrF(f...) }
