package nox

/*
#include "server__script__script.h"
#include "server__script__internal.h"
#include "GAME4_1.h" // for nox_xxx_scriptPrepareFoundUnit_511D70 and nox_xxx_script_511C50
extern int nox_script_count_xxx_1599640;
extern nox_script_xxx_t* nox_script_arr_xxx_1599636;
*/
import "C"
import (
	"encoding/binary"
	"strings"
	"unsafe"

	"nox/v1/server/script"
)

var (
	activators struct {
		lastID uint32
		head   *Activator
	}
)

type Activator struct {
	frame     uint32
	callback  uint32
	arg       uint32
	id        uint32
	trigger   *Object
	caller    *Object
	triggerID uint32
	callerID  uint32
	next      *Activator
	prev      *Activator
}

func nox_script_activatorNextHandle_51AD20() uint32 {
	activators.lastID++
	id := activators.lastID
	if activators.lastID > 32000 {
		id = 1
		activators.lastID = 1
	}
	return id
}

//export nox_script_activatorCancel_51AD60
func nox_script_activatorCancel_51AD60(id C.int) C.bool {
	return C.bool(nox_script_activatorCancel(uint32(id)))
}

func nox_script_activatorCancel(id uint32) bool {
	for it := activators.head; it != nil; it = it.next {
		if it.id == id {
			nox_script_activatorDoneNext_51AD90(it)
			return true
		}
	}
	return false
}

//export nox_script_activatorCancelAll_51AC60
func nox_script_activatorCancelAll_51AC60() {
	activators.head = nil
}

func newScriptTimer(df int, callback, arg uint32) uint32 {
	act := &Activator{
		id:       nox_script_activatorNextHandle_51AD20(),
		frame:    gameFrame() + uint32(df),
		callback: callback, arg: arg,
	}
	nox_script_activator_append(act)
	return act.id
}

//export nox_script_activatorTimer_51ACA0
func nox_script_activatorTimer_51ACA0(df, callback, arg C.int) {
	id := newScriptTimer(int(df), uint32(callback), uint32(arg))
	nox_script_push(id)
}

func nox_script_activatorClearObj(obj *Object) {
	for it := activators.head; it != nil; {
		if it.trigger == obj {
			it = nox_script_activatorDoneNext_51AD90(it)
		} else {
			if it.caller == obj {
				it.caller = nil
			}
			it = it.next
		}
	}
}

//export nox_script_activatorClearObj_51AE60
func nox_script_activatorClearObj_51AE60(obj *C.nox_object_t) {
	nox_script_activatorClearObj(asObjectC(obj))
}

func nox_script_activator_append(act *Activator) {
	var last *Activator
	for it := activators.head; it != nil; it = it.next {
		last = it
	}
	if last != nil {
		last.next = act
		act.prev = last
	} else {
		activators.head = act
		act.prev = nil
	}
}

//export nox_script_activatorSave_51AEA0
func nox_script_activatorSave_51AEA0() C.int {
	var buf [4]byte
	binary.LittleEndian.PutUint16(buf[:], 1)
	cryptFileReadWrite(buf[:2])
	binary.LittleEndian.PutUint32(buf[:], gameFrame())
	cryptFileReadWrite(buf[:4])

	cnt := 0
	for it := activators.head; it != nil; it = it.next {
		cnt++
	}
	binary.LittleEndian.PutUint32(buf[:], uint32(cnt))
	cryptFileReadWrite(buf[:4])
	for it := activators.head; it != nil; it = it.next {
		binary.LittleEndian.PutUint32(buf[:], it.frame)
		cryptFileReadWrite(buf[:4])
		binary.LittleEndian.PutUint32(buf[:], uint32(it.callback))
		cryptFileReadWrite(buf[:4])
		binary.LittleEndian.PutUint32(buf[:], it.arg)
		cryptFileReadWrite(buf[:4])
		oid := 0
		if it.trigger != nil {
			oid = it.trigger.ScriptID()
		}
		binary.LittleEndian.PutUint32(buf[:], uint32(oid))
		cryptFileReadWrite(buf[:4])
		oid = 0
		if it.caller != nil {
			oid = it.caller.ScriptID()
		}
		binary.LittleEndian.PutUint32(buf[:], uint32(oid))
		cryptFileReadWrite(buf[:4])
	}
	return 1
}

//export nox_script_activatorLoad_51AF80
func nox_script_activatorLoad_51AF80() C.int {
	var buf [4]byte
	cryptFileReadWrite(buf[:2])
	vers := binary.LittleEndian.Uint16(buf[:])
	if vers > 1 || vers <= 0 {
		return 0
	}
	cryptFileReadWrite(buf[:4])
	saveFrame := binary.LittleEndian.Uint32(buf[:])
	cryptFileReadWrite(buf[:4])
	cnt := int(binary.LittleEndian.Uint32(buf[:]))
	for i := 0; i < cnt; i++ {
		cryptFileReadWrite(buf[:4])
		frame := binary.LittleEndian.Uint32(buf[:])
		cryptFileReadWrite(buf[:4])
		callback := binary.LittleEndian.Uint32(buf[:])
		cryptFileReadWrite(buf[:4])
		arg := binary.LittleEndian.Uint32(buf[:])
		cryptFileReadWrite(buf[:4])
		trigger := binary.LittleEndian.Uint32(buf[:])
		cryptFileReadWrite(buf[:4])
		caller := binary.LittleEndian.Uint32(buf[:])

		act := &Activator{
			id:       nox_script_activatorNextHandle_51AD20(),
			frame:    gameFrame() + (frame - saveFrame),
			callback: callback, arg: arg,
			triggerID: trigger, callerID: caller,
		}
		nox_script_activator_append(act)
	}
	return 1
}

func nox_script_activatorRun_51ADF0() {
	scripts := unsafe.Slice((*C.nox_script_xxx_t)(unsafe.Pointer(C.nox_script_arr_xxx_1599636)), int(C.nox_script_count_xxx_1599640))
	for it := activators.head; it != nil; {
		if it.frame > gameFrame() {
			it = it.next
		} else {
			callback := it.callback
			caller := it.caller
			trigger := it.trigger
			if scripts[callback].size_28 != 0 {
				nox_script_push(it.arg)
			}
			it = nox_script_activatorDoneNext_51AD90(it)
			C.nox_script_callByIndex_507310(C.int(callback), unsafe.Pointer(caller.CObj()), unsafe.Pointer(trigger.CObj()))
		}
	}
}

//export nox_script_activatorResolveObjs_51B0C0
func nox_script_activatorResolveObjs_51B0C0() {
	for it := activators.head; it != nil; it = it.next {
		if it.triggerID != 0 {
			it.trigger = nox_server_scriptValToObjectPtr(int(it.triggerID))
			it.triggerID = 0
		}
		if it.callerID != 0 {
			it.caller = nox_server_scriptValToObjectPtr(int(it.callerID))
			it.callerID = 0
		}
	}
}

func nox_script_activatorDoneNext_51AD90(act *Activator) *Activator {
	next := act.next
	if next != nil {
		next.prev = act.prev
	}
	if prev := act.prev; prev != nil {
		prev.next = next
	}

	if act == activators.head {
		activators.head = next
	}
	*act = Activator{}
	return next
}

func nox_script_push(v uint32) {
	C.nox_script_push(C.int(v))
}

//export nox_script_callOnEvent
func nox_script_callOnEvent(cevent *C.char, a1, a2 unsafe.Pointer) {
	if a1 != nil || a2 != nil { // these are never set to anything
		panic("unexpected argument to nox_script_callOnEvent")
	}
	event := script.EventType(GoString(cevent))
	scriptOnEvent(event)
}

func noxscriptOnEvent(event script.EventType) {
	for i := 0; i < int(C.nox_script_count_xxx_1599640); i++ {
		sc := (*C.nox_script_xxx_t)(unsafe.Add(unsafe.Pointer(C.nox_script_arr_xxx_1599636), unsafe.Sizeof(C.nox_script_xxx_t{})*uintptr(i)))
		name := GoString(sc.field_0)
		if strings.HasPrefix(name, string(event)) {
			C.nox_script_callByIndex_507310(C.int(i), nil, nil)
		}
	}
}

//export nox_server_scriptValToObjectPtr_511B60
func nox_server_scriptValToObjectPtr_511B60(val C.int) *C.nox_object_t {
	return nox_server_scriptValToObjectPtr(int(val)).CObj()
}

func nox_server_scriptValToObjectPtr(val int) *Object {
	if val == -1 {
		obj := asObject(C.nox_script_get_caller())
		if obj == nil || (obj.field_4&0x20) != 0 {
			return nil
		}
		return obj
	}
	if val == -2 {
		obj := asObject(C.nox_script_get_trigger())
		if obj == nil || (obj.field_4&0x20) != 0 {
			return nil
		}
		return obj
	}
	if obj := asObjectC(C.nox_xxx_script_511C50(C.int(val))); obj != nil {
		return obj
	}

	for obj := firstServerObject(); obj != nil; obj = obj.Next() {
		if (obj.field_4&0x20) == 0 && obj.ScriptID() == val {
			C.nox_xxx_scriptPrepareFoundUnit_511D70(obj.CObj())
			return obj
		}
		for sub := obj.FirstItem(); sub != nil; sub = sub.NextItem() {
			if (sub.field_4&0x20) == 0 && sub.ScriptID() == val {
				C.nox_xxx_scriptPrepareFoundUnit_511D70(sub.CObj())
				return sub
			}
		}
	}
	for obj := firstServerObjectUninited(); obj != nil; obj = obj.Next() {
		if (obj.field_4&0x20) == 0 && obj.ScriptID() == val {
			C.nox_xxx_scriptPrepareFoundUnit_511D70(obj.CObj())
			return obj
		}
	}
	return nil
}
