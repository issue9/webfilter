// SPDX-FileCopyrightText: 2022-2024 caixw
//
// SPDX-License-Identifier: MIT

// Package sanitizer 内容修正工具
package sanitizer

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"unicode"
)

func Sanitizers[T any](f ...func(*T)) func(*T) {
	return func(v *T) {
		for _, ss := range f {
			ss(v)
		}
	}
}

// Trim 过滤左右空格
func Trim(v *string) { *v = strings.TrimSpace(*v) }

func TrimLeft(v *string) {
	*v = strings.TrimLeftFunc(*v, func(r rune) bool { return unicode.IsSpace(r) })
}

func TrimRight(v *string) {
	*v = strings.TrimRightFunc(*v, func(r rune) bool { return unicode.IsSpace(r) })
}

func Upper(v *string) { *v = strings.ToUpper(*v) }

func Lower(v *string) { *v = strings.ToLower(*v) }

func MD5(v *string) {
	h := md5.New()
	h.Write([]byte(*v))
	*v = hex.EncodeToString(h.Sum(nil))
}
