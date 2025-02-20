package legacy

import (
	"math"
	"unsafe"

	"github.com/gotranspile/cxgo/runtime/libc"
	"github.com/gotranspile/cxgo/runtime/stdio"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/types"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/unit/ai"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/server"
)

var nox_file_8 *FILE = nil
var nox_cheat_charmall int32 = 0
var dword_5d4594_1599480 uint32 = 0
var dword_5d4594_1599476 uint32 = 0
var dword_5d4594_1599540 unsafe.Pointer = nil
var dword_5d4594_1599556 unsafe.Pointer = nil
var dword_5d4594_1599548 unsafe.Pointer = nil
var dword_5d4594_1599588 unsafe.Pointer = nil
var dword_5d4594_1599592 unsafe.Pointer = nil

func nox_xxx_XFerSpellReward_4F5F30(a1p *server.Object, data unsafe.Pointer) int {
	a1 := a1p
	var (
		v1     *uint8
		result int32
		v3     uint8
		v4     uint8
		v5     uint8
		v6     uint8
		v7     int32
		v8     uint8
		v9     uint8
		v10    int32
		v11    int8
		v12    int8
		v13    int32
		v14    [128]byte
	)
	v1 = (*uint8)(a1.UseData)
	v13 = int32(a1.Field34)
	v10 = 60
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v10)), 2)
	if int32(int16(v10)) > 60 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(a1, int32(int16(v10)))
	if result == 0 {
		return int(result)
	}
	if int32(int16(v10)) >= 31 {
		if nox_crypt_IsReadOnly() == 1 {
			if int32(int16(v10)) >= 41 {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v7)), 1)
				if int32(uint8(int8(v7))) >= 0x80 {
					return 0
				}
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v14[0], uint32(uint8(int8(v7))))
				v14[uint8(int8(v7))] = 0
				*v1 = uint8(int8(nox_xxx_spellNameToN_4243F0(&v14[0])))
			} else {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v7)), 1)
				if int32(uint8(int8(v7))) >= 0x80 {
					return 0
				}
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v14[0], uint32(uint8(int8(v7))))
				v14[uint8(int8(v7))] = 0
				nox_xxx_spellNameToN_4243F0(&v14[0])
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v7)), 1)
				if int32(uint8(int8(v7))) >= 0x80 {
					return 0
				}
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v14[0], uint32(uint8(int8(v7))))
				v14[uint8(int8(v7))] = 0
				v5 = uint8(int8(nox_xxx_spellNameToN_4243F0(&v14[0])))
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v7)), 1)
				if int32(uint8(int8(v7))) >= 0x80 {
					return 0
				}
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v14[0], uint32(uint8(int8(v7))))
				v14[uint8(int8(v7))] = 0
				v6 = uint8(int8(nox_xxx_spellNameToN_4243F0(&v14[0])))
				*v1 = v5
				if int32(v6) != 0 {
					*v1 = v6
				}
			}
			goto LABEL_28
		}
	} else if nox_crypt_IsReadOnly() == 1 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v12)), 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v9, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v8, 1)
		v3 = v9
		if int32(v9) >= 0x89 {
			v3 = 0
			v9 = 0
		}
		v4 = v8
		if int32(v8) >= 0x89 {
			v4 = 0
			v8 = 0
		}
		*v1 = v3
		if int32(v4) != 0 {
			*v1 = v4
		}
		if int32(uint16(int16(v10))) == 10 {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v11)), 1)
		}
		goto LABEL_28
	}
	libc.StrCpy(&v14[0], nox_xxx_spellNameByN_424870(int32(*v1)))
	*(*uint8)(unsafe.Pointer(&v7)) = uint8(int8(libc.StrLen(&v14[0])))
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v7)), 1)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v14[0], uint32(uint8(int8(v7))))
LABEL_28:
	if a1.Field34 == 0 || nox_crypt_IsReadOnly() != 1 || (func() int32 {
		result = nox_xxx_xfer_4F3E30(uint16(int16(v10)), a1, int32(a1.Field34))
		return result
	}()) != 0 {
		a1.Field34 = uint32(v13)
		result = 1
	}
	return int(result)
}
func Nox_xxx_XFerAbilityReward_4F6240(a1p *server.Object, data unsafe.Pointer) int {
	a1 := a1p
	var (
		v1     *uint8
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     [128]byte
	)
	v1 = (*uint8)(a1.UseData)
	v5 = int32(a1.Field34)
	v4 = 61
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v4)), 2)
	if int32(int16(v4)) > 61 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(a1, int32(int16(v4)))
	if result == 0 {
		return int(result)
	}
	libc.StrCpy(&v6[0], nox_xxx_abilityGetName_425250(int32(*v1)))
	*(*uint8)(unsafe.Pointer(&v3)) = uint8(int8(libc.StrLen(&v6[0])))
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v3)), 1)
	if int32(uint8(int8(v3))) < 0x80 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v6[0], uint32(uint8(int8(v3))))
		v6[uint8(int8(v3))] = 0
		*v1 = uint8(int8(nox_xxx_abilityNameToN_424D80(&v6[0])))
		if a1.Field34 == 0 || nox_crypt_IsReadOnly() != 1 || (func() int32 {
			result = nox_xxx_xfer_4F3E30(uint16(int16(v4)), a1, int32(a1.Field34))
			return result
		}()) != 0 {
			a1.Field34 = uint32(v5)
			result = 1
		}
	} else {
		result = 0
	}
	return int(result)
}
func Nox_xxx_XFerFieldGuide_4F6390(a1p *server.Object, data unsafe.Pointer) int {
	var (
		v2     *byte
		v3     int32
		result int32
		v5     int32
	)
	v1 := a1p
	v2 = (*byte)(a1p.UseData)
	v3 = int32(a1p.Field34)
	v5 = 60
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v5)), 2)
	if int32(int16(v5)) > 60 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(v1, int32(int16(v5)))
	if result == 0 {
		return int(result)
	}
	if nox_crypt_IsReadOnly() != 0 {
		var a1 int32
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&a1)), 1)
		if a1 >= 0x40 {
			return 0
		}
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v2, uint32(uint8(a1)))
		*(*byte)(unsafe.Add(unsafe.Pointer(v2), uint8(a1))) = 0
	} else {
		var a1 int32
		*(*uint8)(unsafe.Pointer(&a1)) = uint8(int8(libc.StrLen(v2)))
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&a1)), 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v2, uint32(uint8(a1)))
	}
	if v1.Field34 == 0 || nox_crypt_IsReadOnly() != 1 || (func() int32 {
		result = nox_xxx_xfer_4F3E30(uint16(int16(v5)), v1, int32(v1.Field34))
		return result
	}()) != 0 {
		v1.Field34 = uint32(v3)
		return 1
	}
	return int(result)
}
func nox_xxx_XFerWeapon_4F64A0(a1p *server.Object, data unsafe.Pointer) int {
	a1 := a1p
	var (
		result int32
		v2     ***byte
		v3     int32
		v4     int32
		v5     *byte
		v6     int32
		v7     int32
		v9     uint8
		v10    int32
		v11    int32
		v12    uint32
		v14    uint8
		v15    uint8
		v16    int32
		v17    int32
		v18    int32
		v19    int32
		v20    [20]byte
		v21    [256]byte
	)
	v19 = int32(a1.Field34)
	v17 = 64
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v17)), 2)
	if int32(int16(v17)) > 64 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(a1, int32(int16(v17)))
	if result == 0 {
		return 0
	}
	if int32(int16(v17)) < 11 && nox_crypt_IsReadOnly() == 1 {
		*(*uint32)(unsafe.Pointer(&v20[0])) = 0
		*(*uint32)(unsafe.Pointer(&v20[4])) = 0
		*(*uint32)(unsafe.Pointer(&v20[8])) = 0
		*(*uint32)(unsafe.Pointer(&v20[12])) = 0
		nox_xxx_modifSetItemAttrs_4E4990(a1, (*int32)(unsafe.Pointer(&v20[0])))
		return 1
	}
	if nox_crypt_IsReadOnly() != 0 {
		v4 = 0
		v5 = &v20[0]
		for {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v18)), 1)
			if int32(uint8(int8(v18))) >= 256 {
				return 0
			}
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v21[0], uint32(uint8(int8(v18))))
			v21[uint8(int8(v18))] = 0
			v6 = nox_xxx_modifGetIdByName_413290(&v21[0])
			*(**server.ModifierEff)(unsafe.Pointer(v5)) = nox_xxx_modifGetDescById_413330(v6)
			v4++
			v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 4))
			if v4 >= 4 {
				*(*uint16)(unsafe.Pointer(&v20[16])) = math.MaxUint16
				*(*uint16)(unsafe.Pointer(&v20[18])) = math.MaxUint16
				nox_xxx_modifSetItemAttrs_4E4990(a1, (*int32)(unsafe.Pointer(&v20[0])))
				goto LABEL_18
			}
		}
	}
	v2 = (***byte)(a1.InitData)
	v3 = 4
	for {
		if *v2 != nil {
			*(*uint8)(unsafe.Pointer(&v18)) = uint8(int8(libc.StrLen(**v2)))
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v18)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(**v2, uint32(uint8(int8(v18))))
		} else {
			*(*uint8)(unsafe.Pointer(&v18)) = 0
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v18)), 1)
		}
		v2 = (***byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((**byte)(nil))*1))
		v3--
		if v3 == 0 {
			break
		}
	}
LABEL_18:
	if int32(int16(v17)) >= 41 {
		v7 = 0
		if a1.ObjClass&0x1000 != 0 && a1.ObjSubClass&0x47F0000 != 0 {
			v7 = 1
		}
		if (int32(int16(v17)) >= 62 || (a1.ObjClass&0x1000) == 0 || (a1.ObjSubClass&0x4000000) == 0) && v7 != 0 {
			v8 := a1.UseData
			v14 = *(*uint8)(unsafe.Add(v8, 108))
			v15 = *(*uint8)(unsafe.Add(v8, 109))
			v9 = v15
			v16 = int32(*(*uint32)(unsafe.Add(v8, 112)))
			v10 = v16
			v14 = *(*uint8)(unsafe.Add(v8, 108))
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v14, 1)
			v15 = *(*uint8)(unsafe.Add(v8, 109))
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v15, 1)
			if int32(int16(v17)) >= 61 {
				v16 = int32(*(*uint32)(unsafe.Add(v8, 112)))
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v16)), 4)
			}
			if !noxflags.HasGame(4096) {
				*(*uint8)(unsafe.Add(v8, 108)) = v14
				*(*uint8)(unsafe.Add(v8, 109)) = v15
				*(*uint32)(unsafe.Add(v8, 112)) = uint32(v16)
				goto LABEL_37
			}
			if int32(v14) <= int32(v9) && v16 >= 0 && v16 <= v10 && int32(v9) == int32(v15) {
				*(*uint8)(unsafe.Add(v8, 108)) = v14
				*(*uint8)(unsafe.Add(v8, 109)) = v15
				*(*uint32)(unsafe.Add(v8, 112)) = uint32(v16)
				goto LABEL_37
			}
			*(*uint8)(unsafe.Add(v8, 108)) = 0
			*(*uint8)(unsafe.Add(v8, 109)) = v9
			*(*uint32)(unsafe.Add(v8, 112)) = 0
		}
	}
LABEL_37:
	if int32(int16(v17)) >= 42 {
		*(*uint16)(unsafe.Add(unsafe.Pointer(&v11), unsafe.Sizeof(uint16(0))*0)) = uint16(nox_xxx_unitGetHP_4EE780(a1))
		v16 = v11
		v12 = nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v16)), 2)
		*(*uint16)(unsafe.Add(unsafe.Pointer(&v12), unsafe.Sizeof(uint16(0))*0)) = a1.HealthData.Max
		if int32(uint16(int16(v16))) > int32(uint16(v12)) {
			v16 = int32(v12)
		}
		if nox_crypt_IsReadOnly() == 1 {
			if nox_xxx_gameIsSwitchToSolo_4DB240() == 1 || nox_xxx_gameIsNotMultiplayer_4DB250() == 1 || noxflags.HasGame(4096) && sub_419EA0() != 0 {
				nox_xxx_unitSetHP_4E4560(a1, uint16(int16(v16)))
			} else {
				v13 := nox_xxx_getProjectileClassById_413250(int32(a1.TypeInd))
				if v13 != nil {
					a1.HealthData.Max = *(*uint16)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint16(0))*26))
					a1.HealthData.Field2 = *(*uint16)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint16(0))*26))
					nox_xxx_unitSetHP_4E4560(a1, *(*uint16)(unsafe.Add(unsafe.Pointer(v13), unsafe.Sizeof(uint16(0))*26)))
				}
			}
		}
	}
	if int32(uint16(int16(v17))) == 63 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v14, 1)
	}
	if int32(int16(v17)) >= 64 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(a1.UpdateData, 4)), 4)
	}
	if a1.Field34 != 0 && nox_crypt_IsReadOnly() == 1 && nox_xxx_xfer_4F3E30(uint16(int16(v17)), a1, int32(a1.Field34)) == 0 {
		return 0
	}
	a1.Field34 = uint32(v19)
	return 1
}
func nox_xxx_XFerArmor_4F6860(a1p *server.Object, data unsafe.Pointer) int {
	a1 := a1p
	var (
		result int32
		v2     ***byte
		v3     int32
		v4     int32
		v5     *byte
		v6     int32
		v7     int32
		v8     uint32
		v10    int32
		v11    int32
		v12    uint32
		v13    int8
		v14    int32
		v15    [20]byte
		v16    [256]byte
	)
	v14 = int32(a1.Field34)
	v10 = 62
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v10)), 2)
	if int32(int16(v10)) > 62 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(a1, int32(int16(v10)))
	if result == 0 {
		return int(result)
	}
	if int32(int16(v10)) < 11 && nox_crypt_IsReadOnly() == 1 {
		*(*uint32)(unsafe.Pointer(&v15[0])) = 0
		*(*uint32)(unsafe.Pointer(&v15[4])) = 0
		*(*uint32)(unsafe.Pointer(&v15[8])) = 0
		*(*uint32)(unsafe.Pointer(&v15[12])) = 0
		nox_xxx_modifSetItemAttrs_4E4990(a1, (*int32)(unsafe.Pointer(&v15[0])))
		return 1
	}
	if nox_crypt_IsReadOnly() != 0 {
		v4 = 0
		v5 = &v15[0]
		for {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v11)), 1)
			if int32(uint8(int8(v11))) >= 256 {
				return 0
			}
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v16[0], uint32(uint8(int8(v11))))
			v16[uint8(int8(v11))] = 0
			v6 = nox_xxx_modifGetIdByName_413290(&v16[0])
			*(**server.ModifierEff)(unsafe.Pointer(v5)) = nox_xxx_modifGetDescById_413330(v6)
			v4++
			v5 = (*byte)(unsafe.Add(unsafe.Pointer(v5), 4))
			if v4 >= 4 {
				*(*uint16)(unsafe.Pointer(&v15[16])) = math.MaxUint16
				*(*uint16)(unsafe.Pointer(&v15[18])) = math.MaxUint16
				nox_xxx_modifSetItemAttrs_4E4990(a1, (*int32)(unsafe.Pointer(&v15[0])))
				goto LABEL_18
			}
		}
	}
	v2 = (***byte)(a1.InitData)
	v3 = 4
	for {
		if *v2 != nil {
			*(*uint8)(unsafe.Pointer(&v11)) = uint8(int8(libc.StrLen(**v2)))
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v11)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(**v2, uint32(uint8(int8(v11))))
		} else {
			*(*uint8)(unsafe.Pointer(&v11)) = 0
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v11)), 1)
		}
		v2 = (***byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((**byte)(nil))*1))
		v3--
		if v3 == 0 {
			break
		}
	}
LABEL_18:
	if int32(int16(v10)) >= 41 {
		*(*uint16)(unsafe.Add(unsafe.Pointer(&v7), unsafe.Sizeof(uint16(0))*0)) = uint16(nox_xxx_unitGetHP_4EE780(a1))
		v12 = uint32(v7)
		v8 = nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v12)), 2)
		*(*uint16)(unsafe.Add(unsafe.Pointer(&v8), unsafe.Sizeof(uint16(0))*0)) = a1.HealthData.Max
		if int32(uint16(v12)) > int32(uint16(v8)) {
			v12 = v8
		}
		if nox_crypt_IsReadOnly() == 1 {
			if nox_xxx_gameIsSwitchToSolo_4DB240() == 1 || nox_xxx_gameIsNotMultiplayer_4DB250() == 1 || noxflags.HasGame(4096) && sub_419EA0() != 0 {
				nox_xxx_unitSetHP_4E4560(a1, uint16(v12))
			} else {
				v9 := nox_xxx_equipClothFindDefByTT_413270(int32(a1.TypeInd))
				if v9 != nil {
					a1.HealthData.Max = *(*uint16)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint16(0))*26))
					a1.HealthData.Field2 = *(*uint16)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint16(0))*26))
					nox_xxx_unitSetHP_4E4560(a1, *(*uint16)(unsafe.Add(unsafe.Pointer(v9), unsafe.Sizeof(uint16(0))*26)))
				}
			}
		}
	}
	if int32(uint16(int16(v10))) == 61 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v13)), 1)
	}
	if int32(int16(v10)) >= 62 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(a1.UpdateData, 4)), 4)
	}
	if a1.Field34 != 0 && nox_crypt_IsReadOnly() == 1 && nox_xxx_xfer_4F3E30(uint16(int16(v10)), a1, int32(a1.Field34)) == 0 {
		return 0
	}
	a1.Field34 = uint32(v14)
	return 1
}
func nox_xxx_XFerAmmo_4F6B20(a1p *server.Object, data unsafe.Pointer) int {
	a1 := a1p
	var (
		v1     int32
		result int32
		v3     ***byte
		v4     int32
		v5     *byte
		v6     int32
		v7     *byte
		v8     int32
		v9     bool
		v10    *byte
	)
	var v11 int8
	var v12 int8
	var v13 int8
	var v14 int32
	var v15 *byte
	var v16 int32
	var v17 int32
	var v18 [20]byte
	var v19 [256]byte
	v1 = int32(a1.Field34)
	v15 = (*byte)(a1.UseData)
	v17 = v1
	v16 = 60
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v16)), 2)
	if int32(int16(v16)) > 60 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(a1, int32(int16(v16)))
	if result == 0 {
		return int(result)
	}
	if nox_crypt_IsReadOnly() != 0 {
		v6 = 0
		v7 = &v18[0]
		for {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v14)), 1)
			if int32(uint8(int8(v14))) >= 256 {
				return 0
			}
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v19[0], uint32(uint8(int8(v14))))
			v19[uint8(int8(v14))] = 0
			v8 = nox_xxx_modifGetIdByName_413290(&v19[0])
			*(**server.ModifierEff)(unsafe.Pointer(v7)) = nox_xxx_modifGetDescById_413330(v8)
			v6++
			v7 = (*byte)(unsafe.Add(unsafe.Pointer(v7), 4))
			if v6 >= 4 {
				*(*uint16)(unsafe.Pointer(&v18[16])) = math.MaxUint16
				*(*uint16)(unsafe.Pointer(&v18[18])) = math.MaxUint16
				nox_xxx_modifSetItemAttrs_4E4990(a1, (*int32)(unsafe.Pointer(&v18[0])))
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v13)), 1)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v12)), 1)
				v9 = !noxflags.HasGame(4096)
				v10 = v15
				if !v9 {
					*(*byte)(unsafe.Add(unsafe.Pointer(v15), 2)) = 0
				}
				v11 = v12
				*(*byte)(unsafe.Add(unsafe.Pointer(v10), 1)) = byte(v13)
				*v10 = byte(v11)
				goto LABEL_17
			}
		}
	}
	v3 = (***byte)(a1.InitData)
	v4 = 4
	for {
		if *v3 != nil {
			*(*uint8)(unsafe.Pointer(&v14)) = uint8(int8(libc.StrLen(**v3)))
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v14)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(**v3, uint32(uint8(int8(v14))))
		} else {
			*(*uint8)(unsafe.Pointer(&v14)) = 0
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v14)), 1)
		}
		v3 = (***byte)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof((**byte)(nil))*1))
		v4--
		if v4 == 0 {
			break
		}
	}
	v5 = v15
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v15), 1)), 1)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v5, 1)
LABEL_17:
	if a1.Field34 != 0 && nox_crypt_IsReadOnly() == 1 && nox_xxx_xfer_4F3E30(uint16(int16(v16)), a1, int32(a1.Field34)) == 0 {
		return 0
	}
	a1.Field34 = uint32(v17)
	result = 1
	return int(result)
}
func nox_xxx_XFerTeam_4F6D20(a1p *server.Object, data unsafe.Pointer) int {
	a1 := a1p
	var (
		result int32
		v2     ***byte
		v3     int32
		v4     *byte
		v5     int32
		v6     int32
		v7     *uint32
		v8     int32
		v9     int32
		v10    int32
		v11    [20]byte
		v12    [256]byte
	)
	v10 = int32(a1.Field34)
	v9 = 60
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v9)), 2)
	if int32(int16(v9)) > 60 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(a1, int32(int16(v9)))
	if result == 0 {
		return int(result)
	}
	if nox_crypt_IsReadOnly() != 0 {
		v4 = &v11[0]
		v5 = 4
		for {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v8)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v12[0], uint32(uint8(int8(v8))))
			v12[uint8(int8(v8))] = 0
			v6 = nox_xxx_modifGetIdByName_413290(&v12[0])
			*(**server.ModifierEff)(unsafe.Pointer(v4)) = nox_xxx_modifGetDescById_413330(v6)
			v4 = (*byte)(unsafe.Add(unsafe.Pointer(v4), 4))
			v5--
			if v5 == 0 {
				break
			}
		}
		*(*uint16)(unsafe.Pointer(&v11[16])) = math.MaxUint16
		*(*uint16)(unsafe.Pointer(&v11[18])) = math.MaxUint16
		nox_xxx_modifSetItemAttrs_4E4990(a1, (*int32)(unsafe.Pointer(&v11[0])))
		if uint32(a1.ObjClass)&0x10000000 != 0 {
			v7 = (*uint32)(a1.UpdateData)
			*v7 = uint32(a1.PosVec.X)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*1)) = uint32(a1.PosVec.Y)
		}
	} else {
		v2 = (***byte)(a1.InitData)
		v3 = 4
		for {
			if *v2 != nil {
				*(*uint8)(unsafe.Pointer(&v8)) = uint8(int8(libc.StrLen(**v2)))
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v8)), 1)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(**v2, uint32(uint8(int8(v8))))
			} else {
				*(*uint8)(unsafe.Pointer(&v8)) = 0
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v8)), 1)
			}
			v2 = (***byte)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof((**byte)(nil))*1))
			v3--
			if v3 == 0 {
				break
			}
		}
	}
	if a1.Field34 == 0 || nox_crypt_IsReadOnly() != 1 || (func() int32 {
		result = nox_xxx_xfer_4F3E30(uint16(int16(v9)), a1, int32(a1.Field34))
		return result
	}()) != 0 {
		a1.Field34 = uint32(v10)
		result = 1
	}
	return int(result)
}
func nox_xxx_XFerGold_4F6EC0(a1p *server.Object, data unsafe.Pointer) int {
	var (
		v2     *uint8
		v3     int32
		result int32
	)
	v1 := a1p
	v2 = (*uint8)(a1p.InitData)
	v3 = int32(a1p.Field34)
	a1 := int32(60)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&a1)), 2)
	if int32(int16(a1)) > 60 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(v1, int32(int16(a1)))
	if result == 0 {
		return int(result)
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v2, 4)
	if v1.Field34 == 0 || nox_crypt_IsReadOnly() != 1 || (func() int32 {
		result = nox_xxx_xfer_4F3E30(uint16(int16(a1)), v1, int32(v1.Field34))
		return result
	}()) != 0 {
		v1.Field34 = uint32(v3)
		result = 1
	}
	return int(result)
}
func nox_xxx_XFerObelisk_4F6F60(a1p *server.Object, data unsafe.Pointer) int {
	var (
		v2     *uint8
		v3     int32
		result int32
		v7     float32
		v8     int32
		v9     int32
	)
	v1 := a1p
	v2 = (*uint8)(a1p.UpdateData)
	v3 = int32(a1p.Field34)
	v8 = 61
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v8)), 2)
	if int32(int16(v8)) > 61 {
		return 0
	}
	result = nox_xxx_mapReadWriteObjData_4F4530(v1, int32(int16(v8)))
	if result == 0 {
		return int(result)
	}
	if int32(int16(v8)) >= 61 {
		var a1 int32
		*(*uint8)(unsafe.Pointer(&a1)) = 0
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v2, 4)
		if nox_crypt_IsReadOnly() == 1 {
			v9 = int32(*(*uint32)(unsafe.Pointer(v2)) * 80 / 50)
			v7 = float32(float64(v9))
			nullsub_35(unsafe.Pointer(v1), *(*uint32)(unsafe.Add(unsafe.Pointer(&v7), 4*0)))
		}
		if noxflags.HasGame(2048) {
			v5 := nox_xxx_netSpriteByCodeStatic_45A720(v1.Extent)
			if v5 != nil {
				v6 := nox_xxx_cliFirstMinimapObj_459EB0()
				if v6 != nil {
					for v6 != v5 {
						v6 = nox_xxx_cliNextMinimapObj_459EC0(v6)
						if v6 == nil {
							goto LABEL_14
						}
					}
					*(*uint8)(unsafe.Pointer(&a1)) = 1
				}
			}
		}
	LABEL_14:
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&a1)), 1)
	}
	if v1.Field34 == 0 || nox_crypt_IsReadOnly() != 1 || (func() int32 {
		result = nox_xxx_xfer_4F3E30(uint16(int16(v8)), v1, int32(v1.Field34))
		return result
	}()) != 0 {
		v1.Field34 = uint32(v3)
		return 1
	}
	return int(result)
}
func nox_xxx_XFerToxicCloud_4F70A0(a1p *server.Object, data unsafe.Pointer) int {
	var (
		v3 int32
	)
	v1 := a1p
	v2 := a1p.UpdateData
	v3 = int32(a1p.Field34)
	a1 := int32(61)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&a1)), 2)
	if int32(int16(a1)) > 61 {
		return 0
	}
	if int32(int16(a1)) <= 0 {
		return 0
	}
	if nox_xxx_mapReadWriteObjData_4F4530(v1, int32(int16(a1))) == 0 {
		return 0
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*byte)(v2), 4)
	if v1.Field34 != 0 {
		if nox_crypt_IsReadOnly() == 1 && nox_xxx_xfer_4F3E30(uint16(int16(a1)), v1, int32(v1.Field34)) == 0 {
			return 0
		}
	}
	v1.Field34 = uint32(v3)
	return 1
}
func nox_xxx_XFerMonsterGen_4F7130(a1p *server.Object, data unsafe.Pointer) int {
	a1 := a1p
	var (
		v2  int32
		v4  int32
		v5  *uint8
		v6  *byte
		v7  *byte
		v8  *byte
		v9  *byte
		v10 *int32
		v11 int8
		v12 *int32
		v13 int32
		v15 *uint8
		v16 int32
		v17 int32
		v19 int32
		v20 bool
		v21 int32
		v22 *uint8
		v24 uint32
		v25 int8
		v26 int32
		v27 int32
		v28 int32
		v29 int32
		v30 int32
		v31 int32
		v32 *uint8
		v33 int32
		v34 [256]byte
	)
	v1 := a1.UpdateData
	v2 = int32(a1.Field34)
	v3 := a1.Field189
	v32 = (*uint8)(a1.UpdateData)
	v33 = v2
	v29 = 63
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v29)), 2)
	if !(int32(int16(v29)) <= 63 && int32(int16(v29)) > 0 && nox_xxx_mapReadWriteObjData_4F4530(a1, int32(int16(v29))) != 0) {
		return 0
	}
	*(*uint8)(unsafe.Pointer(&v31)) = 3
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v31)), 1)
	v4 = 0
	if int32(uint8(int8(v31))) != 0 {
		v5 = (*uint8)(unsafe.Add(v1, 80))
		for {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v5, 1)
			v4++
			v5 = (*uint8)(unsafe.Add(unsafe.Pointer(v5), 1))
			if v4 >= int32(uint8(int8(v31))) {
				break
			}
		}
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(v1, 86)), 1)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(v1, 87)), 1)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(v1, 88)), 4)
	if v3 != nil {
		v6 = (*byte)(unsafe.Add(v3, 1920))
	} else {
		v6 = nil
	}
	nox_xxx_xferReadScriptHandler_4F5580(unsafe.Add(v1, 48), v6)
	if v3 != nil {
		v7 = (*byte)(unsafe.Add(v3, 2048))
	} else {
		v7 = nil
	}
	nox_xxx_xferReadScriptHandler_4F5580(unsafe.Add(v1, 56), v7)
	if v3 != nil {
		v8 = (*byte)(unsafe.Add(v3, 2176))
	} else {
		v8 = nil
	}
	nox_xxx_xferReadScriptHandler_4F5580(unsafe.Add(v1, 72), v8)
	if v3 != nil {
		v9 = (*byte)(unsafe.Add(v3, 2304))
	} else {
		v9 = nil
	}
	nox_xxx_xferReadScriptHandler_4F5580(unsafe.Add(v1, 64), v9)
	if nox_crypt_IsReadOnly() != 0 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v27)), 1)
		v16 = 0
		if int32(uint8(int8(v27))) != 0 {
			for {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v30)), 1)
				v17 = 0
				if int32(uint8(int8(v30))) != 0 {
					for {
						nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v28)), 1)
						nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v34[0], uint32(uint8(int8(v28))))
						v34[uint8(int8(v28))] = 0
						v18 := nox_xxx_newObjectByTypeID_4E3810(&v34[0])
						if v18 == nil {
							return 0
						}
						nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v26)), 2)
						nox_xxx_fileCryptReadCrcMB_426C20((*uint8)(unsafe.Pointer(&v32)), 4)
						if v18.Xfer.Get()(v18, nil) == 0 {
							return 0
						}
						v19 = func() int32 {
							p := &v17
							x := *p
							*p++
							return x
						}() + v16*4
						v20 = v17 < int32(uint8(int8(v30)))
						*(*uint32)(unsafe.Add(v1, v19*4)) = uint32(uintptr(unsafe.Pointer(v18)))
						if !v20 {
							break
						}
					}
				}
				if func() int32 {
					p := &v16
					*p++
					return *p
				}() >= int32(uint8(int8(v27))) {
					goto LABEL_37
				}
			}
		}
	} else {
		*(*uint8)(unsafe.Pointer(&v26)) = 3
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v26)), 1)
		v10 = (*int32)(unsafe.Pointer(v1))
		v28 = 3
		for {
			v11 = 0
			v12 = v10
			v25 = 0
			v13 = 4
			for {
				if *v12 != 0 {
					v11++
				}
				v12 = (*int32)(unsafe.Add(unsafe.Pointer(v12), 4*1))
				v13--
				if v13 == 0 {
					break
				}
			}
			v25 = v11
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v25)), 1)
			v30 = 4
			for {
				v14 := *(**server.Object)(unsafe.Pointer(v10))
				if v14 != nil {
					*(*uint8)(unsafe.Pointer(&v27)) = uint8(int8(libc.StrLen(nox_xxx_getUnitName_4E39D0(v14))))
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v27)), 1)
					v24 = uint32(uint8(int8(v27)))
					v15 = nox_xxx_getUnitName_4E39D0(v14)
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v15, v24)
					nox_xxx_xfer_saveObj_51DF90(v14)
				}
				v10 = (*int32)(unsafe.Add(unsafe.Pointer(v10), 4*1))
				v30--
				if v30 == 0 {
					break
				}
			}
			v10 = v12
			v28--
			if v28 == 0 {
				break
			}
		}
		v1 = unsafe.Pointer(v32)
	}
LABEL_37:
	if int32(int16(v29)) >= 62 {
		*(*uint8)(unsafe.Pointer(&v26)) = 3
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v26)), 1)
		v21 = 0
		if int32(uint8(int8(v26))) != 0 {
			v22 = (*uint8)(unsafe.Add(v1, 83))
			for {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v22, 1)
				v21++
				v22 = (*uint8)(unsafe.Add(unsafe.Pointer(v22), 1))
				if v21 >= int32(uint8(int8(v26))) {
					break
				}
			}
		}
	}
	if int32(int16(v29)) >= 63 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(v1, 92)), 4)
	}
	if a1.Field34 == 0 || nox_crypt_IsReadOnly() != 1 || nox_xxx_xfer_4F3E30(uint16(int16(v29)), a1, int32(a1.Field34)) != 0 {
		a1.Field34 = uint32(v33)
		return 1
	}
	return 0
}
func nox_xxx_XFerRewardMarker_4F74D0(a1p *server.Object, data unsafe.Pointer) int {
	a1 := a1p
	var (
		v1  *uint8
		v2  int32
		i   int32
		v4  *byte
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 *byte
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 *byte
		v16 int32
		v18 uint32
		v19 uint32
		v20 uint32
		v21 int32
		v22 int32
		v23 int32
		v24 int32
		v25 [256]byte
	)
	v1 = (*uint8)(a1.InitData)
	v24 = int32(a1.Field34)
	v23 = 63
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v23)), 2)
	if int32(int16(v23)) > 63 {
		return 0
	}
	v2 = 0
	if int32(int16(v23)) <= 0 {
		return 0
	}
	if nox_xxx_mapReadWriteObjData_4F4530(a1, int32(int16(v23))) == 0 {
		return 0
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v1, 4)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v1), 4)), 4)
	v22 = 0
	for i = 0; i < 137; i++ {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), i+8))) == 1 {
			v22++
		}
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v22)), 2)
	if nox_crypt_IsReadOnly() != 0 {
		v5 = 0
		if int32(uint16(int16(v22))) != 0 {
			for {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 1)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v25[0], uint32(uint8(int8(v21))))
				v25[uint8(int8(v21))] = 0
				v6 = nox_xxx_spellNameToN_4243F0(&v25[0])
				if v6 == 0 {
					return 0
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), v6+8)) = 1
				if func() int32 {
					p := &v5
					*p++
					return *p
				}() >= int32(uint16(int16(v22))) {
					break
				}
			}
		}
	} else {
		for {
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), v2+8))) != 0 {
				*(*uint8)(unsafe.Pointer(&v21)) = uint8(int8(libc.StrLen(nox_xxx_spellNameByN_424870(v2))))
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 1)
				v18 = uint32(uint8(int8(v21)))
				v4 = nox_xxx_spellNameByN_424870(v2)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v4, v18)
			}
			v2++
			if v2 >= 137 {
				break
			}
		}
	}
	v7 = 0
	v8 = 0
	v22 = 0
	for {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), v8+145))) == 1 {
			v22++
		}
		v8++
		if v8 >= 6 {
			break
		}
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v22)), 2)
	if nox_crypt_IsReadOnly() != 0 {
		if int32(uint16(int16(v22))) != 0 {
			for {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 1)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v25[0], uint32(uint8(int8(v21))))
				v25[uint8(int8(v21))] = 0
				v11 = nox_xxx_abilityNameToN_424D80(&v25[0])
				if v11 == 0 {
					return 0
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), v11+145)) = 1
				if func() int32 {
					p := &v7
					*p++
					return *p
				}() >= int32(uint16(int16(v22))) {
					break
				}
			}
		}
	} else {
		v9 = 0
		for {
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), v9+145))) != 0 {
				*(*uint8)(unsafe.Pointer(&v21)) = uint8(int8(libc.StrLen(nox_xxx_abilityGetName_425250(v9))))
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 1)
				v19 = uint32(uint8(int8(v21)))
				v10 = nox_xxx_abilityGetName_425250(v9)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v10, v19)
			}
			v9++
			if v9 >= 6 {
				break
			}
		}
	}
	v12 = 0
	v13 = 0
	v22 = 0
	for {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), v13+151))) == 1 {
			v22++
		}
		v13++
		if v13 >= 41 {
			break
		}
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v22)), 2)
	if nox_crypt_IsReadOnly() != 0 {
		if int32(uint16(int16(v22))) != 0 {
			for {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 1)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v25[0], uint32(uint8(int8(v21))))
				v25[uint8(int8(v21))] = 0
				v16 = nox_xxx_guide_427010(&v25[0])
				if v16 == 0 {
					return 0
				}
				*(*uint8)(unsafe.Add(unsafe.Pointer(v1), v16+151)) = 1
				if func() int32 {
					p := &v12
					*p++
					return *p
				}() >= int32(uint16(int16(v22))) {
					break
				}
			}
		}
	} else {
		v14 = 0
		for {
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), v14+151))) != 0 {
				*(*uint8)(unsafe.Pointer(&v21)) = uint8(int8(libc.StrLen(nox_xxx_guideNameByN_427230(v14))))
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 1)
				v20 = uint32(uint8(int8(v21)))
				v15 = nox_xxx_guideNameByN_427230(v14)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v15, v20)
			}
			v14++
			if v14 >= 41 {
				break
			}
		}
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v1), 196)), 4)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v1), 192)), 4)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v1), 200)), 4)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v1), 204)), 4)
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v1), 208)), 4)
	if int32(int16(v23)) >= 62 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v1), 212)), 4)
	}
	if int32(int16(v23)) >= 63 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v1), 216)), 1)
	}
	if a1.Field34 == 0 || nox_crypt_IsReadOnly() != 1 || nox_xxx_xfer_4F3E30(uint16(int16(v23)), a1, int32(a1.Field34)) != 0 {
		a1.Field34 = uint32(v24)
		return 1
	}
	return 0
}
func nox_xxx_equipedItemByCode_4F7920(a1 *server.Object, a2 int32) *server.Object {
	result := a1.InvFirstItem
	if result == nil {
		return nil
	}
	for result.NetCode != uint32(a2) {
		result = result.InvNextItem
		if result == nil {
			return nil
		}
	}
	return result
}
func sub_4F7950(a1 *server.Object) {
	var (
		v2 int32
	)
	v1 := a1.UpdateDataPlayer()
	v2 = 3
	v3 := &v1.Field42
	for {
		if *v3 != 0 {
			nox_xxx_delayedDeleteObject_4E5CC0((*server.Object)(unsafe.Pointer(uintptr(*v3))))
		}
		*v3 = 0
		v3 = (*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*1))
		v2--
		if v2 == 0 {
			break
		}
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 181)) = 0
	*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 180)) = 0
}
func nox_xxx_playerSetCustomWP_4F79A0(a1 *server.Object, a2 int32, a3 int32) {
	var (
		v5 types.Pointf
	)
	v3 := a1.UpdateData
	if (int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 3680))) & 3) == 0 {
		v4 := *(**server.Object)(unsafe.Add(v3, int32(*(*uint8)(unsafe.Add(v3, 180)))*4+168))
		if v4 != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(&v5.X), 4*0)) = uint32(a2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(&v5.Y), 4*0)) = uint32(a3)
			nox_xxx_unitMove_4E7010(v4, &v5)
		} else {
			*(**server.Object)(unsafe.Add(v3, int32(*(*uint8)(unsafe.Add(v3, 180)))*4+168)) = nox_xxx_newObjectByTypeID_4E3810(internCStr("PlayerWaypoint"))
			nox_xxx_createAt_4DAA50((*server.Object)(*(*unsafe.Pointer)(unsafe.Add(v3, int32(*(*uint8)(unsafe.Add(v3, 180)))*4+168))), a1, float32(a2), float32(a3))
		}
	}
}
func nox_xxx_playerConfusedGetDirection_4F7A40(a1p *server.Object) int32 {
	var (
		a1 = a1p
		v1 int32
		v2 int32
		v3 int32
		v4 int32
	)
	v1 = int32(a1.Direction2)
	v2 = int32(uint8(nox_xxx_buffGetPower_4FF570(a1, 3)))
	v3 = int32((gameFrame() + a1.NetCode) % 0x28)
	if v3 > 20 {
		v3 = 40 - v3
	}
	v4 = (v2+3)*(v3-10) + v1
	if v4 < 0 {
		v4 += int32(uint32(math.MaxUint8-v4) >> 8 << 8)
	}
	if v4 >= 256 {
		v4 += int32((uint32(v4) >> 8) * 4294967040)
	}
	return v4
}
func nox_xxx_mapFindPlayerStart_4F7AB0(a1 *types.Pointf, a2p *server.Object) {
	var (
		a2  = a2p
		v3  int32
		v4  int32
		v8  int32
		v11 float64
		v12 float64
		v13 float64
		v14 int32
		v15 float32
		v16 int32
		v17 float32
		v18 int32
	)
	v3 = 0
	v16 = 0
	if dword_5d4594_1568868 == 0 {
		result := nox_xxx_getNameId_4E3AA0(internCStr("PlayerStart"))
		dword_5d4594_1568868 = uint32(result)
	}
	if a2 != nil {
		if nox_xxx_servObjectHasTeam_419130((*server.ObjectTeam)(unsafe.Add(unsafe.Pointer(a2), 48))) != 0 {
			v16 = int32(a2.TeamVal.ID)
			nox_xxx_getTeamByID_418AB0(v16)
			v3 = v16
		}
		v4 = 0
		var v5 *server.Object
		v6 := nox_server_getFirstObject_4DA790()
		if v6 == nil {
			a1.X = 2000.0
			a1.Y = 2000.0
			return
		}
		for {
			if uint32(v6.TypeInd) == dword_5d4594_1568868 {
				v5 = v6
				if sub_4F7CE0(unsafe.Pointer(v6), v3) != 0 {
					v4++
				}
			}
			v6 = nox_server_getNextObject_4DA7A0(v6)
			if v6 == nil {
				break
			}
		}
		v18 = v4
		if v4 != 0 {
			v17 = 0.0
			var j *server.Object
			v8 = 1
			v9 := nox_server_getFirstObject_4DA790()
			for v9 != nil {
				if uint32(v9.TypeInd) == dword_5d4594_1568868 && sub_4F7CE0(unsafe.Pointer(v9), v16) != 0 {
					v15 = 1e+07
					for i := nox_xxx_getFirstPlayerUnit_4DA7C0(); i != nil; i = nox_xxx_getNextPlayerUnit_4DA7F0(i) {
						if i != a2 && nox_xxx_unitIsEnemyTo_5330C0(a2, i) != 0 {
							v11 = float64(v9.PosVec.X - i.PosVec.X)
							v12 = float64(v9.PosVec.Y - i.PosVec.Y)
							v13 = v12*v12 + v11*v11
							if v13 < float64(v15) {
								v15 = float32(v13)
							}
							v8 = 0
						}
					}
					if float64(v15) > float64(v17) {
						j = v9
						v17 = v15
					}
				}
				v9 = nox_server_getNextObject_4DA7A0(v9)
			}
			if v8 != 0 || j == nil {
				v14 = nox_common_randomInt_415FA0(0, v18-1)
				for j = nox_server_getFirstObject_4DA790(); j != nil; j = nox_server_getNextObject_4DA7A0(j) {
					if uint32(j.TypeInd) == dword_5d4594_1568868 && sub_4F7CE0(unsafe.Pointer(j), v16) != 0 {
						if v14 == 0 {
							break
						}
						v14--
					}
				}
			}
			*a1 = *(*types.Pointf)(unsafe.Add(unsafe.Pointer(j), 56))
		} else {
			if v5 == nil {
				a1.X = 2000.0
				a1.Y = 2000.0
				return
			}
			*a1 = *(*types.Pointf)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(types.Pointf{})*7))
		}
	}
}
func sub_4F7CE0(a1 unsafe.Pointer, a2 int32) int32 {
	return bool2int32(*(*uint32)(unsafe.Add(a1, 16))&0x1000000 != 0 && (a2 == 0 || nox_xxx_servObjectHasTeam_419130((*server.ObjectTeam)(unsafe.Add(a1, 48))) == 0 || nox_xxx_teamCompare2_419180((*server.ObjectTeam)(unsafe.Add(a1, 48)), uint8(int8(a2))) != 0))
}
func nox_xxx_playerSubStamina_4F7D30(a1p *server.Object, a2 int32) int32 {
	var (
		a1 = a1p
		v2 int32
	)
	v2 = int32(a1.ObjClass)
	if v2&4 != 0 {
		v3 := a1.UpdateData
		if int32(*(*uint8)(unsafe.Add(v3, 91))) >= a2 {
			*(*uint8)(unsafe.Add(v3, 91)) -= uint8(int8(a2))
			nox_xxx_netReportStamina_4D8800(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 2064))), a1)
			return 1
		}
	} else {
		if (v2 & 2) == 0 {
			return 1
		}
		v5 := a1.UpdateData
		if int32(*(*uint8)(unsafe.Add(v5, 1128))) >= a2 {
			*(*uint8)(unsafe.Add(v5, 1128)) -= uint8(int8(a2))
			return 1
		}
	}
	return 0
}
func sub_4F7DB0(a1 *server.Object, a2 int8) {
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 8)))&4 != 0 {
		v2 := a1.UpdateData
		*(*uint8)(unsafe.Add(v2, 91)) -= uint8(a2)
		nox_xxx_netReportStamina_4D8800(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v2, 276)), 2064))), a1)
	}
}
func nox_xxx_checkWinkFlags_4F7DF0(a1p *server.Object) int32 {
	var (
		a1 = a1p
		v1 int32
		v4 int32
	)
	v1 = int32(*memmap.PtrUint32(0x5D4594, 1568872))
	if *memmap.PtrUint32(0x5D4594, 1568872) == 0 {
		v1 = nox_xxx_getNameId_4E3AA0(internCStr("GameBall"))
		*memmap.PtrUint32(0x5D4594, 1568872) = uint32(v1)
	}
	v2 := a1.Field129
	if v2 == nil {
		return 0
	}
	for int32(v2.TypeInd) != v1 {
		v2 = v2.Field128
		if v2 == nil {
			return 0
		}
	}
	v4 = int32(v2.ObjFlags)
	*(*uint8)(unsafe.Pointer(&v4)) = uint8(int8(v4 & 0xBF))
	v2.ObjFlags = uint32(v4)
	nox_xxx_objectApplyForce_52DF80(&a1.PosVec, v2, 100.0)
	v2.Obj130 = nil
	nox_xxx_unitClearOwner_4EC300(v2)
	nox_xxx_aud_501960(926, a1, 0, 0)
	sub_4E8290(1, 0)
	return 1
}
func nox_xxx_weaponGetStaminaByType_4F7E80(a1 int32) int32 {
	if a1&0x200 != 0 {
		return 70
	}
	if a1&0x4000 != 0 {
		return 100
	}
	if a1&0x800 != 0 {
		return 50
	}
	if a1&0x100 != 0 {
		return 45
	}
	if a1&0x1000 != 0 {
		return 75
	}
	if a1&0x2000 != 0 {
		return 100
	}
	if uint32(a1)&0x7FF8000 != 0 {
		return 45
	}
	if (a1 & 0x400) != 0 {
		return 75
	}
	return 10
}
func Nox_xxx_playerRespawn_4F7EF0(a1p *server.Object) {
	var (
		a1 = a1p
		v2 *byte
		v9 types.Pointf
	)
	v1 := (*byte)(sub_416640())
	v2 = v1
	if a1 == nil {
		return
	}
	v3 := a1.UpdateDataPlayer()
	v4 := v3.Player
	if !(!noxflags.HasGame(4096) || (func() *byte {
		v1 = (*byte)(unsafe.Pointer(uintptr(v3.Field137)))
		return v1
	}()) == nil) {
		return
	}
	if v4 != nil {
		v4.Field4700 = 0
	}
	if noxflags.HasGame(4096) {
		nox_xxx_playerMakeDefItems_4EF7D0(a1, 1, 1)
		*(*uint8)(unsafe.Add(unsafe.Add(unsafe.Pointer(v3), v3.Player.PlayerInd), 452)) = 250
		nox_xxx_netPriMsgToPlayer_4DA2C0(a1, internCStr("GeneralPrint:Respawn"), 0)
	} else {
		nox_xxx_playerMakeDefItems_4EF7D0(a1, 1, 0)
	}
	if noxflags.HasGame(4096) {
		nox_xxx_aud_501960(1006, a1, 0, 0)
	} else {
		nox_xxx_aud_501960(148, a1, 0, 0)
	}
	if dword_5d4594_2650652 == 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(v2), 58)) != 0 {
		nox_xxx_respawnPlayerImpl_53FBC0(&a1.PosVec, int32(a1.Direction1))
	}
	v9 = a1.PosVec
	if noxflags.HasGame(4096) && v3.Field77 != nil {
		sub_4F80C0(v3.Field77, &v9)
	} else {
		nox_xxx_mapFindPlayerStart_4F7AB0(&v9, a1)
	}
	nox_xxx_unitMove_4E7010(a1, &v9)
	if noxflags.HasGame(16) {
		if nox_xxx_CheckGameplayFlags_417DA0(4) {
			v7 := *(**server.Object)(unsafe.Add(unsafe.Pointer(nox_xxx_getTeamByID_418AB0(int32(v4.PlayerUnit.TeamVal.ID))), 4*19))
			if v7 != nil {
				if v7.InvHolder == nil {
					sub_4F3400(v4.PlayerUnit, v7, 1, 0)
				}
			}
		}
	}
	if noxflags.HasGame(0x2000) {
		nox_xxx_buffApplyTo_4FF380(a1, 23, int16(int32(uint16(gameFPS()))*5), 5)
	}
}
func sub_4F80C0(a1 unsafe.Pointer, a3 *types.Pointf) {
	var (
		v2 int32
	)
	*a3 = *(*types.Pointf)(unsafe.Add(a1, 56))
	v2 = 32
	for {
		sub_4ED970(60.0, (*types.Pointf)(unsafe.Add(a1, 56)), a3)
		if nox_xxx_mapTileAllowTeleport_411A90(a3) == 0 {
			break
		}
		v2--
		if v2 == 0 {
			break
		}
	}
}
func sub_4F9A80(a1 *server.Object) int32 {
	return bool2int32(*(*uint32)(unsafe.Add(a1.UpdateData, uint32(int32(*(*uint8)(unsafe.Add(a1.UpdateData, 181)))*4)+168)) != 0)
}
func sub_4F9AB0(a1p *server.Object) int32 {
	var (
		a1     = a1p
		result int32
		v4     float64
		v5     int32
		v6     float64
		v7     types.Pointf
	)
	v1 := a1.UpdateData
	v2 := *(**server.Object)(unsafe.Add(v1, int32(*(*uint8)(unsafe.Add(v1, 181)))*4+168))
	if v2 == nil {
		return 0
	}
	v7.X = v2.PosVec.X - a1.PosVec.X
	v4 = float64(v2.PosVec.Y - a1.PosVec.Y)
	v7.Y = float32(v4)
	if v4*float64(v7.Y)+float64(v7.X*v7.X) >= 100.0 {
		a1.Direction2 = server.Dir16(uint16(int16(nox_xxx_math_509ED0(&v7))))
		if nox_xxx_testUnitBuffs_4FF350(a1, 3) != 0 {
			a1.Direction2 = server.Dir16(uint16(int16(nox_xxx_playerConfusedGetDirection_4F7A40(a1))))
		}
		v5 = int32(a1.Direction2) * 8
		a1.ForceVec.X = *mem_getFloatPtr(0x587000, uint32(v5)+194136)*a1.SpeedCur + a1.ForceVec.X
		v6 = float64(*mem_getFloatPtr(0x587000, uint32(v5)+194140) * a1.SpeedCur)
		result = 1
		a1.ForceVec.Y = float32(v6 + float64(a1.ForceVec.Y))
	} else {
		nox_xxx_delayedDeleteObject_4E5CC0(v2)
		result = 0
		*(*uint32)(unsafe.Add(v1, int32(*(*uint8)(unsafe.Add(v1, 181)))*4+168)) = 0
	}
	return result
}
func nox_xxx_playerCanMove_4F9BC0(a1 *server.Object) int32 {
	v1 := a1.UpdateData
	if nox_xxx_testUnitBuffs_4FF350(a1, 25) != 0 {
		return 0
	}
	if nox_xxx_testUnitBuffs_4FF350(a1, 5) != 0 {
		return 0
	}
	if noxflags.HasGame(4096) && *(*uint32)(unsafe.Add(v1, 280)) != 0 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Add(v1, 88))) != 1 {
		return 1
	}
	v3 := *(*unsafe.Pointer)(unsafe.Add(v1, 104))
	if v3 == nil {
		return 1
	}
	if *(*uint32)(unsafe.Add(v3, 8))&0x1000000 != 0 && int32(*(*uint8)(unsafe.Add(v3, 12)))&8 != 0 {
		return 0
	}
	return 1
}
func nox_xxx_playerCanAttack_4F9C40(a1 *server.Object) int32 {
	v1 := a1.UpdateData
	if nox_xxx_testUnitBuffs_4FF350(a1, 25) != 0 {
		return 0
	} else {
		return bool2int32(int32(*(*uint8)(unsafe.Add(v1, 88))) != 23)
	}
}
func nox_xxx_playerInputAttack_4F9C70(a1p *server.Object) {
	var (
		a1 = a1p
		v2 int32
		v4 int32
		v5 int32
		v6 int8
	)
	if a1 != nil && nox_xxx_playerAimsAtEnemy_4F9DC0(a1) != 0 {
		v1 := a1.UpdateData
		v2 = int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v1, 276)), 4)))
		if v2 != 0 {
			if uint32(v2)&0x47F0000 != 0 && nox_common_mapPlrActionToStateId_4FA2B0(a1) != 29 {
				v3 := *(*unsafe.Pointer)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v1, 104)), 736))
				if int32(*(*uint8)(unsafe.Add(v3, 108))) != 0 || int32(*(*uint8)(unsafe.Add(v3, 109))) == 0 {
					a1.Field34 = gameFrame()
					*(*uint8)(unsafe.Add(v1, 236)) = 0
					nox_xxx_playerSetState_4FA020(a1, 1)
					nox_xxx_useByNetCode_53F8E0(unsafe.Pointer(a1), *(*unsafe.Pointer)(unsafe.Add(v1, 104)))
				} else if nox_xxx_playerSubStamina_4F7D30(a1, 45) != 0 {
					v4 = int32(*(*uint32)(unsafe.Add(v3, 96)))
					*(*uint8)(unsafe.Pointer(&v4)) = uint8(int8(v4 | 2))
					*(*uint32)(unsafe.Add(v3, 96)) = uint32(v4)
					a1.Field34 = gameFrame()
					*(*uint8)(unsafe.Add(v1, 236)) = 0
					nox_xxx_playerSetState_4FA020(a1, 1)
				}
			} else if int32(*(*uint8)(unsafe.Add(v1, 88))) != 1 {
				v5 = nox_xxx_weaponGetStaminaByType_4F7E80(int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v1, 276)), 4))))
				v6 = int8(v5)
				if nox_xxx_playerSubStamina_4F7D30(a1, v5) != 0 {
					a1.Field34 = gameFrame()
					*(*uint8)(unsafe.Add(v1, 236)) = 0
					if nox_xxx_playerSetState_4FA020(a1, 1) == 0 {
						sub_4F7DB0(a1, int8(int32(-v6)))
					}
				}
			}
			nox_xxx_spellBuffOff_4FF5B0(a1, 0)
			nox_xxx_spellBuffOff_4FF5B0(a1, 23)
			nox_xxx_spellCancelDurSpell_4FEB10(67, a1)
		} else if int32(*(*uint8)(unsafe.Add(v1, 88))) != 1 {
			nox_xxx_playerSetState_4FA020(a1, 1)
		}
	}
}
func nox_xxx_playerAimsAtEnemy_4F9DC0(a1 *server.Object) int32 {
	var result int32
	if a1 == nil {
		return 0
	}
	if *(*uint32)(unsafe.Add(a1.UpdateData, 288)) == 0 || nox_xxx_unitIsEnemyTo_5330C0(a1, (*server.Object)(*(*unsafe.Pointer)(unsafe.Add(a1.UpdateData, 288)))) != 0 || (func() int32 {
		result = bool2int32(noxflags.HasGame(4096))
		return result
	}()) != 0 {
		result = 1
	}
	return result
}
func sub_4F9E10(a1p *server.Object) int32 {
	var (
		a1 = a1p
		v2 int32
	)
	if a1 == nil {
		return 0
	}
	if a1.Obj130 == nil {
		return 0
	}
	v1 := nox_xxx_findParentChainPlayer_4EC580(a1.Obj130)
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v1), 16)))&0x20 != 0 {
		return 0
	}
	v2 = int32(v1.ObjClass)
	if v2&2 != 0 || v2&4 != 0 && int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v1.UpdateData, 276)), 3680)))&1 != 0 {
		return 0
	}
	nox_xxx_playerCameraFollow_4E6060(a1, v1)
	return 1
}
func nox_xxx_unitGetStrength_4F9FD0(obj *server.Object) int32 {
	if obj == nil {
		return 0
	}
	if obj.Class().Has(object.ClassPlayer) {
		return int32(*(*uint32)(unsafe.Add(unsafe.Pointer(obj.UpdateDataPlayer().Player), 2239)))
	}
	if !obj.Class().Has(object.ClassMonster) {
		return 0
	}
	if obj.SubClass().AsMonster().Has(object.MonsterNPC) {
		return int32(*(*uint8)(unsafe.Add(unsafe.Pointer(obj.UpdateDataMonster()), 1324)))
	}
	return 30
}
func nox_xxx_playerSetState_4FA020(a1p *server.Object, a2 int32) int32 {
	var (
		a1 = a1p
		v3 int32
		v4 int32
		v7 int32
		v9 int32
	)
	v2 := a2
	v3 = 1
	v4 = int32(a1.ObjFlags)
	v5 := a1.UpdateDataPlayer()
	if (v4&0x8000) != 0 && a2 != 3 && a2 != 4 {
		return 0
	}
	if !noxflags.HasGame(2048) {
		v7 = int32(a1.ObjFlags)
		if v7&0x4000 != 0 {
			if a2 == 30 {
				return 0
			}
		}
	}
	if a2 == 24 || a2 == 25 || a2 == 26 || a2 == 27 || a2 == 28 || a2 == 29 {
		if nox_common_playerIsAbilityActive_4FC250(a1, 1) != 0 {
			return 0
		}
		if int32(v5.State) == 12 {
			return 0
		}
	}
	if int32(v5.State) == 1 {
		if a2 == 1 {
			goto LABEL_26
		}
		if nox_xxx_probablyWarcryCheck_4FC3E0(a1, 2) != 0 && a2 != 4 && a2 != 3 {
			return 0
		}
	}
	if a2 != 1 {
		*(*uint8)(unsafe.Add(unsafe.Pointer(v5.Player), 8)) = 0
		switch a2 {
		case 3, 4:
			v5.Field40_0 = 0
			v5.Field41 = 0
		case 25:
			if int32(v5.State) != a2 {
				nox_xxx_aud_501960(301, a1, 0, 0)
			}
		case 26:
			if int32(v5.State) != a2 {
				nox_xxx_aud_501960(300, a1, 0, 0)
			}
		case 28:
			if int32(v5.State) != a2 {
				nox_xxx_aud_501960(302, a1, 0, 0)
			}
		default:
			goto LABEL_42
		}
		goto LABEL_42
	}
LABEL_26:
	if v5.Field0 <= gameFrame() {
		v8 := v5.Player
		v9 = int32(v8.Field4)
		v5.Field0 = 0
		if v9 != 0 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(v8), 8)) = 0
		} else {
			*(*uint8)(unsafe.Add(unsafe.Pointer(v5.Player), 8)) = uint8(int8(nox_common_randomInt_415FA0(23, 24)))
			if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5.Player), 2251))) == 0 && nox_common_randomInt_415FA0(0, 100) > 75 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(v5.Player), 8)) = 25
			}
			nox_xxx_spellBuffOff_4FF5B0(a1, 0)
			nox_xxx_spellBuffOff_4FF5B0(a1, 23)
			nox_xxx_spellCancelDurSpell_4FEB10(67, a1)
		}
	} else {
		v3 = 0
		v2 = int32(v5.State)
	}
LABEL_42:
	if int32(v5.State) != v2 {
		v5.Field22_1 = v5.State
		v5.State = uint8(int8(v2))
		a1.Field34 = gameFrame()
		v5.Field59_0 = 0
	}
	if v2 == 30 {
		v5.Field41 = gameFrame()
	}
	return v3
}
func sub_4FA280(a1 int32) int32 {
	var v1 int32
	v1 = 2
	for ((1 << v1) & a1) == 0 {
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= 27 {
			return 0
		}
	}
	return int32(*memmap.PtrUint32(0x587000, uintptr(v1*4)+215824))
}
func nox_common_mapPlrActionToStateId_4FA2B0(a1p *server.Object) int32 {
	var (
		a1 = a1p
		v3 int32
	)
	v1 := a1.UpdateData
	switch *(*uint8)(unsafe.Add(v1, 88)) {
	case 0:
		return 4
	case 1, 0xE, 0x16:
		if nox_common_playerIsAbilityActive_4FC250(a1, 2) != 0 && nox_xxx_probablyWarcryCheck_4FC3E0(a1, 2) != 0 {
			return 46
		} else if nox_common_playerIsAbilityActive_4FC250(a1, 1) != 0 {
			return 45
		} else {
			r1 := *(*unsafe.Pointer)(unsafe.Add(v1, 276))
			v3 = int32(*(*uint32)(unsafe.Add(r1, 4)))
			if uint32(v3)&0x47F0000 != 0 {
				result := int32(^uint8(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v1, 104)), 736)), 96)))) & 2
				return result | 0x1D
			} else if v3 != 0 && v3 != 1 || *(*uint8)(unsafe.Add(r1, 8)) == 0 {
				return sub_4FA280(v3)
			}
			return 0
		}
	case 2, 0xA:
		return 21
	case 3:
		return 1
	case 4:
		return 2
	case 5:
		return 6
	case 0xC:
		return 3
	case 0xD:
		if (*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v1, 276)), 4)) & 0x400) != 0 {
			return 0x26
		} else {
			return 0
		}
	case 0xF, 0x10, 0x11:
		return 40
	case 0x12:
		return 48
	case 0x13:
		return 49
	case 0x14:
		return 47
	case 0x15:
		return 30
	case 0x17:
		return 50
	case 0x18:
		return 19
	case 0x19:
		return 20
	case 0x1A:
		return 15
	case 0x1B, 0x1C, 0x1D:
		return 16
	case 0x1E:
		return 52
	case 0x20:
		return 54
	default:
		return 0
	}
}
func nox_xxx_checkInversionEffect_4FA4F0(a1 *server.Object, a2 *server.Object) int32 {
	var (
		v3  int32
		v4  int32
		v10 int32
	)
	v2 := a1.InvFirstItem
	if v2 == nil {
		return 0
	}
	for {
		v3 = int32(v2.ObjFlags)
		if v3&0x100 != 0 {
			if v2.ObjClass&0x13001000 != 0 {
				v4 = 2
				v5 := unsafe.Add(v2.InitData, 8)
				for {
					v6 := *(*unsafe.Pointer)(v5)
					v10 = 0
					if v6 != nil {
						v7 := ccall.AsFunc[func(int32, int32, int32, int32, int32, *int32) int32](*(*unsafe.Pointer)(unsafe.Add(v6, 88)))
						if v7 != nil {
							if ccall.FuncAddr(v7) == ccall.FuncAddr(nox_xxx_inversionEffect_4E03D0) {
								v8 := nox_xxx_findParentChainPlayer_4EC580(a2)
								ccall.AsFunc[func(int32, *uint32, unsafe.Pointer, unsafe.Pointer, int32, *int32)](*(*unsafe.Pointer)(unsafe.Add(v6, 88)))(v6, v2, unsafe.Pointer(a1), unsafe.Pointer(a2), v8, &v10)
								if v10 == 1 {
									return 1
								}
							}
						}
					}
					v4++
					v5 = unsafe.Add(v5, 4*1)
					if v4 >= 4 {
						goto LABEL_10
					}
				}
			}
		}
	LABEL_10:
		v2 = v2.InvNextItem
		if v2 == nil {
			return 0
		}
	}
}
func nox_xxx_playerAddGold_4FA590(a1 unsafe.Pointer, a2 int32) {
	var v2 int32
	v2 = int32(*(*uint32)(unsafe.Add(a1, 748)))
	*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v2, 276)), 2164)) += uint32(a2)
	sub_56F920(int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v2, 276)), 4588))), a2)
}
func nox_xxx_playerSubGold_4FA5D0(a1 *server.Object, a2 uint32) {
	v2 := a1.UpdateData
	v3 := *(*unsafe.Pointer)(unsafe.Add(v2, 276))
	v4 := *(*uint32)(unsafe.Add(v3, 2164))
	if v4 >= a2 {
		*(*uint32)(unsafe.Add(v3, 2164)) = v4 - a2
	} else {
		*(*uint32)(unsafe.Add(v3, 2164)) = 0
	}
	sub_56F920(int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v2, 276)), 4588))), int32(-a2))
}
func nox_object_setGold_4FA620(a1p *server.Object, a2 int32) {
	var (
		a1 = a1p
		v3 int32
		v4 int32
	)
	if a1 != nil && int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 8)))&4 != 0 {
		v3 = int32(a1.UpdateData)
		if a2 >= 0 || (func() bool {
			v4 = int32(*(*uint32)(unsafe.Add(v3, 276)))
			return *(*uint32)(unsafe.Add(v4, 2164)) >= uint32(-a2)
		}()) {
			*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 2164)) += uint32(a2)
			sub_56F920(int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 4588))), a2)
		} else {
			*(*uint32)(unsafe.Add(v4, 2164)) = 0
			nox_xxx_playerResetProtectionCRC_56F7D0(int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 4588))), 0)
		}
	}
}
func nox_xxx_playerGetGold_4FA6B0(a1 *server.Object) int32 {
	return int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a1.UpdateData, 276)), 2164)))
}
func nox_object_getGold_4FA6D0(a1p *server.Object) int32 {
	var (
		a1     = a1p
		result int32
	)
	if a1 != nil && int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 8)))&4 != 0 {
		result = int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a1.UpdateData, 276)), 2164)))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_playerBotCreate_4FA700(obj *server.Object) {
	ud := obj.UpdateDataPlayer()
	if ud.UpdateDataBot == nil {
		m, _ := alloc.New(server.MonsterUpdateData{})
		ud.UpdateDataBot = m
	}
	m := ud.UpdateDataBot
	if m != nil {
		typ := nox_xxx_getNameId_4E3AA0(internCStr("NPC"))
		*m = server.MonsterUpdateData{
			UpdateDataBot:         ud,
			MonsterDef:            Nox_xxx_monsterDefByTT_517560(int(typ)),
			RetreatLevel:          math.Float32frombits(1048576000),
			Field335:              1,
			ResumeLevel:           math.Float32frombits(1061997773),
			Field337:              1,
			StatusFlags:           0x2D808,
			Field361:              0,
			Field85:               38,
			Aggression2:           math.Float32frombits(1056964608),
			Aggression:            math.Float32frombits(1062501089),
			Field328:              math.Float32frombits(1125515264),
			Field329:              math.Float32frombits(1106247680),
			Field330:              math.Float32frombits(1065353216),
			Field332:              math.Float32frombits(1056964608),
			Field338:              math.Float32frombits(1065353216),
			ScriptLookingForEnemy: server.ScriptCallback{Func: -1},
			ScriptEnemySighted:    server.ScriptCallback{Func: -1},
			ScriptChangeFocus:     server.ScriptCallback{Func: -1},
			ScriptIsHit:           server.ScriptCallback{Func: -1},
			ScriptRetreat:         server.ScriptCallback{Func: -1},
			ScriptDeath:           server.ScriptCallback{Func: -1},
			ScriptCollision:       server.ScriptCallback{Func: -1},
			ScriptHearEnemy:       server.ScriptCallback{Func: -1},
			ScriptEndOfWaypoint:   server.ScriptCallback{Func: -1},
			ScriptLostEnemy:       server.ScriptCallback{Func: -1},
			Field510:              3,
			DialogStartFunc:       -1,
			DialogEndFunc:         -1,
			Field0:                0xDEADFACE,
			Direction94:           uint32(int16(obj.Direction1)),
			Pos95:                 obj.PosVec,
		}
		m.AIStackInd = 0
		m.AIStack[0].Action = ai.ACTION_HUNT
		*(*uint8)(unsafe.Pointer(&m.Field331)) = 30
		*(*uint8)(unsafe.Pointer(&m.Field333)) = math.MaxUint8
		switch ud.Player.PlayerClass() {
		case player.Warrior:
			m.FleeRange = 0
		case player.Wizard:
			m.FleeRange = math.Float32frombits(1112014848)
			m.Field410 = 0x8000000
			m.StatusFlags |= object.MonStatusCanCastSpells
			m.Field423 = 0x10000000
			m.Field408 = 0x10000000
			m.Field411 = 0x10000000
			m.Field384 = 0x20000000
			m.Field405 = 0x20000000
			m.Field362_0 = 0
			m.Field362_2 = uint16(gameFPS() / 2)
			m.Field364_0 = 0
			m.Field364_2 = uint16(gameFPS())
			m.Field366_0 = uint16(gameFPS() * 3)
			m.Field366_2 = uint16(gameFPS() * 30)
			if v := nox_common_randomInt_415FA0(0, 100); v < 33 {
				m.Field399 = 0x40000000
			} else if v < 66 {
				m.Field388 = 0x40000000
			} else {
				m.Field415 = 0x40000000
				m.Field422 = 0x40000000
			}
			m.Field368_0 = 0
			m.Field376 = math.Float32frombits(2147483648)
			m.Field370_0 = 0
			m.Field368_2 = uint16(gameFPS() * 2)
			m.Field370_2 = uint16(gameFPS() * 6)
		case player.Conjurer:
			m.FleeRange = math.Float32frombits(1112014848)
			m.Field410 = 0x8000000
			m.StatusFlags |= object.MonStatusCanCastSpells
			m.Field362_0 = 0
			m.Field362_2 = uint16(gameFPS() / 2)
			m.Field430 = 0x10000000
			m.Field364_0 = 0
			m.Field364_2 = uint16(gameFPS() * 6)
			m.Field432 = 0x20000000
			m.Field446 = 0x20000000
			m.Field366_0 = uint16(gameFPS() * 3)
			m.Field368_0 = 0
			m.Field366_2 = uint16(gameFPS() * 30)
			m.Field401 = 0x40000000
			m.Field424 = 0x40000000
			m.Field456 = 0x40000000
			m.Field455 = 0x40000000
			m.Field464 = 0x40000000
			m.Field376 = math.Float32frombits(2147483648)
			m.Field370_0 = 0
			m.Field368_2 = uint16(gameFPS() * 2)
			m.Field370_2 = uint16(gameFPS() * 6)
		}
	}
}
func nox_xxx_mobMorphFromPlayer_4FAAC0(obj *server.Object) {
	if obj.Class().Has(object.ClassPlayer) {
		ud := obj.UpdateDataPlayer()
		obj.ObjClass &^= uint32(object.ClassPlayer)
		obj.ObjClass |= uint32(object.ClassMonster)
		obj.ObjSubClass = uint32(object.MonStatusUnused5)
		obj.UpdateData = unsafe.Pointer(ud.UpdateDataBot)
	}
}
func nox_xxx_mobMorphToPlayer_4FAAF0(obj *server.Object) {
	if obj.Class().Has(object.ClassMonster) {
		ud := obj.UpdateDataMonster()
		obj.ObjClass &^= uint32(object.ClassMonster)
		obj.ObjClass |= uint32(object.ClassPlayer)
		obj.ObjSubClass = 0
		obj.UpdateData = unsafe.Pointer(ud.UpdateDataBot)
	}
}
func Nox_xxx_updatePlayerMonsterBot_4FAB20(obj *server.Object) {
	ud := obj.UpdateDataPlayer()
	if ud.UpdateDataBot == nil {
		nox_xxx_playerBotCreate_4FA700(obj)
	}
	if ud.UpdateDataBot == nil {
		obj.Update.Set(nox_xxx_updatePlayer_4F8100)
		return
	}
	if !nox_xxx_respawnPlayerBot_4FAC70(obj) {
		mud := ud.UpdateDataBot
		mud.StatusFlags |= object.MonStatusDestroyWhenDead
		nox_xxx_mobMorphFromPlayer_4FAAC0(obj)
		nox_xxx_unitUpdateMonster_50A5C0(obj)
		nox_xxx_mobMorphToPlayer_4FAAF0(obj)
		ud.State = uint8(nox_xxx_monsterActionToPlrState_4FABC0(obj))
		ud.Field59_0 = mud.Field120_1
		ud.Player.Pos3632Vec = obj.PosVec
	}
}
func nox_xxx_monsterActionToPlrState_4FABC0(obj *server.Object) int8 {
	ud := obj.UpdateDataPlayer()
	mud := ud.UpdateDataBot
	act := mud.AIStackHead()
	if act == nil {
		return 13
	}
	switch act.Type() {
	case ai.ACTION_MOVE_TO, ai.ACTION_FAR_MOVE_TO, ai.ACTION_ROAM, ai.ACTION_FIND_OBJECT, ai.ACTION_RANDOM_WALK:
		if mud.StatusFlags.Has(object.MonStatusRunning) {
			return 5
		}
		return 0
	case ai.ACTION_DODGE:
		return 0
	case ai.ACTION_MELEE_ATTACK, ai.ACTION_MISSILE_ATTACK:
		return 1
	case ai.ACTION_CAST_SPELL_ON_OBJECT, ai.ACTION_CAST_SPELL_ON_LOCATION, ai.ACTION_CAST_DURATION_SPELL:
		return 2
	case ai.ACTION_BLOCK_ATTACK, ai.ACTION_WEAPON_BLOCK:
		return 16
	case ai.ACTION_BLOCK_FINISH:
		return 17
	case ai.ACTION_FLEE:
		return 5
	case ai.ACTION_DYING:
		return 3
	case ai.ACTION_DEAD:
		return 4
	default:
		return 13
	}
}
func nox_xxx_respawnPlayerBot_4FAC70(obj *server.Object) bool {
	a1 := obj
	var (
		v1 int32
		v2 *byte
		v4 types.Pointf
	)
	v1 = int32(*(*uint32)(unsafe.Add(a1.UpdateData, 292)))
	v2 = (*byte)(sub_416640())
	if int32(*a1.HealthData) != 0 {
		return false
	}
	if gameFrame()-*(*uint32)(unsafe.Add(v1, 548)) < (gameFPS() * 2) {
		return true
	}
	nox_xxx_playerBotCreate_4FA700(obj)
	nox_xxx_playerMakeDefItems_4EF7D0(a1, 1, 0)
	if dword_5d4594_2650652 == 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(v2), 58)) != 0 {
		nox_xxx_respawnPlayerImpl_53FBC0(&a1.PosVec, int32(a1.Direction1))
	}
	nox_xxx_mapFindPlayerStart_4F7AB0(&v4, obj)
	nox_xxx_unitMove_4E7010(obj, &v4)
	nox_xxx_aud_501960(148, obj, 0, 0)
	if noxflags.HasGame(0x2000) {
		nox_xxx_buffApplyTo_4FF380(a1, 23, int16(int32(uint16(gameFPS()))*5), 5)
	}
	return false
}
func nox_xxx_netSendRewardNotify_4FAD50(a1 unsafe.Pointer, a2 int32, a3 unsafe.Pointer, a4 int8) int32 {
	var (
		result int32
		v5     int32
		v6     [5]byte
	)
	result = a1
	if !(a1 != nil && int32(*(*uint8)(unsafe.Add(a1, 8)))&4 != 0) {
		return result
	}
	v5 = int32(*(*uint32)(unsafe.Add(a1, 748)))
	v6[0] = 240
	if a2 != 0 {
		if a2 == 1 {
			v6[1] = 31
		} else {
			result = a2 - 2
			if a2 != 2 {
				return result
			}
			v6[1] = 32
		}
	} else {
		v6[1] = 30
	}
	*(*uint16)(unsafe.Pointer(&v6[3])) = *(*uint16)(unsafe.Add(a3, 36))
	v6[2] = byte(a4)
	result = nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v5, 276)), 2064))), unsafe.Pointer(&v6[0]), 5, nil, 1)
	return result
}
func sub_4FADD0(a1 *server.Object, a2 *byte, a3 int8) {
	var (
		v4 uint32
		v5 [52]byte
	)
	if a1 != nil {
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 8)))&4 != 0 {
			if a2 != nil {
				v4 = uint32(libc.StrLen(a2) + 1)
				if int32(uint8(v4)) != 1 && int32(uint8(v4-1)) <= 0x30 {
					v5[51] = byte(a3)
					v5[0] = 240
					v5[1] = 33
					libc.StrCpy(&v5[2], a2)
					nox_xxx_netSendPacket0_4E5420(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a1.UpdateData, 276)), 2064))), unsafe.Pointer(&v5[0]), 52, nil, 1)
				}
			}
		}
	}
}
func sub_4FB000(a1 *server.Object, a2 *server.Object) int32 {
	var (
		v2     unsafe.Pointer
		v3     *byte
		v4     int32
		result int32
	)
	if a1 != nil && a2 != nil && int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 8)))&4 != 0 && int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a2), 8)))&2 != 0 && (func() bool {
		v2 = a1.UpdateData
		v3 = nox_xxx_getUnitName_4E39D0(a2)
		return (func() int32 {
			v4 = nox_xxx_guide_427010(v3)
			return v4
		}()) != 0
	}()) {
		result = int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v2, 276)), uint32(v4*4)+4244)))
	} else {
		result = 0
	}
	return result
}
func sub_4FB050(a1 *server.Object, a2 *server.Object, a3 *int32) {
	var (
		result int32
		v4     float32
	)
	result = sub_4FB000(a1, a2)
	if result == 0 {
		return
	}
	v4 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("FieldGuideDamageBonus"))*float64(*a3) + 0.5)
	result = int32(v4)
	*a3 = result
}
func Nox_xxx_playerDoSchedSpell_4FB0E0(a1p *server.Object, a2 *server.Object) int32 {
	var (
		v5 int32
		v6 *uint32
	)
	v2 := a1p
	ud := a1p.UpdateDataPlayer()
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 212))) == 0 {
		return 0
	}
	a1 := nox_xxx_checkPlrCantCastSpell_4FD150(a1p, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(ud), 192))), 0)
	var sa server.SpellAcceptArg
	sa.Obj = a2
	sa.Pos.X = float32(float64(ud.Field55))
	sa.Pos.Y = float32(float64(ud.Field56))
	if a1 != 0 {
		nox_xxx_netInformTextMsg_4DA0F0(int32(ud.Player.PlayerInd), 0, &a1)
		nox_xxx_aud_501960(231, v2, 0, 0)
	} else {
		nox_xxx_castSpellByUser_4FDD20(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(ud), 192))), v2, &sa)
	}
	v5 = 1
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 212))) > 1 {
		v6 = (*uint32)(unsafe.Add(unsafe.Pointer(ud), 192))
		for {
			v5++
			*v6 = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*1))
			v6 = (*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*1))
			if v5 >= int32(*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 212))) {
				break
			}
		}
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(ud), v5*4+188)) = 0
	*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 212))--
	return 1
}
func Nox_xxx_playerDoSchedSpellQueue_4FB1D0(a1p *server.Object, a2 *server.Object) int32 {
	v2 := a1p
	ud := a1p.UpdateDataPlayer()
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 212))) == 0 {
		return 0
	}
	a1 := nox_xxx_checkPlrCantCastSpell_4FD150(a1p, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(ud), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 212)))*4+188))), 0)
	var v5 server.SpellAcceptArg
	v5.Obj = a2
	v5.Pos.X = float32(float64(ud.Field55))
	v5.Pos.Y = float32(float64(ud.Field56))
	if a1 != 0 {
		nox_xxx_netInformTextMsg_4DA0F0(int32(ud.Player.PlayerInd), 0, &a1)
		nox_xxx_aud_501960(231, v2, 0, 0)
	} else {
		nox_xxx_castSpellByUser_4FDD20(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(ud), int32(*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 212)))*4+188))), v2, (*server.SpellAcceptArg)(unsafe.Pointer(&v5[0])))
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer(ud), 212))--
	return 1
}
func nox_xxx_resetMapInit_4FC570(a1 int32) {
	nox_xxx_resetMapInit_1569652 = uint32(a1)
}
func sub_4FC580(a1 int32) {
	dword_5d4594_1569656 = uint32(a1)
}
func sub_4FC960(a1 int32, a2 int8) {
	var (
		v4 int32
		v5 int32
	)
	result := nox_xxx_getFirstPlayerUnit_4DA7C0()
	for i := result; result != nil; i = result {
		if i != a1 {
			v5 = int32(i.NetCode)
			v4 = nox_xxx_spellGetPhoneme_4FE1C0(int32(*(*uint32)(unsafe.Add(a1, 36))), a2)
			nox_xxx_aud_501960(v4, (*server.Object)(a1), 2, v5)
		}
		result = nox_xxx_getNextPlayerUnit_4DA7F0(i)
	}
}
func nox_xxx_Fn_4FCAC0(a1 int32, a2 int32) int32 {
	sub_4FE8A0(a1)
	nox_alloc_class_free_all(nox_alloc_magicEnt_1569668)
	dword_5d4594_1569672 = nil
	for u := nox_xxx_getFirstPlayerUnit_4DA7C0(); u != nil; u = nox_xxx_getNextPlayerUnit_4DA7F0(u) {
		v3 := u.UpdateDataPlayer()
		v3.Field47_0 = 0
		v3.SpellCastStart = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 192)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 196)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 200)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 204)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 208)) = 0
		*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 212)) = 0
	}
	if a2 == 0 {
		return 1
	}
	if nox_setImaginaryCaster() == 0 {
		return 0
	}
	return 1
}
func nox_xxx_spellCastByBook_4FCB80() {
	var (
		v0  int32
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 *uint32
		v12 int32
		v13 int32
		v14 *uint32
		v15 *uint32
		v16 int32
		v17 int32
		v18 int32
		v19 *byte
		v20 uint8
		v21 int32
		v22 int32
		v23 [2]byte
		v24 int32
		v25 int32
		v26 int32
		v27 int32
	)
	v0 = int32(dword_5d4594_1569672)
	if dword_5d4594_1569672 == nil {
		return
	}
	for {
		v1 = int32(*(*uint32)(unsafe.Add(v0, 4)))
		if *(*uint32)(unsafe.Add(v1, 16))&0x8020 != 0 {
			v2 = int32(*(*uint32)(unsafe.Add(v0, 52)))
			if v2 != 0 {
				*(*uint32)(unsafe.Add(v2, 56)) = *(*uint32)(unsafe.Add(v0, 56))
			}
			v3 = int32(*(*uint32)(unsafe.Add(v0, 56)))
			if v3 == 0 {
				dword_5d4594_1569672 = *(*uint32)(unsafe.Add(v0, 52))
			} else {
				*(*uint32)(unsafe.Add(v3, 52)) = *(*uint32)(unsafe.Add(v0, 52))
			}
			goto LABEL_40
		}
		if gameFrame() < uint32(*(*int32)(unsafe.Add(v0, 40))) {
			v0 = int32(*(*uint32)(unsafe.Add(v0, 52)))
			goto LABEL_48
		}
		v4 = 0
		if int32(*(*uint8)(unsafe.Add(v1, 8)))&4 != 0 {
			v4 = int32(*(*uint32)(unsafe.Add(v1, 748)))
		}
		if int32(*(*uint8)(unsafe.Add(v0, 36))) == 0 {
			v23[0] = 112
			v5 = int32(*(*uint32)(unsafe.Add(v4, 276)))
			v23[1] = *(*uint8)(unsafe.Add(unsafe.Pointer(uintptr(v0+int32(*(*uint8)(unsafe.Add(v0, 28)))*4)), 8))
			nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Add(v5, 2064))), 1, &v23[0], 2)
		}
		v6 = int32(*(*uint8)(unsafe.Add(v0, 28)))
		v7 = int32(**(**uint32)(unsafe.Add(v0, 32)))
		v24 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(v0+v6*4)), 8)))
		v8 = v24
		if v7 != v24 {
			v19 = (*byte)(sub_416640())
			*(*uint8)(unsafe.Pointer(&v27)) = uint8(nox_xxx_spellPhonemes_424A20(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(v0+int32(*(*uint8)(unsafe.Add(v0, 28)))*4)), 8))), int32(*(*uint8)(unsafe.Add(v0, 36)))))
			v20 = uint8(int8(v27))
			if dword_5d4594_2650652 == 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(v19), 62)) != 0 {
				sub_4FC960(int32(*(*uint32)(unsafe.Add(v0, 4))), int8(v27))
			}
			v21 = int32(uintptr(nox_xxx_updateSpellRelated_424830(*(*unsafe.Pointer)(unsafe.Add(v0, 32)), int32(v20))))
			v22 = int32(*(*uint32)(unsafe.Add(v0, 4)))
			*(*uint32)(unsafe.Add(v0, 32)) = uint32(v21)
			if int32(*(*uint8)(unsafe.Add(v22, 8)))&4 != 0 {
				*(*uint32)(unsafe.Add(v4, 184)) = *(*uint32)(unsafe.Add(v0, 32))
			}
			*(*uint8)(unsafe.Add(v0, 36))++
			*(*uint32)(unsafe.Add(v0, 40)) = gameFrame() + *(*uint32)(unsafe.Add(v0, 44))
			v0 = int32(*(*uint32)(unsafe.Add(v0, 52)))
			goto LABEL_48
		}
		v26 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(v0+v6*4)), 12)))
		if int32(*(*uint8)(unsafe.Add(v0, 29))) == 0 {
			if v8 != 34 && v26 != 0 {
				*(*uint8)(unsafe.Add(v0, 36)) = 0
				*(*uint32)(unsafe.Add(v0, 32)) = uint32(uintptr(nox_xxx_spellGetDefArrayPtr_424820()))
				*(*uint32)(unsafe.Add(v0, 40)) = gameFrame() + *(*uint32)(unsafe.Add(v0, 44))
				*(*uint8)(unsafe.Add(v0, 28))++
				v0 = int32(*(*uint32)(unsafe.Add(v0, 52)))
				goto LABEL_48
			}
		} else if v24 != 34 {
			if int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v0, 4)), 8)))&4 != 0 {
				v9 = 0
				v10 = 1
				if int32(*(*uint8)(unsafe.Add(v4, 212))) != 0 {
					v11 = (*uint32)(unsafe.Add(v4, 192))
					for {
						if *v11 == uint32(v8) {
							v12 = int32(*(*uint32)(unsafe.Add(v4, 276)))
							v25 = 6
							nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(v12, 2064))), 0, &v25)
							v8 = v24
							v10 = 0
						}
						v9++
						v11 = (*uint32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
						if v9 >= int32(*(*uint8)(unsafe.Add(v4, 212))) {
							break
						}
					}
				}
				if v10 != 0 {
					if sub_4FCF90((*server.Object)(*(*unsafe.Pointer)(unsafe.Add(v0, 4))), v8, 2) < 0 {
						v13 = int32(*(*uint32)(unsafe.Add(v4, 276)))
						v25 = 11
						nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(v13, 2064))), 0, &v25)
						nox_xxx_aud_501960(232, (*server.Object)(*(*unsafe.Pointer)(unsafe.Add(v0, 4))), 0, 0)
					} else {
						*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(v4+int32(func() uint8 {
							p := (*uint8)(unsafe.Add(v4, 212))
							x := *p
							*p++
							return x
						}())*4)), 192)) = uint32(v24)
					}
					v8 = v24
				}
			}
			if v8 != 34 && v26 != 0 {
				*(*uint8)(unsafe.Add(v0, 36)) = 0
				*(*uint32)(unsafe.Add(v0, 32)) = uint32(uintptr(nox_xxx_spellGetDefArrayPtr_424820()))
				*(*uint32)(unsafe.Add(v0, 40)) = gameFrame() + *(*uint32)(unsafe.Add(v0, 44))
				*(*uint8)(unsafe.Add(v0, 28))++
				v0 = int32(*(*uint32)(unsafe.Add(v0, 52)))
				goto LABEL_48
			}
		}
		v14 = *(**uint32)(unsafe.Add(v0, 4))
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v14), 4*2))&4 != 0 {
			v15 = *(**uint32)(unsafe.Add(v4, 276))
			*(*uint32)(unsafe.Add(v4, 220)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v15), 4*571))
			*(*uint32)(unsafe.Add(v4, 224)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v15), 4*572))
			if *(*uint32)(unsafe.Add(v0, 48)) != 0 {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v15), 4*910)) = *(*uint32)(unsafe.Add(v0, 4))
			} else {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v15), 4*910)) = *(*uint32)(unsafe.Add(v4, 288))
			}
			nox_xxx_playerSpell_4FB2A0_magic_plyrspel((*server.Object)(*(*unsafe.Pointer)(unsafe.Add(v0, 4))))
			*(*uint32)(unsafe.Add(v4, 216)) = 0
			*(*uint8)(unsafe.Add(v4, 188)) = 0
			*(*uint8)(unsafe.Add(v4, 212)) = 0
		} else {
			nox_xxx_castSpellByUser_4FDD20(v8, (*server.Object)(unsafe.Pointer(v14)), nil)
		}
		v16 = int32(*(*uint32)(unsafe.Add(v0, 52)))
		if v16 != 0 {
			*(*uint32)(unsafe.Add(v16, 56)) = *(*uint32)(unsafe.Add(v0, 56))
		}
		v17 = int32(*(*uint32)(unsafe.Add(v0, 56)))
		if v17 == 0 {
			dword_5d4594_1569672 = *(*uint32)(unsafe.Add(v0, 52))
		} else {
			*(*uint32)(unsafe.Add(v17, 52)) = *(*uint32)(unsafe.Add(v0, 52))
		}
	LABEL_40:
		v18 = int32(*(*uint32)(unsafe.Add(v0, 52)))
		nox_alloc_class_free_obj_first(nox_alloc_magicEnt_1569668, v0)
		v0 = v18
	LABEL_48:
		if v0 == 0 {
			return
		}
	}
}
func sub_4FCEB0(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
	)
	result = int32(uintptr(nox_xxx_spellCastedFirst_4FE930()))
	if result == 0 {
		return result
	}
	for {
		v2 = int32(*(*uint32)(unsafe.Add(result, 116)))
		if a1 != 1 || (func() int32 {
			v3 = int32(*(*uint32)(unsafe.Add(result, 48)))
			return v3
		}()) == 0 || (int32(*(*uint8)(unsafe.Add(v3, 8)))&4) != 4 {
			nox_xxx_spellCancelSpellDo_4FE9D0(result)
		}
		result = v2
		if v2 == 0 {
			break
		}
	}
	return result
}
func nox_xxx_spellCheckSmth_4FCEF0(a1 unsafe.Pointer, a2 *int32, a3 int32) int32 {
	var (
		v3 *int32
		v5 int32
		v6 int32
		v7 int32
		v8 int32
	)
	if a1 == nil {
		return 0
	}
	v3 = a2
	if a2 == nil {
		return 0
	}
	if a3 == 0 {
		return 0
	}
	if nox_common_getEngineFlag(NOX_ENGINE_FLAG_GODMODE) {
		return 1
	}
	if int32(*(*uint8)(unsafe.Add(a1, 8)))&2 != 0 {
		return 1
	}
	v5 = int32(uint16(nox_xxx_unitGetOldMana_4EEC80((*server.Object)(a1))))
	v6 = 0
	if a3 <= 0 {
		return 1
	}
	for {
		v7 = *v3
		if *v3 < 75 || v7 > 114 {
			v8 = nox_xxx_spellManaCost_4249A0(v7, 2)
		} else {
			v8 = sub_500CA0(v7, a1)
		}
		if v8 > v5 {
			break
		}
		v5 -= v8
		v6++
		v3 = (*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1))
		if v6 >= a3 {
			return 1
		}
	}
	return 0
}
func sub_4FCF90(a1p *server.Object, a2 int32, a3 int32) int32 {
	var (
		a1     = unsafe.Pointer(a1p)
		v3     *uint16
		result int32
		v5     int32
		v6     int32
	)
	v3 = *(**uint16)(unsafe.Add(a1, 748))
	if (int32(*(*uint8)(unsafe.Add(a1, 8))) & 4) == 0 {
		return -1
	}
	if a2 == 0 {
		return -1
	}
	if nox_common_getEngineFlag(NOX_ENGINE_FLAG_GODMODE) {
		return 0
	}
	if a2 < 75 || a2 > 114 {
		v5 = nox_xxx_spellManaCost_4249A0(a2, a3)
	} else {
		v5 = sub_500CA0(a2, a1)
	}
	v6 = v5
	if int32(*(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*2))) >= v5 {
		nox_xxx_playerManaSub_4EEBF0(a1, v5)
		result = v6
	} else {
		*(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*40)) = uint16(int16(nox_xxx_spellManaCost_4249A0(a2, 1)))
		result = -1
		*(*uint16)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(uint16(0))*41)) = uint16(gameFPS())
	}
	return result
}
func sub_4FD030(a1 int32, a2 int16) {
	if int32(*(*uint8)(unsafe.Add(a1, 8)))&4 != 0 {
		nox_xxx_playerManaAdd_4EEB80((*server.Object)(a1), a2)
	}
}
func sub_4FD0E0(a1p *server.Object, a2 int32) int32 {
	var (
		a1 = a1p
		v2 int32
		v4 int32
	)
	nox_xxx_spellFlags_424A70(a2)
	v2 = nox_xxx_findParentChainPlayer_4EC580(a1)
	if !nox_xxx_spellIsEnabled_424B70(a2) {
		return 10
	}
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 8)))&4 != 0 {
		return nox_xxx_playerCheckSpellClass_57AEA0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a1.UpdateData, 276)), 2251))), a2)
	}
	v4 = -sub_57AEE0(a2, (*server.Object)(v2))
	*(*uint8)(unsafe.Pointer(&v4)) = uint8(int8(v4 & 0xF6))
	return v4 + 10
}
func nox_xxx_checkPlrCantCastSpell_4FD150(a1p *server.Object, a2 int32, a3 int32) int32 {
	var (
		a1     = a1p
		v3     int32
		result int32
	)
	var v5 int32
	var v6 int32
	var v7 int32
	var v8 float64
	var v9 int32
	var v10 int32
	var v11 int32
	var v12 int32
	nox_xxx_findParentChainPlayer_4EC580(a1)
	if a3 != 0 {
		goto LABEL_9
	}
	if noxflags.HasGame(16) {
		if *memmap.PtrUint32(0x5D4594, 1569704) == 0 {
			*memmap.PtrUint32(0x5D4594, 1569704) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("Crown")))
		}
		if nox_xxx_spellHasFlags_424A50(a2, 0x80000) {
			v3 = int32(a1.Field129)
			if v3 != 0 {
				for uint32(*(*uint16)(unsafe.Add(v3, 4))) != *memmap.PtrUint32(0x5D4594, 1569704) {
					v3 = int32(*(*uint32)(unsafe.Add(v3, 512)))
					if v3 == 0 {
						goto LABEL_9
					}
				}
				if nox_xxx_servObjectHasTeam_419130((*server.ObjectTeam)(unsafe.Add(unsafe.Pointer(a1), 48))) != 0 {
					return 17
				}
			}
		}
		goto LABEL_9
	}
	if !noxflags.HasGame(64) {
		if !noxflags.HasGame(32) {
			goto LABEL_9
		}
		if !nox_xxx_spellHasFlags_424A50(a2, 0x80000) {
			goto LABEL_9
		}
		v6 = int32(a1.InvFirstItem)
		if v6 == 0 {
			goto LABEL_9
		}
		for (*(*uint32)(unsafe.Add(v6, 8)) & 0x10000000) == 0 {
			v6 = int32(*(*uint32)(unsafe.Add(v6, 496)))
			if v6 == 0 {
				goto LABEL_9
			}
		}
		return 13
	}
	if *memmap.PtrUint32(0x5D4594, 1569708) == 0 {
		*memmap.PtrUint32(0x5D4594, 1569708) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("GameBall")))
	}
	if nox_xxx_spellHasFlags_424A50(a2, 0x80000) {
		v5 = int32(a1.Field129)
		if v5 != 0 {
			for uint32(*(*uint16)(unsafe.Add(v5, 4))) != *memmap.PtrUint32(0x5D4594, 1569708) {
				v5 = int32(*(*uint32)(unsafe.Add(v5, 512)))
				if v5 == 0 {
					goto LABEL_9
				}
			}
			return 16
		}
	}
LABEL_9:
	if nox_xxx_testUnitBuffs_4FF350(a1, 29) != 0 {
		return 14
	}
	if sub_4D7100(a2) == 0 {
		return 10
	}
	switch a2 {
	case 29:
		if nox_xxx_unitIsUnitTT_4E7C80(a1, *memmap.PtrInt32(0x5D4594, 1569692)) > 0 {
			return 3
		}
		if nox_xxx_unitIsUnitTT_4E7C80(a1, *memmap.PtrInt32(0x5D4594, 1569688)) <= 0 {
			v10 = int32(*memmap.PtrUint32(0x5D4594, 1569684))
			v9 = nox_xxx_unitIsUnitTT_4E7C80(a1, v10)
			if v9 <= 0 {
				return 0
			}
		}
		return 3
	case 31:
		v9 = nox_xxx_unitIsUnitTT_4E7C80(a1, *memmap.PtrInt32(0x5D4594, 1569696))
		if v9 <= 0 {
			return 0
		}
		return 3
	case 50:
		v7 = nox_xxx_unitIsUnitTT_4E7C80(a1, *memmap.PtrInt32(0x5D4594, 1569680))
		v12 = nox_xxx_spellGetPower_4FE7B0(a2, a1) - 1
		v8 = nox_xxx_gamedataGetFloatTable_419D70(internCStr("MagicMissileCount"), v12)
		if v7 < int32(int64(v8)) {
			return 0
		}
		return 3
	case 52:
		v10 = int32(*memmap.PtrUint32(0x5D4594, 1569700))
		v9 = nox_xxx_unitIsUnitTT_4E7C80(a1, v10)
		if v9 <= 0 {
			return 0
		}
		return 3
	case 58:
		v7 = nox_xxx_unitIsUnitTT_4E7C80(a1, *memmap.PtrInt32(0x5D4594, 1569676))
		v11 = nox_xxx_spellGetPower_4FE7B0(a2, a1) - 1
		v8 = nox_xxx_gamedataGetFloatTable_419D70(internCStr("PixieCount"), v11)
		if v7 < int32(int64(v8)) {
			return 0
		}
		return 3
	default:
		return 0
	}
}
func nox_xxx_gameCaptureMagic_4FDC10(a1 int32, a2p *server.Object) int32 {
	var (
		a2 = unsafe.Pointer(a2p)
		v3 int32
		v4 int32
		v5 int32
	)
	if *memmap.PtrUint32(0x5D4594, 1569712) == 0 {
		*memmap.PtrUint32(0x5D4594, 1569712) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("GameBall")))
		*memmap.PtrUint32(0x5D4594, 1569716) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("Crown")))
	}
	if a2 == nil {
		return 0
	}
	if !(nox_xxx_spellHasFlags_424A50(a1, 0x80000) && int32(*(*uint8)(unsafe.Add(a2, 8)))&4 != 0) {
		return 1
	}
	v3 = int32(*(*uint32)(unsafe.Add(a2, 748)))
	if noxflags.HasGame(32) {
		if int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 4)))&1 != 0 {
			return 0
		}
	} else if noxflags.HasGame(64) {
		v4 = int32(*(*uint32)(unsafe.Add(a2, 516)))
		if v4 != 0 {
			for uint32(*(*uint16)(unsafe.Add(v4, 4))) != *memmap.PtrUint32(0x5D4594, 1569712) {
				v4 = int32(*(*uint32)(unsafe.Add(v4, 512)))
				if v4 == 0 {
					return 1
				}
			}
			return 0
		}
	} else if noxflags.HasGame(16) {
		v5 = int32(*(*uint32)(unsafe.Add(a2, 516)))
		if v5 != 0 {
			for uint32(*(*uint16)(unsafe.Add(v5, 4))) != *memmap.PtrUint32(0x5D4594, 1569716) || nox_xxx_servObjectHasTeam_419130((*server.ObjectTeam)(unsafe.Add(a2, 48))) == 0 {
				v5 = int32(*(*uint32)(unsafe.Add(v5, 512)))
				if v5 == 0 {
					return 1
				}
			}
			return 0
		}
	}
	return 1
}
func nox_xxx_createSpellFly_4FDDA0(a1p *server.Object, a2p *server.Object, a3 int32) *uint32 {
	var (
		a1     = a1p
		a2     = a2p
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     float32
		v8     float64
		v9     float64
		v10    float64
		result *uint32
		v12    *uint32
		v13    *int32
		v14    int16
		v15    int32
		v16    int16
		v17    int32
		v18    int32
		v19    int32
		v20    int8
		v21    Point32
		v22    float4
		a2a    float32
	)
	v3 = a1
	a2a = float32(float64(a1.Shape.Circle) + 4.0)
	if a2 == nil {
		if int32(*(*uint8)(unsafe.Add(v3, 8)))&4 != 0 {
			v4 = int32(*(*uint32)(unsafe.Add(v3, 748)))
			*(*float32)(unsafe.Pointer(&v21.X)) = float32(float64(*(*int32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v4, 276)), 2284))))
			*(*float32)(unsafe.Pointer(&v21.Y)) = float32(float64(*(*int32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v4, 276)), 2288))))
			v18 = int32(nox_xxx_spellFlags_424A70(a3))
			v5 = int32(uintptr(unsafe.Pointer(nox_xxx_spellFlySearchTarget_540610((*types.Pointf)(unsafe.Pointer(&v21)), (*server.Object)(v3), v18, 600.0, 0, (*server.Object)(v3)))))
		} else {
			v19 = int32(nox_xxx_spellFlags_424A70(a3))
			v5 = int32(uintptr(unsafe.Pointer(nox_xxx_spellFlySearchTarget_540610(nil, (*server.Object)(v3), v19, 600.0, 0, (*server.Object)(v3)))))
		}
		a2 = v5
	}
	v6 = int32(*(*int16)(unsafe.Add(v3, 124))) * 8
	v7 = *(*float32)(unsafe.Add(v3, 60))
	v8 = float64(a2a * *mem_getFloatPtr(0x587000, uint32(v6)+194136))
	v22.field_0 = *(*float32)(unsafe.Add(v3, 56))
	v9 = v8 + float64(*(*float32)(unsafe.Add(v3, 56)))
	v10 = float64(a2a * *mem_getFloatPtr(0x587000, uint32(v6)+194140))
	v22.field_4 = v7
	v22.field_C = float32(v10 + float64(*(*float32)(unsafe.Add(v3, 60))))
	v22.field_8 = float32(v9 + float64(*(*float32)(unsafe.Add(v3, 80))))
	v22.field_C = v22.field_C + *(*float32)(unsafe.Add(v3, 84))
	result = (*uint32)(unsafe.Pointer(uintptr(nox_xxx_mapTraceRay_535250(&v22, nil, nil, 5))))
	if result == nil {
		return result
	}
	result = nox_xxx_newObjectByTypeID_4E3810(internCStr("Magic"))
	v12 = result
	if result == nil {
		return result
	}
	v13 = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(result), 4*187)))
	*(*int32)(unsafe.Add(unsafe.Pointer(v13), 4*4)) = nox_xxx_spellGetPower_4FE7B0(a3, (*server.Object)(v3))
	nox_xxx_createAt_4DAA50((*server.Object)(unsafe.Pointer(v12)), (*server.Object)(v3), v22.field_8, v22.field_C)
	v14 = int16(*(*uint16)(unsafe.Add(v3, 124)))
	*(*uint16)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint16(0))*62)) = uint16(v14)
	*(*uint16)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(uint16(0))*63)) = uint16(v14)
	*v13 = v3
	*(*int32)(unsafe.Add(unsafe.Pointer(v13), 4*1)) = a2
	*(*int32)(unsafe.Add(unsafe.Pointer(v13), 4*2)) = v3
	*(*int32)(unsafe.Add(unsafe.Pointer(v13), 4*3)) = a3
	nox_xxx_xferIndexedDirection_509E20(int32(*(*int16)(unsafe.Add(v3, 124))), &v21)
	v15 = int32(*(*int16)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(int16(0))*62)))
	*(*float32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(float32(0))*20)) = *mem_getFloatPtr(0x587000, uint32(v15*8)+194136) * *(*float32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(float32(0))*136))
	*(*float32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(float32(0))*21)) = *mem_getFloatPtr(0x587000, uint32(v15*8)+194140) * *(*float32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(float32(0))*136))
	*(*float32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(float32(0))*20)) = *(*float32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(float32(0))*20)) + *(*float32)(unsafe.Add(v3, 80))
	*(*float32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(float32(0))*21)) = *(*float32)(unsafe.Add(unsafe.Pointer(v12), unsafe.Sizeof(float32(0))*21)) + *(*float32)(unsafe.Add(v3, 84))
	if nox_xxx_testUnitBuffs_4FF350((*server.Object)(v3), 21) != 0 {
		v20 = nox_xxx_buffGetPower_4FF570((*server.Object)(v3), 21)
		v16 = int16(nox_xxx_unitGetBuffTimer_4FF550((*server.Object)(v3), 21))
		nox_xxx_buffApplyTo_4FF380((*server.Object)(unsafe.Pointer(v12)), 21, v16, v20)
	}
	v17 = nox_xxx_spellGetAud44_424800(a3, 0)
	nox_xxx_aud_501960(v17, (*server.Object)(v3), 0, 0)
	return result
}
func nox_xxx_collide_4FDF90(a1 int32, a2 int32) {
	var (
		v2     int32
		v3     int32
		result int32
		v5     float32
	)
	if nox_xxx_testUnitBuffs_4FF350((*server.Object)(a1), 22) != 0 && (*(*uint32)(unsafe.Add(a2, 16))&0x8008) == 0 && int32(*(*uint8)(unsafe.Add(a2, 8)))&6 != 0 && nox_xxx_unitIsEnemyTo_5330C0((*server.Object)(a1), (*server.Object)(a2)) != 0 {
		v2 = int32(*(*uint8)(unsafe.Add(a1, 430))) - 1
		nox_xxx_aud_501960(135, (*server.Object)(a1), 0, 0)
		nox_xxx_spellBuffOff_4FF5B0((*server.Object)(a1), 22)
		v5 = float32(nox_xxx_gamedataGetFloatTable_419D70(internCStr("ShockDamage"), v2))
		v3 = int32(v5)
		ccall.AsFunc[func(int32, int32, int32, int32, int32)](*(*unsafe.Pointer)(unsafe.Add(a2, 716)))(a2, a1, a1, v3, 9)
	}
	result = int32(*(*uint32)(unsafe.Add(a2, 8)))
	if uint32(result)&0x20006 != 0 {
		if (*(*uint32)(unsafe.Add(a2, 16)) & 0x8020) == 0 {
			if nox_xxx_unitsHaveSameTeam_4EC520((*server.Object)(a2), (*server.Object)(a1)) == 0 {
				nox_xxx_spellBuffOff_4FF5B0((*server.Object)(a1), 0)
			}
		}
	}
	if int32(*(*uint8)(unsafe.Add(a1, 8)))&4 != 0 && *(*uint32)(unsafe.Add(a2, 8))&0x20000 != 0 && (*(*uint32)(unsafe.Add(a2, 16))&0x8020) == 0 {
		nox_xxx_spellBuffOff_4FF5B0((*server.Object)(a1), 0)
	}
}
func nox_xxx_spellGetPhoneme_4FE1C0(a1 int32, a2 int8) int32 {
	var v2 *byte
	if noxflags.HasGame(1) {
		if (int32(*(*uint8)(unsafe.Add(unsafe.Pointer(nox_server_getObjectFromNetCode_4ECCB0(uint32(a1))), unsafe.Sizeof(server.Object{})*8))) & 4) == 0 {
			switch a2 {
			case 0:
				return 193
			case 1:
				return 186
			case 2:
				return 187
			case 3:
				return 192
			case 5:
				return 188
			case 6:
				return 191
			case 7:
				return 190
			case 8:
				return 189
			default:
				return 0
			}
		}
	} else if (nox_xxx_netSpriteByCodeDynamic_45A6F0(uint32(a1)).Flags28Val & 4) == 0 {
		switch a2 {
		case 0:
			return 193
		case 1:
			return 186
		case 2:
			return 187
		case 3:
			return 192
		case 5:
			return 188
		case 6:
			return 191
		case 7:
			return 190
		case 8:
			return 189
		default:
			return 0
		}
	}
	v2 = (*byte)(unsafe.Pointer(nox_common_playerInfoGetByID_417040(a1)))
	if v2 == nil {
		return 0
	}
	if *(*byte)(unsafe.Add(unsafe.Pointer(v2), 2252)) == 0 {
		switch a2 {
		case 0:
			return 193
		case 1:
			return 186
		case 2:
			return 187
		case 3:
			return 192
		case 5:
			return 188
		case 6:
			return 191
		case 7:
			return 190
		case 8:
			return 189
		default:
			return 0
		}
	}
	switch a2 {
	case 0:
		return 201
	case 1:
		return 194
	case 2:
		return 195
	case 3:
		return 200
	case 5:
		return 196
	case 6:
		return 199
	case 7:
		return 198
	case 8:
		return 197
	default:
		return 0
	}
}
func nox_xxx_spellByBookInsert_4FE340(a1p unsafe.Pointer, a2 *int32, a3 int32, a4 int32, a5 int32) int32 {
	var (
		a1  int32
		v5  *uint32
		v6  *int32
		v7  int32
		v8  *int32
		v9  int32
		v10 int32
		v11 *int32
		v12 *int32
		v13 int32
		v15 uint8
		v16 int32
		v17 *int32
		v18 int32
		v19 int32
		v20 *uint32
		v21 int32
		v22 int32
		v23 int32
		v24 *int32
		v25 *int32
		v26 int32
	)
	v5 = (*uint32)(a1p)
	if *(*uint32)(unsafe.Add(a1p, 16))&0x8022 != 0 {
		return 0
	}
	v6 = a2
	v7 = 0
	v8 = a2
	for {
		if *v8 < 0 || *v8 >= 137 {
			return 0
		}
		v7++
		v8 = (*int32)(unsafe.Add(unsafe.Pointer(v8), 4*1))
		if v7 >= 5 {
			break
		}
	}
	if (int32(*(*uint8)(unsafe.Add(a1p, 8))) & 4) == 0 {
		return 0
	}
	v9 = int32(*(*uint32)(unsafe.Add(a1p, 748)))
	if *(*uint32)(unsafe.Add(v9, 280)) != 0 {
		return 0
	}
	v10 = 0
	v11 = a2
	for {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(v9, 276))+uint32(*v11*4))), 3696)) == 0 && *v11 != 0 {
			return 0
		}
		v10++
		v11 = (*int32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
		if v10 >= 5 {
			break
		}
	}
	if *(*uint32)(unsafe.Add(v9, 216)) != 0 {
		return 0
	}
	v26 = 0
	if a3 > 0 {
		v12 = a2
		v13 = a3
		for {
			if *v12 == 34 {
				v26 = 1
			}
			v12 = (*int32)(unsafe.Add(unsafe.Pointer(v12), 4*1))
			v13--
			if v13 == 0 {
				break
			}
		}
		if v26 != 0 {
			if nox_xxx_spellCheckSmth_4FCEF0(a1p, a2, a3) == 0 {
				a1 := int32(12)
				nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
				nox_xxx_aud_501960(232, (*server.Object)(unsafe.Pointer(v5)), 0, 0)
				return 0
			}
			if int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2251))) == 2 {
				v15 = uint8(int8(bool2int32(nox_xxx_checkSummonedCreaturesLimit_500D70((*server.Object)(unsafe.Pointer(v5)), 5))))
				if int32(v15) == 0 {
					a1 := int32(4)
					nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
					nox_xxx_aud_501960(231, (*server.Object)(unsafe.Pointer(v5)), 0, 0)
					return 0
				}
				v16 = nox_xxx_unitCountSlaves_4E7CF0(int32(uintptr(unsafe.Pointer(v5))), 2, 0x2000)
				if v16 >= int32(int64(nox_xxx_gamedataGetFloat_419D40(internCStr("MaxBomberCount")))) {
					a1 := int32(5)
					nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
					nox_xxx_aud_501960(231, (*server.Object)(unsafe.Pointer(v5)), 0, 0)
					return 0
				}
			} else if int32(*(*uint8)(unsafe.Add(v9, 244))) >= int32(int64(nox_xxx_gamedataGetFloat_419D40(internCStr("MaxTrapCount")))) {
				a1 := int32(5)
				nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
				nox_xxx_aud_501960(231, (*server.Object)(unsafe.Pointer(v5)), 0, 0)
				return 0
			}
			v17 = a2
			v18 = 0
			for {
				a1 := sub_4FD0E0((*server.Object)(unsafe.Pointer(v5)), *v17)
				if a1 != 0 {
					nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
					nox_xxx_aud_501960(231, (*server.Object)(unsafe.Pointer(v5)), 0, 0)
					return 0
				}
				a1 = nox_xxx_checkPlrCantCastSpell_4FD150((*server.Object)(unsafe.Pointer(v5)), *v17, v26)
				if a1 != 0 {
					nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
					nox_xxx_aud_501960(231, (*server.Object)(unsafe.Pointer(v5)), 0, 0)
					return 0
				}
				v18++
				v17 = (*int32)(unsafe.Add(unsafe.Pointer(v17), 4*1))
				if v18 >= a3 {
					v6 = a2
					goto LABEL_36
				}
			}
		}
	}
	a1 = sub_4FD0E0((*server.Object)(a1p), *a2)
	if a1 != 0 {
		nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
		nox_xxx_aud_501960(231, (*server.Object)(unsafe.Pointer(v5)), 0, 0)
		return 0
	}
	a1 = nox_xxx_checkPlrCantCastSpell_4FD150((*server.Object)(unsafe.Pointer(v5)), *v6, 0)
	if a1 != 0 {
		nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
		nox_xxx_aud_501960(231, (*server.Object)(unsafe.Pointer(v5)), 0, 0)
		return 0
	}
LABEL_36:
	nox_xxx_playerSetState_4FA020((*server.Object)(unsafe.Pointer(v5)), 2)
	v19 = int32(gameFrame())
	*(*uint8)(unsafe.Add(v9, 188)) = 1
	*(*uint32)(unsafe.Add(v9, 216)) = uint32(v19)
	v20 = (*uint32)(nox_alloc_class_new_obj_zero(nox_alloc_magicEnt_1569668))
	if v20 == nil {
		return 0
	}
	v21 = a5
	v22 = a4
	*(*uint32)(unsafe.Add(unsafe.Pointer(v20), 4*1)) = uint32(uintptr(unsafe.Pointer(v5)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v20), 4*12)) = uint32(v21)
	*(*uint8)(unsafe.Add(unsafe.Pointer(v20), 36)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v20), 4*10)) = gameFrame()
	*(*uint32)(unsafe.Add(unsafe.Pointer(v20), 4*11)) = uint32(v22)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v20), 4*8)) = uint32(uintptr(nox_xxx_spellGetDefArrayPtr_424820()))
	*(*uint8)(unsafe.Add(unsafe.Pointer(v20), 28)) = 0
	*(*uint8)(unsafe.Add(unsafe.Pointer(v20), 29)) = 0
	v23 = 0
	v24 = v6
	v25 = (*int32)(unsafe.Add(unsafe.Pointer(v20), 4*2))
	for {
		if v23 >= a3 {
			*v25 = 0
		} else {
			*v25 = *v24
			if *v24 == 34 {
				*(*uint8)(unsafe.Add(unsafe.Pointer(v20), 29)) = 1
			}
		}
		v23++
		v24 = (*int32)(unsafe.Add(unsafe.Pointer(v24), 4*1))
		v25 = (*int32)(unsafe.Add(unsafe.Pointer(v25), 4*1))
		if v23 >= 5 {
			break
		}
	}
	*(*uint32)(unsafe.Add(unsafe.Pointer(v20), 4*14)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(v20), 4*13)) = dword_5d4594_1569672
	if dword_5d4594_1569672 != nil {
		*(*uint32)(unsafe.Add(dword_5d4594_1569672, 56)) = uint32(uintptr(unsafe.Pointer(v20)))
	}
	dword_5d4594_1569672 = unsafe.Pointer(v20)
	return 1
}
func nox_xxx_spell_4FE680(a1p *server.Object, a2 float32) {
	var (
		a1  = unsafe.Pointer(a1p)
		v2  int32
		v5  int32
		v6  float64
		v7  float64
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
	)
	v2 = int32(dword_5d4594_1569672)
	if dword_5d4594_1569672 != nil {
		v3 := a1
		for {
			v4 := *(*unsafe.Pointer)(unsafe.Add(v2, 4))
			if ((int32(*(*uint8)(unsafe.Add(v4, 8)))&4) != 4 || nox_xxx_servCompareTeams_419150((*server.ObjectTeam)(unsafe.Add(v3, 48)), unsafe.Add(v4, 48)) == 0) && (func() bool {
				v5 = int32(*(*uint32)(unsafe.Add(v2, 4)))
				v6 = float64(*(*float32)(unsafe.Add(v5, 56)) - *(*float32)(unsafe.Add(v3, 56)))
				v7 = float64(*(*float32)(unsafe.Add(v5, 60)) - *(*float32)(unsafe.Add(v3, 60)))
				return math.Sqrt(v7*v7+v6*v6)+0.1 < float64(a2)
			}()) && nox_xxx_mapCheck_537110((*server.Object)(v3), (*server.Object)(*(*unsafe.Pointer)(unsafe.Add(v2, 4)))) != 0 {
				v8 = int32(*(*uint32)(unsafe.Add(v2, 4)))
				if (int32(*(*uint8)(unsafe.Add(v8, 8))) & 4) == 4 {
					v9 = int32(*(*uint32)(unsafe.Add(v8, 748)))
					*(*uint32)(unsafe.Add(v9, 216)) = 0
					*(*uint8)(unsafe.Add(v9, 188)) = 0
					a1 := int32(15)
					nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v9, 276)), 2064))), 0, &a1)
					nox_xxx_aud_501960(231, (*server.Object)(*(*unsafe.Pointer)(unsafe.Add(v2, 4))), 0, 0)
					nox_xxx_playerSetState_4FA020((*server.Object)(unsafe.Pointer(*(**uint32)(unsafe.Add(v2, 4)))), 13)
				}
				v10 = int32(*(*uint32)(unsafe.Add(v2, 52)))
				if v10 != 0 {
					*(*uint32)(unsafe.Add(v10, 56)) = *(*uint32)(unsafe.Add(v2, 56))
				}
				v11 = int32(*(*uint32)(unsafe.Add(v2, 56)))
				if v11 != 0 {
					*(*uint32)(unsafe.Add(v11, 52)) = *(*uint32)(unsafe.Add(v2, 52))
				} else {
					dword_5d4594_1569672 = *(*uint32)(unsafe.Add(v2, 52))
				}
				v12 = int32(*(*uint32)(unsafe.Add(v2, 52)))
				nox_alloc_class_free_obj_first(nox_alloc_magicEnt_1569668, v2)
				v2 = v12
			} else {
				v2 = int32(*(*uint32)(unsafe.Add(v2, 52)))
			}
			if v2 == 0 {
				break
			}
		}
	}
}
func nox_xxx_spellGetPower_4FE7B0(a1 int32, a2p *server.Object) int32 {
	var (
		a2 = a2p
		v2 int32
		v4 int32
	)
	v2 = int32(*memmap.PtrUint32(0x5D4594, 1569720))
	if *memmap.PtrUint32(0x5D4594, 1569720) == 0 {
		v2 = nox_xxx_getNameId_4E3AA0(internCStr("ImaginaryCaster"))
		*memmap.PtrUint32(0x5D4594, 1569720) = uint32(v2)
	}
	if int32(a2.TypeInd) == v2 {
		return 1
	}
	if noxflags.HasGame(1392) {
		return 3
	}
	if a2 == nil {
		return 2
	}
	v4 = int32(a2.ObjClass)
	if v4&4 != 0 {
		return int32(*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(a2.UpdateData, 276))+uint32(a1*4))), 3696)))
	}
	if (v4 & 2) == 0 {
		return 3
	} else {
		return int32(*(*uint32)(unsafe.Add(a2.UpdateData, 2040)))
	}
}
func nox_xxx_spellCancelSpellDo_4FE9D0(a1p unsafe.Pointer) int8 {
	var (
		a1     int32 = int32(uintptr(a1p))
		v1     int32
		v2     int32
		v3     int32
		result int8
		i      int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(a1, 16)))
	if v1 != 0 && int32(*(*uint8)(unsafe.Add(v1, 8)))&4 != 0 {
		v2 = int32(*(*uint32)(unsafe.Add(a1, 4)))
		v3 = int32(*(*uint32)(unsafe.Add(v1, 748)))
		if v2 == 43 {
			nox_xxx_netReportSpellStat_4D9630(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 2064))), 43, 0)
		} else {
			nox_xxx_netReportSpellStat_4D9630(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 2064))), v2, 15)
		}
	}
	if *(*uint32)(unsafe.Add(a1, 4)) == 43 {
		for i = int32(*(*uint32)(unsafe.Add(a1, 108))); i != 0; i = int32(*(*uint32)(unsafe.Add(i, 116))) {
			if *(*uint32)(unsafe.Add(i, 48)) != 0 {
				nox_xxx_netStopRaySpell_4FEF90(i, *(**uint32)(unsafe.Add(i, 48)))
			}
		}
	} else if *(*uint32)(unsafe.Add(a1, 48)) != 0 {
		nox_xxx_netStopRaySpell_4FEF90(a1, *(**uint32)(unsafe.Add(a1, 48)))
		result = int8(int32(*(*uint8)(unsafe.Add(a1, 88))) | 1)
		*(*uint8)(unsafe.Add(a1, 88)) = uint8(result)
		return result
	}
	result = int8(int32(*(*uint8)(unsafe.Add(a1, 88))) | 1)
	*(*uint8)(unsafe.Add(a1, 88)) = uint8(result)
	return result
}
func sub_4FEA70(a1 int32, a2 *types.Pointf) int32 {
	var (
		v2 float64
		v3 float64
		v5 float32
	)
	v2 = float64(a2.X - *(*float32)(unsafe.Add(a1, 56)))
	if v2 < 0.0 {
		v2 = -v2
	}
	v5 = float32(v2)
	v3 = float64(a2.Y - *(*float32)(unsafe.Add(a1, 60)))
	if v3 < 0.0 {
		v3 = -v3
	}
	return bool2int32(float64(v5) >= 5.0 || v3 >= 5.0)
}
func nox_xxx_playerCancelSpells_4FEAE0(a1p *server.Object) int32 {
	var (
		a1     = a1p
		result int32
		v2     int32
	)
	result = int32(uintptr(nox_xxx_spellCastedFirst_4FE930()))
	if result == 0 {
		return result
	}
	for {
		v2 = int32(*(*uint32)(unsafe.Add(result, 116)))
		if *(*uint32)(unsafe.Add(result, 16)) == uint32(a1) {
			nox_xxx_spellCancelSpellDo_4FE9D0(result)
		}
		result = v2
		if v2 == 0 {
			break
		}
	}
	return result
}
func sub_4FEB60(a1 *server.Object, a2 *server.Object) {
	if a2.ObjClass&0x1000 != 0 {
		if a2.ObjSubClass&0x40000 != 0 {
			nox_xxx_spellCancelDurSpell_4FEB10(43, a1)
		}
		if a2.ObjSubClass&0x4000000 != 0 {
			nox_xxx_spellCancelDurSpell_4FEB10(59, a1)
		}
	}
}
func Nox_xxx_plrCastSmth_4FEDA0(sp *server.DurSpell) {
	if sp.Obj16 != nil {
		snd := nox_xxx_spellGetAud44_424800(int32(sp.Spell), 2)
		nox_xxx_aud_501960(snd, sp.Obj16, 0, 0)
	}
	if destroy := sp.Destroy.Get(); destroy != nil {
		destroy(sp)
	}
	if u := sp.Obj16; u != nil {
		if u.Class().Has(object.ClassPlayer) {
			ud := u.UpdateDataPlayer()
			if ud.Player.PlayerClass() != player.Warrior || nox_common_playerIsAbilityActive_4FC250(u, 1) == 0 {
				nox_xxx_playerSetState_4FA020(u, 13)
				sub_4FE900(unsafe.Pointer(sp))
				sub_4FE980(unsafe.Pointer(sp))
				return
			}
		} else if u.Class().Has(object.ClassMonster) {
			sub_541630(u, int32(sp.Spell))
		}
	}
	sub_4FE900(unsafe.Pointer(sp))
	sub_4FE980(unsafe.Pointer(sp))
}
func nox_xxx_cancelAllSpells_4FEE90(a1p *server.Object) {
	var (
		a1 = a1p
		v1 int32
		v2 int32
		v3 int32
	)
	v1 = int32(uintptr(nox_xxx_spellCastedFirst_4FE930()))
	v2 = v1
	if v1 != 0 {
		for {
			v3 = int32(uintptr(nox_xxx_spellCastedNext_4FE940(v2)))
			v1 = int32(*(*uint32)(unsafe.Add(v2, 16)))
			if v1 == a1 {
				v1 = int32(*(*uint32)(unsafe.Add(v2, 4)))
				if v1 == 24 || v1 == 43 || v1 == 35 || v1 == 8 || v1 == 22 || v1 == 59 || v1 == 67 {
					*(*uint8)(unsafe.Pointer(&v1)) = uint8(nox_xxx_spellCancelSpellDo_4FE9D0(v2))
				}
			}
			v2 = v3
			if v3 == 0 {
				break
			}
		}
	}
}
func nox_xxx_netStopRaySpell_4FEF90(a1 int32, a2 *server.Object) {
	var (
		v2  int32
		v3  int32
		v4  int8
		i   int32
		v6  int8
		v7  int8
		v8  int8
		v9  int16
		v10 int8
		v11 [7]byte
	)
	if a1 != 0 {
		v2 = int32(*(*uint32)(unsafe.Add(a1, 16)))
		if v2 != 0 {
			if a2 != nil {
				v3 = int32(*(*uint32)(unsafe.Add(a1, 4)))
				v11[0] = 158
				switch v3 {
				case 7:
					v6 = int8(*(*uint8)(unsafe.Add(a1, 8)))
					v11[1] = 10
					v11[2] = byte(v6)
					*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(a2))
					*(*uint16)(unsafe.Pointer(&v11[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*server.Object)(unsafe.Pointer(*(**uint32)(unsafe.Add(a1, 16))))))
				case 9:
					v4 = int8(*(*uint8)(unsafe.Add(a1, 8)))
					v11[1] = 9
					v11[2] = byte(v4)
					*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(a2))
					*(*uint16)(unsafe.Pointer(&v11[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*server.Object)(unsafe.Pointer(*(**uint32)(unsafe.Add(a1, 16))))))
				case 22:
					v8 = int8(*(*uint8)(unsafe.Add(a1, 8)))
					v11[1] = 12
					v11[2] = byte(v8)
					*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(a2))
					*(*uint16)(unsafe.Pointer(&v11[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*server.Object)(unsafe.Pointer(*(**uint32)(unsafe.Add(a1, 16))))))
				case 24:
					v7 = int8(*(*uint8)(unsafe.Add(a1, 8)))
					v11[1] = 11
					v11[2] = byte(v7)
					*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(a2))
					*(*uint16)(unsafe.Pointer(&v11[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*server.Object)(unsafe.Pointer(*(**uint32)(unsafe.Add(a1, 16))))))
				case 35:
					if uint32(v2) == *(*uint32)(unsafe.Add(a1, 48)) {
						return
					}
					v11[1] = 13
					*(*uint16)(unsafe.Pointer(&v11[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(a2))
					v9 = int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0((*server.Object)(unsafe.Pointer(*(**uint32)(unsafe.Add(a1, 16)))))))
					v10 = int8(*(*uint8)(unsafe.Add(a1, 8)))
					*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(v9)
					v11[2] = byte(v10)
				case 43:
					for i = int32(*(*uint32)(unsafe.Add(a1, 108))); i != 0; i = int32(*(*uint32)(unsafe.Add(i, 116))) {
						nox_xxx_netStopRaySpell_4FEF90(i, *(**uint32)(unsafe.Add(i, 48)))
					}
					return
				case 59:
					v11[1] = 8
					v11[2] = *(*uint8)(unsafe.Add(v2, 124))
					*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(a2))
					*(*uint16)(unsafe.Pointer(&v11[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0((*server.Object)(unsafe.Pointer(*(**uint32)(unsafe.Add(a1, 16))))))
				default:
					return
				}
				nox_xxx_netSendPacket1_4E5390(math.MaxUint8, unsafe.Pointer(&v11[0]), 7, nil, 1)
				nox_xxx_netUnmarkMinimapSpec_417470(*(*unsafe.Pointer)(unsafe.Add(a1, 16)), 2)
				nox_xxx_netUnmarkMinimapSpec_417470(a2, 2)
			}
		}
	}
}
func nox_xxx_netStartDurationRaySpell_4FF130(a1 unsafe.Pointer) {
	var v11 [7]byte
	v11[0] = 158
	switch *(*uint32)(unsafe.Add(a1, 4)) {
	case 7:
		v4 := int8(*(*uint8)(unsafe.Add(a1, 8)))
		v11[1] = 3
		v11[2] = byte(v4)
	case 9:
		v11[1] = 2
		v11[2] = *(*uint8)(unsafe.Add(a1, 8))
	case 0x16:
		v11[1] = 5
		v11[2] = *(*uint8)(unsafe.Add(a1, 8))
	case 0x18:
		v5 := int8(*(*uint8)(unsafe.Add(a1, 8)))
		v11[1] = 4
		v11[2] = byte(v5)
	case 0x23:
		if *(*uint32)(unsafe.Add(a1, 16)) != *(*uint32)(unsafe.Add(a1, 48)) {
			v10 := *(**server.Object)(unsafe.Add(a1, 48))
			v11[1] = 6
			*(*uint16)(unsafe.Pointer(&v11[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(v10))
			v8 := int16(uint16(nox_xxx_netGetUnitCodeServ_578AC0(*(**server.Object)(unsafe.Add(a1, 16)))))
			v9 := int8(*(*uint8)(unsafe.Add(a1, 8)))
			*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(v8)
			v11[2] = byte(v9)
			nox_xxx_netSendPacket1_4E5390(math.MaxUint8, unsafe.Pointer(&v11[0]), 7, nil, 1)
			nox_xxx_netMarkMinimapForAll_4174B0((*server.Object)(*(*unsafe.Pointer)(unsafe.Add(a1, 16))), 2)
			nox_xxx_netMarkMinimapForAll_4174B0((*server.Object)(*(*unsafe.Pointer)(unsafe.Add(a1, 48))), 2)
		}
		return
	case 0x2B:
		for i := *(*unsafe.Pointer)(unsafe.Add(a1, 108)); i != nil; i = *(*unsafe.Pointer)(unsafe.Add(i, 116)) {
			nox_xxx_netStartDurationRaySpell_4FF130(i)
		}
		return
	case 0x3B:
		v2 := *(*unsafe.Pointer)(unsafe.Add(a1, 16))
		v11[1] = 1
		v11[2] = *(*uint8)(unsafe.Add(v2, 124))
	default:
		return
	}
	if *(*uint32)(unsafe.Add(a1, 48)) != 0 {
		v6 := nox_xxx_netGetUnitCodeServ_578AC0(*(**server.Object)(unsafe.Add(a1, 48)))
		v7 := *(**server.Object)(unsafe.Add(a1, 16))
		*(*uint16)(unsafe.Pointer(&v11[5])) = uint16(v6)
		*(*uint16)(unsafe.Pointer(&v11[3])) = uint16(nox_xxx_netGetUnitCodeServ_578AC0(v7))
		nox_xxx_netSendPacket1_4E5390(math.MaxUint8, unsafe.Pointer(&v11[0]), 7, nil, 1)
		nox_xxx_netMarkMinimapForAll_4174B0((*server.Object)(*(*unsafe.Pointer)(unsafe.Add(a1, 16))), 2)
		nox_xxx_netMarkMinimapForAll_4174B0((*server.Object)(*(*unsafe.Pointer)(unsafe.Add(a1, 48))), 2)
	}
}
func sub_4FF2D0(a1 int32, a2 unsafe.Pointer) unsafe.Pointer {
	it := nox_xxx_spellCastedFirst_4FE930()
	if it == nil {
		return nil
	}
	for {
		if (int32(*(*uint8)(unsafe.Add(it, 88)))&1) == 0 && *(*uint32)(unsafe.Add(it, 4)) == uint32(a1) {
			v3 := *(*unsafe.Pointer)(unsafe.Add(it, 48))
			if v3 != nil {
				if v3 == a2 {
					break
				}
			}
		}
		it = nox_xxx_spellCastedNext_4FE940(it)
		if it == nil {
			return nil
		}
	}
	return it
}
func nox_xxx_testUnitBuffs_4FF350(obj *server.Object, buff server.EnchantID) int32 {
	return bool2int32(obj.HasEnchant(buff))
}
func nox_xxx_buffApplyTo_4FF380(unit *server.Object, buff server.EnchantID, dur int16, power int8) {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(unit)))
		v5 int32
		v6 int32
	)
	if *memmap.PtrUint32(0x5D4594, 1569740) == 0 {
		*memmap.PtrUint32(0x5D4594, 1569740) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("Hecubah")))
		*memmap.PtrUint32(0x5D4594, 1569744) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("Necromancer")))
	}
	if unit == nil {
		return
	}
	var v4w uint16 = *(*uint16)(unsafe.Add(a1, 4))
	if uint32(v4w) == *memmap.PtrUint32(0x5D4594, 1569740) && buff == 29 {
		return
	}
	if noxflags.HasGame(4096) && uint32(*(*uint16)(unsafe.Add(a1, 4))) == *memmap.PtrUint32(0x5D4594, 1569740) && buff == 3 {
		nox_xxx_aud_501960(582, unit, 0, 0)
		return
	}
	var v4 int32 = bool2int32(noxflags.HasGame(4096))
	if v4 != 0 && (func() bool {
		*(*uint16)(unsafe.Add(unsafe.Pointer(&v4), unsafe.Sizeof(uint16(0))*0)) = *memmap.PtrUint16(0x5D4594, 1569744)
		return uint32(*(*uint16)(unsafe.Add(a1, 4))) == *memmap.PtrUint32(0x5D4594, 1569744)
	}()) && buff == 3 {
		nox_xxx_aud_501960(595, unit, 0, 0)
	} else if int32(*(*uint8)(unsafe.Add(a1, 8)))&2 != 0 && (func() int32 {
		v4 = int32(*(*uint32)(unsafe.Add(a1, 12)))
		return int32(*(*uint8)(unsafe.Add(unsafe.Pointer(&v4), 1))) & 0x10
	}()) != 0 && buff == 11 && (func() int32 {
		v4 = bool2int32(noxflags.HasGame(2048))
		return v4
	}()) == 0 {
		v4 = int32(*(*uint16)(unsafe.Add(a1, 4)))
		if uint32(uint16(int16(v4))) == *memmap.PtrUint32(0x5D4594, 1569740) {
			nox_xxx_aud_501960(582, unit, 0, 0)
		} else if uint32(v4) == *memmap.PtrUint32(0x5D4594, 1569744) {
			nox_xxx_aud_501960(595, unit, 0, 0)
		}
	} else if (*(*uint32)(unsafe.Add(a1, 16)) & 0x8022) == 0 {
		if nox_xxx_testUnitBuffs_4FF350(unit, buff) == 0 || (func() int32 {
			v4 = nox_xxx_unitGetBuffTimer_4FF550(unit, buff)
			return v4
		}()) != 0 {
			if buff != 0 {
				nox_xxx_spellBuffOff_4FF5B0(unit, 0)
			}
			unit.BuffsDur[buff] = uint16(dur)
			unit.BuffsPower[buff] = uint8(power)
			nox_xxx_setUnitBuffFlags_4E48F0(unsafe.Pointer(unit), int32(uint32(1<<buff)|unit.Buffs))
			v5 = nox_xxx_getEnchantSpell_424920(int32(buff))
			v6 = nox_xxx_spellGetAud44_424800(v5, 1)
			nox_xxx_aud_501960(v6, unit, 0, 0)
		}
	}
}
func nox_xxx_unitGetBuffTimer_4FF550(unit *server.Object, buff server.EnchantID) int32 {
	return int32(unit.EnchantDur(buff))
}
func nox_xxx_buffGetPower_4FF570(unit *server.Object, buff server.EnchantID) int8 {
	return int8(unit.EnchantPower(buff))
}
func nox_xxx_unitClearBuffs_4FF580(unit *server.Object) {
	nox_xxx_setUnitBuffFlags_4E48F0(unsafe.Pointer(unit), 0)
	for i := int32(0); i < 32; i++ {
		unit.BuffsDur[i] = 0
		unit.BuffsPower[i] = 0
	}
}
func nox_xxx_spellBuffOff_4FF5B0(a1p *server.Object, a2 int32) int32 {
	var (
		a1     = unsafe.Pointer(a1p)
		result int32
		v3     int32
		v4     int32
		v5     int32
	)
	result = 1 << a2
	v3 = int32(*(*uint32)(unsafe.Add(a1, 340)))
	if v3&(1<<a2) == 0 {
		return result
	}
	nox_xxx_setUnitBuffFlags_4E48F0(a1, v3 & ^result)
	result = 0
	*(*uint16)(unsafe.Add(a1, a2*2+344)) = 0
	*(*uint8)(unsafe.Add(a1, a2+408)) = 0
	if !(a2 != 16 && a2 != 30) {
		return result
	}
	v4 = nox_xxx_getEnchantSpell_424920(a2)
	v5 = nox_xxx_spellGetAud44_424800(v4, 2)
	nox_xxx_aud_501960(v5, (*server.Object)(a1), 0, 0)
	return result
}
func nox_xxx_updateUnitBuffs_4FF620(a1p *server.Object) {
	var (
		a1 = unsafe.Pointer(a1p)
		v1 int32
		v2 uint16
		v3 int16
		v4 int32
	)
	if *(*uint32)(unsafe.Add(a1, 340)) != 0 {
		v1 = 0
		for {
			if uint32(1<<v1)&*(*uint32)(unsafe.Add(a1, 340)) != 0 {
				if v1 == 16 && uint32(*(*uint16)(unsafe.Add(a1, 376)))%gameFPS() == gameFPS()-1 {
					nox_xxx_aud_501960(26, (*server.Object)(a1), 0, 0)
				}
				v2 = *(*uint16)(unsafe.Add(a1, v1*2+344))
				if int32(v2) > 0 {
					v3 = int16(int32(v2) - 1)
					*(*uint16)(unsafe.Add(a1, v1*2+344)) = uint16(v3)
					if int32(v3) == 0 {
						if v1 == 7 {
							v4 = int32(*(*uint32)(unsafe.Add(a1, 16)))
							*(*uint8)(unsafe.Pointer(&v4)) = uint8(int8(v4 & 0xBF))
							*(*uint32)(unsafe.Add(a1, 16)) = uint32(v4)
						} else if v1 == 16 {
							*(*uint32)(unsafe.Add(a1, 520)) = 0
							*(*uint32)(unsafe.Add(a1, 524)) = 13
							nox_xxx_unitDamageClear_4EE5E0((*server.Object)(a1), 9999999)
							nox_xxx_aud_501960(779, (*server.Object)(a1), 0, 0)
							if int32(*(*uint8)(unsafe.Add(a1, 8)))&4 != 0 {
								nox_xxx_playerIncrementElimDeath_4D8D40(a1)
								nox_xxx_netReportLesson_4D8EF0((*server.Object)(a1))
							}
						}
						nox_xxx_spellBuffOff_4FF5B0((*server.Object)(a1), v1)
						*(*uint8)(unsafe.Add(a1, v1+408)) = 0
					}
				}
			}
			v1++
			if v1 >= 32 {
				break
			}
		}
		if nox_xxx_testUnitBuffs_4FF350((*server.Object)(a1), 9) != 0 {
			*(*float32)(unsafe.Add(a1, 544)) = float32(float64(*(*float32)(unsafe.Add(a1, 544))) * 1.25)
		}
	}
}

type magicWall struct {
	Field0  uint32       // 0, 0
	Field4  uint32       // 1, 4
	Field8  *server.Wall // 2, 8
	Field12 byte         // 3, 12
	Field13 byte         // 3, 13
	Field14 byte         // 3, 14
	Field15 byte         // 3, 15
	Field16 uint32       // 4, 16
	Field20 uint32       // 5, 20
	Field24 *magicWall   // 6, 24
	Field28 *magicWall   // 7, 28
}

func nox_xxx_allocMagicWallArray_4FF730() int32 {
	dword_5d4594_1569756 = 0
	nox_alloc_magicWall_1569748 = alloc.NewClassT("MagicWall", magicWall{}, int(int32((*memmap.PtrUint32(0x587000, 217844)<<6)+32)))
	return bool2int32(nox_alloc_magicWall_1569748.Class != nil)
}
func sub_4FF770() int32 {
	nox_alloc_magicWall_1569748.Free()
	dword_5d4594_1569752 = nil
	dword_5d4594_1569756 = 0
	return 0
}
func nox_xxx_mapWall_4FF790() {
	nox_alloc_magicWall_1569748.FreeAllObjects()
	dword_5d4594_1569752 = nil
}
func sub_4FF7B0(a1p *server.Player) {
	var (
		a1 = a1p
		v1 int8
		v2 int32
		v3 *uint32
		v4 *byte
		v6 int32
		v7 [6]byte
	)
	v1 = int8(a1.PlayerInd)
	v2 = 1 << int32(v1)
	if int32(v1) != 31 {
		v3 = (*uint32)(dword_5d4594_1569752)
		if dword_5d4594_1569752 != nil {
			for {
				if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 3680)))&0x10 != 0 {
					if (uint32(v2) & *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*4))) == 0 {
						v4 = (*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), 4*2)))
						v7[0] = 61
						v7[1] = *(*byte)(unsafe.Add(unsafe.Pointer(v4), 1))
						v7[2] = *v4
						v7[3] = *(*byte)(unsafe.Add(unsafe.Pointer(v4), 2))
						v7[4] = *(*byte)(unsafe.Add(unsafe.Pointer(v4), 5))
						v6 = int32(a1.PlayerInd)
						v7[5] = *(*byte)(unsafe.Add(unsafe.Pointer(v4), 6))
						nox_xxx_netSendPacket0_4E5420(v6, unsafe.Pointer(&v7[0]), 6, nil, 1)
						*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*4)) = uint32(v2) | *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*4))
					}
				}
				v3 = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), 4*6)))
				if v3 == nil {
					break
				}
			}
		}
	}
}
func nox_xxx_wallDestroyMagicwallSysuse_4FF840(a1 int32) int32 {
	var (
		result int32
		v2     int32
	)
	result = int32(dword_5d4594_1569752)
	if dword_5d4594_1569752 == nil {
		return result
	}
	for {
		v2 = int32(*(*uint32)(unsafe.Add(result, 24)))
		if *(*uint32)(unsafe.Add(result, 8)) == uint32(a1) {
			nox_xxx_wallDestroy_4FF870(result)
		}
		result = v2
		if v2 == 0 {
			break
		}
	}
	return result
}
func nox_xxx_wallDestroy_4FF870(a1 *magicWall) {
	sub_4FF900(a1)
	if noxflags.HasGame(1) {
		v1 := a1.Field8
		if a1.Field0 != 0 {
			v1.Tile1 = a1.Field12
			a1.Field8.Dir0 = a1.Field13
			a1.Field8.Field2 = a1.Field14
		} else {
			nox_xxx_mapDelWallAtPt_410430(int32(v1.X5), int32(v1.Y6))
		}
	}
	v2 := a1.Field28
	if v2 != nil {
		v2.Field24 = a1.Field24
	} else {
		dword_5d4594_1569752 = a1.Field24
	}
	v3 := a1.Field24
	if v3 != nil {
		v3.Field28 = a1.Field28
	}
	nox_alloc_magicWall_1569748.FreeObjectFirst(a1)
}
func sub_4FF900(a1 *magicWall) {
	var (
		i  uint32
		v4 int8
		v5 int8
		v8 [6]byte
	)
	v1 := a1
	for i = 0; i < 0x20; i++ {
		if (1<<i)&v1.Field16 != 0 {
			if v1.Field0 != 0 {
				v4 = int8(v1.Field12)
				v5 = int8(v1.Field13)
				v8[3] = v1.Field14
				v6 := v1.Field8
				v8[0] = 61
				v8[1] = byte(v4)
				v8[2] = byte(v5)
				v8[4] = v6.X5
				v8[5] = v6.Y6
				nox_xxx_netSendPacket0_4E5420(int32(i), unsafe.Pointer(&v8[0]), 6, nil, 1)
			} else {
				v7 := v1.Field8
				*(*uint8)(unsafe.Pointer(&a1)) = 62
				*(*uint16)(unsafe.Add(unsafe.Pointer(&a1), 1)) = *(*uint16)(unsafe.Add(unsafe.Pointer(v7), 5))
				nox_xxx_netSendPacket0_4E5420(int32(i), unsafe.Pointer(&a1), 3, nil, 1)
			}
		}
	}
}
func sub_4FF990(a1 int32) int32 {
	var result int32
	for result = int32(dword_5d4594_1569752); result != 0; result = int32(*(*uint32)(unsafe.Add(result, 24))) {
		*(*uint32)(unsafe.Add(result, 16)) &= uint32(^a1)
	}
	return result
}
func nox_xxx_spellWallCreateCalcDirMB_4FF9B0(a1 int32, a2 int32, a3 int32, a4 int32, a5 int32) int8 {
	var result int8
	result = sub_4FFA40(a1)
	switch a1 {
	case 1:
		if a5 != a3 {
			result = sub_4FFA40(7)
		}
	case 7:
		if a5 != a3 {
			result = sub_4FFA40(1)
		}
	case 3:
		if a4 != a2 {
			result = sub_4FFA40(5)
		}
	default:
		if a1 == 5 && a4 != a2 {
			result = sub_4FFA40(3)
		}
	}
	return result
}
func sub_4FFA40(a1 int32) int8 {
	var result int8
	switch a1 {
	case 1:
		result = 7
	case 2, 6:
		result = 1
	case 3:
		result = 8
	case 5:
		result = 10
	case 7:
		result = 9
	default:
		result = 0
	}
	return result
}
func Nox_xxx_spellWallCreate_4FFA90(sp *server.DurSpell) int32 {
	a1 := sp
	var (
		v1  int32
		v2  int32
		v3  float32
		v4  float32
		v5  float32
		v6  int32
		v7  int32
		v9  int32
		v10 int32
		v11 int32
		v13 int32
		v14 int32
		v15 int32
		v16 int32
		v17 uint8
		v18 uint8
		v19 uint8
		i   int32
		j   int32
		v22 int32
		v23 types.Pointf
		a1a float4
	)
	v1 = a1
	v22 = int32(gameFPS() * 60 * (a1.Level + 5))
	v2 = int32(a1.Obj16)
	if v2 == 0 || *(*uint32)(unsafe.Add(v2, 16))&0x8020 != 0 {
		return 1
	}
	v3 = a1.Pos.Y
	a1a.field_0 = a1.Pos.X
	v4 = a1.Pos2.X
	a1a.field_4 = v3
	v5 = a1.Pos2.Y
	a1a.field_8 = v4
	a1a.field_C = v5
	if int32(uint8(int8(nox_xxx_traceRay_5374B0(&a1a)))) == 0 {
		v6 = int32(*(*uint32)(unsafe.Add(v1, 16)))
		if int32(*(*uint8)(unsafe.Add(v6, 8)))&4 == 0 {
			return 1
		}
		v7 = int32(*(*uint32)(unsafe.Add(v6, 748)))
		a1 = 2
		nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v7, 276)), 2064))), 0, &a1)
		return 1
	}
	v23.X = a1a.field_8 - a1a.field_0
	v23.Y = a1a.field_C - a1a.field_4
	v9 = nox_xxx_math_509ED0(&v23)
	v10 = nox_xxx_math_509EA0(v9)
	v11 = int32(int64(*(*float32)(unsafe.Add(v1, 52))) / 23)
	v13 = int32(int64(*(*float32)(unsafe.Add(v1, 56))) / 23)
	if ((int32(uint8(int8(v11))) + int32(uint8(int8(v13)))) & 1) == 1 {
		v11++
	}
	v16 = v11
	a1 = int32(int64(*(*float32)(unsafe.Add(v1, 56))) / 23)
	v17 = uint8(sub_4FFA40(v10))
	if sub_4FFD00(v1, v11, v13, v17) != 0 {
		for i = 0; i < *memmap.PtrInt32(0x587000, uintptr(*(*uint32)(unsafe.Add(v1, 8)))*4+217844); i++ {
			v11 = int32(uint8(nox_xxx_spellWallCreateCalcXMB_4FFEF0(v10, v16, v11, 0)))
			v13 = int32(uint8(nox_xxx_spellWallCreateCalcYMB_4FFFB0(v10, a1, v13, 0)))
			v18 = uint8(nox_xxx_spellWallCreateCalcDirMB_4FF9B0(v10, v16, a1, v11, v13))
			if sub_4FFD00(v1, v11, v13, v18) == 0 {
				break
			}
		}
		v14 = v16
		v15 = a1
		for j = 0; j < *memmap.PtrInt32(0x587000, uintptr(*(*uint32)(unsafe.Add(v1, 8)))*4+217844); j++ {
			v14 = int32(uint8(nox_xxx_spellWallCreateCalcXMB_4FFEF0(v10, v16, v14, 1)))
			v15 = int32(uint8(nox_xxx_spellWallCreateCalcYMB_4FFFB0(v10, a1, v15, 1)))
			v19 = uint8(nox_xxx_spellWallCreateCalcDirMB_4FF9B0(v10, v16, a1, v14, v15))
			if sub_4FFD00(v1, v14, v15, v19) == 0 {
				break
			}
		}
	}
	*(*uint32)(unsafe.Add(v1, 68)) = uint32(v22) + gameFrame()
	return 0
}
func sub_4FFD00(a1 int32, a2 int32, a3 int32, a4 uint8) int32 {
	var (
		v4  int32
		v7  int8
		v9  *uint8
		v10 int32
		v11 int8
		v12 int8
		v13 int8
	)
	v4 = 0
	v13 = 0
	v12 = 0
	v11 = 0
	if int32(*memmap.PtrUint8(0x5D4594, 1570004)) == 0 {
		*memmap.PtrUint8(0x5D4594, 1570004) = uint8(int8(nox_xxx_wallTileByName_410D60(internCStr("MagicWallSystemUseOnly"))))
		*memmap.PtrUint8(0x5D4594, 1570005) = uint8(int8(nox_xxx_wallTileByName_410D60(internCStr("InvisibleWallSet"))))
		*memmap.PtrUint8(0x5D4594, 1570006) = uint8(int8(nox_xxx_wallTileByName_410D60(internCStr("InvisibleBlockingWallSet"))))
	}
	v5 := nox_server_getWallAtGrid_410580(a2, a3)
	v6 := v5
	if v5 != nil {
		v7 = int8(v5.Tile1)
		if int32(v7) == int32(*memmap.PtrUint8(0x5D4594, 1570004)) {
			return 0
		}
		if int32(v7) == int32(*memmap.PtrUint8(0x5D4594, 1570005)) || int32(v7) == int32(*memmap.PtrUint8(0x5D4594, 1570006)) {
			return 0
		}
		if int32(v6.Flags4)&0x1C != 0 {
			return 0
		}
		v13 = int8(v6.Tile1)
		v4 = 1
		v12 = int8(*v6)
		v11 = int8(v6.Field2)
		*v6 = nox_xxx_wall_42A6C0(*v6, a4)
	} else {
		v9 = (*uint8)(nox_xxx_wallCreateAt_410250(a2, a3))
		v6 = v9
		if v9 == nil {
			return 0
		}
		*(*uint8)(unsafe.Add(unsafe.Pointer(v9), 1)) = *memmap.PtrUint8(0x5D4594, 1570004)
		*v9 = a4
		if int32(a4) == 0 || int32(a4) == 1 {
			*(*uint8)(unsafe.Add(unsafe.Pointer(v9), 2)) = uint8(int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v9), 5))) % int32(int16(nox_xxx_map_410E00(int32(*memmap.PtrUint8(0x5D4594, 1570004)))))))
			goto LABEL_12
		}
	}
	v6.Field2 = 0
	if v4 != 0 {
		nox_xxx_netWallCreate_4FFE80(a1, v6, v4, v13, v12, v11)
		return bool2int32(v4 == 0)
	}
LABEL_12:
	v10 = int32(v6.Tile1)
	v6.Flags4 |= 8
	v6.Health7 = nox_xxx_mapWallGetHpByTile_410E20(v10)
	nox_xxx_netWallCreate_4FFE80(a1, v6, v4, v13, v12, v11)
	return bool2int32(v4 == 0)
}
func nox_xxx_netWallCreate_4FFE80(a1 int32, a2 *server.Wall, a3 int32, a4 int8, a5 int8, a6 int8) *magicWall {
	p := nox_alloc_magicWall_1569748.NewObject()
	if p == nil {
		return nil
	}
	p.Field8 = a2
	p.Field0 = uint32(a3)
	*(*uint8)(unsafe.Pointer(&p.Field4)) = a2.Dir0
	p.Field12 = uint8(a4)
	p.Field13 = uint8(a5)
	p.Field14 = uint8(a6)
	p.Field20 = uint32(a1)
	p.Field16 = 0
	p.Field28 = nil
	p.Field24 = dword_5d4594_1569752
	if dword_5d4594_1569752 != nil {
		dword_5d4594_1569752.Field28 = p
	}
	dword_5d4594_1569752 = p
	return p
}
func nox_xxx_spellWallCreateCalcXMB_4FFEF0(a1 int32, a2 int32, a3 int32, a4 int8) int8 {
	var v5 int8
	if int32(a4) == 0 {
		switch a1 {
		case 0:
			return int8(a3 + 1)
		case 1:
			return int8(a3 + 1)
		case 2:
			return int8(a3 + 1)
		case 3:
			v5 = int8(a3)
			if a3 != a2 {
				return int8(int32(v5) - 1)
			}
			return int8(int32(v5) + 1)
		case 5:
			v5 = int8(a3)
			if a3 == a2 {
				return int8(int32(v5) - 1)
			}
			return int8(int32(v5) + 1)
		case 6:
			return int8(a3 - 1)
		case 7:
			return int8(a3 - 1)
		case 8:
			return int8(a3 - 1)
		default:
			return int8(a3)
		}
	}
	switch a1 {
	case 0:
		return int8(a3 - 1)
	case 1:
		return int8(a3 - 1)
	case 2:
		return int8(a3 - 1)
	case 3:
		v5 = int8(a3)
		if a3 != a2 {
			return int8(int32(v5) - 1)
		}
		return int8(int32(v5) + 1)
	case 5:
		v5 = int8(a3)
		if a3 == a2 {
			return int8(int32(v5) - 1)
		}
		return int8(int32(v5) + 1)
	case 6:
		return int8(a3 + 1)
	case 7:
		return int8(a3 + 1)
	case 8:
		return int8(a3 + 1)
	default:
		return int8(a3)
	}
}
func nox_xxx_spellWallCreateCalcYMB_4FFFB0(a1 int32, a2 int32, a3 int32, a4 int8) int8 {
	var v4 int8
	if int32(a4) == 0 {
		switch a1 {
		case 0:
			v4 = int8(a3)
			return int8(int32(v4) - 1)
		case 1:
			v4 = int8(a3)
			if a3 != a2 {
				return int8(int32(v4) + 1)
			}
			return int8(a3 - 1)
		case 2:
			return int8(a3 + 1)
		case 3:
			return int8(a3 - 1)
		case 5:
			return int8(a3 + 1)
		case 6:
			return int8(a3 - 1)
		case 7:
			v4 = int8(a3)
			if a3 != a2 {
				return int8(int32(v4) - 1)
			}
			return int8(int32(v4) + 1)
		case 8:
			return int8(a3 + 1)
		default:
			return int8(a3)
		}
	}
	switch a1 {
	case 0:
		return int8(a3 + 1)
	case 1:
		v4 = int8(a3)
		if a3 != a2 {
			return int8(int32(v4) + 1)
		}
		return int8(a3 - 1)
	case 2:
		return int8(a3 - 1)
	case 3:
		return int8(a3 + 1)
	case 5:
		return int8(a3 - 1)
	case 6:
		return int8(a3 + 1)
	case 7:
		v4 = int8(a3)
		if a3 != a2 {
			return int8(int32(v4) - 1)
		}
		return int8(int32(v4) + 1)
	case 8:
		v4 = int8(a3)
		return int8(int32(v4) - 1)
	default:
		return int8(a3)
	}
}
func Nox_xxx_spellWallUpdate_500070(sp *server.DurSpell) int32 {
	return 0
}
func Nox_xxx_spellWallDestroy_500080(sp *server.DurSpell) int32 {
	a1 := sp
	var (
		result int32
		v2     int32
	)
	result = int32(dword_5d4594_1569752)
	if dword_5d4594_1569752 == nil {
		return result
	}
	for {
		v2 = int32(*(*uint32)(unsafe.Add(result, 24)))
		if *(*uint32)(unsafe.Add(result, 20)) == uint32(a1) {
			nox_xxx_wallDestroy_4FF870(result)
		}
		result = v2
		if v2 == 0 {
			break
		}
	}
	return result
}
func sub_5000B0(a1 unsafe.Pointer) int32 {
	var (
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 int32
		v14 int32
		v15 int32
	)
	v6 = 1
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v6)), 2)
	if int32(int16(v6)) > 1 || int32(int16(v6)) <= 0 {
		return 0
	}
	v1 := dword_5d4594_1569752
	for *(*uint8)(unsafe.Pointer(&v5)) = 0; v1 != nil; v1 = v1.Field24 {
		*(*uint8)(unsafe.Pointer(&v5)) = uint8(int8(v5 + 1))
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v5)), 1)
	if nox_crypt_IsReadOnly() != 0 {
		if int32(uint8(int8(v5))) == 0 {
			return 1
		}
		sub_5002D0((*server.Object)(a1))
		v4 = 0
		dword_5d4594_1569756 = 0
		if int32(uint8(int8(v5))) == 0 {
			return 1
		}
		for {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v13)), 4)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v15)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v14)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v12)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v11)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v10)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v9)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v8)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v7)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&a1)), 1)
			sub_500330(int8(v15), int8(v14), v13, int8(v12), int8(v11), int8(v10), int8(v9), int8(v8), int8(v7), int8(uintptr(a1)))
			v4++
			if v4 >= int32(uint8(int8(v5))) {
				break
			}
		}
		return 1
	}
	v2 := dword_5d4594_1569752
	if dword_5d4594_1569752 == nil {
		return 1
	}
	for {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(v2), 4)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v2.Field8.X5, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v2.Field8.Y6, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v2.Field12, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v2.Field13, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v2.Field14, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v2.Field8.Tile1, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v2.Field8.Field2, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v2.Field8, 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v2.Field8.Health7, 1)
		v2 = v2.Field24
		if v2 == nil {
			break
		}
	}
	return 1
}
func sub_500330(a1 int8, a2 int8, a3 int32, a4 int8, a5 int8, a6 int8, a7 int8, a8 int8, a9 int8, a10 int8) {
	var (
		v11    int32
		result int32
	)
	if dword_5d4594_1569756 < 15 {
		result = int32(dword_5d4594_1569756 * 16)
		v11 = int32(dword_5d4594_1569756 + 1)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569764) = uint8(a1)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569765) = uint8(a2)
		*memmap.PtrUint32(0x5D4594, uintptr(result)+1569768) = uint32(a3)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569772) = uint8(a4)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569773) = uint8(a5)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569774) = uint8(a6)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569775) = uint8(a7)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569776) = uint8(a8)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569777) = uint8(a9)
		*memmap.PtrUint8(0x5D4594, uintptr(result)+1569778) = uint8(a10)
		dword_5d4594_1569756 = uint32(v11)
	}
}
func nox_xxx_mapSetWallInGlobalDir0pr1_5004D0() int32 {
	var result int32
	for result = int32(dword_5d4594_1569752); result != 0; result = int32(*(*uint32)(unsafe.Add(result, 24))) {
		if *(*uint32)(result) == 1 {
			**(**uint8)(unsafe.Add(result, 8)) = *(*uint8)(unsafe.Add(result, 13))
		}
	}
	return result
}
func nox_xxx_map_5004F0() int32 {
	var result int32
	for result = int32(dword_5d4594_1569752); result != 0; result = int32(*(*uint32)(unsafe.Add(result, 24))) {
		if *(*uint32)(result) == 1 {
			**(**uint8)(unsafe.Add(result, 8)) = *(*uint8)(unsafe.Add(result, 4))
		}
	}
	return result
}
func nox_xxx_journalQuestSet_500540(a1 *byte, a2 int32) *byte {
	var (
		result *byte
		v3     *byte
	)
	result = nox_xxx_scriptGetJournal_5005E0(a1)
	if result != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*34)) = uint32(a2)
	} else {
		result = (*byte)(alloc.Calloc1(1, 0x94))
		v3 = result
		if result != nil {
			libc.StrCpy(result, (*byte)(memmap.PtrOff(0x5D4594, 1570140)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*33)) = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*34)) = uint32(a2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*36)) = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*35)) = dword_5d4594_1570272
			result = dword_5d4594_1570272
			if dword_5d4594_1570272 != 0 {
				*(*uint32)(unsafe.Add(dword_5d4594_1570272, 144)) = uint32(uintptr(unsafe.Pointer(v3)))
			}
			dword_5d4594_1570272 = uint32(uintptr(unsafe.Pointer(v3)))
		}
	}
	return result
}
func nox_xxx_scriptGetJournal_5005E0(a1 *byte) *byte {
	var (
		v1 uint32
		v2 *uint8
		v3 *uint8
		v4 *byte
		v5 int8
		v6 uint32
		v7 int8
		i  int32
	)
	if libc.StrChr(a1, 58) != nil {
		v6 = uint32(libc.StrLen(a1) + 1)
		v7 = int8(uint8(v6))
		v6 >>= 2
		alloc.Memcpy(memmap.PtrOff(0x5D4594, 1570140), unsafe.Pointer(a1), uintptr(v6*4))
		v4 = (*byte)(unsafe.Add(unsafe.Pointer(a1), v6*4))
		v3 = (*uint8)(memmap.PtrOff(0x5D4594, uintptr(v6)*4+1570140))
		v5 = v7
	} else {
		libc.StrCpy((*byte)(memmap.PtrOff(0x5D4594, 1570140)), (*byte)(memmap.PtrOff(0x5D4594, 1570008)))
		*memmap.PtrUint16(0x5D4594, uintptr(libc.StrLen((*byte)(memmap.PtrOff(0x5D4594, 1570140)))+1570140)) = *memmap.PtrUint16(0x587000, 217952)
		v1 = uint32(libc.StrLen(a1) + 1)
		v2 = (*uint8)(memmap.PtrOff(0x5D4594, uintptr(libc.StrLen((*byte)(memmap.PtrOff(0x5D4594, 1570140)))+1570140)))
		alloc.Memcpy(unsafe.Pointer(v2), unsafe.Pointer(a1), uintptr((v1>>2)*4))
		v4 = (*byte)(unsafe.Add(unsafe.Pointer(a1), (v1>>2)*4))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), (v1>>2)*4))
		v5 = int8(uint8(v1))
	}
	alloc.Memcpy(unsafe.Pointer(v3), unsafe.Pointer(v4), uintptr(int32(v5)&3))
	for i = int32(dword_5d4594_1570272); i != 0; i = int32(*(*uint32)(unsafe.Add(i, 140))) {
		if nox_strcmpi((*byte)(i), (*byte)(memmap.PtrOff(0x5D4594, 1570140))) == 0 {
			break
		}
	}
	return (*byte)(i)
}
func nox_xxx_journalQuestSetBool_5006B0(a1 *byte, a2 int32) *byte {
	var (
		result *byte
		v3     *byte
	)
	result = nox_xxx_scriptGetJournal_5005E0(a1)
	if result != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*34)) = uint32(a2)
	} else {
		result = (*byte)(alloc.Calloc1(1, 0x94))
		v3 = result
		if result != nil {
			libc.StrCpy(result, (*byte)(memmap.PtrOff(0x5D4594, 1570140)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*33)) = 1
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*34)) = uint32(a2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*36)) = 0
			*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*35)) = dword_5d4594_1570272
			result = dword_5d4594_1570272
			if dword_5d4594_1570272 != 0 {
				*(*uint32)(unsafe.Add(dword_5d4594_1570272, 144)) = uint32(uintptr(unsafe.Pointer(v3)))
			}
			dword_5d4594_1570272 = uint32(uintptr(unsafe.Pointer(v3)))
		}
	}
	return result
}
func sub_500750(a1 *byte) int32 {
	var (
		v1     *byte
		result int32
	)
	v1 = nox_xxx_scriptGetJournal_5005E0(a1)
	if v1 != nil {
		result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*34)))
	} else {
		result = 0
	}
	return result
}
func sub_500770(a1 *byte) float64 {
	var (
		v1     *byte
		result float64
	)
	v1 = nox_xxx_scriptGetJournal_5005E0(a1)
	if v1 != nil {
		result = float64(*(*float32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(float32(0))*34)))
	} else {
		result = 0.0
	}
	return result
}
func sub_500790(lpMem unsafe.Pointer) {
	var (
		v1 int32
		v2 int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(lpMem, 4*36)))
	if v1 != 0 {
		*(*uint32)(unsafe.Add(v1, 140)) = *(*uint32)(unsafe.Add(lpMem, 4*35))
	}
	v2 = int32(*(*uint32)(unsafe.Add(lpMem, 4*35)))
	if v2 != 0 {
		*(*uint32)(unsafe.Add(v2, 144)) = *(*uint32)(unsafe.Add(lpMem, 4*36))
	}
	if lpMem == dword_5d4594_1570272 {
		dword_5d4594_1570272 = *(*uint32)(unsafe.Add(lpMem, 4*35))
	}
	alloc.FreePtr(lpMem)
}
func sub_5007E0(a1 *byte) *byte {
	var (
		v1     *uint8
		result *byte
		v3     uint32
		v4     *byte
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     *byte
		v10    *byte
		v11    *byte
		v12    int32
		v13    uint32
		v14    int32
		v15    int32
		v16    uint32
		v17    int32
	)
	sub_5009B0(a1)
	v1 = libc.StrChr((*byte)(memmap.PtrOff(0x5D4594, 1570140)), 42)
	if v1 != nil {
		v3 = uint32(libc.StrLen((*byte)(memmap.PtrOff(0x5D4594, 1570140))) + 1)
		result = nil
		if libc.StrCmp((*byte)(memmap.PtrOff(0x5D4594, 1570140)), internCStr("*:*")) == 0 {
			result = dword_5d4594_1570272
			if dword_5d4594_1570272 != 0 {
				for {
					v4 = (*byte)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*35)))))
					sub_500790(unsafe.Pointer(result))
					result = v4
					if v4 == nil {
						break
					}
				}
			}
		} else if unsafe.Pointer(v1) == memmap.PtrOff(0x5D4594, uintptr(v3)+1570138) {
			v5 = int32(dword_5d4594_1570272)
			if dword_5d4594_1570272 != 0 {
				for {
					v6 = int32(*(*uint32)(unsafe.Add(v5, 140)))
					result = (*byte)(unsafe.Pointer(uintptr(nox_strnicmp((*byte)(v5), (*byte)(memmap.PtrOff(0x5D4594, 1570140)), int32(v3-2)))))
					if result == nil {
						sub_500790(v5)
					}
					v5 = v6
					if v6 == 0 {
						break
					}
				}
			}
		} else if unsafe.Pointer(v1) == memmap.PtrOff(0x5D4594, 1570140) {
			v7 = int32(dword_5d4594_1570272)
			if dword_5d4594_1570272 != 0 {
				for {
					v8 = int32(*(*uint32)(unsafe.Add(v7, 140)))
					result = libc.StrStr((*byte)(v7), (*byte)(memmap.PtrOff(0x5D4594, 1570141)))
					if result != nil {
						v9 = result
						result = nil
						if v3-2 == uint32(libc.StrLen(v9)) {
							sub_500790(v7)
						}
					}
					v7 = v8
					if v8 == 0 {
						break
					}
				}
			}
		} else {
			v10 = libc.StrChr((*byte)(memmap.PtrOff(0x5D4594, 1570140)), 58)
			result = nil
			v11 = (*byte)(unsafe.Add(unsafe.Pointer(v10), 2))
			v12 = int32(dword_5d4594_1570272)
			v13 = uint32(libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(v10), 2))) + 1)
			if dword_5d4594_1570272 != 0 {
				v14 = int32(uintptr(unsafe.Pointer(v10)) - uintptr(memmap.PtrOff(0x5D4594, 1570140)))
				v17 = int32(uintptr(unsafe.Pointer(v10)) - uintptr(memmap.PtrOff(0x5D4594, 1570140)))
				for {
					v15 = int32(*(*uint32)(unsafe.Add(v12, 140)))
					result = (*byte)(unsafe.Pointer(uintptr(nox_strnicmp((*byte)(v12), (*byte)(memmap.PtrOff(0x5D4594, 1570140)), v14+1))))
					if result == nil {
						result = libc.StrStr((*byte)(unsafe.Add(unsafe.Pointer(uintptr(v14+v12)), 2)), v11)
						if result != nil {
							v16 = uint32(libc.StrLen(result) + 1)
							result = (*byte)(unsafe.Pointer(uintptr(v13 - 1)))
							if v13-1 == v16-1 {
								sub_500790(v12)
							}
							v14 = v17
						}
					}
					v12 = v15
					if v15 == 0 {
						break
					}
				}
			}
		}
	} else {
		result = nox_xxx_scriptGetJournal_5005E0(a1)
		if result != nil {
			sub_500790(unsafe.Pointer(result))
		}
	}
	return result
}
func sub_5009B0(a1 *byte) uint32 {
	var (
		v1     uint32
		v2     int8
		v3     *uint8
		v4     *uint8
		result uint32
	)
	if libc.StrChr(a1, 58) != nil {
		result = uint32(libc.StrLen(a1) + 1)
		alloc.Memcpy(memmap.PtrOff(0x5D4594, 1570140), unsafe.Pointer(a1), uintptr(result))
	} else {
		v1 = uint32(libc.StrLen((*byte)(memmap.PtrOff(0x5D4594, 1570008))) + 1)
		v2 = int8(uint8(v1))
		v1 >>= 2
		alloc.Memcpy(memmap.PtrOff(0x5D4594, 1570140), memmap.PtrOff(0x5D4594, 1570008), uintptr(v1*4))
		v4 = (*uint8)(memmap.PtrOff(0x5D4594, uintptr(v1)*4+1570008))
		v3 = (*uint8)(memmap.PtrOff(0x5D4594, uintptr(v1)*4+1570140))
		*(*uint8)(unsafe.Pointer(&v1)) = uint8(v2)
		result = 0
		alloc.Memcpy(unsafe.Pointer(v3), unsafe.Pointer(v4), uintptr(v1&3))
		*memmap.PtrUint16(0x5D4594, uintptr(libc.StrLen((*byte)(memmap.PtrOff(0x5D4594, 1570140)))+1570140)) = *memmap.PtrUint16(0x587000, 217960)
		libc.StrCat((*byte)(memmap.PtrOff(0x5D4594, 1570140)), a1)
	}
	return result
}
func sub_500A60() int32 {
	var (
		result int32
		v1     int32
		j      int32
		v3     int32
		i      int32
		v5     int32
		v6     int32
	)
	v5 = 1
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v5)), 2)
	if int32(int16(v5)) > 1 {
		return 0
	}
	v1 = int32(dword_5d4594_1570272)
	for i = 0; v1 != 0; i++ {
		v1 = int32(*(*uint32)(unsafe.Add(v1, 140)))
	}
	if noxflags.HasGame(2048) {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&i)), 4)
		for j = int32(dword_5d4594_1570272); j != 0; j = int32(*(*uint32)(unsafe.Add(j, 140))) {
			*(*uint8)(unsafe.Pointer(&v6)) = uint8(int8(libc.StrLen((*byte)(j))))
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v6)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(j), uint32(uint8(int8(v6))))
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(j, 132)), 4)
			v3 = int32(*(*uint32)(unsafe.Add(j, 132)))
			if v3 != 0 {
				if v3 == 1 {
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(j, 136)), 4)
				}
			} else {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(j, 136)), 4)
			}
		}
		result = 1
	} else {
		i = 0
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&i)), 4)
		result = 1
	}
	return result
}
func sub_500B70() int32 {
	var (
		i  uint32
		v2 int32
		v3 int32
		v4 uint32
		v5 int32
		v6 int32
		v7 int32
		v8 [256]byte
	)
	sub_5007E0(internCStr("*:*"))
	v3 = 1
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v3)), 2)
	if int32(int16(v3)) > 1 {
		return 0
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v4)), 4)
	for i = 0; i < v4; i++ {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v2)), 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v8[0], uint32(uint8(int8(v2))))
		v8[uint8(int8(v2))] = 0
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v5)), 4)
		if v5 != 0 {
			if v5 == 1 {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v7)), 4)
				nox_xxx_journalQuestSetBool_5006B0(&v8[0], v7)
			}
		} else {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v6)), 4)
			nox_xxx_journalQuestSet_500540(&v8[0], v6)
		}
	}
	return 1
}
func nox_xxx_orderUnitLocal_500C70(owner int32, orderType int32) int32 {
	nox_common_playerInfoFromNum_417090(owner).SummonOrderAll = uint32(orderType)
	return nox_xxx_netCreatureCmd_4D7EE0(owner, int8(orderType))
}
func sub_500CA0(a1 int32, a2 unsafe.Pointer) int32 {
	var result int32
	if a2 != nil && int32(*(*uint8)(unsafe.Add(a2, 8)))&4 != 0 {
		result = int32(*memmap.PtrUint32(0x587000, uintptr(a1*4)+217668))
	} else {
		result = 0
	}
	return result
}
func nox_xxx_creatureIsMonitored_500CC0(a1p *server.Object, a2p *server.Object) int32 {
	var (
		a1     = unsafe.Pointer(a1p)
		a2     = unsafe.Pointer(a2p)
		v2     int32
		result int32
	)
	if (int32(*(*uint8)(unsafe.Add(a2, 8)))&2 != 0 && (func() bool {
		v2 = int32(*(*uint32)(unsafe.Add(a2, 16)))
		return (v2 & 0x8000) == 0
	}()) || nox_xxx_unitIsZombie_534A40((*server.Object)(a2)) != 0) && (int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a2, 748)), 1440)))&0x80) != 0 {
		result = nox_xxx_unitHasThatParent_4EC4F0((*server.Object)(a2), (*server.Object)(a1))
	} else {
		result = 0
	}
	return result
}
func Nox_xxx_summonStart_500DA0(sp *server.DurSpell) int32 {
	a1 := sp
	var (
		v1  int32
		v2  int32
		v3  int32
		v5  uint8
		v6  *byte
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 float64
		v12 int32
		v13 int32
		v14 int8
		v15 float32
		v16 int16
		v17 [16]uint8
	)
	v1 = int32(a1.Obj16)
	v2 = int32(a1.Spell - 74)
	if v1 == 0 || *(*uint32)(unsafe.Add(v1, 16))&0x8020 != 0 || a1.Flag20 != 0 {
		return 1
	}
	if int32(*(*uint8)(unsafe.Add(v1, 8)))&4 != 0 {
		v3 = int32(*(*uint32)(unsafe.Add(v1, 748)))
		if noxflags.HasGame(4608) && *(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(v3, 276))+uint32(v2*4))), 4244)) == 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0(a1.Obj16, internCStr("Summon.c:NeedGuideToSummon"), 0)
			return 1
		}
		v5 = uint8(int8(bool2int32(nox_xxx_checkSummonedCreaturesLimit_500D70(a1.Obj16, v2))))
		if int32(v5) == 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0(a1.Obj16, internCStr("Summon.c:CreatureControlFailed"), 0)
			return 1
		}
	}
	if sub_500F40(a1, COERCE_FLOAT(uint32(uintptr(unsafe.Pointer(&v17[2]))))) == 0 {
		return 1
	}
	v17[10] = *(*uint8)(unsafe.Add(unsafe.Pointer(a1.Obj16), 124))
	v6 = nox_xxx_guideNameByN_427230(v2)
	*(*uint16)(unsafe.Pointer(&v17[0])) = uint16(int16(nox_xxx_getNameId_4E3AA0(v6)))
	*(*uint16)(unsafe.Pointer(&v17[11])) = func() uint16 {
		p := mem_getU16Ptr(0x5D4594, 1570276)
		x := *p
		*p++
		return x
	}()
	if int32(*memmap.PtrUint16(0x5D4594, 1570276)) >= 0xFDE8 {
		*memmap.PtrUint16(0x5D4594, 1570276) = 0
	}
	v7 = int32(*(*uint32)(unsafe.Pointer(&v17[4])))
	v17[13] = 0
	a1.Field72 = int32(*(*uint32)(unsafe.Pointer(&v17[0])))
	v8 = int32(*(*uint32)(unsafe.Pointer(&v17[8])))
	a1.Field76 = uint32(v7)
	*(*uint16)(unsafe.Add(unsafe.Pointer(&v7), unsafe.Sizeof(uint16(0))*0)) = *(*uint16)(unsafe.Pointer(&v17[12]))
	a1.Field80 = uint32(v8)
	*(*uint16)(unsafe.Add(unsafe.Pointer(a1), 84)) = uint16(int16(v7))
	v9 = int32(nox_xxx_guideGetUnitSize_427460(v2)) - 1
	if v9 != 0 {
		v10 = v9 - 1
		if v10 != 0 {
			if v10 != 2 {
				v12 = a1
				goto LABEL_22
			}
			v11 = nox_xxx_gamedataGetFloatTable_419D70(internCStr("SummonDuration"), 2)
		} else {
			v11 = nox_xxx_gamedataGetFloatTable_419D70(internCStr("SummonDuration"), 1)
		}
	} else {
		v11 = nox_xxx_gamedataGetFloatTable_419D70(internCStr("SummonDuration"), 0)
	}
	v15 = float32(v11)
	v12 = int32(v15)
LABEL_22:
	v13 = int32(uint32(v12) + gameFrame())
	v16 = int16(v12)
	v14 = int8(v17[10])
	a1.Frame68 = uint32(v13)
	nox_xxx_sendSummonStartFX_5236F0(*(*int16)(unsafe.Pointer(&v17[11])), (*float32)(unsafe.Pointer(&v17[2])), v14, *(*int16)(unsafe.Pointer(&v17[0])), v16)
	return 0
}
func sub_500F40(a1 int32, a2 float32) int32 {
	var (
		v2     *uint32
		v3     int32
		v4     float32
		v5     float32
		v6     float32
		v7     float64
		v8     float64
		v9     float64
		v10    float32
		result int32
		v12    int32
		v13    int32
		v14    float32
		v15    float32
		v16    float32
		v17    int32
		v18    int32
		v19    float4
	)
	v2 = (*uint32)(a1)
	if float64(a1) == 0.0 {
		return 0
	}
	v3 = int32(*(*uint32)(unsafe.Add(a1, 16)))
	if v3 == 0 {
		return 0
	}
	v4 = a2
	if float64(a2) == 0.0 {
		return 0
	}
	if int32(*(*uint8)(unsafe.Add(v3, 8)))&4 != 0 {
		v19.field_0 = *(*float32)(unsafe.Add(v3, 56))
		v5 = *(*float32)(unsafe.Add(v3, 60))
		v6 = *(*float32)(unsafe.Add(a1, 56))
		v19.field_8 = *(*float32)(unsafe.Add(a1, 52))
		v19.field_C = v6
		v7 = float64(v19.field_8 - v19.field_0)
		v19.field_4 = v5
		v8 = float64(v6 - v5)
		a1 = int32(float32(v8))
		v9 = math.Sqrt(v8*float64(a1) + v7*v7)
		a2 = float32(v9)
		if v9 > 50.0 {
			v19.field_8 = float32(v7*50.0/float64(a2) + float64(v19.field_0))
			v19.field_C = float32(float64(a1)*50.0/float64(a2) + float64(v19.field_4))
		}
		if nox_xxx_mapTraceRay_535250(&v19, nil, nil, 9) != 0 && nox_xxx_mapTileAllowTeleport_411A90((*types.Pointf)(unsafe.Pointer(&v19.field_8))) == 0 {
			v10 = v19.field_C
			*(*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&v4), 4*0))) = *(*uint32)(unsafe.Add(unsafe.Pointer(&v19.field_8), 4*0))
			*(*float32)(unsafe.Add(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer(&v4), 4*0))))), 4)) = v10
			return 1
		}
		v12 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*4)))
		if int32(*(*uint8)(unsafe.Add(v12, 8)))&4 == 0 {
			return 0
		}
		v13 = int32(*(*uint32)(unsafe.Add(v12, 748)))
		a1 = 2
		nox_xxx_netInformTextMsg_4DA0F0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v13, 276)), 2064))), 0, &a1)
		return 0
	}
	v19.field_0 = *(*float32)(unsafe.Add(v3, 56))
	v14 = *(*float32)(unsafe.Add(v3, 60))
	v15 = *(*float32)(unsafe.Add(a1, 52))
	v16 = *(*float32)(unsafe.Add(a1, 56))
	v19.field_4 = v14
	v19.field_8 = v15
	v19.field_C = v16
	if nox_xxx_mapTraceRay_535250(&v19, nil, nil, 9) != 0 {
		*(*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&v4), 4*0))) = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*13))
		*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer(&v4), 4*0))))), 4)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*14))
		result = 1
	} else {
		v17 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*4)))
		*(*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(&v4), 4*0))) = *(*uint32)(unsafe.Add(v17, 56))
		v18 = int32(*(*uint32)(unsafe.Add(v17, 60)))
		result = 1
		*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr((*(*uint32)(unsafe.Add(unsafe.Pointer(&v4), 4*0))))), 4)) = uint32(v18)
	}
	return result
}
func Nox_xxx_summonFinish_5010D0(sp *server.DurSpell) int32 {
	a1 := sp
	var (
		v1 int32
		v2 int32
		v3 int32
		v5 uint8
		v6 *uint32
	)
	v1 = int32(a1.Obj16)
	if v1 == 0 || *(*uint32)(unsafe.Add(v1, 16))&0x8020 != 0 {
		return 1
	}
	if a1.Frame68-1 != gameFrame() {
		return 0
	}
	v2 = int32(a1.Spell - 74)
	if (int32(*(*uint8)(unsafe.Add(v1, 8))) & 4) == 0 {
		goto LABEL_17
	}
	v3 = int32(*(*uint32)(unsafe.Add(v1, 748)))
	if noxflags.HasGame(512) && *(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(v3, 276))+uint32(v2*4))), 4244)) == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0(a1.Obj16, internCStr("Summon.c:NeedGuideToSummon"), 0)
		return 1
	}
	v5 = uint8(int8(bool2int32(nox_xxx_checkSummonedCreaturesLimit_500D70(a1.Obj16, v2))))
	if int32(v5) == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0(a1.Obj16, internCStr("Summon.c:CreatureControlFailed"), 0)
		return 1
	}
LABEL_17:
	v6 = (*uint32)(unsafe.Pointer(nox_xxx_unitDoSummonAt_5016C0(int32(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), 72))), (*float32)(unsafe.Add(unsafe.Pointer(a1), 74)), a1.Obj16, *(*uint8)(unsafe.Add(unsafe.Pointer(a1), 82)))))
	if v6 != nil {
		nox_xxx_aud_501960(899, (*server.Object)(unsafe.Pointer(v6)), 0, 0)
	}
	*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 85)) = 1
	return 1
}
func Nox_xxx_summonCancel_5011C0(sp *server.DurSpell) int32 {
	a1 := sp
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 85))) == 0 {
		nox_xxx_sendSummonCancelFX_523760(int16(*(*uint16)(unsafe.Add(unsafe.Pointer(a1), 83))))
		nox_xxx_audCreate_501A30(900, (*types.Pointf)(unsafe.Add(unsafe.Pointer(a1), 74)), 0, 0)
	}
	return 0
}
func Nox_xxx_charmCreature1_5011F0(sp *server.DurSpell) int32 {
	a1 := sp
	var (
		v1  int16
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  float64
		v10 *int32
		v11 *server.Object
		v12 int32
		v13 int32
		v14 float32
		v15 int32
		v16 float32
	)
	if a1.Flag20 != 0 {
		v14 = float32(nox_xxx_gamedataGetFloat_419D40(internCStr("ConfuseEnchantDuration")))
		v1 = int16(int32(v14))
		nox_xxx_buffApplyTo_4FF380(a1.Obj48, 3, v1, int8(a1.Level))
		sub_4E7540(a1.Obj16, a1.Obj48)
		return 1
	}
	v15 = a1.Obj16
	v3 = int32(nox_xxx_spellFlags_424A70(int32(a1.Spell)))
	v4 = int32(uintptr(unsafe.Pointer(nox_xxx_spellFlySearchTarget_540610((*types.Pointf)(unsafe.Add(unsafe.Pointer(a1), 4*13)), a1.Obj16, v3, 300.0, 0, (*server.Object)(v15)))))
	a1.Obj48 = v4
	if v4 == 0 {
		nox_xxx_netPriMsgToPlayer_4DA2C0(a1.Obj16, internCStr("Summon.c:CharmNoCreatureCloseEnough"), 0)
		nox_xxx_aud_501960(16, a1.Obj16, 0, 0)
		return 1
	}
	if int32(*(*uint8)(unsafe.Add(v4, 8)))&2 != 0 && nox_xxx_creatureIsMonitored_500CC0(a1.Obj16, (*server.Object)(v4)) == 0 {
		v5 = nox_xxx_creatureIsCharmableByTT_4272B0(int32(a1.Obj48.TypeInd))
		v6 = a1.Obj16
		if int32(*(*uint8)(unsafe.Add(v6, 8)))&4 != 0 && nox_cheat_charmall == 0 {
			if v5 == 0 {
				nox_xxx_netPriMsgToPlayer_4DA2C0((*server.Object)(v6), internCStr("Summon.c:CreatureNotCharmable"), 0)
				v12 = a1.Obj16
				a1.Obj48 = nil
				nox_xxx_aud_501960(16, (*server.Object)(v12), 0, 0)
				return 1
			}
			if *(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v6, 748)), 276))+uint32(v5*4))), 4244)) == 0 {
				a1.Obj48 = nil
				nox_xxx_netPriMsgToPlayer_4DA2C0((*server.Object)(v6), internCStr("Summon.c:NeedGuideToCharm"), 0)
				nox_xxx_aud_501960(16, a1.Obj16, 0, 0)
				return 1
			}
		}
		v7 = int32(nox_xxx_guideGetUnitSize_427460(v5)) - 1
		if v7 <= 0 {
			v9 = nox_xxx_gamedataGetFloatTable_419D70(internCStr("CharmSmallDuration"), int32(a1.Level-1))
		} else if v7 == 1 {
			v9 = nox_xxx_gamedataGetFloatTable_419D70(internCStr("CharmMediumDuration"), int32(a1.Level-1))
		} else {
			v8 = v7 - 1
			if v8 != 2 {
				v10 = a1
				goto LABEL_20
			}
			v9 = nox_xxx_gamedataGetFloatTable_419D70(internCStr("CharmLargeDuration"), int32(a1.Level-1))
		}
		v16 = float32(v9)
		v10 = (*int32)(unsafe.Pointer(uintptr(int32(v16))))
	LABEL_20:
		v11 = a1.Obj48
		a1.Frame68 = uint32(int32(uintptr(unsafe.Pointer(v10)))) + gameFrame()
		nox_xxx_buffApplyTo_4FF380(v11, 28, int16(int32(uint16(uintptr(unsafe.Pointer(v10))))+1), 5)
		nox_xxx_netStartDurationRaySpell_4FF130(unsafe.Pointer(a1))
		return 0
	}
	v13 = a1.Obj16
	a1.Obj48 = nil
	nox_xxx_aud_501960(16, (*server.Object)(v13), 0, 0)
	return 1
}
func Nox_xxx_charmCreatureFinish_5013E0(sp *server.DurSpell) int32 {
	a1 := sp
	var (
		v1  int32
		v2  int32
		v4  int32
		v5  int32
		v6  int32
		v7  uint8
		v8  int32
		v9  int32
		v10 int32
		v11 int32
		v12 int32
		v13 uint16
		v14 int32
		v15 uint16
		v16 int32
		v17 int32
		v18 int32
		v19 int32
		v20 int32
		v21 int32
		v22 int32
		v23 int32
	)
	v1 = a1.Obj48
	if v1 == 0 || *(*uint32)(unsafe.Add(v1, 16))&0x8020 != 0 {
		nox_xxx_aud_501960(16, a1.Obj16, 0, 0)
		return 1
	}
	if nox_xxx_calcDistance_4E6C00(a1.Obj16, a1.Obj48) > 300.0 {
		v2 = a1.Obj16
		if int32(*(*uint8)(unsafe.Add(v2, 8)))&4 != 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*server.Object)(v2), internCStr("Summon.c:CharmBrokenDistance"), 0)
		}
		nox_xxx_aud_501960(16, a1.Obj16, 0, 0)
		return 1
	}
	if uint32(a1.Frame68-1) != gameFrame() {
		return 0
	}
	v4 = a1.Obj48
	v5 = int32(*(*uint32)(unsafe.Add(v4, 12)))
	if nox_cheat_charmall == 0 {
		if v5&0x2000 != 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0(a1.Obj16, internCStr("Summon.c:CreatureControlImpossible"), 0)
			nox_xxx_aud_501960(16, a1.Obj16, 0, 0)
			return 1
		}
		if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1.Obj16), 8)))&4 != 0 {
			v6 = nox_xxx_creatureIsCharmableByTT_4272B0(int32(*(*uint16)(unsafe.Add(v4, 4))))
			v7 = uint8(int8(bool2int32(nox_xxx_checkSummonedCreaturesLimit_500D70(a1.Obj16, v6))))
			if int32(v7) == 0 {
				nox_xxx_netPriMsgToPlayer_4DA2C0(a1.Obj16, internCStr("Summon.c:CreatureControlFailed"), 0)
				nox_xxx_aud_501960(16, a1.Obj16, 0, 0)
				return 1
			}
		}
	}
	nox_xxx_spellBuffOff_4FF5B0(a1.Obj48, 28)
	v8 = nox_xxx_findParentChainPlayer_4EC580(a1.Obj48)
	v9 = a1.Obj16
	if v8 == v9 {
		if int32(*(*uint8)(unsafe.Add(v9, 8)))&4 != 0 {
			nox_xxx_netPriMsgToPlayer_4DA2C0((*server.Object)(v9), internCStr("Summon.c:CreatureAlreadyOwned"), 0)
		}
		nox_xxx_aud_501960(16, a1.Obj16, 0, 0)
		return 1
	}
	nox_xxx_unitClearOwner_4EC300(a1.Obj48)
	nox_xxx_unitSetOwner_4EC290(a1.Obj16, a1.Obj48)
	if nox_xxx_servObjectHasTeam_419130((*server.ObjectTeam)(unsafe.Add(unsafe.Pointer(a1.Obj48), 48))) != 0 {
		nox_xxx_netChangeTeamMb_419570(unsafe.Add(unsafe.Pointer(a1.Obj48), 48), int32(a1.Obj48.NetCode))
	}
	v10 = int32(a1.Obj48.UpdateData)
	v11 = int32(*(*uint32)(unsafe.Add(v10, 1440)))
	*(*uint8)(unsafe.Pointer(&v11)) = uint8(int8(v11 | 0x80))
	*(*uint32)(unsafe.Add(v10, 1440)) = uint32(v11)
	if noxflags.HasGame(4096) {
		v12 = int32(uintptr(nox_xxx_objectTypeByIndHealthData(int32(a1.Obj48.TypeInd))))
		v13 = uint16(nox_xxx_unitGetHP_4EE780(a1.Obj48))
		v14 = int32(*(*uint32)(unsafe.Add(v10, 484)))
		if v14 != 0 {
			v15 = *(*uint16)(unsafe.Add(v14, 72))
		} else {
			v15 = *(*uint16)(unsafe.Add(v12, 4))
		}
		a1.Obj48.HealthData.Max = v15
		if int32(v13) > int32(v15) {
			nox_xxx_unitSetHP_4E4560(a1.Obj48, v15)
		}
	}
	v16 = a1.Obj16
	if int32(*(*uint8)(unsafe.Add(v16, 8)))&4 != 0 {
		v17 = int32(*(*uint32)(unsafe.Add(v16, 748)))
		nox_xxx_orderUnit_533900((*server.Object)(v16), a1.Obj48, int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v17, 276)), 3648))))
		v18 = a1.Obj48
		v19 = int32(*(*uint32)(unsafe.Add(v18, 12)))
		*(*uint8)(unsafe.Pointer(&v19)) = uint8(int8(v19 | 0x80))
		*(*uint32)(unsafe.Add(v18, 12)) = uint32(v19)
		v20 = a1.Obj48
		v21 = int32(*(*uint32)(unsafe.Add(v20, 12)))
		*(*uint8)(unsafe.Add(unsafe.Pointer(&v21), 1)) |= 1
		*(*uint32)(unsafe.Add(v20, 12)) = uint32(v21)
		nox_xxx_netReportAcquireCreature_4D91A0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v17, 276)), 2064))), a1.Obj48)
		nox_xxx_netMarkMinimapObject_417190(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v17, 276)), 2064))), a1.Obj48, 1)
		nox_xxx_netSendSimpleObject2_4DF360(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v17, 276)), 2064))), a1.Obj48)
		if noxflags.HasGame(4096) {
			sub_50E140(unsafe.Pointer(a1.Obj48))
		}
	} else {
		nox_xxx_orderUnit_533900((*server.Object)(v16), a1.Obj48, 4)
	}
	v23 = a1.Obj16
	v22 = nox_xxx_spellGetAud44_424800(int32(a1.Spell), 1)
	nox_xxx_aud_501960(v22, (*server.Object)(v23), 0, 0)
	return 1
}
func Nox_xxx_charmCreature2_501690(sp *server.DurSpell) int32 {
	a1 := sp
	var result int32
	result = int32(a1.Obj48)
	if result == 0 {
		return result
	}
	if (*(*uint32)(unsafe.Add(result, 16)) & 0x8020) == 0 {
		nox_xxx_spellBuffOff_4FF5B0((*server.Object)(result), 5)
		result = nox_xxx_spellBuffOff_4FF5B0(a1.Obj48, 28)
	}
	return result
}
func nox_xxx_banishUnit_5017F0(unit unsafe.Pointer) {
	var (
		v1            int32
		v2            int32
		unitEventData int32
	)
	if *memmap.PtrUint32(0x5D4594, 1570280) == 0 {
		*memmap.PtrUint32(0x5D4594, 1570280) = uint32(nox_xxx_getNameId_4E3AA0(internCStr("Glyph")))
	}
	if unit != nil {
		v1 = int32(*(*uint32)(unsafe.Add(unit, 504)))
		if v1 != 0 {
			for {
				v2 = int32(*(*uint32)(unsafe.Add(v1, 496)))
				if uint32(*(*uint16)(unsafe.Add(v1, 4))) == *memmap.PtrUint32(0x5D4594, 1570280) {
					nox_xxx_delayedDeleteObject_4E5CC0((*server.Object)(v1))
				}
				v1 = v2
				if v2 == 0 {
					break
				}
			}
		}
		nox_xxx_netSendPointFx_522FF0(-127, (*types.Pointf)(unsafe.Add(unit, 56)))
		unitEventData = int32(*(*uint32)(unsafe.Add(unit, 748)))
		nox_xxx_scriptCallByEventBlock_502490(unsafe.Add(unitEventData, 1264), nil, unit, 7)
		nox_xxx_delayedDeleteObject_4E5CC0((*server.Object)(unit))
	}
}
func sub_501C00(a1 *float32, a2p *server.Object) int8 {
	var (
		a2 = a2p
		v2 int8
		v3 int32
		v4 *byte
		v5 int32
		v6 float32
		v7 *nox_player_polygon_check_data
		v9 Point32
	)
	v2 = 0
	if a2 != nil {
		v3 = int32(a2.ObjClass)
		if v3&4 != 0 {
			v2 = int8(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a2.UpdateData, 276)), 3668)))
			if int32(v2) != 0 {
				return v2
			}
			goto LABEL_8
		}
		if v3&2 != 0 {
			v4 = nox_xxx_polygonGetByIdx_4214A0(int32(*a2.UpdateData))
			if v4 != nil {
				v2 = int8(*(*byte)(unsafe.Add(unsafe.Pointer(v4), 130)))
				if int32(v2) != 0 {
					return v2
				}
				goto LABEL_8
			}
		}
	}
LABEL_8:
	v5 = int32(*a1)
	v6 = *(*float32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(float32(0))*1))
	v9.X = v5
	v9.Y = int32(v6)
	v7 = nox_xxx_polygonIsPlayerInPolygon_4217B0(&v9, 0)
	if v7 != nil {
		return int8(*(*uint8)(unsafe.Add(unsafe.Pointer(&v7.field_0[32]), 2)))
	}
	return v2
}
func nox_xxx_netUpdateRemotePlr_501CA0(a1p *server.Object) {
	var (
		a1  = a1p
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  *nox_player_polygon_check_data
		v17 Point32
		v18 int8
	)
	v2 = int32(a1.UpdateData)
	v3 = int32(*(*uint32)(unsafe.Add(v2, 276)))
	if (int32(*(*uint8)(unsafe.Add(v3, 3680)))&3) == 0 || (func() int32 {
		v4 = int32(*(*uint32)(unsafe.Add(v3, 3628)))
		return v4
	}()) == 0 {
		if *(*int32)(unsafe.Add(v3, 3664)) == -559023410 {
			nox_xxx_questCheckSecretArea_421C70(a1)
		}
		v5 = int32(*(*uint32)(unsafe.Add(v2, 276)))
		v18 = int8(*(*uint8)(unsafe.Add(v5, 3668)))
	} else if int32(*(*uint8)(unsafe.Add(v4, 8)))&4 != 0 {
		v5 = int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v4, 748)), 276)))
		v18 = int8(*(*uint8)(unsafe.Add(v5, 3668)))
	} else {
		v6 = int32(*(*float32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 3628)), 56)))
		v7 = int32(*(*uint32)(unsafe.Add(v2, 276)))
		v17.X = v6
		v17.Y = int32(*(*float32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v7, 3628)), 60)))
		v8 = nox_xxx_polygonIsPlayerInPolygon_4217B0(&v17, 0)
		if v8 != nil {
			v18 = int8(*(*uint8)(unsafe.Add(unsafe.Pointer(&v8.field_0[32]), 2)))
		} else {
			v18 = 0
		}
	}
	nox_xxx_netUpdateRemotePlr_501CA0_B(a1, v2, v18)
}
func nox_server_scriptExecuteFnForEachGroupObj_502670(groupPtr *uint8, expectedType int32, a3 func(unsafe.Pointer, unsafe.Pointer), a4 unsafe.Pointer) {
	if groupPtr == nil {
		return
	}
	switch *groupPtr {
	case 0:
		if expectedType != 0 {
			break
		}
		for i := (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(groupPtr), 4*21))))); i != nil; i = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(i), 4*2))) {
			v5 := nox_xxx_netGetUnitByExtent_4ED020(*i)
			if v5 != nil {
				a3(unsafe.Pointer(v5), a4)
			}
		}
	case 1:
		if expectedType != 1 {
			break
		}
		for j := (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(groupPtr), 4*21))))); j != nil; j = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(j), 4*2))) {
			v7 := (*uint32)(unsafe.Pointer(nox_server_getWaypointById_579C40(*j)))
			if v7 != nil {
				a3(unsafe.Pointer(v7), a4)
			}
		}
	case 2:
		if expectedType != 2 {
			break
		}
		for k := (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(groupPtr), 4*21))))); k != nil; k = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(k), 4*2))) {
			v9 := nox_server_getWallAtGrid_410580(*k, *(*int32)(unsafe.Add(unsafe.Pointer(k), 4*1)))
			if v9 != nil {
				a3(unsafe.Pointer(v9), a4)
			}
		}
		fallthrough
	case 3:
		for l := (*int32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(groupPtr), 4*21))))); l != nil; l = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(l), 4*2))) {
			v11 := (*uint8)(nox_server_scriptGetGroup_57C0A0(*l))
			if v11 != nil {
				nox_server_scriptExecuteFnForEachGroupObj_502670(v11, expectedType, a3, a4)
			}
		}
	default:
	}
}
func nox_xxx_mapgenMakeScript_502790(a1 *FILE, a2 *byte) int32 {
	var (
		result int32
		i      int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     int32
		v9     int32
		v10    [1024]byte
	)
	nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&v8)), 4, 1, a1)
	nox_binfile_fread_408E40(&v10[0], 1, v8, a1)
	nox_binfile_fread_408E40(a2, 4, 1, a1)
	nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&v7)), 4, 1, a1)
	result = v7
	for i = 0; i < v7; i++ {
		nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&v6)), 1, 1, a1)
		nox_binfile_fseek_409050(a1, 1, stdio.SEEK_CUR)
		v4 = 0
		v5 = int32(uint8(int8(v6))) * 268
		if int32(*memmap.PtrUint8(0x587000, uintptr(v5)+218640)) != 0 {
			for {
				switch *memmap.PtrUint32(0x587000, uintptr(v4)*8+218648+uintptr(v5)) {
				case 0, 3, 4, 5, 6:
					nox_binfile_fseek_409050(a1, 4, stdio.SEEK_CUR)
				case 1:
					nox_binfile_fseek_409050(a1, 8, stdio.SEEK_CUR)
				case 2, 7:
					nox_binfile_fread_408E40((*byte)(unsafe.Pointer(&v9)), 1, 1, a1)
					nox_binfile_fseek_409050(a1, int32(uint8(int8(v9))), stdio.SEEK_CUR)
				default:
				}
				v4++
				v5 = int32(uint8(int8(v6))) * 268
				if v4 >= int32(*memmap.PtrUint8(0x587000, uintptr(v5)+218640)) {
					break
				}
			}
		}
		result = v7
	}
	return result
}
func sub_5029A0(a1 *byte) int32 {
	var (
		v1 int32
		i  int32
	)
	v1 = 0
	if dword_5d4594_1599596 <= 0 {
		return -1
	}
	for i = 0; nox_strcmpi(a1, (*byte)(unsafe.Pointer(uintptr(uint32(i)+dword_5d4594_1599576)))) != 0; i += 76 {
		if func() int32 {
			p := &v1
			*p++
			return *p
		}() >= dword_5d4594_1599596 {
			return -1
		}
	}
	return v1
}
func sub_5029F0(a1 int32) int32 {
	var result int32
	if a1 < 0 || a1 > dword_5d4594_1599596 {
		result = 0
	} else {
		result = int32(dword_5d4594_1599576 + uint32(a1*76))
	}
	return result
}
func sub_502A20() int32 {
	return int32(dword_5d4594_1599596)
}
func sub_502A50(a1 *byte) int32 {
	var result int32
	sub_502DF0()
	if a1 != nil {
		libc.StrNCpy((*byte)(dword_5d4594_1599588), a1, 0x7FF)
		result = 1
	} else {
		*dword_5d4594_1599588 = *memmap.PtrUint8(0x5D4594, 1599608)
		result = 0
	}
	return result
}
func sub_502AB0(a1 *byte) int32 {
	var result int32
	if a1 != nil {
		libc.StrNCpy((*byte)(dword_5d4594_1599592), a1, 0x7FF)
		result = 1
	} else {
		*dword_5d4594_1599592 = *memmap.PtrUint8(0x5D4594, 1599612)
		result = 0
	}
	return result
}
func sub_502B10() int32 {
	var (
		result int32
		v1     int32
		v2     int32
		v3     int32
		v4     int8
		v5     int8
		v6     int32
		v7     int32
		v8     int32
		v9     float32
		v10    float32
		v11    [64]byte
	)
	dword_5d4594_1599596 = 0
	if dword_5d4594_1599588 == nil {
		dword_5d4594_1599588 = alloc.Calloc1(1, 0x800)
	}
	if dword_5d4594_1599592 == nil {
		dword_5d4594_1599592 = alloc.Calloc1(1, 0x800)
	}
	if dword_5d4594_1599576 == 0 {
		dword_5d4594_1599576 = uint32(uintptr(alloc.Calloc1(1, 0x26000)))
	}
	result = 0
	if libc.StrLen((*byte)(dword_5d4594_1599588)) == 0 {
		return result
	}
	result = int32(uintptr(unsafe.Pointer(sub_502DA0((*byte)(dword_5d4594_1599588)))))
	if result == 0 {
		return result
	}
	nox_fs_fread(nox_file_8, unsafe.Pointer(&v8), 4)
	if v8 == -889266515 {
		for {
			v6 = 0
			nox_fs_fread(nox_file_8, unsafe.Pointer(&v6), 4)
			v1 = v6
			if v6 == 0 {
				break
			}
			if dword_5d4594_1599596 >= 2048 {
				sub_502DF0()
				return 0
			}
			*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+dword_5d4594_1599596*76)), 72)) = uint32(nox_fs_ftell(nox_file_8) - 4)
			nox_fs_fread(nox_file_8, unsafe.Pointer(&v7), 1)
			nox_fs_fread(nox_file_8, unsafe.Pointer(&v11[0]), int32(uint8(int8(v7))))
			v2 = -1 - int32(uint8(int8(v7)))
			v11[uint8(int8(v7))] = 0
			v3 = v2 + v1
			libc.StrCpy((*byte)(unsafe.Pointer(uintptr(dword_5d4594_1599576+dword_5d4594_1599596*76))), &v11[0])
			nox_fs_fread(nox_file_8, unsafe.Pointer(&v4), 1)
			nox_fs_fread(nox_file_8, unsafe.Pointer(&v5), 1)
			nox_fs_fread(nox_file_8, unsafe.Pointer(&v9), 4)
			nox_fs_fread(nox_file_8, unsafe.Pointer(&v10), 4)
			*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+dword_5d4594_1599596*76)), 64)) = v9
			*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+func() uint32 {
				p := &dword_5d4594_1599596
				x := *p
				*p++
				return x
			}()*76)), 68)) = v10
			nox_fs_fseek(nox_file_8, v3-10, stdio.SEEK_CUR)
		}
		sub_502DF0()
		return 1
	} else {
		sub_502DF0()
		return 0
	}
	return result
}
func sub_502D70(a1 int32) int32 {
	if a1 < 0 || a1 >= dword_5d4594_1599596 {
		return 0
	}
	dword_5d4594_3835396 = uint32(a1)
	return bool2int32(nox_xxx_mapgenSaveMap_503830(a1) == 0)
}
func sub_502DA0(a1 *byte) *FILE {
	var result *FILE
	result = nox_file_8
	if nox_file_8 != nil || (func() *FILE {
		result = (*FILE)(unsafe.Pointer(uintptr(nox_xxx_cryptOpen_426910(a1, 1, -1))))
		return result
	}()) != nil && (func() bool {
		result = nox_xxx_mapgenGetSomeFile_426A60()
		return (func() *FILE {
			nox_file_8 = result
			return nox_file_8
		}()) != nil
	}()) {
		nox_fs_fseek(result, 0, stdio.SEEK_SET)
		result = (*FILE)(unsafe.Pointer(uintptr(1)))
	}
	return result
}
func sub_502DF0() *FILE {
	var result *FILE
	result = nox_file_8
	if nox_file_8 != nil {
		nox_xxx_cryptClose_4269F0()
		nox_file_8 = nil
	}
	return result
}
func sub_502E10(a1 int32) *FILE {
	if nox_file_8 == nil || a1 < 0 || a1 >= dword_5d4594_1599596 {
		return nil
	}
	nox_fs_fseek(nox_file_8, int32(*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+uint32(a1*76))), 72))), stdio.SEEK_SET)
	return nox_file_8
}
func sub_502E70(a1 int32) float64 {
	var result float64
	if a1 < 0 || a1 >= dword_5d4594_1599596 {
		result = -1.0
	} else {
		result = float64(*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+uint32(a1*76))), 64)))
	}
	return result
}
func sub_502EA0(a1 int32) float64 {
	var result float64
	if a1 < 0 || a1 >= dword_5d4594_1599596 {
		result = -1.0
	} else {
		result = float64(*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+uint32(a1*76))), 68)))
	}
	return result
}
func nox_xxx_mapgenSaveMap_503830(a1 int32) int32 {
	var (
		v1  *FILE
		v2  *uint32
		v3  int32
		v5  uint8
		v6  int32
		v7  int8
		v8  int32
		v9  int32
		v10 int32
		v11 [8]int32
		v19 [2]int32
		v21 int32
		v22 int32
		v23 int32
		v24 [4]byte
		v25 int4
		v26 [64]byte
		v27 [256]byte
	)
	if a1 < 0 {
		return 0
	}
	if a1 >= dword_5d4594_1599596 {
		return 0
	}
	nox_xxx_free_503F40()
	*memmap.PtrUint32(0x5D4594, 1599572) = math.MaxUint32
	dword_5d4594_1599644 = 0
	sub_502DA0((*byte)(dword_5d4594_1599588))
	if sub_502E10(a1) == nil {
		return 0
	}
	v1 = nox_xxx_mapgenGetSomeFile_426A60()
	nox_fs_fread(v1, unsafe.Pointer(&v22), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v9), 1)
	nox_fs_fread(v1, unsafe.Pointer(&v26[0]), int32(uint8(int8(v9))))
	nox_fs_fread(v1, unsafe.Pointer(&v7), 1)
	nox_fs_fread(v1, unsafe.Pointer(&v5), 1)
	nox_fs_fread(v1, unsafe.Pointer(&v21), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v23), 4)
	if int32(v5) > 1 {
		nox_fs_fread(v1, unsafe.Pointer(&v6), 4)
		nox_fs_fseek(v1, v6, stdio.SEEK_CUR)
	}
	nox_fs_fread(v1, unsafe.Pointer(&v10), 4)
	if v10 != -889266515 {
		return 0
	}
	nox_fs_fread(v1, unsafe.Pointer(&v19[0]), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v19[1]), 4)
	nox_xxx_mapWall_426A80(&v19[0])
	nox_fs_fread(v1, unsafe.Pointer(&v11[0]), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v11[1]), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v11[6]), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v11[7]), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v11[2]), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v11[3]), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v11[4]), 4)
	nox_fs_fread(v1, unsafe.Pointer(&v11[5]), 4)
	sub_4D3C80((*uint32)(unsafe.Pointer(&v11[0])))
	alloc.Memcpy(memmap.PtrOff(0x5D4594, 1599500), unsafe.Pointer(&v11[0]), 0x20)
	sub_428170(unsafe.Pointer(&v11[0]), &v25)
	nox_xxx_cryptSetTypeMB_426A50(1)
	for {
		v6 = 0
		*(*uint8)(unsafe.Pointer(&v8)) = 0
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v8)), 1)
		if int32(uint8(int8(v8))) == 0 {
			break
		}
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v27[0], uint32(uint8(int8(v8))))
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v24[0], 4)
		if nox_xxx_mapReadSection_426EA0(unsafe.Pointer(&v11[0]), &v27[0], (*uint32)(unsafe.Pointer(&v6))) == 0 {
			if v6 == 1 {
				sub_502DF0()
				return 0
			}
			v2 = nox_xxx_newObjectByTypeID_4E3810(&v27[0])
			v3 = int32(uintptr(unsafe.Pointer(v2)))
			if v2 == nil {
				return 0
			}
			if (ccall.AsFunc[func(*uint32, *int4) int32](*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), 4*176))))(v2, &v25) == 0 {
				nox_xxx_objectFreeMem_4E38A0((*server.Object)(v3))
				sub_502DF0()
				return 0
			}
			nox_xxx_servMapLoadPlaceObj_4F3F50((*server.Object)(v3), nil, unsafe.Pointer(&v25))
		}
	}
	nox_xxx_cryptSetTypeMB_426A50(0)
	dword_5d4594_1599480 = uint32(a1)
	dword_5d4594_1599476 = 0
	dword_5d4594_3835396 = uint32(a1)
	sub_502DF0()
	return 1
}
func sub_503B30(a1 *types.Pointf) int32 {
	var (
		result int32
		v2     int32
		v3     float64
		v4     float32
		v5     *byte
		v6     *byte
		v7     int32
		v8     int32
		v9     int32
		v13    types.Pointf
		v14    types.Pointf
		a2     types.Pointf
		v16    Point32
		v17    int4
		v18    [8]int32
	)
	result = nox_xxx_mapGenFixCoords_4D3D90(a1, &a2)
	if result == 0 {
		return result
	}
	v2 = int32(dword_5d4594_3835396)
	if dword_5d4594_1599480 != dword_5d4594_3835396 || dword_5d4594_1599480 == -1 || dword_5d4594_1599476 == 1 {
		result = nox_xxx_mapgenSaveMap_503830(int32(dword_5d4594_3835396))
		if result == 0 {
			return result
		}
		v2 = int32(dword_5d4594_3835396)
	}
	v18[2] = int32(int64(a2.X))
	v18[3] = int32(int64(a2.Y))
	v13.X = *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+uint32(v2*76))), 64)) + a1.X
	v13.Y = *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+uint32(v2*76))), 68)) + a1.Y
	nox_xxx_mapGenFixCoords_4D3D90(&v13, &v14)
	v18[4] = int32(int64(v14.X))
	v18[5] = int32(int64(v14.Y))
	v3 = float64(*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+dword_5d4594_3835396*76)), 64)) + a1.X)
	v13.Y = a1.Y
	v13.X = float32(v3)
	nox_xxx_mapGenFixCoords_4D3D90(&v13, &v14)
	v18[0] = int32(int64(v14.X))
	v4 = a1.X
	v18[1] = int32(int64(v14.Y))
	v13.X = v4
	v13.Y = *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(dword_5d4594_1599576+dword_5d4594_3835396*76)), 68)) + a1.Y
	nox_xxx_mapGenFixCoords_4D3D90(&v13, &v14)
	v18[6] = int32(int64(v14.X))
	v18[7] = int32(int64(v14.Y))
	sub_4D3C80((*uint32)(unsafe.Pointer(&v18[0])))
	sub_428170(unsafe.Pointer(&v18[0]), &v17)
	v5 = nox_xxx_mapGetWallSize_426A70()
	v6 = v5
	v7 = int32(*(*uint32)(unsafe.Pointer(v5)))
	*memmap.PtrUint32(0x5D4594, 1599484) = uint32(v7)
	*memmap.PtrUint32(0x5D4594, 1599488) = *(*uint32)(unsafe.Add(unsafe.Pointer(v6), 4*1))
	*mem_getFloatPtr(0x5D4594, 1599492) = float32(float64(v7 * 23))
	*mem_getFloatPtr(0x5D4594, 1599496) = float32(float64(int32(*memmap.PtrUint32(0x5D4594, 1599488) * 23)))
	v8 = int32(int64(float64(a2.X) - float64(*memmap.PtrInt32(0x5D4594, 1599508))))
	v9 = int32(int64(float64(a2.Y) - float64(*memmap.PtrInt32(0x5D4594, 1599512))))
	result = nox_xxx_tileInit_504150(v8, v9)
	if result == 0 {
		return result
	}
	result = sub_504330(v8, v9)
	if result == 0 {
		return result
	}
	result = sub_504560(v8, v9)
	if result == 0 {
		return result
	}
	result = sub_504910(v8, v9)
	if result == 0 {
		return result
	}
	sub_579D20()
	for i := sub_579890(); i != nil; i = sub_5798A0(i) {
		i.Flags |= 0x80000000
	}
	dword_5d4594_3835392 = uint32(nox_xxx_interesting_xfer_4D0010((*uint32)(unsafe.Pointer(&v17)), int32(dword_5d4594_3835392)))
	result = sub_504720(uint32(v8), uint32(v9))
	if result == 0 {
		return result
	}
	for j := sub_579890(); j != nil; j = sub_5798A0(j) {
		j.Field1 = 0
	}
	for k := nox_server_getFirstObjectUninited_4DA870(); k != nil; k = nox_server_getNextObjectUninited_4DA880(k) {
		k.ScriptIDVal = 0
	}
	nox_xxx_waypoint_5799C0()
	nox_xxx_unitClearPendingMB_4DB030()
	dword_5d4594_1599476 = 1
	if dword_5d4594_1599644 != 0 {
		*memmap.PtrUint32(0x973F18, 35880)++
		sub_542BF0(int32(dword_5d4594_3835312), v8, v9)
		v16.X = v8
		v16.Y = v9
		sub_543110((*byte)(memmap.PtrOff(0x973F18, 30760)), &v16.X)
		if *memmap.PtrUint32(0x5D4594, 1599580) != 0 {
			nox_fs_remove((*byte)(memmap.PtrOff(0x973F18, 36008)))
			nox_fs_move((*byte)(memmap.PtrOff(0x973F18, 38056)), (*byte)(memmap.PtrOff(0x973F18, 36008)))
			nox_script_readWriteZzz_541670((*byte)(memmap.PtrOff(0x973F18, 36008)), (*byte)(memmap.PtrOff(0x973F18, 30760)), (*byte)(memmap.PtrOff(0x973F18, 38056)))
		} else {
			*memmap.PtrUint32(0x5D4594, 1599580) = 1
			nox_fs_move((*byte)(memmap.PtrOff(0x973F18, 30760)), (*byte)(memmap.PtrOff(0x973F18, 38056)))
		}
	}
	dword_5d4594_3835312++
	result = 1
	return result
}
func sub_503EC0(a1 int32, a2 *float32) int32 {
	var (
		a1a types.Pointf
		v4  types.Pointf
		a2a types.Pointf
	)
	if dword_5d4594_1599480 != dword_5d4594_3835396 || dword_5d4594_1599480 == -1 || dword_5d4594_1599476 == 1 {
		return 0
	}
	a1a.X = float32(float64(*memmap.PtrInt32(0x5D4594, 1599508)))
	a1a.Y = float32(float64(*memmap.PtrInt32(0x5D4594, 1599512)))
	sub_4D3E30(&a1a, &a2a)
	sub_4D3E30((*types.Pointf)(unsafe.Add(a1, 56)), &v4)
	*a2 = v4.X - a2a.X
	*(*float32)(unsafe.Add(unsafe.Pointer(a2), unsafe.Sizeof(float32(0))*1)) = v4.Y - a2a.Y
	return 1
}
func nox_xxx_tileAllocTileInCoordList_5040A0(a1 int32, a2 int32, a3 float32) *unsafe.Pointer {
	var (
		result *uint32
		v4     *uint32
		v5     unsafe.Pointer
		v6     float64
		v7     bool
		v8     float32
	)
	result = (*uint32)(alloc.Calloc1(1, 0x18))
	v4 = result
	if result == nil {
		return result
	}
	v5 = alloc.Calloc1(1, 0x14)
	*v4 = uint32(uintptr(v5))
	if v5 != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*5)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*4)) = uint32(uintptr(dword_5d4594_1599556))
		if dword_5d4594_1599556 != nil {
			*(*uint32)(unsafe.Add(dword_5d4594_1599556, 20)) = uint32(uintptr(unsafe.Pointer(v4)))
		}
		dword_5d4594_1599556 = unsafe.Pointer(v4)
		v6 = float64(a1) * 46.0
		*memmap.PtrUint32(0x5D4594, 1599560)++
		v7 = int32(*(*uint8)(unsafe.Pointer(&a3))) == 1
		*(*uint8)(unsafe.Add(unsafe.Pointer(v4), 12)) = *(*uint8)(unsafe.Pointer(&a3))
		*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*1)) = float32(v6)
		v8 = float32(float64(a2) * 46.0)
		*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*2)) = v8
		result = v4
		if v7 {
			*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*1)) = float32(v6 + 23.0)
		} else {
			*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*2)) = float32(float64(v8) + 23.0)
		}
	} else {
		alloc.Free(v4)
		result = nil
	}
	return result
}
func nox_xxx_tileInit_504150(a1 int32, a2 int32) int32 {
	var (
		v5  int32
		i   *int32
		a1a types.Pointf
		v8  [72]byte
		v9  float32
		v10 float32
	)
	if *memmap.PtrInt32(0x587000, 229704) == -1 {
		if nox_tile_def_cnt > 0 {
			var v2 int32 = 0
			for i := int32(0); uint32(i) < nox_tile_def_cnt; i++ {
				var p *nox_tileDef_t = &nox_tile_defs_arr[i]
				if libc.StrCmp(&p.name[0], internCStr("TransparentFloor")) == 0 {
					*memmap.PtrUint32(0x587000, 229704) = uint32(i)
					v2 = i
					break
				}
			}
			if v2 == -1 {
				return 0
			}
		}
	}
	alloc.Memcpy(unsafe.Pointer(&v8[0]), unsafe.Pointer(sub_4D3C70()), 72)
	v5 = int32(uintptr(dword_5d4594_1599556))
	if dword_5d4594_1599556 != nil {
		v9 = float32(float64(a1))
		v10 = float32(float64(a2))
		for {
			a1a.X = v9 + *(*float32)(unsafe.Add(v5, 4))
			a1a.Y = v10 + *(*float32)(unsafe.Add(v5, 8))
			nox_xxx_tileCheckImage_51D540(int32(**(**uint32)(v5)))
			nox_xxx_tileCheckImageVari_51D570(int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(v5), 4))))
			nox_xxx_tile_51D5C0(1)
			if **(**uint32)(v5) != *memmap.PtrUint32(0x587000, 229704) {
				sub_51D8F0(&a1a)
			}
			for i = *(**int32)(unsafe.Add(*(*unsafe.Pointer)(v5), 16)); i != nil; i = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(i), 4*4))) {
				nox_xxx_tileCheckByte3_544070(*(*int32)(unsafe.Add(unsafe.Pointer(i), 4*2)))
				nox_xxx_tileCheckByte4_5440A0(*(*int32)(unsafe.Add(unsafe.Pointer(i), 4*3)))
				nox_xxx_tileCheckImage_51D540(*i)
				nox_xxx_tileCheckImageVari_51D570(*(*int32)(unsafe.Add(unsafe.Pointer(i), 4*1)))
				nox_xxx_tileSubtile_544310(&a1a)
			}
			v5 = int32(*(*uint32)(unsafe.Add(v5, 16)))
			if v5 == 0 {
				break
			}
		}
	}
	nox_xxx_tileInitdataClear_4D3C50(unsafe.Pointer(&v8[0]))
	return 1
}

type wallListItem struct {
	Wall *server.Wall
	Next *wallListItem
	Prev *wallListItem
}

var dword_5d4594_1599532 *wallListItem

func sub_504290(x, y int8) *wallListItem {
	p, _ := alloc.New(wallListItem{})
	if p == nil {
		return nil
	}
	p.Prev = nil
	p.Next = dword_5d4594_1599532
	if dword_5d4594_1599532 != nil {
		dword_5d4594_1599532.Prev = p
	}
	dword_5d4594_1599532 = p
	wl, _ := alloc.New(server.Wall{})
	p.Wall = wl
	wl.X5 = uint8(x)
	wl.Y6 = uint8(y)
	return p
}
func nox_xxx_cliWallGet_5042F0(a1 int32, a2 int32) *wallListItem {
	p := dword_5d4594_1599532
	if dword_5d4594_1599532 == nil {
		return nil
	}
	for int32(p.Wall.X5) != a1 || int32(p.Wall.Y6) != a2 {
		p = p.Next
		if p == nil {
			return nil
		}
	}
	return p
}
func sub_504330(a1 int32, a2 int32) int32 {
	for v2 := dword_5d4594_1599532; v2 == nil; v2 = v2.Next {
		v3 := (a1 + int32(v2.Wall.X5)*23) / 23
		v4 := a2 + int32(v2.Wall.Y6)*23
		v6 := v4 / 23
		v7 := nox_server_getWallAtGrid_410580(v3, v6)
		if v7 != nil {
			v8 := v2.Wall.Dir0
			if dword_5d4594_3835368 != 0 {
				v7.Dir0 = nox_xxx_wall_42A6C0(v7.Dir0, v8)
			} else {
				v7.Dir0 = v8
			}
		} else {
			v7 = nox_xxx_wallCreateAt_410250(v3, v6)
			if v7 == nil {
				return 0
			}
			v7.Dir0 = v2.Wall.Dir0
		}
		v7.Tile1 = v2.Wall.Tile1
		v7.Field2 = v2.Wall.Field2
		v7.Health7 = v2.Wall.Health7
		if (int32(v2.Wall.Flags4) & 0x80) != 0 {
			v7.Flags4 |= 0x80
		}
		if int32(v7.Field2) >= int32(nox_xxx_mapWallMaxVariation_410DD0(int32(v7.Tile1), int32(v7.Dir0), 0)) {
			v7.Field2 = 0
		}
		if int32(v7.Flags4)&4 != 0 {
			sub_4107A0(v7.Data28)
		}
		if int32(v2.Wall.Flags4)&4 != 0 {
			v7.Flags4 |= 4
			v11 := v2.Wall.Data28
			v7.Data28 = v11
			nox_xxx_wallSecretBlock_410760(v11)
		}
		if int32(v2.Wall.Flags4)&8 != 0 {
			v9 := v7.Flags4
			if (int32(v9) & 8) == 0 {
				v12 := v7.Data28
				v7.Flags4 = server.WallFlags(uint8(int8(int32(v9) | 8)))
				nox_xxx_wallBreackableListAdd_410840((*server.Wall)(v12))
			}
		}
		if int32(v2.Wall.Flags4)&0x40 != 0 {
			v7.Flags4 |= 0x40
		}
	}
	return 1
}
func sub_5044B0(a1 int32, a2 float32, a3 float32) *uint32 {
	var (
		result *uint32
		v4     *uint32
		v6     *uint32
	)
	result = (*uint32)(alloc.Calloc1(1, 0xC))
	v4 = result
	if result == nil {
		return nil
	}
	v5 := sub_579E70()
	*v4 = uint32(uintptr(unsafe.Pointer(v5)))
	if v5 != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*2)) = 0
		*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*1)) = uint32(uintptr(dword_5d4594_1599548))
		if dword_5d4594_1599548 != nil {
			*(*uint32)(unsafe.Add(dword_5d4594_1599548, 8)) = uint32(uintptr(unsafe.Pointer(v4)))
		}
		dword_5d4594_1599548 = unsafe.Pointer(v4)
		*(*uint32)(unsafe.Pointer(uintptr(*v4))) = uint32(a1)
		*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v4)), 8)) = a2
		*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v4)), 12)) = a3
		*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*v4)), 488)) = 0
		v6 = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), 4*1)))
		if v6 != nil {
			*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*v4)), 484)) = *v6
			*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v4), 4*1))), 488)) = *v4
			result = v4
		} else {
			result = v4
			*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*v4)), 484)) = 0
		}
	} else {
		alloc.Free(v4)
		result = nil
	}
	return result
}
func sub_504560(a1 int32, a2 int32) int32 {
	var (
		v2 *int32
		v4 float32
		v5 float32
	)
	v2 = (*int32)(dword_5d4594_1599548)
	if dword_5d4594_1599548 == nil {
		return 1
	}
	v4 = float32(float64(a1))
	v5 = float32(float64(a2))
	for {
		*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 8)) = v4 + *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 8))
		*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 12)) = v5 + *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 12))
		sub_579E90(*v2)
		v2 = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), 4*1)))
		if v2 == nil {
			break
		}
	}
	return 1
}
func nox_xxx_unitAddToList_5048A0(a1 *server.Object) {
	var (
		v2 *uint32
	)
	result := alloc.Calloc1(1, 0xC)
	if result == nil {
		return
	}
	*(*uint32)(unsafe.Add(result, 4*2)) = 0
	*(**server.Object)(result) = a1
	*(*uint32)(unsafe.Add(result, 4*1)) = uint32(uintptr(dword_5d4594_1599540))
	if dword_5d4594_1599540 != nil {
		*(*uint32)(unsafe.Add(dword_5d4594_1599540, 8)) = uint32(uintptr(unsafe.Pointer(result)))
	}
	dword_5d4594_1599540 = result
	*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*result)), 448)) = 0
	v2 = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(result, 4*1)))
	if v2 != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*result)), 444)) = *v2
		*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(*(*unsafe.Pointer)(unsafe.Add(result, 4*1))), 448)) = *(*uint32)(result)
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*result)), 444)) = 0
	}
}
func sub_504910(a1 int32, a2 int32) int32 {
	var (
		v2 *int32
		v4 float32
		v5 float32
	)
	v2 = (*int32)(dword_5d4594_1599540)
	if dword_5d4594_1599540 == nil {
		return 1
	}
	v4 = float32(float64(a1))
	v5 = float32(float64(a2))
	for {
		*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 56)) = v4 + *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 56))
		*(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 60)) = v5 + *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 60))
		nox_xxx_createAt_4DAA50((*server.Object)(unsafe.Pointer(uintptr(*v2))), nil, *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 56)), *(*float32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 60)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*v2)), 16)) |= 0x80000000
		v2 = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v2), 4*1)))
		if v2 == nil {
			break
		}
	}
	return 1
}
func sub_504980() int32 {
	var result int32
	if (dword_5d4594_1599480 == dword_5d4594_3835396 && dword_5d4594_1599480 != -1 && dword_5d4594_1599476 != 1 || nox_xxx_mapgenSaveMap_503830(int32(dword_5d4594_3835396)) != 0) && dword_5d4594_1599540 != nil {
		result = int32(*dword_5d4594_1599540)
	} else {
		result = 0
	}
	return result
}
func sub_5049C0(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Add(a1, 444)))
	}
	return result
}
func sub_5049D0() unsafe.Pointer {
	return dword_5d4594_1599540
}
func sub_5049E0(a1 int32) int32 {
	var result int32
	result = a1
	if a1 != 0 {
		result = int32(*(*uint32)(unsafe.Add(a1, 4)))
	}
	return result
}
func sub_504A10(a1 int32) int32 {
	var (
		v1 *int32
		v3 int32
		v4 int32
		v5 int32
		v6 int32
	)
	if a1 == 0 {
		return 0
	}
	v1 = (*int32)(dword_5d4594_1599540)
	if dword_5d4594_1599540 == nil {
		return 0
	}
	for *v1 != a1 {
		v1 = (*int32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), 4*1)))
		if v1 == nil {
			return 0
		}
	}
	v3 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1))
	if v3 != 0 {
		*(*uint32)(unsafe.Add(v3, 8)) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*2)))
	}
	v4 = *(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*2))
	if v4 != 0 {
		*(*uint32)(unsafe.Add(v4, 4)) = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(v1), 4*1)))
	}
	if unsafe.Pointer(v1) == dword_5d4594_1599540 {
		dword_5d4594_1599540 = *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v1), 4*1))
	}
	v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*v1)), 444)))
	if v5 != 0 {
		*(*uint32)(unsafe.Add(v5, 448)) = *(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*v1)), 448))
	}
	v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*v1)), 448)))
	if v6 != 0 {
		*(*uint32)(unsafe.Add(v6, 444)) = *(*uint32)(unsafe.Add(unsafe.Pointer(uintptr(*v1)), 444))
	}
	nox_xxx_objectFreeMem_4E38A0((*server.Object)(unsafe.Pointer(uintptr(*v1))))
	alloc.Free(v1)
	return 1
}
func sub_505060() unsafe.Pointer {
	var result unsafe.Pointer
	result = dword_5d4594_1599616
	if dword_5d4594_1599616 != 0 {
		alloc.FreePtr(dword_5d4594_1599616)
		dword_5d4594_1599616 = 0
	}
	return result
}
func nox_server_mapRWMapIntro_505080() int32 {
	var (
		v0     *FILE
		v1     int32
		v2     *byte
		v3     int16
		v4     uint8
		v5     *byte
		v6     *byte
		v7     uint8
		v8     int16
		v9     int32
		v10    *FILE
		v11    *FILE
		result int32
		v13    *uint8
		v14    *uint8
		i      int32
		v16    int8
		v17    uint32
		v18    int32
		v19    [1024]byte
	)
	v0 = nil
	v18 = 1
	v17 = 0
	v1 = bool2int32(noxflags.HasGame(0x200000))
	sub_505060()
	v2 = nox_fs_root()
	v3 = int16(*memmap.PtrUint16(0x587000, 229832))
	libc.StrCpy(&v19[0], v2)
	v4 = *memmap.PtrUint8(0x587000, 229834)
	v5 = &v19[libc.StrLen(&v19[0])]
	*(*uint32)(unsafe.Pointer(v5)) = *memmap.PtrUint32(0x587000, 229828)
	*(*uint16)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(uint16(0))*2)) = uint16(v3)
	*(*byte)(unsafe.Add(unsafe.Pointer(v5), 6)) = v4
	libc.StrCat(&v19[0], nox_xxx_mapGetMapName_409B40())
	*(*uint16)(unsafe.Pointer(&v19[libc.StrLen(&v19[0])])) = *memmap.PtrUint16(0x587000, 229836)
	libc.StrCat(&v19[0], nox_xxx_mapGetMapName_409B40())
	v6 = &v19[libc.StrLen(&v19[0])+1]
	v7 = *memmap.PtrUint8(0x587000, 229844)
	v8 = int16(v18)
	*(*uint32)(unsafe.Pointer(func() *byte {
		p := &v6
		*p = (*byte)(unsafe.Add(unsafe.Pointer(*p), -1))
		return *p
	}())) = *memmap.PtrUint32(0x587000, 229840)
	*(*byte)(unsafe.Add(unsafe.Pointer(v6), 4)) = v7
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v18)), 2)
	if int32(v8) > int32(int16(v18)) {
		return 0
	}
	v9 = 0
	if nox_crypt_IsReadOnly() != 0 {
		if nox_crypt_IsReadOnly() != 1 {
			return 0
		}
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v17)), 4)
		if int32(v17) <= 0 {
			return 1
		}
		if noxflags.HasGame(0x400000) {
			nox_xxx_cryptSeekCur_40E0A0(int32(v17))
			return 1
		}
		if v1 != 0 {
			v0 = nox_fs_create(&v19[0])
			if v0 == nil {
				return 0
			}
			v14 = (*uint8)(v18)
		} else {
			v13 = (*uint8)(alloc.Calloc1(1, uintptr(v17)))
			dword_5d4594_1599616 = uint32(uintptr(unsafe.Pointer(v13)))
			if v13 == nil {
				return 0
			}
			v14 = v13
		}
		for i = 0; i < int32(v17); i++ {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v16)), 1)
			if v1 != 0 {
				nox_fs_fwrite(v0, unsafe.Pointer(&v16), 1)
			} else {
				*func() *uint8 {
					p := &v14
					x := *p
					*p = (*uint8)(unsafe.Add(unsafe.Pointer(*p), 1))
					return x
				}() = uint8(v16)
			}
		}
		if v0 != nil {
			nox_fs_close(v0)
		}
		return 1
	}
	if v1 != 0 && (func() bool {
		v10 = nox_fs_open(&v19[0])
		return (func() *FILE {
			v11 = v10
			return v11
		}()) != nil
	}()) {
		v17 = uint32(nox_fs_fsize(v11))
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v17)), 4)
		if int32(v17) > 0 {
			for {
				nox_fs_fread(v11, unsafe.Pointer(&v16), 1)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v16)), 1)
				v9++
				if v9 >= int32(v17) {
					break
				}
			}
		}
		nox_fs_close(v11)
		result = 1
	} else {
		v17 = 0
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v17)), 4)
		result = 1
	}
	return result
}
func nox_server_mapRWGroupData_505C30() int32 {
	var (
		v0  int8
		i   int32
		j   int32
		v7  int8
		v8  bool
		v9  *byte
		v10 *byte
		v12 int32
		v14 [2]byte
		v15 int32
		v16 int32
		v17 int32
		v18 int32
		v19 [2]int32
		v21 int32
	)
	var v22 int32
	var v23 [76]byte
	var v24 [76]byte
	var v25 [76]byte
	var v26 [76]byte
	v15 = 3
	v0 = int8(nox_xxx_wallGet_426A30())
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v15)), 2)
	if int32(int16(v15)) > 3 {
		return 0
	}
	if nox_crypt_IsReadOnly() != 0 {
		v21 = 0
		v22 = 0
		var v13 int32
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v13)), 4)
		if v13 <= 0 {
			return 1
		}
		for v2 := int32(0); v2 < v13; v2++ {
			v17 = 0
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v17)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v23[0], uint32(uint8(int8(v17))))
			v23[v17] = 0
			v8 = int32(uint16(int16(v15))) == 2
			if int32(int16(v15)) < 2 {
				if libc.StrLen(nox_server_currentMapGetFilename_409B30())+1+libc.StrLen(&v23[0]) >= 0x35 {
					return 0
				}
				v9 = nox_server_currentMapGetFilename_409B30()
				nox_sprintf(&v25[0], internCStr("%s.map:%s"), v9, &v23[0])
				libc.StrCpy(&v23[0], &v25[0])
				v8 = int32(uint16(int16(v15))) == 2
			}
			if v8 {
				libc.StrCpy(&v14[0], internCStr(":"))
				libc.StrCpy(&v24[0], nox_server_currentMapGetFilename_409B30())
				*(*uint16)(unsafe.Pointer(&v24[libc.StrLen(&v24[0])])) = *memmap.PtrUint16(0x587000, 229976)
				libc.StrCpy(&v26[0], &v23[0])
				libc.StrTok(&v26[0], &v14[0])
				v10 = libc.StrTok(nil, &v14[0])
				if libc.StrLen(nox_server_currentMapGetFilename_409B30())+1+libc.StrLen(v10) >= 0x35 {
					return 0
				}
				libc.StrCat(&v24[0], v10)
				libc.StrCpy(&v23[0], &v24[0])
			}
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v18)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v16)), 4)
			if (int32(v0) & 4) == 0 {
				if noxflags.HasGame(0x400000) {
					sub_504600(&v23[0], uint32(v16), uint8(int8(v18)))
				} else if noxflags.HasGame(2097153) {
					nox_server_mapLoadAddGroup_57C0C0(&v23[0], uint32(v16), uint8(int8(v18)))
				}
			}
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v12)), 4)
			for v11 := int32(0); v11 < v12; v11++ {
				if int32(uint8(int8(v18))) != 0 {
					if int32(uint8(int8(v18))) == 1 {
						nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v19[0])), 4)
					} else if int32(uint8(int8(v18))) != 2 {
						if int32(uint8(int8(v18))) != 3 {
							return 0
						}
						nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v19[0])), 4)
					} else {
						nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v19[0])), 4)
						nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v19[1])), 4)
					}
				} else {
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v19[0])), 4)
				}
				if (int32(v0) & 4) == 0 {
					if noxflags.HasGame(0x400000) {
						sub_5046A0((*uint32)(unsafe.Pointer(&v19[0])), uint32(v16))
					} else {
						sub_57C130((*uint32)(unsafe.Pointer(&v19[0])), v16)
					}
				}
			}
		}
		return 1
	}
	var v13 int32 = 0
	for i = int32(uintptr(nox_server_getFirstMapGroup_57C080())); i != 0; i = nox_server_getNextMapGroup_57C090(i) {
		v13++
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v13)), 4)
	for v4 := int32(uintptr(nox_server_getFirstMapGroup_57C080())); v4 != 0; v4 = nox_server_getNextMapGroup_57C090(v4) {
		*(*uint8)(unsafe.Pointer(&v17)) = uint8(int8(libc.StrLen((*byte)(unsafe.Add(v4, 8))) + 1))
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v17)), 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(v4, 8)), uint32(uint8(int8(v17))))
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(v4), 1)
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(v4, 4)), 4)
		v12 = 0
		for j = int32(*(*uint32)(unsafe.Add(v4, 84))); j != 0; j = int32(*(*uint32)(unsafe.Add(j, 8))) {
			v12++
		}
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v12)), 4)
		for v6 := int32(*(*uint32)(unsafe.Add(v4, 84))); v6 != 0; v6 = int32(*(*uint32)(unsafe.Add(v6, 8))) {
			v7 = int8(*(*uint8)(v4))
			if int32(*(*uint8)(v4)) == 0 || int32(v7) == 1 {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(v6), 4)
			} else if int32(v7) != 2 {
				if int32(v7) != 3 {
					return 0
				}
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(v6), 4)
			} else {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(v6), 4)
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(v6, 4)), 4)
			}
		}
	}
	return 1
}
func nox_server_mapRWWaypoints_506260(a1 *uint32) int32 {
	var (
		v2  *float32
		v3  *uint32
		v4  *byte
		v5  int32
		v6  *uint8
		v7  *byte
		v8  **byte
		v9  *byte
		v10 *uint8
		v11 int32
		v12 *uint8
		v13 *uint8
		v14 int32
		v15 *uint8
		v16 *uint8
		v17 *float32
		v18 int32
		v19 int32
		v20 int32
		v21 uint32
		v22 float32
		v23 float32
		v24 int32
		v25 int32
		v26 Point32
		v27 int32
		v28 int64
	)
	var v29 int64
	var v30 int4
	var v31 [76]byte
	v18 = 4
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v18)), 2)
	if int32(int16(v18)) > 4 {
		return 0
	}
	if nox_crypt_IsReadOnly() != 0 {
		nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v19)), 4)
		v24 = 0
		if v19 > 0 {
			for {
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v25)), 4)
				if int32(int16(v18)) < 4 {
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 4)
					v28 = int64(v21)
					v22 = float32(float64(v21))
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 4)
					v29 = int64(v21)
					v23 = float32(float64(v21))
				} else {
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v22)), 4)
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v23)), 4)
				}
				if int32(int16(v18)) >= 3 {
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v20)), 1)
					if int32(uint8(int8(v20))) != 0 {
						nox_xxx_fileReadWrite_426AC0_file3_fread_impl(&v31[0], uint32(uint8(int8(v20))))
						v31[uint8(int8(v20))] = 0
					} else {
						v31[0] = *memmap.PtrUint8(0x5D4594, 1599648)
					}
				}
				if a1 != nil {
					v7 = nox_xxx_mapGetWallSize_426A70()
					sub_428170(unsafe.Pointer(a1), &v30)
					v22 = float32(float64(v22) - float64(int32(*(*uint32)(unsafe.Pointer(v7))*23)) + float64(v30.field_0) - 11.0)
					v23 = float32(float64(v23) - float64(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*1))*23)) + float64(v30.field_4) - 11.0)
				}
				if noxflags.HasGame(0x400000) {
					v8 = (**byte)(unsafe.Pointer(sub_5044B0(v25, v22, v23)))
					v9 = *v8
					v17 = (*float32)(unsafe.Pointer(*v8))
				} else {
					v17 = nox_xxx_waypointNewNotMap_579970(v25, v22, v23)
					v9 = (*byte)(unsafe.Pointer(v17))
				}
				if v9 == nil {
					break
				}
				if int32(int16(v18)) >= 3 {
					libc.StrCpy((*byte)(unsafe.Add(unsafe.Pointer(v9), 16)), &v31[0])
				}
				nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v9), 480)), 4)
				if int32(int16(v18)) < 4 {
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v21)), 4)
					v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 476))
					*(*byte)(unsafe.Add(unsafe.Pointer(v9), 476)) = byte(v21)
				} else {
					v10 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 476))
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v9), 476)), 1)
				}
				if int32(int16(v18)) >= 2 {
					v14 = 0
					if int32(*v10) != 0 {
						v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 96))
						v16 = (*uint8)(unsafe.Add(unsafe.Pointer(v17), unsafe.Sizeof(float32(0))*87))
						for {
							nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v16, 4)
							nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v15, 1)
							v14++
							v16 = (*uint8)(unsafe.Add(unsafe.Pointer(v16), 4))
							v15 = (*uint8)(unsafe.Add(unsafe.Pointer(v15), 8))
							if v14 >= int32(*v10) {
								break
							}
						}
					}
				} else {
					v11 = 0
					if int32(*v10) != 0 {
						v12 = (*uint8)(unsafe.Add(unsafe.Pointer(v9), 96))
						v13 = (*uint8)(unsafe.Add(unsafe.Pointer(v17), unsafe.Sizeof(float32(0))*87))
						for {
							nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v13, 4)
							*v12 = 2
							v11++
							v13 = (*uint8)(unsafe.Add(unsafe.Pointer(v13), 4))
							v12 = (*uint8)(unsafe.Add(unsafe.Pointer(v12), 8))
							if v11 >= int32(*v10) {
								break
							}
						}
					}
				}
				if func() int32 {
					p := &v24
					*p++
					return *p
				}() >= v19 {
					return 1
				}
			}
			return 0
		}
		return 1
	}
	v19 = 0
	v2 = (*float32)(nox_xxx_waypointGetList_579860())
	if v2 != nil {
		for {
			v3 = a1
			if a1 == nil || (func() int32 {
				v26.X = int32(int64(*(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*2))))
				v26.Y = int32(int64(*(*float32)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(float32(0))*3))))
				return nox_xxx_wallMath_427F30(&v26, (*int32)(unsafe.Pointer(a1)))
			}()) != 0 {
				v19++
			}
			v2 = (*float32)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(v2)))))))
			if v2 == nil {
				break
			}
		}
	} else {
		v3 = a1
	}
	nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v19)), 4)
	v4 = (*byte)(nox_xxx_waypointGetList_579860())
	if v4 == nil {
		return 1
	}
	for {
		if v3 == nil || (func() int32 {
			v26.X = int32(int64(*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*2))))
			v26.Y = int32(int64(*(*float32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(float32(0))*3))))
			return nox_xxx_wallMath_427F30(&v26, (*int32)(unsafe.Pointer(v3)))
		}()) != 0 {
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v4, 4)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v4), 8)), 4)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v4), 12)), 4)
			*(*uint8)(unsafe.Pointer(&v20)) = uint8(int8(libc.StrLen((*byte)(unsafe.Add(unsafe.Pointer(v4), 16)))))
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v20)), 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v4), 16)), uint32(uint8(int8(v20))))
			v27 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*120)) & 1)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Pointer(&v27)), 4)
			nox_xxx_fileReadWrite_426AC0_file3_fread_impl((*uint8)(unsafe.Add(unsafe.Pointer(v4), 476)), 1)
			v5 = 0
			if *(*byte)(unsafe.Add(unsafe.Pointer(v4), 476)) != 0 {
				v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v4), 96))
				for {
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl(*(**uint8)(unsafe.Add(unsafe.Pointer(v6), -int(unsafe.Sizeof((*uint8)(nil))*1))), 4)
					nox_xxx_fileReadWrite_426AC0_file3_fread_impl(v6, 1)
					v5++
					v6 = (*uint8)(unsafe.Add(unsafe.Pointer(v6), 8))
					if v5 >= int32(*(*byte)(unsafe.Add(unsafe.Pointer(v4), 476))) {
						break
					}
				}
			}
			v3 = a1
		}
		v4 = (*byte)(unsafe.Pointer(uintptr(nox_xxx_waypointNext_579870(int32(uintptr(unsafe.Pointer(v4)))))))
		if v4 == nil {
			break
		}
	}
	return 1
}
func nox_xxx_allocVoteArray_5066D0() int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(nox_new_alloc_class(internCStr("VoteClass"), 52, 64))))
	nox_alloc_vote_1599652 = result
	if result != 0 {
		dword_5d4594_1599656 = 0
		result = 1
	}
	return result
}
func sub_506720() int32 {
	var result int32
	nox_free_alloc_class((*nox_alloc_class)(nox_alloc_vote_1599652))
	result = 0
	nox_alloc_vote_1599652 = nil
	dword_5d4594_1599656 = 0
	return result
}
func sub_506740(a1p *server.Object) int32 {
	var (
		a1     = a1p
		result int32
		v2     int32
		v3     int32
		v4     int32
	)
	result = a1
	if a1 == nil {
		return result
	}
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(a1), 8)))&4 == 0 {
		return result
	}
	result = int32(dword_5d4594_1599656)
	v2 = 1 << int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a1.UpdateData, 276)), 2064)))
	if dword_5d4594_1599656 == 0 {
		return result
	}
	for {
		v3 = int32(*(*uint32)(unsafe.Add(result, 8)))
		v4 = int32(*(*uint32)(unsafe.Add(result, 44)))
		if v3&v2 != 0 {
			*(*uint32)(unsafe.Add(result, 8)) = uint32(^v2 & v3)
			*(*uint8)(unsafe.Add(result, 4))--
		}
		if int32(*(*uint8)(unsafe.Add(result, 4))) == 0 {
			sub_5067B0(result)
		}
		result = v4
		if v4 == 0 {
			break
		}
	}
	return result
}
func sub_5067B0(a1 int32) {
	var v1 int32
	if a1 != 0 {
		if *(*uint32)(a1) == 2 {
			v1 = 0
			for {
				if uint32(1<<v1)&*(*uint32)(unsafe.Add(a1, 8)) != 0 {
					nox_xxx_netSendVote_506840(v1)
				}
				v1++
				if v1 >= 32 {
					break
				}
			}
		}
		sub_506810(a1)
		nox_alloc_class_free_obj_first((*nox_alloc_class)(nox_alloc_vote_1599652), a1)
		if dword_5d4594_1599656 == 0 {
			sub_507190(math.MaxUint8, 0)
		}
	}
}
func sub_506810(a1 int32) int32 {
	var (
		result int32
		v2     int32
		v3     int32
	)
	result = a1
	v2 = int32(*(*uint32)(unsafe.Add(a1, 44)))
	if v2 != 0 {
		*(*uint32)(unsafe.Add(v2, 48)) = *(*uint32)(unsafe.Add(a1, 48))
	}
	v3 = int32(*(*uint32)(unsafe.Add(a1, 48)))
	if v3 != 0 {
		result = int32(*(*uint32)(unsafe.Add(a1, 44)))
		*(*uint32)(unsafe.Add(v3, 44)) = uint32(result)
	} else {
		dword_5d4594_1599656 = *(*uint32)(unsafe.Add(a1, 44))
	}
	return result
}
func nox_xxx_netSendVote_506840(a1 int32) int32 {
	var v2 [2]byte
	v2[0] = 238
	v2[1] = 7
	return nox_xxx_netSendPacket1_4E5390(a1, unsafe.Pointer(&v2[0]), 2, nil, 1)
}
func sub_506870(a1 int32, a2 unsafe.Pointer, a3 *wchar2_t) {
	if a2 != nil && int32(*(*uint8)(unsafe.Add(a2, 8)))&4 != 0 {
		switch a1 {
		case 0:
			sub_5068E0(0, a2, a3)
		case 1:
			sub_5068E0(1, a2, a3)
		case 2:
			sub_506B00(2, a2)
		case 3:
			sub_506B80(3, a2, a3)
		}
	}
}
func sub_5068E0(a1 int32, a2 unsafe.Pointer, a3 *wchar2_t) int8 {
	var (
		v3 int32
		v4 int32
		v5 int32
		v7 int32
	)
	*(*uint8)(unsafe.Pointer(&v3)) = *memmap.PtrUint8(0x587000, 229980)
	if *memmap.PtrUint32(0x587000, 229980) > 0x20 {
		return int8(v3)
	}
	if *memmap.PtrUint32(0x587000, 229980) == 0 {
		return int8(v3)
	}
	if a3 == nil {
		return int8(v3)
	}
	if a2 == nil {
		return int8(v3)
	}
	v4 = 1 << int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a2, 748)), 276)), 2064)))
	v3 = nox_common_playerInfoGetFirst_416EA0()
	v5 = v3
	if v3 == 0 {
		return int8(v3)
	}
	for {
		if *(*uint32)(unsafe.Add(v5, 2092)) == 1 {
			v3 = nox_wcscmp((*wchar2_t)(unsafe.Add(v5, 4704)), a3)
			if v3 == 0 {
				break
			}
		}
		v3 = nox_common_playerInfoGetNext_416EE0((*server.Player)(v5))
		v5 = v3
		if v3 == 0 {
			return int8(v3)
		}
	}
	if int32(*(*uint8)(unsafe.Add(v5, 2064))) == 31 {
		return int8(v3)
	}
	v6 := *(*unsafe.Pointer)(unsafe.Add(v5, 2056))
	if v6 == nil {
		return int8(v3)
	}
	if a2 == v6 {
		return int8(v3)
	}
	if !(!nox_xxx_CheckGameplayFlags_417DA0(4) || (func() int32 {
		v3 = nox_xxx_servCompareTeams_419150((*server.ObjectTeam)(unsafe.Add(a2, 48)), unsafe.Add(v6, 48))
		return v3
	}()) != 0) {
		return int8(v3)
	}
	v7 = int32(dword_5d4594_1599656)
	if dword_5d4594_1599656 != 0 {
		for *(*uint32)(v7) != uint32(a1) || *(*unsafe.Pointer)(unsafe.Add(v7, 28)) != v6 {
			v7 = int32(*(*uint32)(unsafe.Add(v7, 44)))
			if v7 == 0 {
				break
			}
		}
	}
	if v7 == 0 {
		v3 = int32(uintptr(unsafe.Pointer(sub_506A20(a1, a2))))
		v7 = v3
		if v3 == 0 {
			return int8(v3)
		}
		*(*unsafe.Pointer)(unsafe.Add(v3, 28)) = v6
		if nox_xxx_CheckGameplayFlags_417DA0(4) {
			*(*uint32)(unsafe.Add(v7, 20)) = 1
		}
	}
	v3 = int32(*(*uint32)(unsafe.Add(v7, 8)))
	if (v4 & v3) != 0 {
		return int8(v3)
	}
	*(*uint8)(unsafe.Pointer(&v3)) = uint8(int8(int32(*(*uint8)(unsafe.Add(v7, 4))) + 1))
	*(*uint32)(unsafe.Add(v7, 8)) |= uint32(v4)
	*(*uint8)(unsafe.Add(v7, 4)) = uint8(int8(v3))
	return int8(v3)
}
func sub_506A20(a1 int32, a2 unsafe.Pointer) *uint32 {
	var (
		v2 int32
		v3 *uint32
	)
	v2 = 0
	if a2 == nil || (int32(*(*uint8)(unsafe.Add(a2, 8)))&4) == 0 {
		return nil
	}
	if dword_5d4594_1599656 == 0 {
		v2 = 1
	}
	v3 = (*uint32)(nox_alloc_class_new_obj_zero((*nox_alloc_class)(nox_alloc_vote_1599652)))
	if v3 == nil {
		return nil
	}
	*v3 = uint32(a1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*6)) = gameFrame()
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v3), 4*4)) = unsafe.Add(a2, 48)
	switch a1 {
	case 0, 1:
		*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 12)) = *memmap.PtrUint8(0x587000, 229980)
	case 2, 3:
		*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 12)) = 6
	default:
		*(*uint8)(unsafe.Add(unsafe.Pointer(v3), 12)) = *memmap.PtrUint8(0x587000, 229984)
	}
	nox_xxx_voteAddMB_506AD0(int32(uintptr(unsafe.Pointer(v3))))
	if v2 != 0 {
		sub_507190(math.MaxUint8, 1)
	}
	return v3
}
func nox_xxx_voteAddMB_506AD0(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Add(a1, 48)) = 0
	*(*uint32)(unsafe.Add(a1, 44)) = dword_5d4594_1599656
	if dword_5d4594_1599656 != 0 {
		*(*uint32)(unsafe.Add(dword_5d4594_1599656, 48)) = uint32(a1)
	}
	dword_5d4594_1599656 = uint32(a1)
	return result
}
func sub_506B00(a1 int32, a2 unsafe.Pointer) *uint32 {
	var (
		result *uint32
		v3     int32
		v4     int8
	)
	result = nox_server_resetQuestMinVotes_229988
	if nox_server_resetQuestMinVotes_229988 == 0 {
		return result
	}
	if a2 == nil {
		return result
	}
	result = *(**uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a2, 748)), 276))
	v3 = 1 << int32(*(*uint8)(unsafe.Add(unsafe.Pointer(result), 2064)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*1198)) == 0 {
		return result
	}
	result = dword_5d4594_1599656
	if dword_5d4594_1599656 != 0 {
		for *result != uint32(a1) {
			result = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(result), 4*11)))
			if result == nil {
				break
			}
		}
	}
	if result == nil {
		result = sub_506A20(a1, a2)
		if result == nil {
			return result
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*5)) = 0
	}
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*2)) & uint32(v3)) != 0 {
		return result
	}
	v4 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(result), 4))) + 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*2)) |= uint32(v3)
	*(*uint8)(unsafe.Add(unsafe.Pointer(result), 4)) = uint8(v4)
	return result
}
func sub_506B80(a1 int32, a2 unsafe.Pointer, a3 *wchar2_t) *uint32 {
	var (
		result *uint32
		v4     int32
		v5     *wchar2_t
		v7     int8
	)
	result = nox_server_kickQuestPlayerMinVotes_229992
	if nox_server_kickQuestPlayerMinVotes_229992 == 0 {
		return result
	}
	if a3 == nil {
		return result
	}
	result = (*uint32)(a2)
	if a2 == nil {
		return result
	}
	result = *(**uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a2, 748)), 276))
	v4 = 1 << int32(*(*uint8)(unsafe.Add(unsafe.Pointer(result), 2064)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*1198)) == 0 {
		return result
	}
	result = nox_common_playerInfoGetFirst_416EA0()
	v5 = (*wchar2_t)(unsafe.Pointer(result))
	if result == nil {
		return result
	}
	for {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*523)) == 1 {
			result = (*uint32)(unsafe.Pointer(uintptr(nox_wcscmp((*wchar2_t)(unsafe.Add(unsafe.Pointer(v5), unsafe.Sizeof(wchar2_t(0))*2352)), a3))))
			if result == nil {
				break
			}
		}
		result = nox_common_playerInfoGetNext_416EE0((*server.Player)(unsafe.Pointer(v5)))
		v5 = (*wchar2_t)(unsafe.Pointer(result))
		if result == nil {
			return result
		}
	}
	if int32(*(*uint8)(unsafe.Add(unsafe.Pointer(v5), 2064))) == 31 {
		return result
	}
	result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*1198)))))
	if result == nil {
		return result
	}
	v6 := *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(v5), 4*514))
	if v6 == nil {
		return result
	}
	if a2 == v6 {
		return result
	}
	result = dword_5d4594_1599656
	if dword_5d4594_1599656 != 0 {
		for *result != uint32(a1) || *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(result), 4*7)) != v6 {
			result = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(result), 4*11)))
			if result == nil {
				break
			}
		}
	}
	if result == nil {
		result = sub_506A20(a1, a2)
		if result == nil {
			return result
		}
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(result), 4*7)) = v6
	}
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*2)) & uint32(v4)) != 0 {
		return result
	}
	v7 = int8(int32(*(*uint8)(unsafe.Add(unsafe.Pointer(result), 4))) + 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*2)) |= uint32(v4)
	*(*uint8)(unsafe.Add(unsafe.Pointer(result), 4)) = uint8(v7)
	return result
}
func sub_506C90(a1 int32, a2 unsafe.Pointer, a3 *wchar2_t) {
	if a2 != nil && int32(*(*uint8)(unsafe.Add(a2, 8)))&4 != 0 {
		switch a1 {
		case 0:
			sub_506D00(a2, a3)
		case 1:
			sub_506D00(a2, a3)
		case 2:
			sub_506DE0(a2)
		case 3:
			sub_506E50(a2, a3)
		default:
			return
		}
	}
}
func sub_506D00(a1 unsafe.Pointer, a2 *wchar2_t) {
	var (
		v2 *byte
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 bool
	)
	if a1 != nil {
		if a2 != nil {
			if int32(*(*uint8)(unsafe.Add(a1, 8)))&4 != 0 {
				v2 = nox_common_playerInfoGetFirst_416EA0()
				if v2 != nil {
					for *(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*523)) != 1 || nox_wcscmp((*wchar2_t)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(wchar2_t(0))*2352)), a2) != 0 {
						v2 = nox_common_playerInfoGetNext_416EE0((*server.Player)(unsafe.Pointer(v2)))
						if v2 == nil {
							return
						}
					}
					if *(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)) != 31 {
						v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*514)))
						if v3 != 0 {
							v4 = int32(dword_5d4594_1599656)
							v5 = 1 << int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a1, 748)), 276)), 2064)))
							if dword_5d4594_1599656 != 0 {
								for *(*uint32)(v4) != 0 || *(*uint32)(unsafe.Add(v4, 28)) != uint32(v3) || (uint32(v5)&*(*uint32)(unsafe.Add(v4, 8))) == 0 {
									v4 = int32(*(*uint32)(unsafe.Add(v4, 44)))
									if v4 == 0 {
										return
									}
								}
								if v4 != 0 {
									v6 = int32(uint32(^v5) & *(*uint32)(unsafe.Add(v4, 8)))
									v7 = int32(func() uint8 {
										p := (*uint8)(unsafe.Add(v4, 4))
										x := *p
										*p--
										return x
									}()) == 1
									*(*uint32)(unsafe.Add(v4, 8)) = uint32(v6)
									if v7 {
										sub_5067B0(v4)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func sub_506DE0(a1 unsafe.Pointer) {
	var (
		result int32
		v2     int32
		v3     int8
		v4     int32
	)
	if a1 != nil {
		if int32(*(*uint8)(unsafe.Add(a1, 8)))&4 != 0 {
			result = int32(dword_5d4594_1599656)
			v2 = 1 << int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a1, 748)), 276)), 2064)))
			if dword_5d4594_1599656 != 0 {
				for *(*uint32)(result) != 2 {
					result = int32(*(*uint32)(unsafe.Add(result, 44)))
					if result == 0 {
						return
					}
				}
				if result != 0 && uint32(v2)&*(*uint32)(unsafe.Add(result, 8)) != 0 {
					v3 = int8(int32(*(*uint8)(unsafe.Add(result, 4))) - 1)
					v4 = int32(uint32(^v2) & *(*uint32)(unsafe.Add(result, 8)))
					*(*uint8)(unsafe.Add(result, 4)) = uint8(v3)
					*(*uint32)(unsafe.Add(result, 8)) = uint32(v4)
					if int32(v3) == 0 {
						sub_5067B0(result)
					}
				}
			}
		}
	}
}
func sub_506E50(a1 unsafe.Pointer, a2 *wchar2_t) {
	var (
		v2 *byte
		v3 int32
		v4 int32
		v5 int32
		v6 int32
		v7 bool
	)
	if a1 != nil {
		if a2 != nil {
			if int32(*(*uint8)(unsafe.Add(a1, 8)))&4 != 0 {
				v2 = nox_common_playerInfoGetFirst_416EA0()
				if v2 != nil {
					for *(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*523)) != 1 || nox_wcscmp((*wchar2_t)(unsafe.Add(unsafe.Pointer(v2), unsafe.Sizeof(wchar2_t(0))*2352)), a2) != 0 {
						v2 = nox_common_playerInfoGetNext_416EE0((*server.Player)(unsafe.Pointer(v2)))
						if v2 == nil {
							return
						}
					}
					if *(*byte)(unsafe.Add(unsafe.Pointer(v2), 2064)) != 31 {
						v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v2), 4*514)))
						if v3 != 0 {
							v4 = int32(dword_5d4594_1599656)
							v5 = 1 << int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(a1, 748)), 276)), 2064)))
							if dword_5d4594_1599656 != 0 {
								for *(*uint32)(v4) != 3 || *(*uint32)(unsafe.Add(v4, 28)) != uint32(v3) || (uint32(v5)&*(*uint32)(unsafe.Add(v4, 8))) == 0 {
									v4 = int32(*(*uint32)(unsafe.Add(v4, 44)))
									if v4 == 0 {
										return
									}
								}
								if v4 != 0 {
									v6 = int32(uint32(^v5) & *(*uint32)(unsafe.Add(v4, 8)))
									v7 = int32(func() uint8 {
										p := (*uint8)(unsafe.Add(v4, 4))
										x := *p
										*p--
										return x
									}()) == 1
									*(*uint32)(unsafe.Add(v4, 8)) = uint32(v6)
									if v7 {
										sub_5067B0(v4)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func sub_506F80(a1 int32) {
	var (
		v1 int32
		v3 int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(a1, 28)))
	if int32(*(*uint8)(unsafe.Add(v1, 16)))&0x20 != 0 {
		sub_5067B0(a1)
		return
	}
	*(*uint32)(unsafe.Add(a1, 16)) = uint32(v1 + 48)
	if sub_507000(a1) == 1 {
		v3 = int32(*(*uint32)(unsafe.Add(v1, 748)))
		nox_xxx_playerCallDisconnect_4DEAB0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 2064))), 4)
		sub_416770(15, (*wchar2_t)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 4704)), (*byte)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 2112)))
		sub_5067B0(a1)
	}
}
func sub_507000(a1 int32) int32 {
	var (
		v1 int32
		j  int32
	)
	v1 = 0
	if int32(*(*uint8)(unsafe.Add(a1, 4))) >= int32(*(*uint8)(unsafe.Add(a1, 12))) {
		return 1
	}
	if *(*uint32)(unsafe.Add(a1, 20)) == 1 {
		for i := nox_xxx_getFirstPlayerUnit_4DA7C0(); i != nil; i = nox_xxx_getNextPlayerUnit_4DA7F0(i) {
			if nox_xxx_servCompareTeams_419150(*(*unsafe.Pointer)(unsafe.Add(a1, 16)), (*server.ObjectTeam)(unsafe.Add(unsafe.Pointer(i), 48))) != 0 {
				v1++
			}
		}
	} else {
		for j = nox_xxx_getFirstPlayerUnit_4DA7C0(); j != 0; j = nox_xxx_getNextPlayerUnit_4DA7F0((*server.Object)(j)) {
			v1++
		}
	}
	return bool2int32(uint32(*(*uint8)(unsafe.Add(a1, 4))) >= uint32(v1-1) && int32(*(*uint8)(unsafe.Add(a1, 4))) >= 2)
}
func sub_507090(a1 int32) {
	var (
		i  int32
		v3 int32
		v4 *byte
	)
	nox_xxx_player_4E3CE0()
	if int32(*(*uint8)(unsafe.Add(a1, 4))) >= nox_xxx_player_4E3CE0() {
		for i = nox_xxx_getFirstPlayerUnit_4DA7C0(); i != 0; i = nox_xxx_getNextPlayerUnit_4DA7F0((*server.Object)(i)) {
			v3 = int32(*(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(i, 748)), 276)))
			if *(*uint32)(unsafe.Add(v3, 4792)) == 1 {
				Nox_xxx_playerRespawn_4F7EF0((*server.Object)(*(*unsafe.Pointer)(unsafe.Add(v3, 2056))))
			}
		}
		nox_game_setQuestStage_4E3CD0(0)
		v4 = nox_xxx_getQuestMapFile_4D0F60()
		nox_xxx_mapLoad_4D2450(v4)
		sub_5067B0(a1)
	}
}
func sub_507100(a1 int32) {
	var (
		v1     int32
		v2     int32
		v3     uint32
		result uint32
	)
	v1 = int32(*(*uint32)(unsafe.Add(a1, 28)))
	if v1 == 0 {
		sub_5067B0(a1)
		return
	}
	if int32(*(*uint8)(unsafe.Add(v1, 16)))&0x20 != 0 {
		sub_5067B0(a1)
		return
	}
	v2 = int32(*(*uint32)(unsafe.Add(v1, 748)))
	if *(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v2, 276)), 4792)) == 0 {
		sub_5067B0(a1)
		return
	}
	if int32(*(*uint8)(unsafe.Add(a1, 4))) >= int32(*(*uint8)(unsafe.Add(a1, 12))) {
		goto LABEL_8
	}
	v3 = uint32(nox_xxx_player_4E3CE0())
	if v3 <= 1 {
		sub_5067B0(a1)
		return
	}
	result = v3 - 1
	if uint32(*(*uint8)(unsafe.Add(a1, 4))) < result || int32(*(*uint8)(unsafe.Add(a1, 4))) < 2 {
		return
	}
LABEL_8:
	sub_4DCFB0((*server.Object)(v1))
	sub_416770(15, (*wchar2_t)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v2, 276)), 4704)), (*byte)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v2, 276)), 2112)))
	sub_5067B0(a1)
	return
}
func sub_507190(a1 int32, a2 int8) int32 {
	var v4 [3]byte
	v4[0] = 238
	v4[1] = 6
	v4[2] = byte(a2)
	return nox_xxx_netSendPacket1_4E5390(a1, unsafe.Pointer(&v4[0]), 3, nil, 1)
}
func sub_5071C0() int32 {
	return bool2int32(dword_5d4594_1599656 != 0)
}
func sub_509120(a1 *uint32, a2 int32, a3 *byte) {
	var (
		v3 *byte
		v4 int32
		v5 *uint32
		v6 *byte
		v7 *uint32
		v8 int32
		v9 *uint32
	)
	v3 = (*byte)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), 4*189)))
	if v3 != nil {
		if a2 == 14 {
			if noxflags.HasGame(6291456) {
				libc.StrCpy(v3, a3)
			} else {
				*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*192)) = uint32(nox_script_indexByEvent(a3))
			}
			return
		}
		v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)))
		if v4&0x200 != 0 {
			v5 = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
			if a2 != 0 {
				if a2 == 1 {
					if !noxflags.HasGame(6291456) {
						*(*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*6)) = uint32(nox_script_indexByEvent(a3))
						return
					}
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 256))
				} else {
					if a2 != 2 {
						return
					}
					if !noxflags.HasGame(6291456) {
						*(*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*8)) = uint32(nox_script_indexByEvent(a3))
						return
					}
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 384))
				}
			} else {
				if !noxflags.HasGame(6291456) {
					*(*uint32)(unsafe.Add(unsafe.Pointer(v5), 4*4)) = uint32(nox_script_indexByEvent(a3))
					return
				}
				v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 512))
			}
			libc.StrCpy(v6, a3)
			return
		}
		if v4&2 != 0 {
			v7 = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
			switch a2 {
			case 3:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 640))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*309)) = uint32(nox_script_indexByEvent(a3))
			case 4:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 768))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*307)) = uint32(nox_script_indexByEvent(a3))
			case 5:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 896))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*317)) = uint32(nox_script_indexByEvent(a3))
			case 6:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1024))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*311)) = uint32(nox_script_indexByEvent(a3))
			case 7:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1152))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*313)) = uint32(nox_script_indexByEvent(a3))
			case 8:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1280))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*315)) = uint32(nox_script_indexByEvent(a3))
			case 9:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1408))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*319)) = uint32(nox_script_indexByEvent(a3))
			case 10:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1536))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*321)) = uint32(nox_script_indexByEvent(a3))
			case 11:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1664))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*323)) = uint32(nox_script_indexByEvent(a3))
			case 13:
				if noxflags.HasGame(6291456) {
					v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1792))
					libc.StrCpy(v6, a3)
					return
				}
				*(*uint32)(unsafe.Add(unsafe.Pointer(v7), 4*325)) = uint32(nox_script_indexByEvent(a3))
			default:
				return
			}
		} else {
			if v4&0x800 != 0 {
				v8 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*175)))
				if a2 != 12 {
					return
				}
				if !noxflags.HasGame(6291456) {
					*(*uint32)(unsafe.Add(v8, 4)) = uint32(nox_script_indexByEvent(a3))
					return
				}
				v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 128))
				libc.StrCpy(v6, a3)
				return
			}
			if uint32(v4)&0x20000 != 0 {
				v9 = (*uint32)(*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), 4*187)))
				switch a2 {
				case 15:
					if noxflags.HasGame(6291456) {
						v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 1920))
						libc.StrCpy(v6, a3)
						return
					}
					*(*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*13)) = uint32(nox_script_indexByEvent(a3))
				case 16:
					if noxflags.HasGame(6291456) {
						v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 2048))
						libc.StrCpy(v6, a3)
						return
					}
					*(*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*15)) = uint32(nox_script_indexByEvent(a3))
				case 17:
					if noxflags.HasGame(6291456) {
						v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 2304))
						libc.StrCpy(v6, a3)
						return
					}
					*(*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*17)) = uint32(nox_script_indexByEvent(a3))
				case 18:
					if noxflags.HasGame(6291456) {
						v6 = (*byte)(unsafe.Add(unsafe.Pointer(v3), 2176))
						libc.StrCpy(v6, a3)
						return
					}
					*(*uint32)(unsafe.Add(unsafe.Pointer(v9), 4*19)) = uint32(nox_script_indexByEvent(a3))
				default:
					return
				}
			}
		}
	}
}
func sub_5095E0() {
	var (
		v0  *server.Team
		v1  int32
		v2  int32
		v4  int32
		v7  int32
		v8  int32
		v10 *server.Team
		v11 unsafe.Pointer
	)
	v2 = math.MaxInt32
	for i := nox_server_teamFirst_418B10(); i != nil; i = nox_server_teamNext_418B60(i) {
		v4 = int32(i.Lessons)
		if v4 <= v2 {
			v1 = bool2int32(v4 == v2 && v0 != nil)
			v2 = int32(i.Lessons)
			v10 = i
			v0 = i
		}
	}
	v5 := nox_xxx_getFirstPlayerUnit_4DA7C0()
	if v5 != nil {
		for {
			v6 := v5.UpdateData
			if nox_xxx_servObjectHasTeam_419130((*server.ObjectTeam)(unsafe.Add(unsafe.Pointer(v5), 48))) == 0 {
				v7 = int32(*(*uint32)(unsafe.Add(v6, 276)))
				if (int32(*(*uint8)(unsafe.Add(v7, 3680))) & 1) == 0 {
					v8 = int32(*(*uint32)(unsafe.Add(v7, 2140)))
					if v8 <= v2 {
						v1 = bool2int32(v8 == v2 && v11 != nil)
						v2 = v8
						v11 = unsafe.Pointer(v5)
					}
				}
			}
			v5 = nox_xxx_getNextPlayerUnit_4DA7F0(v5)
			if v5 == nil {
				break
			}
		}
		v0 = v10
	}
	nox_xxx_setGameFlags_40A4D0(8)
	if v1 != 0 {
		nox_xxx_netSendDMTeamWinner_4D8BF0(nil, 1)
		return
	}
	if v11 != nil {
		nox_xxx_netSendDMWinner_4D8B90((*server.Object)(v11), 1)
		return
	}
	if v0 != nil {
		nox_xxx_netSendDMTeamWinner_4D8BF0(v0, 1)
	}
}
func sub_5096F0() int32 {
	var (
		result int32
		v1     *uint8
		v2     int32
		v3     int32
		v4     int32
		v5     int32
		v6     [5]byte
	)
	result = sub_40A1A0()
	if result == 0 {
		return result
	}
	if noxflags.HasGame(4096) {
		v1 = sub_4E8E50()
		nox_xxx_mapLoad_4D2450(v1)
		nox_xxx_netPrintLineToAll_4DA390(internCStr("chklimit.c:AutoExitToNextMap"))
		v2 = nox_xxx_getFirstPlayerUnit_4DA7C0()
		if v2 != 0 {
			for {
				v3 = int32(*(*uint32)(unsafe.Add(v2, 748)))
				if *(*uint32)(unsafe.Add(v3, 312)) == 0 && *(*uint32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v3, 276)), 4792)) == 1 {
					sub_4D60E0(v2)
				}
				v2 = nox_xxx_getNextPlayerUnit_4DA7F0((*server.Object)(v2))
				if v2 == 0 {
					break
				}
			}
			return sub_40A1F0(0)
		}
		return sub_40A1F0(0)
	}
	if sub_40A300() == 0 {
		if noxflags.HasGame(96) {
			sub_5099B0()
			return sub_40A1F0(0)
		}
		if noxflags.HasGame(272) {
			sub_5098A0()
			return sub_40A1F0(0)
		}
		if noxflags.HasGame(1024) {
			sub_5095E0()
		}
		return sub_40A1F0(0)
	}
	nox_xxx_setGameFlags_40A4D0(0x4000000)
	v6[0] = 154
	v4 = nox_xxx_getFirstPlayerUnit_4DA7C0()
	if v4 == 0 {
		return sub_40A1F0(0)
	}
	for {
		v5 = int32(*(*uint32)(unsafe.Add(v4, 748)))
		if v5 != 0 {
			*(*uint16)(unsafe.Pointer(&v6[1])) = uint16(int16(int64(*(*float32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v5, 276)), 3632)))))
			*(*uint16)(unsafe.Pointer(&v6[3])) = uint16(int16(int64(*(*float32)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v5, 276)), 3636)))))
			nox_netlist_addToMsgListCli_40EBC0(int32(*(*uint8)(unsafe.Add(*(*unsafe.Pointer)(unsafe.Add(v5, 276)), 2064))), 1, &v6[0], 5)
		}
		nox_xxx_aud_501960(582, (*server.Object)(v4), 2, int32(*(*uint32)(unsafe.Add(v4, 36))))
		v4 = nox_xxx_getNextPlayerUnit_4DA7F0((*server.Object)(v4))
		if v4 == 0 {
			break
		}
	}
	return sub_40A1F0(0)
}
