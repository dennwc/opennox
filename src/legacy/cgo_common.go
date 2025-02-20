package legacy

import (
	"math"
	"unsafe"

	"github.com/gotranspile/cxgo/runtime/libc"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
)

var _ = [1]struct{}{}[16-unsafe.Sizeof(binfile.MemFile{})]

type size_t = uint32
type long = int32

type int4 struct {
	field_0 int32
	field_4 int32
	field_8 int32
	field_C int32
}

type float4 struct {
	field_0 float32
	field_4 float32
	field_8 float32
	field_C float32
}

type Point32 struct {
	X int32
	Y int32
}

func bool2int32(v bool) int32 {
	if v {
		return 1
	}
	return 0
}

func COERCE_FLOAT(x uint32) float32 {
	return math.Float32frombits(x)
}

func memset32(x *uint32, y uint32, z uint32) {
	arr := unsafe.Slice(x, z)
	for i := range arr {
		arr[i] = y
	}
}

func atoi(s *byte) int {
	return libc.Atoi(GoString(s))
}

func atof(s *byte) float64 {
	return libc.Atof(GoString(s))
}
