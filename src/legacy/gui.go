package legacy

import (
	"os"
	"unsafe"

	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/noxworld-dev/opennox-lib/log"

	"github.com/noxworld-dev/opennox/v1/client/gui"
)

var (
	guiLog   = log.New("gui")
	guiDebug = os.Getenv("NOX_DEBUG_GUI") == "true"
)

var (
	Nox_client_gui_set_flag_815132 func(v int)
	Nox_client_onClientStatusA     func(v int)
	Nox_client_setRenderGUI        func(v int)
	Nox_client_getRenderGUI        func() int
)

var _ = [1]struct{}{}[332-unsafe.Sizeof(gui.WindowData{})]

func asWindowDataP(data unsafe.Pointer) *gui.WindowData {
	return (*gui.WindowData)(data)
}

// nox_client_gui_set_flag_815132
func nox_client_gui_set_flag_815132(v int32) { Nox_client_gui_set_flag_815132(int(v)) }

// nox_client_onClientStatusA
func nox_client_onClientStatusA(v int32) { Nox_client_onClientStatusA(int(v)) }

// nox_client_setRenderGUI
func nox_client_setRenderGUI(v int32) { Nox_client_setRenderGUI(int(v)) }

// nox_client_getRenderGUI
func nox_client_getRenderGUI() int32 { return int32(Nox_client_getRenderGUI()) }

// nox_xxx_wndGetFocus_46B4F0
func nox_xxx_wndGetFocus_46B4F0() *gui.Window {
	return (*gui.Window)(GetClient().Cli().GUI.Focused().C())
}

// nox_xxx_windowFocus_46B500
func nox_xxx_windowFocus_46B500(win *gui.Window) int32 {
	GetClient().Cli().GUI.Focus(win)
	return 0
}

// nox_client_getWin1064916_46C720
func nox_client_getWin1064916_46C720() *gui.Window {
	return (*gui.Window)(GetClient().Cli().GUI.WinYYY.C())
}

// nox_xxx_wndSetCaptureMain_46ADC0
func nox_xxx_wndSetCaptureMain_46ADC0(win *gui.Window) int32 {
	if !win.Capture(true) {
		return -4
	}
	return 0
}

// nox_xxx_wndClearCaptureMain_46ADE0
func nox_xxx_wndClearCaptureMain_46ADE0(win *gui.Window) int32 {
	win.Capture(false)
	return 0
}

// nox_xxx_wndGetCaptureMain_46AE00
func nox_xxx_wndGetCaptureMain_46AE00() *gui.Window {
	return (*gui.Window)(GetClient().Cli().GUI.Captured().C())
}

// nox_gui_draw
func nox_gui_draw() {
	GetClient().Cli().GUI.Draw()
}

// nox_color_rgb_4344A0
func nox_color_rgb_4344A0(r, g, b int32) uint32 {
	return noxcolor.RGB5551Color(byte(r), byte(g), byte(b)).Color32()
}

// nox_set_color_rgb_434430
func nox_set_color_rgb_434430(r, g, b int32) {
	GetClient().R2().Data().SetColor2(noxcolor.RGB5551Color(byte(r), byte(g), byte(b)))
}

// nox_xxx_wndWddSetTooltip_46B000
func nox_xxx_wndWddSetTooltip_46B000(draw *gui.WindowData, str *wchar2_t) {
	sm := GetClient().Cli().Strings()
	if str == nil {
		draw.SetTooltip(sm, "")
		return
	}
	draw.SetTooltip(sm, GoWString(str))
}

func sub_46B120(a1, a2 *gui.Window) {
	a1.SetParent(a2)
}

func Sub_46A4A0() int {
	return int(sub_46A4A0())
}

func Nox_xxx_wndEditProc_487D70(a1 *gui.Window, ev gui.WindowEvent) gui.RawEventResp {
	a2 := ev.EventCode()
	a3, a4 := ev.EventArgsC()
	return gui.RawEventResp(nox_xxx_wndEditProc_487D70(a1, uintptr(a2), a3, a4))
}

func Nox_gui_xxx_check_446360() int {
	return int(nox_gui_xxx_check_446360())
}
