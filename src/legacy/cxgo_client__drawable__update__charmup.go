package legacy

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
)

func sub_4CD150(a1 int32, a2 *noxrender.Viewport, a3 *client.Drawable, a4 int32) {
	var (
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 *uint32
		v13 int32
		v14 *uint32
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 int32
		v23 int8
		v25 int32
		v26 int32
		v27 int32
	)
	v4 = nox_common_randomIntMinMax_415FF0(0, 100, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 16)
	if v4 >= 50 {
		return
	}
	*((*uint8)(unsafe.Pointer(&v4))) = a3.Field_108_0
	if int32(uint8(int8(v4))) == 0 {
		if a3 == -432 {
			return
		}
		if a4 != 0 {
			v5 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a3), 441))) + nox_common_randomIntMinMax_415FF0(-20, 20, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 28)
			v6 = nox_common_randomIntMinMax_415FF0(-20, 20, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 29)
			*(*uint16)(unsafe.Add(unsafe.Pointer(&v7), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Add(unsafe.Pointer(a3), 437))
			v8 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a3), 443))) + v6
			*(*uint16)(unsafe.Add(unsafe.Pointer(&v9), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Add(unsafe.Pointer(a3), 439))
		} else {
			v5 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a3), 437))) + nox_common_randomIntMinMax_415FF0(-20, 20, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 36)
			v10 = nox_common_randomIntMinMax_415FF0(-20, 20, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 37)
			*(*uint16)(unsafe.Add(unsafe.Pointer(&v7), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Add(unsafe.Pointer(a3), 441))
			v8 = int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a3), 439))) + v10
			*(*uint16)(unsafe.Add(unsafe.Pointer(&v9), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Add(unsafe.Pointer(a3), 443))
		}
		*(*uint16)(unsafe.Add(unsafe.Pointer(&v25), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v9))
		goto LABEL_20
	}
	if a3 == -432 {
		return
	}
	if nox_xxx_netTestHighBit_578B70(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 437))) != 0 {
		v11 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 437)))))
		v12 = (*uint32)(nox_xxx_netSpriteByCodeStatic_45A720(uint32(v11)).C())
	} else {
		v13 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 437)))))
		v12 = (*uint32)(nox_xxx_netSpriteByCodeDynamic_45A6F0(uint32(v13)).C())
	}
	v14 = v12
	if nox_xxx_netTestHighBit_578B70(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 441))) != 0 {
		v15 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 441)))))
		v4 = int32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeStatic_45A720(uint32(v15)))))
	} else {
		v16 = nox_xxx_netClearHighBit_578B30(int16(uint16(*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 441)))))
		v4 = int32(uintptr(unsafe.Pointer(nox_xxx_netSpriteByCodeDynamic_45A6F0(uint32(v16)))))
	}
	v17 = v4
	if v14 == nil || v4 == 0 {
		return
	}
	if a4 == 0 {
		v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), 4*3)) + uint32(nox_common_randomIntMinMax_415FF0(-20, 20, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 80)))
		v19 = nox_common_randomIntMinMax_415FF0(-20, 20, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 81)
		v7 = int32(*(*uint32)(unsafe.Add(v17, 12)))
		v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), 4*4)) + uint32(v19))
		v9 = int32(*(*uint32)(unsafe.Add(v17, 16)))
		*(*uint16)(unsafe.Add(unsafe.Pointer(&v25), unsafe.Sizeof(uint16(0))*0)) = uint16(int16(v9))
		goto LABEL_20
	}
	v5 = int32(*(*uint32)(unsafe.Add(v4, 12)) + uint32(nox_common_randomIntMinMax_415FF0(-20, 20, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 70)))
	v18 = nox_common_randomIntMinMax_415FF0(-20, 20, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 71)
	v7 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), 4*3)))
	v8 = int32(*(*uint32)(unsafe.Add(v17, 16)) + uint32(v18))
	v25 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v14), 4*4)))
LABEL_20:
	v20 = int32(uint32(v5) + *a2 - a2.World.Min.X)
	v27 = int32(a3.ZVal)
	v26 = int32(a2.World.Min.Y)
	v21 = int32(uint32(v8) + a2.Screen.Min.Y - uint32(v26) - uint32(v27))
	if v20 < 0 {
		v5 = int32(a2.World.Min.X + *a2 + 1)
	}
	if v21 < 0 {
		v8 = int32(a2.Screen.Min.Y + int(uint32(v26)) - int(uint32(v27)) + 1)
	}
	if uint32(v20) >= a2.Size.X {
		v5 = int32(a2.Screen.Max.X + a2.World.Min.X - 1)
	}
	if uint32(v21) >= a2.Size.Y {
		v8 = int32(a2.Screen.Max.Y - int(uint32(v27)) + int(uint32(v26)) - 1)
	}
	v4 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(a1, v5, v8))))
	v22 = v4
	if v4 != 0 {
		v23 = int8(nox_common_randomIntMinMax_415FF0(6, 12, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 112))
		*(*uint16)(unsafe.Add(v22, 432)) = uint16(int16(v7))
		*(*uint16)(unsafe.Add(v22, 434)) = uint16(int16(v25))
		*(*uint8)(unsafe.Add(v22, 443)) = uint8(v23)
		*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(nox_common_randomIntMinMax_415FF0(3, 10, internCStr("C:\\NoxPost\\src\\client\\Drawable\\Update\\charmup.c"), 119)))
		*(*uint8)(unsafe.Add(v22, 444)) = uint8(int8(v4))
	}
	return
}
