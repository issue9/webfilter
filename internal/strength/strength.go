// SPDX-License-Identifier: MIT

package strength

import (
	"unicode"

	"github.com/issue9/web"

	"github.com/issue9/validator/internal/convert"
)

type strength struct {
	length int
	upper  int
	lower  int
	punct  int
}

func New(length, upper, lower, punct int) web.Validator {
	return &strength{
		length: length,
		upper:  upper,
		lower:  lower,
		punct:  punct,
	}
}

func (s *strength) IsValid(v any) (ok bool) {
	if s.length == 0 && s.upper == 0 && s.lower == 0 && s.punct == 0 {
		return true
	}

	rs, ok := convert.Runes(v)
	if !ok {
		return false
	}

	cnt := &strength{}
	for _, r := range rs {
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

	ok = true
	if s.length > 0 {
		ok = cnt.length >= s.length
	}
	if ok && s.lower > 0 {
		ok = cnt.lower >= s.lower
	}
	if ok && s.upper > 0 {
		ok = cnt.upper >= s.upper
	}
	if ok && s.punct > 0 {
		ok = cnt.punct >= s.punct
	}
	return ok
}
