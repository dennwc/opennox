package tainted

/*
#include "defs.h"
#include "client__gui__window.h"
extern unsigned int dword_587000_87404;
extern unsigned int dword_5d4594_1064868;
extern unsigned int dword_5d4594_1316972;
extern nox_window* dword_5d4594_1064896;
extern unsigned int gameex_flags;
int sub_4BDFD0();
int  gameex_sendPacket(char* buf, int len, int smth);
char  mix_MouseKeyboardWeaponRoll(int playerObj, char a2);
int getFlagValueFromFlagIndex(signed int a1);
int  modifyWndInputHandler(int a1, int a2, int a3, int a4);
int  nox_xxx_clientUpdateButtonRow_45E110(int a1);
unsigned int*  nox_xxx_objGetTeamByNetCode_418C80(int a1);
void  nox_xxx_printCentered_445490(wchar_t* a1);
char  playerDropATrap(int playerObj);
*/
import "C"
import (
	"bufio"
	"encoding/binary"
	"math"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"nox/v1/client/input/keybind"
	"nox/v1/common/alloc"
	"nox/v1/common/datapath"
	noxflags "nox/v1/common/flags"
	"nox/v1/common/fs"
	"nox/v1/common/log"
	"nox/v1/common/memmap"
	"nox/v1/common/types"
)

//export gameexSomeWeirdCheckFixmePlease
func gameexSomeWeirdCheckFixmePlease() C.bool {
	// FIXME: no idea what is supposed to do... just checking if both are nil?
	//        previously checked in asm: (cmp ds:6D8555, eax)
	//        although we now know that offsets in Mix was wrong compared to our base binary
	return (uintptr(unsafe.Pointer(C.dword_5d4594_1064896))>>8)|(uintptr(unsafe.Pointer(nox_win_activeWindow_1064900))<<24) == 0
}

var modifyWndPntr *Window
var gameex = struct {
	Log        *log.Logger
	configPath string
}{
	Log: log.New("gameex"),
}

func gameexReadConfig(path string) error {
	path = datapath.Path(path)
	gameex.configPath = path
	gameex.Log.Println("reading config", path)
	f, err := fs.Open(path)
	if os.IsNotExist(err) {
		gameex.Log.Println("no config file")
		return nil
	} else if err != nil {
		return err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) == 0 || line[0] == '-' || line[0] == '#' {
			continue
		}
		sub := strings.SplitN(line, "=", 2)
		if len(sub) < 2 {
			gameex.Log.Printf("malformed option: %q", line)
			continue
		}
		key := strings.TrimSpace(sub[0])
		val := strings.TrimSpace(sub[1])
		switch key {
		case "AUTO_SHIELD":
			v, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			if v {
				C.gameex_flags |= 0x2
			} else {
				C.gameex_flags &^= 0x2
			}
		case "GREAT_SWORD_BLOKING_WALK", "GREAT_SWORD_BLOCKING_WALK":
			v, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			if v {
				C.gameex_flags |= 0x4
			} else {
				C.gameex_flags &^= 0x4
			}
		case "MOUSE_KEYBOARD_ROLL":
			v, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			if v {
				C.gameex_flags |= 0x8
			} else {
				C.gameex_flags &^= 0x8
			}
		case "BERSERKER_SHIED_BLOCK", "BERSERKER_SHIELD_BLOCK":
			v, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			if v {
				C.gameex_flags |= 0x10
			} else {
				C.gameex_flags &^= 0x10
			}
		case "EXTENSION_MESSAGES":
			v, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			if v {
				C.gameex_flags |= 0x20
			} else {
				C.gameex_flags &^= 0x20
			}
		case "PANEL1", "PANEL2", "PANEL3", "PANEL4", "PANEL5", "TRAPKEY":
		default:
			gameex.Log.Printf("unrecognized option: %q", key)
		}
	}
	return sc.Err()
}

func gameex_makeExtensionPacket(buf []byte, opcode uint16, needsPlayer bool) {
	binary.LittleEndian.PutUint16(buf[0:], 0xF13A) // extension packet code
	binary.LittleEndian.PutUint16(buf[2:], opcode)
	if needsPlayer {
		binary.LittleEndian.PutUint32(buf[4:], memmap.Uint32(0x85319C, 0)) // playerNetCode
	}
}

var wndEntryNames = [5][35]uint16{
	{119, 97, 114, 114, 105, 111, 114, 115, 32, 108, 105, 107, 101, 32, 115, 104, 105, 101, 108, 100, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{119, 97, 108, 107, 105, 110, 103, 32, 119, 105, 116, 104, 32, 115, 119, 111, 114, 100, 32, 98, 108, 111, 99, 107, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{119, 101, 97, 112, 111, 110, 32, 114, 111, 108, 108, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{115, 104, 105, 101, 108, 100, 32, 38, 32, 98, 101, 114, 115, 101, 114, 107, 101, 114, 32, 98, 108, 111, 99, 107, 105, 110, 103, 0, 0, 0, 0, 0, 0, 0, 0},
	{101, 120, 116, 101, 110, 115, 105, 111, 110, 32, 109, 101, 115, 115, 97, 103, 101, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func gameexDropTrap() {
	if noxflags.HasGame(516) {
		if C.dword_5d4594_1064868 != 0 || nox_win_unk3 != nil {
			return
		}
		if noxflags.HasGame(1) { // checkGameFlags isServer
			v9 := C.nox_xxx_objGetTeamByNetCode_418C80(C.int(memmap.Uint32(0x85319C, 0)))
			C.playerDropATrap(C.int(uintptr(unsafe.Pointer(v9)) - 12*4)) // TODO: this doesn't look right
		} else {
			// TODO: this currently relies on extension packets, which should not be required for this
			//       it can be done the "natural way": find the trap in the client-side data structures
			//       and ask server to drop that item, as the client does when doing the same manually
			buf := alloc.Bytes(10)
			gameex_makeExtensionPacket(buf, 9, true)
			C.gameex_sendPacket((*C.char)(unsafe.Pointer(&buf[0])), 8, 0)
			alloc.FreeBytes(buf)
		}
	}
}

func gameexOnKeyboardPress(kcode keybind.Key) {
	if ((C.gameex_flags>>3)&1 != 0) && (kcode == keybind.KeyLBracket || kcode == keybind.KeyRBracket) {
		v8 := byte(bool2int(kcode == keybind.KeyLBracket))
		// checks some gameFlags that are yet undiscovered
		if noxflags.HasGame(0x204) {
			if C.dword_5d4594_1064868 != 0 || nox_win_unk3 != nil {
				return
			}
			if noxflags.HasGame(1) { // isServer
				if u := HostPlayerUnit(); u != nil && C.mix_MouseKeyboardWeaponRoll(C.int(uintptr(unsafe.Pointer(u.CObj()))), C.char(v8)) != 0 {
					clientPlaySoundSpecial(895, 100)
				}
			} else {
				buf := alloc.Bytes(10)
				gameex_makeExtensionPacket(buf, 0, true)
				buf[8] = v8 | 0x10 // TODO: should it be just v8?
				C.gameex_sendPacket((*C.char)(unsafe.Pointer(&buf[0])), 9, 0)
				alloc.FreeBytes(buf)
			}
		}
	}
	if kcode == keybind.KeyF8 { // TODO: should be configurable, or actually just set defaults and remove this window
		if !noxflags.HasGame(1) {
			nox_xxx_printCentered_445490("only server can change these options")
			clientPlaySoundSpecial(231, 100)
			return
		}
		if !noxflags.HasGame(516) {
			return
		}
		clientPlaySoundSpecial(921, 100)
		if modifyWndPntr != nil {
			destroyGameExWindow()
			modifyWndPntr = nil
		} else {
			modifyWndPntr = newWindowFromString(gameexModifWnd, unsafe.Pointer(C.modifyWndInputHandler))
			if modifyWndPntr == nil {
				return
			}
			if noxflags.HasGame(512) {
				modifyWndPntr.ChildByID(1938).Hide()
				v11 := modifyWndPntr.ChildByID(1524)
				C.nox_xxx_wnd_46ABB0(C.int(uintptr(unsafe.Pointer(v11.C()))), 0)
			}
			sub46B120(modifyWndPntr, nil)
			a2b := modifyWndPntr.ChildByID(1981)
			for i := 0; i < 5; i++ {
				wstr := GoWStringSlice(wndEntryNames[i][:])
				id := uint(1520 + i)
				a2b.Func94(16397, uintptr(unsafe.Pointer(internWStr(wstr))), math.MaxUint32)
				if uint32(C.getFlagValueFromFlagIndex(C.int(id)-1519))&uint32(C.gameex_flags) != 0 {
					v14 := modifyWndPntr.ChildByID(id)
					v14.DrawData().field_0 |= 0x4
				} else {
					v15 := modifyWndPntr.ChildByID(id)
					v15.DrawData().field_0 &= 0xFFFFFFFB
				}
			}
		}
	}
}

func destroyGameExWindow() {
	C.nox_xxx_wnd_46C6E0(modifyWndPntr.C())
	modifyWndPntr.Destroy()
	modifyWndPntr = nil
}

//export modifyWndInputHandler
func modifyWndInputHandler(a1, a2, a3, a4 C.int) C.int {
	if a2 != 16391 {
		return 0
	}
	clientPlaySoundSpecial(766, 100)
	a3p := asWindow((*C.nox_window)(unsafe.Pointer(uintptr(a3))))
	switch a3p.ID() {
	case 1937:
		destroyGameExWindow()
	case 1938:
		if !noxflags.HasGame(512) {
			C.sub_4BDFD0()
			asWindow((*C.nox_window)(unsafe.Pointer(uintptr(C.dword_5d4594_1316972)))).SetPos(types.Point{X: 200, Y: 100})
		}
	case 1520:
		if (C.gameex_flags>>1)&1 != 0 {
			C.gameex_flags &= 0xFFFFFFFD
		} else {
			C.gameex_flags |= 0x2
		}
	case 1521:
		if (C.gameex_flags>>2)&1 != 0 {
			C.gameex_flags &= 0xFFFFFFFB
		} else {
			C.gameex_flags |= 0x4
		}
	case 1522:
		if (C.gameex_flags>>3)&1 != 0 {
			C.gameex_flags &= 0xFFFFFFF7
		} else {
			C.gameex_flags |= 0x8
		}
	case 1523:
		if (C.gameex_flags>>4)&1 != 0 {
			C.gameex_flags &= 0xFFFFFFEF
		} else {
			C.gameex_flags |= 0x10
		}
	case 1524:
		if (C.gameex_flags>>5)&1 != 0 {
			C.gameex_flags &= 0xFFFFFFDF
		} else {
			C.gameex_flags |= 0x20
		}
	}
	return 0
}

const gameexModifWnd = `
ENABLEDCOLOR = TRANSPARENT;
HILITECOLOR = 255 192 192;
DISABLEDCOLOR = 128 128 128;
SELECTEDCOLOR = 255 128 128;
TEXTCOLOR = 255 255 255;
BACKGROUNDCOLOR = 0 0 0;

WINDOW
  1930 50 55 200 150 USER;
  STATUS = ENABLED+ABOVE+BORDER;

  CHILD
	WINDOW
	  1520 8 22 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  STYLE = MOUSETRACK;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

	WINDOW
	  1521 8 32 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  STYLE = MOUSETRACK;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

	WINDOW
	  1522 8 42 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  STYLE = MOUSETRACK;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

	WINDOW
	  1523 8 52 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  STYLE = MOUSETRACK;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

	WINDOW
	  1524 8 62 15 10 CHECKBOX (Object allowed);
      STATUS = ENABLED+IMAGE;
	  SELECTEDIMAGE = UICheckBoxLit;
	  HILITEIMAGE = NULL;
	  BACKGROUNDIMAGE = UICheckBox;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UICheckBoxDis;
	  STYLE = MOUSETRACK;
	  TOOLTIP = objlst.wnd:ToggleTT;
	END

        WINDOW
          1937 65 120 70 20 PUSHBUTTON;
          STATUS = ENABLED+IMAGE;
          STYLE = MOUSETRACK+TABSTOP;
	  SELECTEDIMAGE = UIButtonLgLit;
	  HILITEIMAGE = UIButtonLgLit;
	  BACKGROUNDIMAGE = UIButtonLg;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UIButtonLgDis;
          TEXT = WolFind.wnd:Done;
       END

        WINDOW
          1938 115 80 70 20 PUSHBUTTON;
          STATUS = ENABLED+IMAGE;
          STYLE = MOUSETRACK+TABSTOP;
	  SELECTEDIMAGE = UIButtonLgLit;
	  HILITEIMAGE = UIButtonLgLit;
	  BACKGROUNDIMAGE = UIButtonLg;
	  ENABLEDIMAGE = NULL;
	  DISABLEDIMAGE = UIButtonLgDis;
          TEXT = servopts.wnd:Advanced;
       END

        WINDOW
          1981 22 21 150 60 SCROLLLISTBOX
          STATUS = ONE_LINE;
          DISABLEDCOLOR = TRANSPARENT;
          DATA = 26 10 0 1 0 0 1;
          FONT = default;
       END
END
`
