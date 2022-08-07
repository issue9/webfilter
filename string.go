// SPDX-License-Identifier: MIT

package validator

import (
	"net/mail"
	"net/netip"
	"net/url"

	"github.com/issue9/validator/gb11643"
	"github.com/issue9/validator/gb32100"
	"github.com/issue9/validator/internal/isbn"
	"github.com/issue9/validator/luhn"
)

func URL(val any) bool {
	if s, ok := any2String(val); ok {
		_, err := url.ParseRequestURI(s)
		return err == nil
	}
	return false
}

func Email(val any) bool {
	if s, ok := any2String(val); ok {
		_, err := mail.ParseAddress(s)
		return err == nil
	}
	return false
}

func IP4(val any) bool {
	if s, ok := any2String(val); ok {
		if ip, err := netip.ParseAddr(s); err == nil {
			return ip.Is4()
		}
	}
	return false
}

func IP6(val any) bool {
	if s, ok := any2String(val); ok {
		if ip, err := netip.ParseAddr(s); err == nil {
			return ip.Is6()
		}
	}
	return false
}

func IP(val any) bool {
	if s, ok := any2String(val); ok {
		_, err := netip.ParseAddr(s)
		return err == nil
	}
	return false
}

// ISBN 判断是否为合法的 [ISBN] 串号
//
// 可以同时判断 ISBN10 和 ISBN13
//
// [ISBN]: https://zh.wikipedia.org/wiki/%E5%9B%BD%E9%99%85%E6%A0%87%E5%87%86%E4%B9%A6%E5%8F%B7
func ISBN(val any) bool { return validBytes(isbn.IsValid, val) }

func ISBN10(val any) bool { return validBytes(isbn.ISBN10, val) }

func ISBN13(val any) bool { return validBytes(isbn.ISBN13, val) }

// GB11643 判断一个身份证是否符合 gb11643 标准
//
// 若是 15 位则当作一代身份证，仅简单地判断各位是否都是数字；
// 若是 18 位则当作二代身份证，会计算校验位是否正确；
func GB11643(val any) bool { return validBytes(gb11643.IsValid, val) }

// GB32100 统一信用代码校验
func GB32100(val any) bool { return validBytes(gb32100.IsValid, val) }

// BankCard 是否为正确的银行卡号
func BankCard(val any) bool { return Luhn(val) }

// Luhn 验证 [luhn] 算法
//
// [luhn]: https://en.wikipedia.org/wiki/Luhn_algorithm
func Luhn(val any) bool { return validBytes(luhn.IsValid, val) }

// HexColor 判断一个字符串是否为合法的 16 进制颜色表示法
func HexColor(val any) bool { return validBytes(hexColor, val) }

func hexColor(val []byte) bool {
	if len(val) != 4 && len(val) != 7 {
		return false
	}

	if val[0] != '#' {
		return false
	}

	for _, v := range val[1:] {
		switch {
		case '0' <= v && v <= '9':
		case 'a' <= v && v <= 'f':
		case 'A' <= v && v <= 'F':
		default:
			return false
		}
	}
	return true
}

func validBytes(f func([]byte) bool, val any) bool {
	switch v := val.(type) {
	case string:
		return f([]byte(v))
	case []byte:
		return f(v)
	case []rune:
		return f([]byte(string(v)))
	default:
		return false
	}
}

func any2String(val any) (string, bool) {
	switch v := val.(type) {
	case string:
		return v, true
	case []byte:
		return string(v), true
	case []rune:
		return string(v), true
	default:
		return "", false
	}
}
