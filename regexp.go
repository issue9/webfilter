// SPDX-License-Identifier: MIT

package validator

import "regexp"

// Match 为正则生成验证函数
func Match(exp *regexp.Regexp) F {
	return func(val any) bool {
		switch v := val.(type) {
		case []rune:
			return exp.MatchString(string(v))
		case []byte:
			return exp.Match(v)
		case string:
			return exp.MatchString(v)
		default:
			return false
		}
	}
}
