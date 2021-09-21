#include "client__gui__guicon.h"
#include "client__gui__window.h"

#include "common__telnet__telnetd.h"
#include "client__system__parsecmd.h"
#include "client__system__ctrlevnt.h"
#include "client__gui__gadgets__listbox.h"

#include "GAME1.h"
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2_2.h"
#include "GAME2_3.h"
#include "GAME3_2.h"

nox_window* nox_gui_console_win = 0;
nox_window* nox_gui_console_scrollbox = 0;
nox_window* nox_gui_console_input = 0;

int nox_gui_console_enabled = 0;
int nox_gui_console_locked = 0;
unsigned int nox_gui_console_translucent = 1;
int nox_gui_console_wantsPass = 0;
extern int dword_5d4594_3799524;

wchar_t nox_gui_console_password[64] = {0};

//----- (00450BE0) --------------------------------------------------------
void nox_gui_console_Enable_450BE0() { nox_gui_console_enabled = 1; }

//----- (00450BF0) --------------------------------------------------------
void nox_gui_console_Disable_450BF0() { nox_gui_console_enabled = 0; }

//----- (00450B70) --------------------------------------------------------
int nox_gui_console_Clear_450B70() {
	return nox_window_call_field_94(nox_gui_console_scrollbox, 16399, 0, 0);
}

//----- (00451410) --------------------------------------------------------
int nox_gui_console_flagXxx_451410() {
	return ((unsigned __int8)~nox_xxx_wndGetFlags_46ADA0(nox_gui_console_win) >> 4) & 1;
}

//----- (00450B20) --------------------------------------------------------
void nox_gui_console_Lock_450B20(wchar_t* pass) {
	if (pass) {
		if (*pass) {
			nox_wcscpy(nox_gui_console_password, pass);
			nox_gui_console_locked = 1;
		}
	}
}

//----- (00450B50) --------------------------------------------------------
void nox_gui_console_Unlock_450B50() {
	nox_gui_console_password[0] = 0;
	nox_gui_console_locked = 0;
	nox_gui_console_wantsPass = 0;
}

//----- (00450B90) --------------------------------------------------------
int nox_gui_console_Print_450B90(unsigned char cl, wchar_t* str) {
	if (nox_gui_console_enabled)
		nox_window_call_field_94(nox_gui_console_scrollbox, 16397, str, cl);
	if (nox_telnet_isActive_579740())
		nox_telnet_broadcast_579750(str);
	return 1;
}

//----- (00450C30) --------------------------------------------------------
void nox_gui_console_PrintOrError_450C30(unsigned char cl, wchar_t* str) {
	if (str) {
		nox_gui_console_Print_450B90(cl, str);
		return;
	}
	str = nox_strman_loadString_40F1D0("InternalError", 0, "C:\\NoxPost\\src\\Client\\Gui\\guicon.c", 97);
	if (str)
		nox_gui_console_Print_450B90(cl, str);
}

//----- (00450C00) --------------------------------------------------------
wchar_t nox_gui_console_printBuf[512] = {0};
int nox_gui_console_Printf_450C00(unsigned char cl, wchar_t* fmt, ...) {
	va_list va;
	va_start(va, fmt);
	nox_vswprintf(nox_gui_console_printBuf, fmt, va);
	return nox_gui_console_Print_450B90(cl, nox_gui_console_printBuf);
}

//----- (00450FD0) --------------------------------------------------------
int nox_gui_console_Enter_450FD0() {
	if (!*(_DWORD*)((char*)nox_gui_console_input->field_8 + 1044)) {
		wchar_t* line = nox_window_call_field_94(nox_gui_console_input, 16413, 0, 0);
		if (nox_gui_console_wantsPass && nox_gui_console_password[0]) {
			if (_nox_wcsicmp(line, nox_gui_console_password)) {
				wchar_t* v4 = nox_strman_loadString_40F1D0("INVALIDPASSWORD", 0, "C:\\NoxPost\\src\\Client\\Gui\\guicon.c", 124);
				nox_gui_console_PrintOrError_450C30(NOX_CONSOLE_WHITE, v4);
			} else {
				nox_gui_console_wantsPass = 0;
				wchar_t* v3 = nox_strman_loadString_40F1D0("ConsoleUnlocked", 0, "C:\\NoxPost\\src\\Client\\Gui\\guicon.c", 121);
				nox_gui_console_PrintOrError_450C30(NOX_CONSOLE_RED, (int)v3);
			}
			nox_window_call_field_94(nox_gui_console_input, 16414, (int)getMemAt(0x5D4594, 833744), 0);
		} else {
			wchar_t buf[256];
			nox_wcscpy(buf, L"> ");
			nox_wcscat(buf, line);
			nox_gui_console_Print_450B90(NOX_CONSOLE_WHITE, buf);
			if (line) {
				if (nox_common_getEngineFlag(NOX_ENGINE_FLAG_REPLAY_WRITE)) {
					nox_xxx_replaySaveConsole_4D33E0(line, nox_wcslen(line) + 1);
				}
				nox_server_parseCmdText_443C80(line, 0);
			}
			nox_window_call_field_94(nox_gui_console_input, 16414, (int)getMemAt(0x5D4594, 833748), 0);
		}
	}
	return 1;
}

//----- (00451350) --------------------------------------------------------
void nox_client_toggleConsole_451350() {
	if (!(nox_xxx_wndGetFlags_46ADA0(nox_gui_console_win) & NOX_WIN_HIDDEN)) {
		nox_gui_console_Hide_4512B0();
		return;
	}
	if (!nox_gui_xxx_check_446360()) {
		nox_xxx_wndShowModalMB_46A8C0(nox_gui_console_win);
		nox_gui_console_win->flags |= 8u;
		nox_gui_console_scrollbox->flags |= 8u;
		nox_gui_console_input->flags |= 8u;
		nox_gui_console_input->flags |= 1u;
		nox_xxx_windowFocus_46B500(nox_gui_console_input);
		if (nox_gui_console_locked) {
			nox_gui_console_Clear_450B70();
			wchar_t* v1 = nox_strman_loadString_40F1D0("ENTERPASSWORD", 0, "C:\\NoxPost\\src\\Client\\Gui\\guicon.c", 459);
			nox_gui_console_Printf_450C00(NOX_CONSOLE_WHITE, v1);
			nox_gui_console_wantsPass = 1;
		}
	}
}

//----- (00450F40) --------------------------------------------------------
#ifdef NOX_CGO
int  nox_xxx_consoleEditProc_450F40(void* a1, int a2, int a3, int a4);
#else // NOX_CGO
int  nox_xxx_consoleEditProc_450F40(_DWORD* a1, int a2, int a3, int a4) {
	_DWORD* v4; // eax

	if (a2 != 21)
		return nox_xxx_wndEditProc_487D70(a1, a2, a3, a4);
	v4 = nox_xxx_getBindKeysBuf_42CD70();
	if (v4) {
		while (v4[9] != 11 || a3 != *v4) {
			v4 = (_DWORD*)v4[19];
			if (!v4)
				goto LABEL_6;
		}
		if (a4 == 2) {
			nox_client_toggleConsole_451350();
			return 1;
		}
		return 1;
	}
LABEL_6:
	if (a3 == 1) {
		if (a4 == 2)
			nox_xxx_consoleEsc_49B7A0();
	} else {
		if (a3 != 28)
			return nox_xxx_wndEditProc_487D70(a1, a2, a3, a4);
		if (a4 == 2) {
			nox_gui_console_Enter_450FD0();
			return 1;
		}
	}
	return 1;
}
#endif // NOX_CGO

//----- (00450E80) --------------------------------------------------------
int nox_xxx_consoleWndFn_450E80() { return 0; }

//----- (00450E90) --------------------------------------------------------
int  nox_xxx_consoleWndFn_450E90(int a1, int a2) {
	int xLeft; // [esp+4h] [ebp-8h]
	int yTop;  // [esp+8h] [ebp-4h]

	nox_client_wndGetPosition_46AA60((_DWORD*)a1, &xLeft, &yTop);
	if (*(_BYTE*)(a1 + 4) & 0x80) {
		nox_client_drawImageAt_47D2C0(*(_DWORD*)(a2 + 24), xLeft, yTop);
		return 1;
	}
	if (*(_DWORD*)(a2 + 20) != 0x80000000) {
		if (nox_gui_console_translucent) {
			nox_client_drawRectFilledAlpha_49CF10(xLeft, yTop, *(_DWORD*)(a1 + 8), *(_DWORD*)(a1 + 12));
			return 1;
		}
		nox_client_drawSetColor_434460(*(_DWORD*)(a2 + 20));
		nox_client_drawRectFilledOpaque_49CE30(xLeft, yTop, *(_DWORD*)(a1 + 8), *(_DWORD*)(a1 + 12));
	}
	return 1;
}

//----- (00450C70) --------------------------------------------------------
nox_window* nox_gui_console_Create_450C70(int win_width, int win_height) {
	int v0 = *getMemU32Ptr(0x85B3FC, 956);
	*getMemU32Ptr(0x5D4594, 833704) = win_width - 1;
	*getMemU32Ptr(0x5D4594, 833708) = win_height / 2;
	nox_gui_console_win = nox_window_new(0, 56, 0, 0, win_width - 1, win_height / 2, 0);
	nox_window_set_all_funcs(nox_gui_console_win, nox_xxx_consoleWndFn_450E80, nox_xxx_consoleWndFn_450E90, 0);
	*(_DWORD*)((char*)nox_gui_console_win + 56) = *getMemU32Ptr(0x85B3FC, 952);

	char drawData[332];
	memset(drawData, 0, sizeof(drawData));
	*(_DWORD*)&drawData[36] = 0x80000000;
	*(_DWORD*)&drawData[52] = 0x80000000;
	*(_DWORD*)&drawData[24] = 0;
	*(_DWORD*)&drawData[48] = 0;
	*(_DWORD*)&drawData[28] = v0;
	*(_DWORD*)&drawData[68] = v0;
	nox_wcsncpy((wchar_t*)&drawData[72], nox_xxx_getNoxVer_401020(), 0x40u);
	*(_WORD*)&drawData[198] = 0;
	*(_DWORD*)&drawData[8] = 32;

	char scrollData[56];
	memset(scrollData, 0, sizeof(scrollData));
	*(_WORD*)&scrollData[2] = 0x0a;
	*(_WORD*)&scrollData[0] = 0x80;
	*(_DWORD*)&scrollData[4] = 1;
	*(_DWORD*)&scrollData[8] = 1;
	*(_DWORD*)&scrollData[12] = 1;
	*(_DWORD*)&scrollData[16] = 0;
	nox_gui_console_scrollbox = nox_gui_newScrollListBox_4A4310(nox_gui_console_win, 1152, 0, 0, *getMemIntPtr(0x5D4594, 833704),
									 *getMemU32Ptr(0x5D4594, 833708) - 20, drawData, scrollData);
	*(_DWORD*)&drawData[44] = v0;
	*(_DWORD*)&drawData[28] = v0;
	*(_DWORD*)&drawData[36] = v0;
	*(_DWORD*)&drawData[52] = v0;
	*(_DWORD*)&drawData[68] = v0;

	unsigned __int8 inpData[1056];
	*(_WORD*)&inpData[1042] = *getMemU16Ptr(0x5D4594, 833704);
	nox_wcscpy(inpData, (const wchar_t*)getMemAt(0x5D4594, 833736));
	*(_WORD*)&inpData[1040] = 128;
	*(_DWORD*)&inpData[1024] = 0;
	*(_DWORD*)&inpData[1028] = 0;
	*(_DWORD*)&inpData[1032] = 0;
	*(_DWORD*)&inpData[1036] = 0;

	nox_wcscpy((wchar_t*)&drawData[72], (const wchar_t*)getMemAt(0x5D4594, 833740));
	*(_DWORD*)&drawData[8] = 128;
	nox_gui_console_input = nox_gui_newEntryField_488500(nox_gui_console_win, 16393, 0, *getMemU32Ptr(0x5D4594, 833708) - 20,
									 *getMemIntPtr(0x5D4594, 833704), 20, drawData, inpData);
	nox_xxx_wndSetWindowProc_46B300(nox_gui_console_input, nox_xxx_consoleEditProc_450F40);
	nox_gui_console_locked = 0;
	nox_gui_console_wantsPass = 0;
	nox_gui_console_password[0] = 0;
	return nox_gui_console_win;
}

//----- (00451100) --------------------------------------------------------
void nox_gui_console_reloadColors_451100() {
	unsigned int color = *getMemU32Ptr(0x85B3FC, 956);
	unsigned int color2 = *getMemU32Ptr(0x85B3FC, 952);
	if (nox_gui_console_win) {
		nox_gui_console_win->draw_data.bg_color = color2;
	}
	if (nox_gui_console_scrollbox) {
		_DWORD* v1 = nox_gui_console_scrollbox->field_8;
		nox_gui_console_scrollbox->draw_data.bg_color = 0x80000000;
		nox_gui_console_scrollbox->draw_data.dis_color = 0x80000000;
		nox_gui_console_scrollbox->draw_data.en_color = color;
		nox_gui_console_scrollbox->draw_data.hl_color = color;
		nox_gui_console_scrollbox->draw_data.sel_color = 0x80000000;
		nox_gui_console_scrollbox->draw_data.text_color = color;
		int v2 = v1[9];
		if (v2) {
			*(_DWORD*)(v2 + 56) = color2;
			*(_DWORD*)(v1[9] + 80) = color2;
			*(_DWORD*)(v1[9] + 64) = color;
			*(_DWORD*)(v1[9] + 72) = color2;
			*(_DWORD*)(v1[9] + 88) = color;
			int v3 = *(_DWORD*)(v1[9] + 400);
			if (v3) {
				*(_DWORD*)(v3 + 56) = color2;
				*(_DWORD*)(*(_DWORD*)(v1[9] + 400) + 80) = color2;
				*(_DWORD*)(*(_DWORD*)(v1[9] + 400) + 64) = color;
				*(_DWORD*)(*(_DWORD*)(v1[9] + 400) + 72) = color2;
				*(_DWORD*)(*(_DWORD*)(v1[9] + 400) + 88) = color2;
			}
		}
		int v4 = v1[7];
		if (v4) {
			*(_DWORD*)(v4 + 56) = color2;
			*(_DWORD*)(v1[7] + 80) = color2;
			*(_DWORD*)(v1[7] + 64) = color;
			*(_DWORD*)(v1[7] + 72) = *getMemU32Ptr(0x5D4594, 2523948);
			*(_DWORD*)(v1[7] + 88) = color;
			*(_DWORD*)(v1[7] + 104) = color;
		}
		int v5 = v1[8];
		if (v5) {
			*(_DWORD*)(v5 + 56) = color2;
			*(_DWORD*)(v1[8] + 80) = color2;
			*(_DWORD*)(v1[8] + 64) = color;
			*(_DWORD*)(v1[8] + 72) = *getMemU32Ptr(0x5D4594, 2523948);
			*(_DWORD*)(v1[8] + 88) = color;
			*(_DWORD*)(v1[8] + 104) = color;
		}
	}
	if (nox_gui_console_input) {
		nox_gui_console_input->draw_data.bg_color = color2;
		nox_gui_console_input->draw_data.dis_color = color;
		nox_gui_console_input->draw_data.en_color = color;
		nox_gui_console_input->draw_data.hl_color = color;
		nox_gui_console_input->draw_data.sel_color = color;
		nox_gui_console_input->draw_data.text_color = color;
	}
}

//----- (004512B0) --------------------------------------------------------
int nox_gui_console_Hide_4512B0() {
	if (wndIsShown_nox_xxx_wndIsShown_46ACC0(nox_gui_console_win))
		return 0;
	if (nox_xxx_wndGetFocus_46B4F0() == nox_gui_console_input)
		nox_xxx_windowFocus_46B500(0);
	nox_window_set_hidden(nox_gui_console_win, 1);
	nox_gui_console_win->flags &= 0xFFFFFFF7;
	nox_gui_console_input->flags &= 0xFFFFFFF7;
	nox_gui_console_scrollbox->flags &= 0xFFFFFFF7;
	nox_gui_console_input->draw_data.field_0 &= 0xFFFFFFFB;
	nox_gui_console_input->draw_data.field_0 &= 0xFFFFFFFD;
	dword_5d4594_3799524 = 1;
	return 1;
}
