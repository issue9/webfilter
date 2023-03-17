// SPDX-License-Identifier: MIT

package sanitizer

import (
	"testing"

	"github.com/issue9/assert/v3"
)

func TestTrim(t *testing.T) {
	a := assert.New(t, false)

	s := " abc\t"
	Sanitizers(TrimLeft, TrimRight)(&s)
	a.Equal(s, "abc")

	s = " abc\t"
	Trim(&s)
	a.Equal(s, "abc")
}

func TestLower_Upper(t *testing.T) {
	a := assert.New(t, false)

	s := "Abc"
	Lower(&s)
	a.Equal(s, "abc")

	s = "Abc"
	Upper(&s)
	a.Equal(s, "ABC")
}
