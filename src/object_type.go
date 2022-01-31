package nox

/*
#include "GAME3_2.h"
#include "GAME3_3.h"
*/
import "C"
import (
	"fmt"
	"math"
	"unsafe"

	"nox/v1/common/object"
	"nox/v1/common/types"
	"nox/v1/server/script"
)

func (s *Server) getObjectTypeByID(id string) *ObjectType { // nox_xxx_objectTypeByID_4E3830
	cstr := CString(id)
	defer StrFree(cstr)
	p := C.nox_xxx_objectTypeByID_4E3830(cstr)
	if p == nil {
		return nil
	}
	return asObjectType(p)
}

func (s *Server) getObjectTypeID(id string) int { // nox_xxx_getNameId_4E3AA0
	typ := s.getObjectTypeByID(id)
	if typ == nil {
		return 0
	}
	return typ.Ind()
}

func (s *Server) getObjectTypeByInd(ind int) *ObjectType {
	if ind == math.MaxUint16 {
		return nil
	}
	if ind < 0 || ind >= int(C.nox_xxx_unitDefGetCount_4E3AC0()) {
		return nil
	}
	p := C.nox_xxx_objectTypeByInd_4E3B70(C.int(ind))
	if p == nil {
		return nil
	}
	return asObjectType(p)
}

func (s *Server) getObjectTypes() (out []*ObjectType) {
	for i := 0; i < int(C.nox_xxx_unitDefGetCount_4E3AC0()); i++ {
		typ := s.getObjectTypeByInd(i)
		if typ == nil {
			continue
		}
		out = append(out, typ)
	}
	return
}

func (s *Server) newObjectByTypeID(id string) *Object { // nox_xxx_newObjectByTypeID_4E3810
	typ := s.getObjectTypeByID(id)
	if typ == nil {
		return nil
	}
	cobj := C.nox_xxx_newObjectWithType_4E3470(typ.C())
	if cobj == nil {
		return nil
	}
	return asObjectC(cobj)
}

type ObjectType C.nox_objectType_t

func asObjectType(p *C.nox_objectType_t) *ObjectType {
	return asObjectTypeP(unsafe.Pointer(p))
}

func asObjectTypeP(p unsafe.Pointer) *ObjectType {
	return (*ObjectType)(p)
}

func (t *ObjectType) C() *C.nox_objectType_t {
	return (*C.nox_objectType_t)(unsafe.Pointer(t))
}

func (t *ObjectType) ID() string {
	return GoString(t.id)
}

func (t *ObjectType) Ind() int {
	if t == nil {
		return 0
	}
	return int(t.ind)
}

func (t *ObjectType) Class() object.Class {
	return object.Class(t.obj_class)
}

func (t *ObjectType) ArmorClass() object.ArmorClass {
	if !t.Class().Has(object.ClassArmor) {
		return 0
	}
	return object.ArmorClass(t.field_7)
}

func (t *ObjectType) String() string {
	return fmt.Sprintf("ObjectType(%d,%q)", t.Ind(), t.ID())
}

func (t *ObjectType) CreateObject(p types.Pointf) script.Object {
	cobj := C.nox_xxx_newObjectWithType_4E3470(t.C())
	if cobj == nil {
		return nil
	}
	obj := asObjectC(cobj)
	nox_xxx_createAt_4DAA50(obj, nil, p)
	if obj.Class().Has(object.MaskUnits) {
		return obj.AsUnit()
	}
	return obj
}

func (t *ObjectType) updateDataRaw() []byte {
	return unsafe.Slice((*byte)(t.data_update), int(t.data_update_size))
}
