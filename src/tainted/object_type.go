package tainted

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

func getObjectTypeByID(id string) *ObjectType {
	cstr := CString(id)
	defer StrFree(cstr)
	p := C.nox_xxx_objectTypeByID_4E3830(cstr)
	if p == nil {
		return nil
	}
	return asObjectType(unsafe.Pointer(p))
}

func getObjectTypeByInd(ind int) *ObjectType {
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
	return asObjectType(unsafe.Pointer(p))
}

type ObjectType C.nox_objectType_t

func asObjectType(p unsafe.Pointer) *ObjectType {
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

func (t *ObjectType) String() string {
	return fmt.Sprintf("ObjectType(%d,%q)", t.Ind(), t.ID())
}

func (t *ObjectType) CreateObject(p types.Pointf) script.Object {
	cobj := C.nox_xxx_newObjectWithType_4E3470(t.C())
	if cobj == nil {
		return nil
	}
	C.nox_xxx_createAt_4DAA50(cobj, 0, C.float(p.X), C.float(p.Y))
	obj := asObjectC(cobj)
	if obj.Class().Has(object.MaskUnits) {
		return obj.AsUnit()
	}
	return obj
}
