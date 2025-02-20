package legacy

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/client/gui"
)

var (
	Sub_44A4A0                        func() int
	Nox_xxx_dialogMsgBoxCreate_449A10 func(win *gui.Window, title, text string, a4 gui.DialogFlags, a5, a6 func())
	Sub_449E00                        func(a1 string) int
	Sub_449E30                        func(a1 string) int
	Sub_449E60                        func(a1 int8) int
	Sub_449EA0                        func(a1 gui.DialogFlags)
	Sub_44A4E0                        func() int
	Sub_44A4B0                        func()
	Sub_44A360                        func(a1 int)
)

// sub_44A4A0
func sub_44A4A0() int { return Sub_44A4A0() }

// nox_xxx_dialogMsgBoxCreate_449A10
func nox_xxx_dialogMsgBoxCreate_449A10(win *gui.Window, title, text *wchar2_t, a4 int, a5, a6 func()) unsafe.Pointer {
	Nox_xxx_dialogMsgBoxCreate_449A10(win, GoWString(title), GoWString(text), gui.DialogFlags(a4), a5, a5)
	return nil
}

// sub_449E00
func sub_449E00(a1 *wchar2_t) int { return Sub_449E00(GoWString(a1)) }

// sub_449E30
func sub_449E30(a1 *wchar2_t) int { return Sub_449E30(GoWString(a1)) }

// sub_449E60
func sub_449E60(a1 int8) int32 { return int32(Sub_449E60(a1)) }

// sub_449EA0
func sub_449EA0(a1 int) { Sub_449EA0(gui.DialogFlags(a1)) }

// sub_44A4E0
func sub_44A4E0() int { return Sub_44A4E0() }

// sub_44A4B0
func sub_44A4B0() { Sub_44A4B0() }

// sub_44A360
func sub_44A360(a1 int) { Sub_44A360(a1) }

func Sub_41DA70(a1, a2 int) {
	sub_41DA70(int32(a1), int16(a2))
}
func Sub_445C20() {
	sub_445C20()
}
