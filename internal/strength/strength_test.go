// SPDX-License-Identifier: MIT

package strength

import (
	"testing"

	"github.com/issue9/assert/v3"
)

func TestStrength(t *testing.T) {
	a := assert.New(t, false)

	v := New(0, 0, 0, 0)
	a.True(v.IsValid(nil))
	a.True(v.IsValid(""))
	a.True(v.IsValid("123"))

	v = New(3, 0, 0, 0)
	a.False(v.IsValid(""))
	a.False(v.IsValid("12"))
	a.True(v.IsValid("123"))
	a.True(v.IsValid("Abcdef"))

	v = New(3, 2, 0, 0)
	a.False(v.IsValid(""))
	a.False(v.IsValid("12"))
	a.False(v.IsValid("123"))
	a.False(v.IsValid("123A"))
	a.True(v.IsValid("123AB"))
	a.False(v.IsValid("AB"))
	a.True(v.IsValid("ABc"))

	v = New(0, 2, 3, 1)
	a.False(v.IsValid(""))
	a.False(v.IsValid(nil))
	a.False(v.IsValid("12345678"))
	a.True(v.IsValid("ABcde>"))
	a.True(v.IsValid("ABcde123><"))

	v = New(4, 2, 0, 1)
	a.False(v.IsValid(""))
	a.False(v.IsValid(nil))
	a.False(v.IsValid("12345678"))
	a.False(v.IsValid("AB>"))
	a.True(v.IsValid("AB=!>"))
}
