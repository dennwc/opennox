package legacy

import (
	"bytes"
	"unicode/utf16"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

type wchar2_t = uint16

func StrFree[T comparable](s *T) {
	alloc.Free(s)
}

func CStringArray(arr []string) []*byte {
	out := make([]*byte, 0, len(arr)+1)
	for _, arg := range arr {
		out = append(out, CString(arg))
	}
	out = append(out, nil)
	return out[:len(arr):len(arr)]
}

func StrLenBytes(s []byte) int {
	i := bytes.IndexByte(s, 0)
	if i < 0 {
		return len(s)
	}
	return i
}

func StrCopy(dst *byte, max int, src string) int {
	d := unsafe.Slice(dst, max)
	return StrCopyBytes(d, src)
}

func StrNCopy(dst *byte, max int, src string) int {
	d := unsafe.Slice(dst, max)
	return StrNCopyBytes(d, src)
}

func StrCopyP(dst unsafe.Pointer, max int, src string) int {
	return StrCopy((*byte)(dst), max, src)
}

func StrCopyBytes(dst []byte, src string) int {
	n := copy(dst[:len(dst)-1], src)
	dst[n] = 0
	return n
}

func StrNCopyBytes(dst []byte, src string) int {
	n := copy(dst, src)
	if n < len(dst) {
		dst[n] = 0
	}
	return n
}

func WStrLenBytes(s []byte) int {
	if s == nil {
		return 0
	}
	n := 0
	for len(s) >= 2 && s[0] != 0 && s[1] != 0 {
		s = s[2:]
		n++
	}
	return n
}

func WStrCopy(p *wchar2_t, max int, src string) int {
	dst := unsafe.Slice(p, max)
	val := utf16.Encode([]rune(src))
	n := copy(dst, val)
	if n < len(dst) {
		dst[n] = 0
	}
	return n
}

func WStrCopyBytes(p []byte, src string) int {
	dst := unsafe.Slice((*uint16)(unsafe.Pointer(&p[0])), len(p)/2)
	val := utf16.Encode([]rune(src))
	n := copy(dst, val)
	if n < len(dst) {
		dst[n] = 0
	}
	return n
}

func WStrCopySlice(dst []uint16, src string) int {
	val := utf16.Encode([]rune(src))
	n := copy(dst, val)
	if n < len(dst) {
		dst[n] = 0
	}
	return n
}

func GoString(s *byte) string {
	return alloc.GoString(s)
}

func GoStringP(s unsafe.Pointer) string {
	return GoString((*byte)(s))
}

func GoStringN(s *byte, n int) string {
	return GoStringNP(unsafe.Pointer(s), n)
}

func GoStringNP(s unsafe.Pointer, n int) string {
	b := unsafe.Slice((*byte)(s), n)
	i := bytes.IndexByte(b, 0)
	if i < 0 {
		return string(b)
	}
	return string(b[:i])
}

func GoStringS(s []byte) string {
	return string(s[:StrLenBytes(s)])
}

func CString(s string) *byte {
	c, _ := alloc.CString(s)
	return c
}

func CBytes(s []byte) unsafe.Pointer {
	b, _ := alloc.Make([]byte{}, len(s))
	copy(b, s)
	return unsafe.Pointer(&b[0])
}

func GoWStringP(s unsafe.Pointer) string {
	return GoWString((*wchar2_t)(s))
}

func GoWString(s *wchar2_t) string {
	n := alloc.StrLen(s)
	if n == 0 {
		return ""
	}
	arr := unsafe.Slice(s, n)
	return GoWStringSlice(arr)
}

func GoWStringN(s *wchar2_t, max int) string {
	n := alloc.StrLenN(s, max)
	if n == 0 {
		return ""
	}
	arr := unsafe.Slice(s, n)
	return GoWStringSlice(arr)
}

func GoWStringSlice(arr []uint16) string {
	return string(utf16.Decode(arr))
}

func GoWStringBytes(arr []byte) string {
	n := WStrLenBytes(arr)
	if n == 0 {
		return ""
	}
	warr := unsafe.Slice((*uint16)(unsafe.Pointer(&arr[0])), n)
	return GoWStringSlice(warr)
}

func CWString(s string) (*wchar2_t, func()) {
	buf := utf16.Encode([]rune(s))
	ptr, free := alloc.Calloc(len(buf)+1, 2)
	copy(unsafe.Slice((*uint16)(ptr), len(buf)), buf)
	return (*wchar2_t)(ptr), free
}

func CWLen(s string) int {
	return len(utf16.Encode([]rune(s)))
}

func CWStringCopyTo(dst *wchar2_t, dstSz int, src string) {
	str := unsafe.Slice(dst, dstSz)
	if len(src) == 0 {
		str[0] = 0
		return
	}
	n := copy(str[:len(str)-1], utf16.Encode([]rune(src)))
	str[n] = 0
}

func GoWStrSlice(arr **wchar2_t) []string {
	n := alloc.ZeroTermLen(arr)
	return GoWStrSliceN(arr, n)
}

func GoWStrSliceN(arr **wchar2_t, n int) []string {
	if n == 0 {
		return nil
	}
	out := make([]string, 0, n)
	for _, c := range unsafe.Slice(arr, n) {
		out = append(out, GoWString(c))
	}
	return out
}

func CWStrSlice(arr []string) ([]*wchar2_t, func()) {
	out, freeList := alloc.Make([]*wchar2_t{}, len(arr)+1)
	out = out[:len(out)-1]
	for i, s := range arr {
		out[i], _ = CWString(s)
	}
	return out, func() {
		for _, p := range out {
			StrFree(p)
		}
		freeList()
	}
}
