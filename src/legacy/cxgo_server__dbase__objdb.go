package legacy

import (
	"unsafe"

	"github.com/gotranspile/cxgo/runtime/libc"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/server"
)

func sub_4E39F0_obj_db(a1p *server.Object) *wchar2_t {
	var (
		a1 **byte = (**byte)(unsafe.Pointer(a1p))
		v1 *byte
		v2 *byte
		v3 *byte
		v4 *byte
		v5 uint8
		v6 *uint8
		i  *uint8
	)
	v1 = *a1
	if *a1 == nil {
		v1 = nox_xxx_getUnitName_4E39D0((*server.Object)(unsafe.Pointer(a1)))
	}
	v2 = libc.StrRChr(v1, 58)
	if v2 != nil {
		v3 = (*byte)(unsafe.Add(unsafe.Pointer(v2), 1))
	} else {
		v3 = v1
	}
	v4 = v3
	libc.StrCpy((*byte)(memmap.PtrOff(0x5D4594, 1563460)), internCStr("NPC:"))
	v5 = *v3
	v6 = (*uint8)(memmap.PtrOff(0x5D4594, uintptr(libc.StrLen((*byte)(memmap.PtrOff(0x5D4594, 1563460)))+1563460)))
	for i = (*byte)(unsafe.Add(unsafe.Pointer(v4), 1)); int32(v5) != 0; i = (*uint8)(unsafe.Add(unsafe.Pointer(i), 1)) {
		if int32(v5) != 95 {
			*func() *uint8 {
				p := &v6
				x := *p
				*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
				return x
			}() = v5
		}
		v5 = *i
	}
	*v6 = 0
	return nox_strman_loadString_40F1D0((*byte)(memmap.PtrOff(0x5D4594, 1563460)), nil, internCStr("C:\\NoxPost\\src\\Server\\DBase\\objdb.c"), 1577)
}
