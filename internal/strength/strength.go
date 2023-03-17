// SPDX-License-Identifier: MIT

package strength

import (
	"unicode"

	"github.com/issue9/web/filter"
)

type strength struct {
	length int
	upper  int
	lower  int
	punct  int
}

func New(length, upper, lower, punct int) filter.ValidatorFuncOf[string] {
	return func(s string) bool {
		if length == 0 && upper == 0 && lower == 0 && punct == 0 {
			return true
		}

		cnt := &strength{}
		for _, r := range s {
			switch {
			case unicode.IsPunct(r) || unicode.IsSymbol(r):
				cnt.punct++
			case unicode.IsUpper(r):
				cnt.upper++
			case unicode.IsLower(r):
				cnt.lower++
			}
			cnt.length++
		}

		ok := true
		if length > 0 {
			ok = cnt.length >= length
		}
		if ok && lower > 0 {
			ok = cnt.lower >= lower
		}
		if ok && upper > 0 {
			ok = cnt.upper >= upper
		}
		if ok && punct > 0 {
			ok = cnt.punct >= punct
		}
		return ok
	}
}
