package main

/*
#include <stddef.h>
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"reflect"
	"unicode/utf16"
	"unsafe"
)

func StrFree(s *C.char) {
	C.free(unsafe.Pointer(s))
}

func BytesFree(s unsafe.Pointer) {
	C.free(s)
}

func WStrFree(s *C.wchar_t) {
	C.free(unsafe.Pointer(s))
}

func WStrSliceFree(arr []*C.wchar_t) {
	for _, p := range arr {
		WStrFree(p)
	}
	C.free(unsafe.Pointer(&arr[0]))
}

func CStringArray(arr []string) []*C.char {
	out := make([]*C.char, 0, len(arr)+1)
	for _, arg := range arr {
		out = append(out, C.CString(arg))
	}
	out = append(out, nil)
	return out[:len(arr):len(arr)]
}

func asByteSlice(p unsafe.Pointer, sz int) (out []byte) {
	*(*reflect.SliceHeader)(unsafe.Pointer(&out)) = reflect.SliceHeader{
		Data: uintptr(p),
		Len:  sz, Cap: sz,
	}
	return
}

func asU16Slice(p unsafe.Pointer, sz int) (out []uint16) {
	*(*reflect.SliceHeader)(unsafe.Pointer(&out)) = reflect.SliceHeader{
		Data: uintptr(p),
		Len:  sz, Cap: sz,
	}
	return
}

func asU32Slice(p unsafe.Pointer, sz int) (out []uint32) {
	*(*reflect.SliceHeader)(unsafe.Pointer(&out)) = reflect.SliceHeader{
		Data: uintptr(p),
		Len:  sz, Cap: sz,
	}
	return
}

func asI32Slice(p unsafe.Pointer, sz int) (out []int32) {
	*(*reflect.SliceHeader)(unsafe.Pointer(&out)) = reflect.SliceHeader{
		Data: uintptr(p),
		Len:  sz, Cap: sz,
	}
	return
}

func asF32Slice(p unsafe.Pointer, sz int) (out []float32) {
	*(*reflect.SliceHeader)(unsafe.Pointer(&out)) = reflect.SliceHeader{
		Data: uintptr(p),
		Len:  sz, Cap: sz,
	}
	return
}

func asWStr(p unsafe.Pointer, sz int) (out []C.wchar_t) {
	*(*reflect.SliceHeader)(unsafe.Pointer(&out)) = reflect.SliceHeader{
		Data: uintptr(p),
		Len:  sz, Cap: sz,
	}
	return
}

func asWStrSlice(p unsafe.Pointer, sz int) (out []*C.wchar_t) {
	*(*reflect.SliceHeader)(unsafe.Pointer(&out)) = reflect.SliceHeader{
		Data: uintptr(p),
		Len:  sz, Cap: sz,
	}
	return
}

func StrLen(s *C.char) int {
	if s == nil {
		return 0
	}
	n := 0
	for *s != 0 {
		n++
		s = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + 1))
	}
	return n
}

func StrLenBytes(s []byte) int {
	i := bytes.IndexByte(s, 0)
	if i < 0 {
		return len(s)
	}
	return i
}

func StrCopy(dst *C.char, max int, src string) int {
	d := asByteSlice(unsafe.Pointer(dst), max)
	return StrCopyBytes(d, src)
}

func StrNCopy(dst *C.char, max int, src string) int {
	d := asByteSlice(unsafe.Pointer(dst), max)
	return StrNCopyBytes(d, src)
}

func StrCopyP(dst unsafe.Pointer, max int, src string) int {
	return StrCopy((*C.char)(dst), max, src)
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

func WStrLen(s *C.wchar_t) int {
	if s == nil {
		return 0
	}
	n := 0
	for *s != 0 {
		s = (*C.wchar_t)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + 2))
		n++
	}
	return n
}

func WStrLenN(s *C.wchar_t, max int) int {
	if s == nil {
		return 0
	}
	n := 0
	for *s != 0 {
		s = (*C.wchar_t)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + 2))
		if n == max {
			return n
		}
		n++
	}
	return n
}

func GoString(s *C.char) string {
	return C.GoString(s)
}

func GoStringP(s unsafe.Pointer) string {
	return GoString((*C.char)(s))
}

func GoStringS(s []byte) string {
	return string(s[:StrLenBytes(s)])
}

func CString(s string) *C.char {
	return C.CString(s)
}

func CBytes(s []byte) unsafe.Pointer {
	return C.CBytes(s)
}

func GoWStringP(s unsafe.Pointer) string {
	return GoWString((*C.wchar_t)(s))
}

func GoWString(s *C.wchar_t) string {
	n := WStrLen(s)
	if n == 0 {
		return ""
	}
	arr := asU16Slice(unsafe.Pointer(s), n)
	return GoWStringSlice(arr)
}

func GoWStringN(s *C.wchar_t, max int) string {
	n := WStrLenN(s, max)
	if n == 0 {
		return ""
	}
	arr := asU16Slice(unsafe.Pointer(s), n)
	return GoWStringSlice(arr)
}

func GoWStringSlice(arr []uint16) string {
	return string(utf16.Decode(arr))
}

func CWString(s string) *C.wchar_t {
	buf := utf16.Encode([]rune(s))
	ptr := C.calloc(C.uint(len(buf)+1), 2)
	copy(asU16Slice(ptr, len(buf)), buf)
	return (*C.wchar_t)(ptr)
}

func CWLen(s string) int {
	return len(utf16.Encode([]rune(s)))
}

func CWStringCopyTo(dst *C.wchar_t, dstSz int, src string) {
	str := asU16Slice(unsafe.Pointer(dst), dstSz)
	if len(src) == 0 {
		str[0] = 0
		return
	}
	n := copy(str[:len(str)-1], utf16.Encode([]rune(src)))
	str[n] = 0
}

func GoWStrSlice(arr **C.wchar_t) []string {
	n := PtrArrLen(unsafe.Pointer(arr))
	return GoWStrSliceN(arr, n)
}

func GoWStrSliceN(arr **C.wchar_t, n int) []string {
	if n == 0 {
		return nil
	}
	out := make([]string, 0, n)
	for _, c := range asWStrSlice(unsafe.Pointer(arr), n) {
		out = append(out, GoWString(c))
	}
	return out
}

func CWStrSlice(arr []string) []*C.wchar_t {
	ptr := C.calloc(C.uint(len(arr)+1), C.uint(ptrSize))
	out := asWStrSlice(ptr, len(arr))
	for i, s := range arr {
		out[i] = CWString(s)
	}
	return out
}
