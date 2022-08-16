// SPDX-License-Identifier: MIT

package convert

import (
	"testing"

	"github.com/issue9/assert/v3"
)

func TestString(t *testing.T) {
	a := assert.New(t, false)

	v, ok := String("abc")
	a.True(ok).Equal(v, "abc")
	v, ok = String([]byte("abc"))
	a.True(ok).Equal(v, "abc")
	v, ok = String([]rune("abc"))
	a.True(ok).Equal(v, "abc")

	v, ok = String(123)
	a.False(ok).Empty(v)
}

func TestBytes(t *testing.T) {
	a := assert.New(t, false)

	v, ok := Bytes("abc")
	a.True(ok).Equal(v, []byte("abc"))
	v, ok = Bytes([]byte("abc"))
	a.True(ok).Equal(v, []byte("abc"))
	v, ok = Bytes([]rune("abc"))
	a.True(ok).Equal(v, []byte("abc"))

	v, ok = Bytes(123)
	a.False(ok).Empty(v)
}

func TestRunes(t *testing.T) {
	a := assert.New(t, false)

	v, ok := Runes("abc")
	a.True(ok).Equal(v, []rune("abc"))
	v, ok = Runes([]byte("abc"))
	a.True(ok).Equal(v, []rune("abc"))
	v, ok = Runes([]rune("abc"))
	a.True(ok).Equal(v, []rune("abc"))

	v, ok = Runes(123)
	a.False(ok).Empty(v)
}
