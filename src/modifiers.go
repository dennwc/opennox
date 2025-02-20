package opennox

import (
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/spell"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	modLog = log.New("modifiers")
)

func nox_xxx_modifGetDescById_413330(a1 int32) *server.ModifierEff {
	return noxServer.Modif.Nox_xxx_modifGetDescById413330(a1)
}

func nox_xxx_getProjectileClassById_413250(a1 int32) *server.Modifier {
	return noxServer.Modif.Nox_xxx_getProjectileClassById413250(a1)
}

func nox_xxx_equipClothFindDefByTT_413270(a1 int32) *server.Modifier {
	return noxServer.Modif.Nox_xxx_equipClothFindDefByTT413270(a1)
}

func sub_4A5E90_A() {
	for it := noxServer.Modif.Dword_5d4594_251608; it != nil; it = it.Next80 {
		switch it.Name() {
		case "StreetPants":
			legacy.Set_dword_5d4594_1308156(it.C())
		case "StreetShirt":
			legacy.Set_dword_5d4594_1308160(it.C())
		case "StreetSneakers":
			legacy.Set_dword_5d4594_1308164(it.C())
		}
	}
}

func nox_xxx_equipWeapon_4131A0() {
	if *memmap.PtrUint32(0x5D4594, 251616) != 1 {
		for it := noxServer.Modif.Dword_5d4594_251600; it != nil; it = it.Next80 {
			var ind int
			if noxflags.HasGame(noxflags.GameHost | noxflags.GameFlag22) {
				ind = noxServer.Types.IndByID(it.Name())
			} else {
				if !noxflags.HasGame(noxflags.GameClient) {
					return
				}
				ind = noxClient.Things.IndByID(it.Name())
			}
			it.Ind4 = uint32(ind)
		}
		for j := noxServer.Modif.Dword_5d4594_251608; j != nil; j = j.Next80 {
			var ind int
			if noxflags.HasGame(noxflags.GameHost | noxflags.GameFlag22) {
				ind = noxServer.Types.IndByID(j.Name())
			} else {
				if !noxflags.HasGame(noxflags.GameClient) {
					return
				}
				ind = noxClient.Things.IndByID(j.Name())
			}
			j.Ind4 = uint32(ind)
		}
		*memmap.PtrUint32(0x5D4594, 251616) = 1
	}
}

func nox_xxx_fireEffect_4E0550(a1 unsafe.Pointer, a2, src, targ *server.Object, a5, a6 unsafe.Pointer) {
	v5 := *(*float32)(unsafe.Add(a1, 56))
	if targ != nil {
		targ.CallDamage(src, a2, int(v5), object.DamageExplosion)
		nox_xxx_netSparkExplosionFx_5231B0(targ.PosVec, byte(int8(int64(float64(v5)*10.0))))
		noxServer.AudioEventObj(sound.SoundWeaponEffectFire, targ, 0, 0)
	}
}

func nox_xxx_fireRingEffect_4E05B0(a1 unsafe.Pointer, a2p, src, a4p *server.Object, a5, a6 unsafe.Pointer) {
	if src != nil {
		sa, free := alloc.New(server.SpellAcceptArg{})
		defer free()
		sa.Obj = nil
		sa.Pos = src.PosVec
		legacy.Nox_xxx_spellCastCleansingFlame_52D5C0(spell.SPELL_CLEANSING_FLAME, src, src, src, sa, int(*(*uint32)(unsafe.Add(a1, 48))))
	}
}

func nox_xxx_blueFREffect_4E05F0(a1 unsafe.Pointer, a2p, src, a4p *server.Object, a5, a6 unsafe.Pointer) {
	if src != nil {
		sa, free := alloc.New(server.SpellAcceptArg{})
		defer free()
		sa.Obj = nil
		sa.Pos = src.PosVec
		legacy.Nox_xxx_spellCastCleansingFlame_52D5C0(spell.SPELL_CLEANSING_MANA_FLAME, src, src, src, sa, int(*(*uint32)(unsafe.Add(a1, 48))))
	}
}
