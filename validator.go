// SPDX-License-Identifier: MIT

// Package validator 符合 [web/validation] 的验证方法
//
// [web/validation]: https://github.com/issue9/web
package validator

import "github.com/issue9/web/validation"

type (
	Validator = validation.Validator
	F         = validation.ValidateFunc
)

func And(v ...Validator) Validator { return validation.And(v...) }

func Or(v ...Validator) Validator { return validation.Or(v...) }
