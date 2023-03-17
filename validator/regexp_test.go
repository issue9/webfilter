// SPDX-License-Identifier: MIT

package validator

import (
	"testing"

	"github.com/issue9/assert/v3"
)

func TestCNPhone(t *testing.T) {
	a := assert.New(t, false)

	a.True(CNPhone("444488888888-4444"))
	a.True(CNPhone("3337777777-1"))
	a.True(CNPhone("333-7777777-1"))
	a.True(CNPhone("7777777"))
	a.True(CNPhone("88888888"))

	a.False(CNPhone("333-7777777-"))      // 尾部没有分机号
	a.False(CNPhone("22-88888888"))       // 区号只有2位
	a.False(CNPhone("33-88888888-55555")) // 分机号超过4位
}

func TestCNMobile(t *testing.T) {
	a := assert.New(t, false)

	a.True(CNMobile("15011111111"))
	a.True(CNMobile("015011111111"))
	a.True(CNMobile("8615011111111"))
	a.True(CNMobile("+8615011111111"))
	a.True(CNMobile("+8619911111111"))

	a.False(CNMobile("+86150111111112")) // 尾部多个2
	a.False(CNMobile("50111111112"))     // 开头少1
	a.False(CNMobile("+8650111111112"))  // 开头少1
	a.False(CNMobile("8650111111112"))   // 开头少1
	a.False(CNMobile("154111111112"))    // 不存在的前缀154
	a.True(CNMobile("15411111111"))
}

func TestCNTel(t *testing.T) {
	a := assert.New(t, false)

	a.True(CNTel("444488888888-4444"))
	a.True(CNTel("3337777777-1"))
	a.True(CNTel("333-7777777-1"))
	a.True(CNTel("7777777"))
	a.True(CNTel("88888888"))
	a.True(CNTel("15011111111"))
	a.True(CNTel("015011111111"))
	a.True(CNTel("8615011111111"))
	a.True(CNTel("+8615011111111"))
}
