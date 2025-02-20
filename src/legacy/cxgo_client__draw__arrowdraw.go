package legacy

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/client"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func nox_thing_arrow_draw(vp *noxrender.Viewport, dr *client.Drawable) int {
	var (
		v2 int32
		v3 int32
		v4 *uint32
		v5 int32
		a2 *uint32 = (*uint32)(dr.C())
	)
	v2 = int32(*memmap.PtrUint32(0x5D4594, 1313720))
	if *memmap.PtrUint32(0x5D4594, 1313720) == 0 {
		v2 = nox_xxx_getTTByNameSpriteMB_44CFC0(internCStr("ArrowTailLink"))
		*memmap.PtrUint32(0x5D4594, 1313720) = uint32(v2)
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*81)))
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))-uint32(v3))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))-uint32(v3))+(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*82)))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*82))) <= 200 {
		return nox_thing_slave_draw(vp, dr)
	}
	v4 = (*uint32)(nox_xxx_spriteLoadAdd_45A360_drawable(v2, v3, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*82)))).C())
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*108)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*109)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4))
	nox_xxx_sprite_45A110_drawable((*client.Drawable)(unsafe.Pointer(v4)))
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*81)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*82)) = uint32(v5)
	nox_xxx_spriteTransparentDecay_49B950((*client.Drawable)(unsafe.Pointer(v4)), int32(gameFPS()/3))
	return nox_thing_slave_draw(vp, dr)
}
func nox_thing_weak_arrow_draw(vp *noxrender.Viewport, dr *client.Drawable) int {
	var (
		v2 int32
		v3 int32
		v4 *uint32
		v5 int32
		a2 *uint32 = (*uint32)(dr.C())
	)
	v2 = int32(*memmap.PtrUint32(0x5D4594, 1313724))
	if *memmap.PtrUint32(0x5D4594, 1313724) == 0 {
		v2 = nox_xxx_getTTByNameSpriteMB_44CFC0(internCStr("WeakArrowTailLink"))
		*memmap.PtrUint32(0x5D4594, 1313724) = uint32(v2)
	}
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*81)))
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))-uint32(v3))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))-uint32(v3))+(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*82)))*(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4))-*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*82))) <= 200 {
		return nox_thing_slave_draw(vp, dr)
	}
	v4 = (*uint32)(nox_xxx_spriteLoadAdd_45A360_drawable(v2, v3, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*82)))).C())
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*108)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*109)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4))
	nox_xxx_sprite_45A110_drawable((*client.Drawable)(unsafe.Pointer(v4)))
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*81)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*3))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*82)) = uint32(v5)
	nox_xxx_spriteTransparentDecay_49B950((*client.Drawable)(unsafe.Pointer(v4)), int32(gameFPS()/3))
	return nox_thing_slave_draw(vp, dr)
}
func nox_thing_arrow_tail_link_draw(vp *noxrender.Viewport, dr *client.Drawable) int {
	a1 := (*uint32)(vp.C())
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		a2 int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v2 = int32(*(*uint32)(unsafe.Add(a2, 16)))
	v3 = int32(*(*uint32)(unsafe.Add(a2, 12)) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	v4 = int32(uint32(v2) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*1)) - uint32(*(*int16)(unsafe.Add(a2, 106))) - uint32(*(*int16)(unsafe.Add(a2, 104))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)))
	v5 = int32(*(*uint32)(unsafe.Add(a2, 436)) - uint32(v2))
	v6 = int32(*(*uint32)(unsafe.Add(a2, 432)) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	if *(*uint32)(unsafe.Add(a2, 356))-gameFrame() <= 0 {
		return 1
	}
	v7 = int32(((*(*uint32)(unsafe.Add(a2, 356)) - gameFrame()) << 6) / (gameFPS() / 3))
	if v7 >= 64 {
		v7 = 63
	}
	nox_client_drawSetColor_434460(int32(*memmap.PtrUint32(0x5D4594, uintptr(v7*4)+1313012)))
	nox_client_drawEnableAlpha_434560(1)
	nox_client_drawSetAlpha_434580(0x80)
	nox_client_drawAddPoint_49F500(v3, v4-4)
	nox_client_drawAddPoint_49F500(v6, v5+v4-4)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawEnableAlpha_434560(0)
	return 1
}
func nox_thing_weak_arrow_tail_link_draw(vp *noxrender.Viewport, dr *client.Drawable) int {
	a1 := (*uint32)(vp.C())
	var (
		v2 int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 int32
		a2 int32 = int32(uintptr(unsafe.Pointer(dr)))
	)
	v2 = int32(*(*uint32)(unsafe.Add(a2, 16)))
	v3 = int32(*(*uint32)(unsafe.Add(a2, 12)) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	v4 = int32(uint32(v2) + *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*1)) - uint32(*(*int16)(unsafe.Add(a2, 106))) - uint32(*(*int16)(unsafe.Add(a2, 104))) - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)))
	v5 = int32(*(*uint32)(unsafe.Add(a2, 436)) - uint32(v2))
	v6 = int32(*(*uint32)(unsafe.Add(a2, 432)) + *a1 - *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	if *(*uint32)(unsafe.Add(a2, 356))-gameFrame() <= 0 {
		return 1
	}
	v7 = int32(((*(*uint32)(unsafe.Add(a2, 356)) - gameFrame()) << 6) / (gameFPS() / 3))
	if v7 >= 64 {
		v7 = 63
	}
	nox_client_drawSetColor_434460(int32(*memmap.PtrUint32(0x5D4594, uintptr(v7*4)+1313268)))
	nox_client_drawEnableAlpha_434560(1)
	nox_client_drawSetAlpha_434580(0x80)
	nox_client_drawAddPoint_49F500(v3, v4-4)
	nox_client_drawAddPoint_49F500(v6, v5+v4-4)
	nox_client_drawLineFromPoints_49E4B0()
	nox_client_drawEnableAlpha_434560(0)
	return 1
}
