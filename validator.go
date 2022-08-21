// SPDX-License-Identifier: MIT

// Package validator 符合 [web.validator] 的验证方法
//
// [web.validator]: https://pkg.go.dev/github.com/issue9/web#Validator
package validator

import "github.com/issue9/web/server"

type (
	Validator = server.Validator

	// Func 是 [Validator] 的函数形式
	Func = server.ValidateFunc
)

func And(v ...Validator) Validator { return server.AndValidator(v...) }

func Or(v ...Validator) Validator { return server.OrValidator(v...) }

func AndFunc(f ...func(any) bool) Validator { return server.AndValidateFunc(f...) }

func OrFunc(f ...func(any) bool) Validator { return server.OrValidateFunc(f...) }
