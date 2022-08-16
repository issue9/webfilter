// SPDX-License-Identifier: MIT

package convert

// String 将 val 转换成 string
func String(val any) (string, bool) {
	switch v := val.(type) {
	case string:
		return v, true
	case []byte:
		return string(v), true
	case []rune:
		return string(v), true
	default:
		return "", false
	}
}

// Bytes 将 val 转换成 []byte
func Bytes(val any) ([]byte, bool) {
	switch v := val.(type) {
	case string:
		return []byte(v), true
	case []byte:
		return v, true
	case []rune:
		return []byte(string(v)), true
	default:
		return nil, false
	}
}

// Runes 将 val 转换成 []rune
func Runes(val any) ([]rune, bool) {
	switch v := val.(type) {
	case string:
		return []rune(v), true
	case []byte:
		return []rune(string(v)), true
	case []rune:
		return v, true
	default:
		return nil, false
	}
}
