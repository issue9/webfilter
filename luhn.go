// SPDX-License-Identifier: MIT

package validator

import "github.com/issue9/validator/luhn"

// BankCard 是否为正确的银行卡号
func BankCard(val any) bool { return Luhn(val) }

// Luhn 验证 [luhn[] 算法
//
// [luhn]: https://en.wikipedia.org/wiki/Luhn_algorithm
func Luhn(val any) bool {
	switch v := val.(type) {
	case []byte:
		return luhn.IsValid(v)
	case string:
		return luhn.IsValid([]byte(v))
	case []rune:
		return luhn.IsValid([]byte(string(v)))
	default:
		return false
	}
}
