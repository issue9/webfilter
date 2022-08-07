// SPDX-License-Identifier: MIT

package validator

import (
	"github.com/issue9/validator/gb11643"
	"github.com/issue9/validator/gb32100"
)

// GB11643 判断一个身份证是否符合 gb11643 标准
//
// 若是 15 位则当作一代身份证，仅简单地判断各位是否都是数字；
// 若是 18 位则当作二代身份证，会计算校验位是否正确；
func GB11643(val any) bool {
	switch v := val.(type) {
	case string:
		return gb11643.IsValid([]byte(v))
	case []byte:
		return gb11643.IsValid(v)
	case []rune:
		return gb11643.IsValid([]byte(string(v)))
	default:
		return false
	}
}

// GB32100 统一信用代码校验
func GB32100(val any) bool {
	switch v := val.(type) {
	case string:
		return gb32100.IsValid([]byte(v))
	case []byte:
		return gb32100.IsValid(v)
	case []rune:
		return gb32100.IsValid([]byte(string(v)))
	default:
		return false
	}
}
