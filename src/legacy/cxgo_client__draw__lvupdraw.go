package legacy

import (
	"unsafe"
)

func nox_thing_falling_sparks_draw_4B7740(a1 int32, a2 int32, dr *nox_drawable) *uint32 {
	var (
		a3     int32 = int32(uintptr(unsafe.Pointer(dr)))
		v3     *int2
		v4     int32
		v5     int32
		v6     int32
		v7     int8
		result *uint32
		a2a    int2
	)
	v3 = (*int2)(unsafe.Pointer(uintptr(a3 + 12)))
	v4 = 2
	for {
		a2a.field_0 = v3.field_0 + nox_common_randomIntMinMax_415FF0(-15, 15, internCStr("C:\\NoxPost\\src\\client\\Draw\\LvUpDraw.c"), 35)
		v5 = int32(*(*uint32)(unsafe.Pointer(uintptr(a3 + 16))) + uint32(nox_common_randomIntMinMax_415FF0(-15, 15, internCStr("C:\\NoxPost\\src\\client\\Draw\\LvUpDraw.c"), 36)))
		a2a.field_4 = v5
		v6 = int32(uint32(v5) - *(*uint32)(unsafe.Pointer(uintptr(a2 + 20))))
		v7 = int8(nox_common_randomIntMinMax_415FF0(8, 12, internCStr("C:\\NoxPost\\src\\client\\Draw\\LvUpDraw.c"), 40))
		result = (*uint32)(unsafe.Pointer(uintptr(sub_499950(a1, &a2a, v3, uint16(int16(v6)), int8(int32(-v7))))))
		v4--
		if v4 == 0 {
			break
		}
	}
	return result
}