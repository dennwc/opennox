package tainted

/*
#include "defs.h"
#include "client__gui__window.h"
int sub_43BE40(int a1);
int sub_4A24F0();

static int nox_gui_call_intvoid_go(int (*f)(void)) { return f(); }
static void nox_gui_call_void2_go(void (*f)(void)) { f(); }
*/
import "C"
import (
	"unsafe"

	"nox/v1/common/alloc"
	"nox/v1/common/env"
	"nox/v1/common/types"
)

var guiAnimSpeed = 1

func init() {
	if env.IsDevMode() || env.IsE2E() {
		guiAnimSpeed = 5
	}
}

type guiAnimState int

func (s guiAnimState) C() C.nox_gui_anim_state {
	return C.nox_gui_anim_state(s)
}

const (
	NOX_GUI_ANIM_IN_DONE  = guiAnimState(0)
	NOX_GUI_ANIM_OUT_DONE = guiAnimState(1)
	NOX_GUI_ANIM_OUT      = guiAnimState(2)
	NOX_GUI_ANIM_IN       = guiAnimState(3)
)

var guiAnimationsHead *guiAnim

func asGUIAnim(p *C.nox_gui_animation) *guiAnim {
	return (*guiAnim)(unsafe.Pointer(p))
}

type guiAnim C.nox_gui_animation

func (a *guiAnim) C() *C.nox_gui_animation {
	return (*C.nox_gui_animation)(unsafe.Pointer(a))
}

func (a *guiAnim) State() guiAnimState {
	if a == nil {
		return -1
	}
	return guiAnimState(a.state)
}

func (a *guiAnim) setState(s guiAnimState) {
	a.state = s.C()
}

func (a *guiAnim) Window() *Window {
	if a == nil {
		return nil
	}
	return asWindow(a.win)
}

func (a *guiAnim) Next() *guiAnim {
	if a == nil {
		return nil
	}
	return asGUIAnim(a.next)
}

func (a *guiAnim) Prev() *guiAnim {
	if a == nil {
		return nil
	}
	return asGUIAnim(a.prev)
}

func (a *guiAnim) Func12() int {
	return int(C.nox_gui_call_intvoid_go(a.field_12))
}

func (a *guiAnim) Func13() int {
	return int(C.nox_gui_call_intvoid_go(a.field_13))
}

func (a *guiAnim) doOut() {
	p := a.Window().Offs()
	p.X += int(a.out_dx) * guiAnimSpeed
	p.Y += int(a.out_dy) * guiAnimSpeed

	maxed := 0

	if mx := int(a.x2); a.out_dx >= 0 {
		if p.X >= mx {
			p.X = mx
			maxed++
		}
	} else {
		if p.X <= mx {
			p.X = mx
			maxed++
		}
	}
	if my := int(a.y2); a.out_dy >= 0 {
		if p.Y >= my {
			p.Y = my
			maxed++
		}
	} else {
		if p.Y <= my {
			p.Y = my
			maxed++
		}
	}

	a.Window().SetPos(p)
	if maxed == 2 {
		a.setState(NOX_GUI_ANIM_OUT_DONE)
		C.sub_43BE40(1)
		if a.fnc_done_out != nil {
			C.nox_gui_call_intvoid_go(a.fnc_done_out)
		}
	}
}

func (a *guiAnim) doIn() {
	p := a.Window().Offs()
	p.X += int(a.in_dx) * guiAnimSpeed
	p.Y += int(a.in_dy) * guiAnimSpeed

	maxed := 0

	if mx := int(a.x1); a.in_dx >= 0 {
		if p.X >= mx {
			p.X = mx
			maxed++
		}
	} else {
		if p.X <= mx {
			p.X = mx
			maxed++
		}
	}
	if my := int(a.y1); a.in_dy >= 0 {
		if p.Y >= my {
			p.Y = my
			maxed++
		}
	} else {
		if p.Y <= my {
			p.Y = my
			maxed++
		}
	}

	a.Window().SetPos(p)
	if maxed == 2 {
		a.setState(NOX_GUI_ANIM_IN_DONE)
		C.sub_43BE40(0)
		if a.fnc_done_in != nil {
			C.nox_gui_call_void2_go(a.fnc_done_in)
		}
		C.sub_4A24F0()
	}
}

func guiAnimationStep() {
	var next *guiAnim
	for a := guiAnimationsHead; a != nil; a = next {
		next = a.Next()
		if a.State() == NOX_GUI_ANIM_OUT {
			a.doOut()
		} else if a.State() == NOX_GUI_ANIM_IN {
			a.doIn()
		}
	}
}

func nox_gui_findAnimationForDest_43C520(dest int) *guiAnim {
	for p := guiAnimationsHead; p != nil; p = p.Next() {
		if int(p.field_0) == dest {
			return p
		}
	}
	return nil
}

func newGUIAnimation(win *Window) *guiAnim {
	p := (*guiAnim)(alloc.Malloc(unsafe.Sizeof(guiAnim{})))
	p.win = win.C()
	p.next = guiAnimationsHead.C()
	if guiAnimationsHead != nil {
		guiAnimationsHead.prev = p.C()
	}
	guiAnimationsHead = p
	return p
}

//export nox_gui_freeAnimation_43C570
func nox_gui_freeAnimation_43C570(a *C.nox_gui_animation) {
	asGUIAnim(a).Free()
}

func (a *guiAnim) Free() {
	if a == nil {
		return
	}
	next := a.Next()
	if next != nil {
		next.prev = a.Prev().C()
	}

	prev := a.Prev()
	if prev != nil {
		prev.next = a.Next().C()
	} else {
		guiAnimationsHead = a.Next()
	}
	alloc.Free(unsafe.Pointer(a))
}

//export nox_gui_makeAnimation_43C5B0
func nox_gui_makeAnimation_43C5B0(win *Window, x1, y1, x2, y2, in_dx, in_dy, out_dx, out_dy int) *C.nox_gui_animation {
	a := newGUIAnimation(win)
	win.SetPos(types.Point{X: x2, Y: y2})
	a.x1 = C.int(x1)
	a.y1 = C.int(y1)
	a.x2 = C.int(x2)
	a.y2 = C.int(y2)
	a.in_dx = C.int(in_dx)
	a.in_dy = C.int(in_dy)
	a.out_dx = C.int(out_dx)
	a.out_dy = C.int(out_dy)
	a.setState(NOX_GUI_ANIM_IN)
	C.sub_43BE40(3)
	clientPlaySoundSpecial(922, 100)
	a.field_12 = nil
	a.fnc_done_out = nil
	a.fnc_done_in = nil
	return a.C()
}
