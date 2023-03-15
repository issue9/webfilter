// SPDX-License-Identifier: MIT

package validator

import (
	"math"
	"testing"

	"github.com/issue9/assert/v3"
)

func TestRange(t *testing.T) {
	a := assert.New(t, false)

	a.Panic(func() {
		Between(100, 5)
	})

	a.True(BetweenEqual(5, math.MaxInt16).IsValid(5))
	a.True(BetweenEqual(5.0, 100.0).IsValid(5.1))
	a.False(BetweenEqual(5, 100).IsValid(200))
	a.False(BetweenEqual(5, 100).IsValid(-1))
	a.False(BetweenEqual(5.0, 100.0).IsValid(-1.1))

	r := GreatEqual(6)
	a.True(r.IsValid(6))
	a.True(r.IsValid(10))
	a.False(r.IsValid(5))

	r = LessEqual(6)
	a.True(r.IsValid(6))
	a.False(r.IsValid(10))
	a.True(r.IsValid(5))
}
