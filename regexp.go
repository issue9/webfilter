// SPDX-License-Identifier: MIT

package validator

import "regexp"

const (
	// 匹配大陆电话
	cnPhonePattern = `((\d{3,4})-?)?` + // 区号
		`\d{5,10}` + // 号码，95500 等 5 位数的，7 位，8 位，以及 400 开头的 10 位数
		`(-\d{1,4})?` // 分机号，分机号的连接符号不能省略。

	// 匹配大陆手机号码
	cnMobilePattern = `(0|\+?86)?` + // 匹配 0,86,+86
		`1[2-9]{1}[0-9]{9}` // 12x-19x

	// 匹配大陆手机号或是电话号码
	cnTelPattern = "(" + cnPhonePattern + ")|(" + cnMobilePattern + ")"
)

var (
	cnPhone  = regexpCompile(cnPhonePattern)
	cnMobile = regexpCompile(cnMobilePattern)
	cnTel    = regexpCompile(cnTelPattern)
)

func regexpCompile(str string) *regexp.Regexp { return regexp.MustCompile("^" + str + "$") }

// CNPhone 验证中国大陆的电话号码
//
// 支持如下格式：
//
//	0578-12345678-1234
//	057812345678-1234
//
// 若存在分机号，则分机号的连接符不能省略。
func CNPhone(val any) bool { return Match(cnPhone)(val) }

// CNMobile 验证中国大陆的手机号码
func CNMobile(val any) bool { return Match(cnMobile)(val) }

// CNTel 验证手机和电话类型
func CNTel(val any) bool { return Match(cnTel)(val) }

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
