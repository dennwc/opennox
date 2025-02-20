package legacy

import (
	"math"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

func sub_4CE200(a1 *noxrender.Viewport, a2 *client.Drawable, a3 int32, a4 int32) int32 {
	var (
		result int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     bool
		v10    int32
		v11    int32
		v12    int32
		v13    int32
	)
	if *memmap.PtrUint32(0x5D4594, 1523016) == 0 {
		*memmap.PtrUint32(0x5D4594, 1523016) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0(internCStr("GreenPuff")))
		*memmap.PtrUint32(0x5D4594, 1523020) = uint32(nox_xxx_getTTByNameSpriteMB_44CFC0(internCStr("GreenSmoke")))
	}
	result = a3
	if a3 <= 0 {
		return result
	}
	for {
		v5 = nox_common_randomIntMinMax_415FF0(0, math.MaxUint8, internCStr("C:\\NoxPost\\src\\Client\\Drawable\\Update\\cloud.c"), 40)
		v6 = nox_common_randomIntMinMax_415FF0(0, a4, internCStr("C:\\NoxPost\\src\\Client\\Drawable\\Update\\cloud.c"), 41)
		v7 = int32(a2.PosVec.X + int(uint32(v6**memmap.PtrInt32(0x587000, uintptr(v5*8)+192088)/16)))
		v8 = int32(a2.PosVec.Y + int(uint32(v6**memmap.PtrInt32(0x587000, uintptr(v5*8)+192092)/16)))
		v9 = nox_common_randomIntMinMax_415FF0(0, 10, internCStr("C:\\NoxPost\\src\\Client\\Drawable\\Update\\cloud.c"), 46) < 3
		v10 = int32(*memmap.PtrUint32(0x5D4594, 1523020))
		if !v9 {
			v10 = int32(*memmap.PtrUint32(0x5D4594, 1523016))
		}
		v11 = int32(uintptr(unsafe.Pointer(nox_xxx_spriteLoadAdd_45A360_drawable(v10, v7, v8))))
		v12 = v11
		if v11 != 0 {
			*(*uint16)(unsafe.Add(v11, 104)) = 0
			nox_xxx_sprite_45A110_drawable((*client.Drawable)(v11))
			*(*uint8)(unsafe.Add(v12, 432)) = uint8(int8(nox_common_randomIntMinMax_415FF0(1, 3, internCStr("C:\\NoxPost\\src\\Client\\Drawable\\Update\\cloud.c"), 57)))
			*(*uint32)(unsafe.Add(v12, 460)) = uint32(uintptr(ccall.FuncAddr(sub_4CE340)))
			v13 = nox_common_randomIntMinMax_415FF0(10, 32, internCStr("C:\\NoxPost\\src\\Client\\Drawable\\Update\\cloud.c"), 61)
			nox_xxx_spriteTransparentDecay_49B950((*client.Drawable)(v12), v13)
			nox_xxx_spriteToList_49BC80_drawable((*client.Drawable)(v12))
			nox_xxx_spriteToSightDestroyList_49BAB0_drawable((*client.Drawable)(v12))
		}
		result = func() int32 {
			p := &a3
			*p--
			return *p
		}()
		if a3 == 0 {
			break
		}
	}
	return result
}
