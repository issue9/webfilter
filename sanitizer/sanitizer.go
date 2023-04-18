// SPDX-License-Identifier: MIT

// Package sanitizer 内容修正工具
package sanitizer

import (
	"strings"
	"unicode"

	"github.com/issue9/web/filter"
)

func Sanitizers[T any](f ...func(*T)) filter.SanitizeFuncOf[T] {
	return filter.Sanitizers(f...)
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
