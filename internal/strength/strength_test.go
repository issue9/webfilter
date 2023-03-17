// SPDX-License-Identifier: MIT

package strength

import (
	"testing"

	"github.com/issue9/assert/v3"
)

func TestStrength(t *testing.T) {
	a := assert.New(t, false)

	v := New(0, 0, 0, 0)
	a.True(v(""))
	a.True(v("123"))

	v = New(3, 0, 0, 0)
	a.False(v(""))
	a.False(v("12"))
	a.True(v("123"))
	a.True(v("Abcdef"))

	v = New(3, 2, 0, 0)
	a.False(v(""))
	a.False(v("12"))
	a.False(v("123"))
	a.False(v("123A"))
	a.True(v("123AB"))
	a.False(v("AB"))
	a.True(v("ABc"))

	v = New(0, 2, 3, 1)
	a.False(v(""))
	a.False(v("12345678"))
	a.True(v("ABcde>"))
	a.True(v("ABcde123><"))

	v = New(4, 2, 0, 1)
	a.False(v(""))
	a.False(v("12345678"))
	a.False(v("AB>"))
	a.True(v("AB=!>"))
}
