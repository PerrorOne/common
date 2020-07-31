package string

import "unsafe"

// 使用更快速、不拷贝内存的方式string转bytes
// 这是不安全的!请不要在生产项目中使用!会导致未知的BUG!
// Deprecated: use []byte()
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 裁剪字符串
func ClipString(text string, number int) string {
	if number <= 0 {
		return ""
	}

	data := []rune(text)
	if len(data) <= number {
		return text
	}
	return string(data[:number])
}
