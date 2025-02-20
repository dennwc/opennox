package legacy

import (
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/strman"

	"github.com/noxworld-dev/opennox/v1/client/gui"
)

var (
	Nox_video_setMenuOptions func(win *gui.Window)
	Nox_gui_menu_proc_ext    func(id int) int
	Sub_4A19F0               func(name strman.ID)
	Sub_4AAA10               func() int
	Sub_4C3A90               func(a1, a2 int, a3 unsafe.Pointer, a4 int) int
	Sub_4CBE70               func(a1, a2 int, a3 unsafe.Pointer, a4 int) int
	Sub_4A1A40               func(a1 int)
)

// nox_video_setMenuOptions
func nox_video_setMenuOptions(cwin *gui.Window) { Nox_video_setMenuOptions(cwin) }

// nox_gui_menu_proc_ext
func nox_gui_menu_proc_ext(id int32) int32 { return int32(Nox_gui_menu_proc_ext(int(id))) }

// sub_4A19F0
func sub_4A19F0(name *byte) { Sub_4A19F0(strman.ID(GoString(name))) }

// sub_4AAA10
func sub_4AAA10() int { return Sub_4AAA10() }

// sub_4C3A90
func sub_4C3A90(win *gui.Window, a2, a3, a4 uintptr) uintptr {
	return uintptr(Sub_4C3A90(int(uintptr(win.C())), int(a2), unsafe.Pointer(a3), int(a4)))
}

// sub_4CBE70
func sub_4CBE70(win *gui.Window, a2, a3, a4 uintptr) uintptr {
	return uintptr(Sub_4CBE70(int(uintptr(win.C())), int(a2), unsafe.Pointer(a3), int(a4)))
}

// sub_4A1A40
func sub_4A1A40(a1 int) {
	Sub_4A1A40(a1)
}

func Sub_42CD90() {
	sub_42CD90()
}

func Sub_4CBD30() {
	sub_4CBD30()
}

func Sub_430AA0(v int) {
	sub_430AA0(int32(v))
}

func Sub_4C35B0(v int) {
	sub_4C35B0(int32(v))
}

func Get_nox_wnd_xxx_1309740() *gui.Anim {
	return nox_wnd_xxx_1309740
}

func Get_dword_5d4594_1309720() *gui.Window {
	return AsWindowP(unsafe.Pointer(dword_5d4594_1309720))
}
