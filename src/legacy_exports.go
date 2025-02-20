package opennox

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy"
)

func init() {
	legacy.WriteConfigLegacy = nox_common_writecfgfile
	legacy.Sub_57C490_2 = sub_57C490_2
	legacy.Nox_thing_debug_draw = nox_thing_debug_draw
	legacy.Sub_4E8290 = sub_4E8290
	legacy.GamedataFloat = gamedataFloat
	legacy.GamedataFloatInd = gamedataFloatInd
	legacy.Nox_client_parseConfigHotkeysLine_42CF50 = nox_client_parseConfigHotkeysLine_42CF50
	legacy.Nox_client_getIntroScreenDuration_44E3B0 = nox_client_getIntroScreenDuration_44E3B0
	legacy.Nox_client_getBriefDuration = nox_client_getBriefDuration
	legacy.Nox_game_exit_xxx2 = nox_game_exit_xxx2
	legacy.Sub_470510 = sub_470510
	legacy.Sub_4703F0 = sub_4703F0
	legacy.Nox_xxx_cliDrawConnectedLoop_43B360 = nox_xxx_cliDrawConnectedLoop_43B360
	legacy.Nox_client_guiXxxDestroy_4A24A0 = nox_client_guiXxxDestroy_4A24A0
	legacy.Nox_client_quit_4460C0 = nox_client_quit_4460C0
	legacy.Nox_server_getWaypointById_579C40 = nox_server_getWaypointById_579C40
	legacy.Nox_gui_makeAnimation_43C5B0 = nox_gui_makeAnimation
	legacy.Nox_client_gui_set_flag_815132 = nox_client_gui_set_flag_815132
	legacy.Nox_client_onClientStatusA = nox_client_onClientStatusA
	legacy.Nox_client_setRenderGUI = nox_client_setRenderGUI
	legacy.Nox_client_getRenderGUI = nox_client_getRenderGUI
	legacy.Nox_gui_console_flagXxx_451410 = nox_gui_console_flagXxx_451410
	legacy.Nox_gui_console_Hide_4512B0 = nox_gui_console_Hide_4512B0
	legacy.Nox_xxx_gLoadImg = nox_xxx_gLoadImg
	legacy.Nox_xxx_gLoadAnim = nox_xxx_gLoadAnim
	legacy.Nox_xxx_dialogMsgBoxCreate_449A10 = NewDialogWindow
	legacy.Sub_449E00 = sub449E00
	legacy.Sub_449E30 = sub449E30
	legacy.Sub_449E60 = sub_449E60
	legacy.Sub_449EA0 = sub449EA0
	legacy.Sub_44A4E0 = sub_44A4E0
	legacy.Sub_44A4B0 = sub_44A4B0
	legacy.Sub_44A360 = sub_44A360
	legacy.Sub_4706C0 = sub_4706C0
	legacy.Sub_473670 = sub_473670
	legacy.Nox_client_toggleMap_473610 = nox_client_toggleMap_473610
	legacy.Nox_client_refreshServerList_4378B0 = nox_client_refreshServerList_4378B0
	legacy.Sub_438770_waitList = sub_438770_waitList
	legacy.Nox_xxx_createSocketLocal = nox_xxx_createSocketLocal
	legacy.Sub_554D10 = sub_554D10
	legacy.Nox_gui_menu_proc_ext = nox_gui_menu_proc_ext
	legacy.Nox_video_setMenuOptions = nox_video_setMenuOptions
	legacy.Sub_4A19F0 = guiSetBackButtonText
	legacy.Sub_4C3A90 = sub_4C3A90
	legacy.Sub_4CBE70 = sub_4CBE70
	legacy.Sub_4AAA10 = sub_4AAA10
	legacy.Nox_new_window_from_file = nox_new_window_from_file
	legacy.Sub_445B40 = clientLoadCoopAuto445B40
	legacy.Nox_xxx_mapGenMakeInfo_4D5DB0 = nox_xxx_mapGenMakeInfo_4D5DB0
	legacy.Nox_common_checkMapFile = nox_common_checkMapFile
	legacy.Nox_xxx_mapWriteSectionsMB_426E20 = nox_xxx_mapWriteSectionsMB_426E20
	legacy.NetworkLogPrint = networkLogPrint
	legacy.ClientSetServerHost = clientSetServerHost
	legacy.Nox_client_joinGame_438A90 = nox_client_joinGame_438A90
	legacy.Nox_xxx_moveUpdateSpecial_517970 = nox_xxx_moveUpdateSpecial_517970
	legacy.Nox_mapToGameFlags = nox_mapToGameFlags
	legacy.Sub_40A1A0 = sub_40A1A0
	legacy.PlatformTicks = platformTicks
	legacy.Nox_ticks_reset_416D40 = nox_ticks_reset_416D40
	legacy.GetPhonemeTree = func() unsafe.Pointer {
		return unsafe.Pointer(getPhonemeTree())
	}
	legacy.Nox_xxx_spellAwardAll1_4EFD80 = nox_xxx_spellAwardAll1_4EFD80
	legacy.Nox_xxx_spellAwardAll2_4EFC80 = nox_xxx_spellAwardAll2_4EFC80
	legacy.Nox_xxx_spellAwardAll3_4EFE10 = nox_xxx_spellAwardAll3_4EFE10
	legacy.Nox_xxx_spellGetAud44_424800 = nox_xxx_spellGetAud44_424800
	legacy.Nox_xxx_spellTitle_424930 = nox_xxx_spellTitle_424930
	legacy.Nox_xxx_spellDescription_424A30 = nox_xxx_spellDescription_424A30
	legacy.Nox_xxx_spellByTitle_424960 = nox_xxx_spellByTitle_424960
	legacy.Nox_xxx_spellManaCost_4249A0 = nox_xxx_spellManaCost_4249A0
	legacy.Nox_xxx_spellPhonemes_424A20 = nox_xxx_spellPhonemes_424A20
	legacy.Nox_xxx_spellIcon_424A90 = nox_xxx_spellIcon_424A90
	legacy.Nox_xxx_spellIconHighlight_424AB0 = nox_xxx_spellIconHighlight_424AB0
	legacy.Nox_xxx_spellFirstValid_424AD0 = nox_xxx_spellFirstValid_424AD0
	legacy.Nox_xxx_spellNextValid_424AF0 = nox_xxx_spellNextValid_424AF0
	legacy.Nox_xxx_spellIsValid_424B50 = nox_xxx_spellIsValid_424B50
	legacy.Nox_xxx_spellIsEnabled_424B70 = nox_xxx_spellIsEnabled_424B70
	legacy.Nox_xxx_spellEnable_424B90 = nox_xxx_spellEnable_424B90
	legacy.Nox_xxx_spellDisable_424BB0 = nox_xxx_spellDisable_424BB0
	legacy.Nox_xxx_spellCanUseInTrap_424BF0 = nox_xxx_spellCanUseInTrap_424BF0
	legacy.Nox_xxx_spellPrice_424C40 = nox_xxx_spellPrice_424C40
	legacy.Nox_xxx_spellEnableAll_424BD0 = nox_xxx_spellEnableAll_424BD0
	legacy.Nox_xxx_castSpellByUser_4FDD20 = nox_xxx_castSpellByUser_4FDD20
	legacy.Nox_xxx_spellCastedFirst_4FE930 = nox_xxx_spellCastedFirst_4FE930
	legacy.Nox_xxx_spellCastedNext_4FE940 = nox_xxx_spellCastedNext_4FE940
	legacy.Sub_4FE8A0 = sub_4FE8A0
	legacy.Sub_4FE900 = sub_4FE900
	legacy.Nox_xxx_spellCastByPlayer_4FEEF0 = nox_xxx_spellCastByPlayer_4FEEF0
	legacy.Nox_xxx_spellCancelDurSpell_4FEB10 = nox_xxx_spellCancelDurSpell_4FEB10
	legacy.Sub_4FE980 = sub_4FE980
	legacy.Sub_4FF310 = sub_4FF310
	legacy.Nox_exit = nox_exit
	legacy.Nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode = nox_xxx_gameGetScreenBoundaries_43BEB0_get_video_mode
	legacy.Sub_4AA9C0 = sub_4AA9C0
	legacy.Nox_server_questAllowDefault = nox_server_questAllowDefault
	legacy.Nox_server_questNextStageThreshold_4D74F0 = nox_server_questNextStageThreshold_4D74F0
	legacy.Sub_4D6F30 = sub_4D6F30
	legacy.GetNetPlayerBufSize = getNetPlayerBufSize
	legacy.Nox_netlist_addToMsgListSrv = nox_netlist_addToMsgListSrv
	legacy.Nox_client_setVersion_409AE0 = nox_client_setVersion_409AE0
	legacy.Nox_video_getCutSize_4766D0 = nox_video_getCutSize_4766D0
	legacy.Nox_video_setCutSize_4766A0 = nox_video_setCutSize_4766A0
	legacy.Nox_video_setGammaSlider = nox_video_setGammaSlider
	legacy.Sub_43BE50_get_video_mode_id = sub_43BE50_get_video_mode_id
	legacy.Get_video_mode_string = get_video_mode_string
	legacy.Nox_getBackbufWidth = nox_getBackbufWidth
	legacy.Nox_getBackbufHeight = nox_getBackbufHeight
	legacy.Nox_video_getFullScreen = nox_video_getFullScreen
	legacy.Nox_video_setFullScreen = nox_video_setFullScreen
	legacy.Sub_430C30_set_video_max = sub_430C30_set_video_max
	legacy.VideoGetMaxSize = videoGetMaxSize
	legacy.Nox_video_callCopyBackBuffer_4AD170 = nox_video_callCopyBackBuffer_4AD170
	legacy.Nox_getBackbufferPitch = nox_getBackbufferPitch
	legacy.Nox_client_clearScreen_440900 = nox_client_clearScreen_440900
	legacy.Nox_draw_setCutSize_476700 = nox_draw_setCutSize_476700
	legacy.Nox_xxx_bookSaveSpellForDragDrop_477640 = nox_xxx_bookSaveSpellForDragDrop_477640
	legacy.Nox_xxx_bookSpellDnDclear_477660 = nox_xxx_bookSpellDnDclear_477660
	legacy.Nox_xxx_bookGetSpellDnDType_477670 = nox_xxx_bookGetSpellDnDType_477670
	legacy.Nox_xxx_cursorSetDraggedItem_477690 = nox_xxx_cursorSetDraggedItem_477690
	legacy.Nox_xxx_cursorResetDraggedItem_4776A0 = nox_xxx_cursorResetDraggedItem_4776A0
	legacy.Sub_478000 = sub_478000
	legacy.Nox_xxx_GetEndgameDialog = nox_xxx_GetEndgameDialog
	legacy.GameGetPlayState = gameGetPlayState
	legacy.ServerCheatGod = serverCheatGod
	legacy.Nox_xxx_serverHost_43B4D0 = nox_xxx_serverHost_43B4D0
	legacy.Nox_xxx_netServerCmd_440950 = nox_xxx_netServerCmd_440950
	legacy.ExecConsoleCmd = execConsoleCmd
	legacy.Nox_xxx_gameIsNotMultiplayer_4DB250 = nox_xxx_gameIsNotMultiplayer_4DB250
	legacy.Nox_xxx_gameSetSwitchSolo_4DB220 = nox_xxx_gameSetSwitchSolo_4DB220
	legacy.Nox_xxx_gameIsSwitchToSolo_4DB240 = nox_xxx_gameIsSwitchToSolo_4DB240
	legacy.Nox_xxx_gameSetWallsDamage_4E25A0 = nox_xxx_gameSetWallsDamage_4E25A0
	legacy.GetDoDamageWalls = func() bool {
		return doDamageWalls
	}
	legacy.ClientPacketFade = clientPacketFade
	legacy.Nox_video_inFadeTransition_44E0D0 = nox_video_inFadeTransition_44E0D0
	legacy.Sub_4373A0 = sub_4373A0
	legacy.Sub_43EA20 = sub_43EA20
	legacy.Sub_43E9F0 = sub_43E9F0
	legacy.Sub_43E940 = sub_43E940
	legacy.Sub_43EFD0 = sub_43EFD0
	legacy.Sub_43EC10 = sub_43EC10
	legacy.Sub_43F130 = sub_43F130
	legacy.Sub_43ED00 = sub_43ED00
	legacy.Sub_43D650 = sub_43D650
	legacy.Sub_43D680 = sub_43D680
	legacy.Sub_43D6A0 = sub_43D6A0
	legacy.Sub_44D640 = sub_44D640
	legacy.Sub_44D7E0 = sub_44D7E0
	legacy.Nox_xxx_musicStartPlay_43D6C0 = nox_xxx_musicStartPlay_43D6C0
	legacy.Sub_44D660 = sub_44D660
	legacy.Sub_43F060 = sub_43F060
	legacy.Sub_43EC30 = sub_43EC30
	legacy.Sub_43ECB0 = sub_43ECB0
	legacy.Nox_xxx_updateSprings_5113A0 = nox_xxx_updateSprings_5113A0
	legacy.Nox_xxx_unitIsUnitTT_4E7C80 = nox_xxx_unitIsUnitTT_4E7C80
	legacy.Nox_xxx_updatePlayer_4F8100 = nox_xxx_updatePlayer_4F8100
	legacy.Nox_xxx_teleportAllPixies_4FD090 = nox_xxx_teleportAllPixies_4FD090
	legacy.Nox_xxx_enemyAggro_5335D0 = nox_xxx_enemyAggro_5335D0
	legacy.Sub_5336D0 = sub_5336D0
	legacy.Nox_xxx_lightningSpellDuration_52FFD0 = nox_xxx_lightningSpellDuration_52FFD0
	legacy.Sub_44A4A0 = sub_44A4A0
	legacy.Nox_script_shouldReadMoreXxx = nox_script_shouldReadMoreXxx
	legacy.Nox_script_shouldReadEvenMoreXxx = nox_script_shouldReadEvenMoreXxx
	legacy.Nox_xxx_xfer_saveObj51DF90 = nox_xxx_xfer_saveObj51DF90
	legacy.Nox_xxx_XFerDefault4F49A0 = nox_xxx_XFerDefault4F49A0
	legacy.Nox_xxx_XFer_ReadShopItem_52A840 = nox_xxx_XFer_ReadShopItem_52A840
	legacy.Nox_xxx_XFer_WriteShopItem_52A5F0 = nox_xxx_XFer_WriteShopItem_52A5F0
	legacy.Sub_5002D0 = sub_5002D0
	legacy.Sub_4FC670 = sub_4FC670
	legacy.Nox_xxx_playerExecuteAbil_4FBB70 = nox_xxx_playerExecuteAbil_4FBB70
	legacy.Sub_4FC0B0 = sub_4FC0B0
	legacy.Nox_xxx_playerCancelAbils_4FC180 = nox_xxx_playerCancelAbils_4FC180
	legacy.Sub_4FC300 = sub_4FC300
	legacy.Nox_xxx_abilityGetName_0_425260 = nox_xxx_abilityGetName_0_425260
	legacy.Nox_xxx_abilityCooldown_4252D0 = nox_xxx_abilityCooldown_4252D0
	legacy.Sub_4252F0 = sub_4252F0
	legacy.Nox_xxx_spellGetAbilityIcon_425310 = nox_xxx_spellGetAbilityIcon_425310
	legacy.Nox_xxx_bookFirstKnownAbil_425330 = nox_xxx_bookFirstKnownAbil_425330
	legacy.Nox_xxx_bookNextKnownAbil_425350 = nox_xxx_bookNextKnownAbil_425350
	legacy.Sub_425450 = sub_425450
	legacy.Nox_xxx_netAbilRepotState_4D8100 = nox_xxx_netAbilRepotState_4D8100
	legacy.Nox_xxx_mapSwitchLevel_4D12E0_end = nox_xxx_mapSwitchLevel_4D12E0_end
	legacy.Nox_script_indexByEvent = nox_script_indexByEvent
	legacy.Nox_script_getString_512E40 = nox_script_getString_512E40
	legacy.Nox_script_callbackName = nox_script_callbackName
	legacy.Nox_script_objCallbackName_508CB0 = nox_script_objCallbackName_508CB0
	legacy.Sub_511E60 = sub_511E60
	legacy.Nox_setImaginaryCaster = nox_setImaginaryCaster
	legacy.Nox_script_readWriteZzz_541670 = nox_script_readWriteZzz_541670
	legacy.Sub_4C26F0 = sub_4C26F0
	legacy.Sub_44E320 = sub_44E320
	legacy.Sub_4A24C0 = sub_4A24C0
	legacy.Sub_43BE40 = sub_43BE40
	legacy.Sub_43BE30 = sub_43BE30
	legacy.GuiSelProc2 = gui_SelProc2
	legacy.Sub_42BFB0 = sub_42BFB0
	legacy.Nox_xxx_objectTOCgetTT = nox_xxx_objectTOCgetTT
	legacy.Sub_42C310 = sub_42C310
	legacy.Sub_42C2E0 = sub_42C2E0
	legacy.Sub_42C300 = sub_42C300
	legacy.Sub_42BFE0 = sub_42BFE0
	legacy.Sub_4E3AD0 = sub_4E3AD0
	legacy.Nox_xxx_modifGetDescById_413330 = nox_xxx_modifGetDescById_413330
	legacy.Nox_xxx_equipClothFindDefByTT_413270 = nox_xxx_equipClothFindDefByTT_413270
	legacy.Sub_4A5E90_A = sub_4A5E90_A
	legacy.Nox_xxx_fireEffect_4E0550 = nox_xxx_fireEffect_4E0550
	legacy.Nox_xxx_fireRingEffect_4E05B0 = nox_xxx_fireRingEffect_4E05B0
	legacy.Nox_xxx_blueFREffect_4E05F0 = nox_xxx_blueFREffect_4E05F0
	legacy.Sub_4E4C90 = sub_4E4C90
	legacy.Sub_4E3B80 = sub_4E3B80
	legacy.Sub415A30 = sub415A30
	legacy.Sub415EC0 = sub415EC0
	legacy.Nox_savegame_rm = deleteSaveDir
	legacy.Nox_client_countPlayerFiles04_4DC7D0 = nox_client_countPlayerFiles04_4DC7D0
	legacy.Nox_xxx_gameGet_4DB1B0 = nox_xxx_gameGet_4DB1B0
	legacy.Sub_4DCC90 = sub_4DCC90
	legacy.Sub_4DB1C0 = sub_4DB1C0
	legacy.Sub_4DCBF0 = sub_4DCBF0
	legacy.Sub_4460A0 = sub_4460A0
	legacy.Nox_xxx_serverIsClosing_446180 = nox_xxx_serverIsClosing_446180
	legacy.Sub_4DCC10 = sub_4DCC10
	legacy.Sub_4DCFB0 = sub_4DCFB0
	legacy.Sub_4DD0B0 = sub_4DD0B0
	legacy.Nox_xxx_monsterCast_540A30 = nox_xxx_monsterCast_540A30
	legacy.InputSetKeyTimeoutLegacy = inputSetKeyTimeoutLegacy
	legacy.InputKeyCheckTimeoutLegacy = inputKeyCheckTimeoutLegacy
	legacy.Sub_416120 = sub_416120
	legacy.Sub_416170 = sub_416170
	legacy.Sub_416150 = sub_416150
	legacy.Nox_xxx_keybind_nameByTitle_42E960 = nox_xxx_keybind_nameByTitle_42E960
	legacy.Nox_xxx_bindevent_bindNameByTitle_42EA40 = nox_xxx_bindevent_bindNameByTitle_42EA40
	legacy.Sub_4C3B70 = sub_4C3B70
	legacy.Sub_4CBBF0 = sub_4CBBF0
	legacy.Nox_input_reset_430140 = nox_input_reset_430140
	legacy.Nox_xxx_updateSpellRelated_424830 = nox_xxx_updateSpellRelated_424830
	legacy.Nox_xxx_playerDisconnByPlrID_4DEB00 = nox_xxx_playerDisconnByPlrID_4DEB00
	legacy.Nox_xxx_playerCallDisconnect_4DEAB0 = nox_xxx_playerCallDisconnect_4DEAB0
	legacy.Nox_xxx_playerCameraUnlock_4E6040 = nox_xxx_playerCameraUnlock_4E6040
	legacy.Nox_xxx_playerCameraFollow_4E6060 = nox_xxx_playerCameraFollow_4E6060
	legacy.Nox_xxx_playerGetPossess_4DDF30 = nox_xxx_playerGetPossess_4DDF30
	legacy.Nox_xxx_playerGoObserver_4E6860 = nox_xxx_playerGoObserver_4E6860
	legacy.Nox_xxx_playerObserveClear_4DDEF0 = nox_xxx_playerObserveClear_4DDEF0
	legacy.Nox_client_onClassStats = nox_client_onClassStats
	legacy.Nox_xxx_playerObserveMonster_4DDE80 = nox_xxx_playerObserveMonster_4DDE80
	legacy.Nox_ai_debug_print = nox_ai_debug_print
	legacy.Sub_545E60 = sub_545E60
	legacy.Sub_50D1C0 = sub_50D1C0
	legacy.Nox_xxx_gameSetAudioFadeoutMb_501AC0 = nox_xxx_gameSetAudioFadeoutMb_501AC0
	legacy.Nox_xxx_monsterPopAction_50A160 = nox_xxx_monsterPopAction_50A160
	legacy.Nox_xxx_monsterPushAction_50A260_impl = nox_xxx_monsterPushAction_50A260_impl
	legacy.Nox_xxx_unitUpdateMonster_50A5C0 = nox_xxx_unitUpdateMonster_50A5C0
	legacy.Sub_446380 = sub_446380
	legacy.Nox_xxx_cliUpdateCameraPos_435600 = nox_xxx_cliUpdateCameraPos_435600
	legacy.Sub_437260 = sub_437260
	legacy.Get_nox_client_texturedFloors_154956 = get_nox_client_texturedFloors_154956
	legacy.Sub_480860 = sub_480860
	legacy.Nox_xxx_drawList1096512_Append_4754C0 = nox_xxx_drawList1096512_Append_4754C0
	legacy.Sub_473970 = sub_473970
	legacy.Nox_client_isConnected = nox_client_isConnected
	legacy.Sub_4A1BE0 = sub_4A1BE0
	legacy.Sub_41E300 = sub_41E300
	legacy.Nox_client_resetScreenParticles_431510 = nox_client_resetScreenParticles_431510
	legacy.Sub_46D6F0 = sub_46D6F0
	legacy.Sub_413A00 = sub_413A00
	legacy.Sub_44A400 = sub_44A400
	legacy.Nox_game_showSelChar_4A4DB0 = nox_game_showSelChar_4A4DB0
	legacy.Nox_savegame_sub_46D580 = nox_savegame_sub_46D580
	legacy.Sub_450580 = sub_450580
	legacy.Sub_4DB170 = sub_4DB170
	legacy.Nox_setSaveFileName_4DB130 = setSaveFileName
	legacy.Sub_44D8F0 = sub_44D8F0
	legacy.Nox_xxx_harpoonBreakForPlr_537520 = nox_xxx_harpoonBreakForPlr_537520
	legacy.Nox_xxx_collideHarpoon_4EB6A0 = nox_xxx_collideHarpoon_4EB6A0
	legacy.Nox_xxx_updateHarpoon_54F380 = nox_xxx_updateHarpoon_54F380
	legacy.SendXXX_5550D0 = sendXXX_5550D0
	legacy.Sub_5545A0 = getServerPort
	legacy.Sub_554230 = getOwnIP
	legacy.Nox_xxx_netStatsMultiplier_4D9C20 = nox_xxx_netStatsMultiplier_4D9C20
	legacy.Sub_554240 = sub_554240
	legacy.Nox_xxx_net_getIP_554200 = nox_xxx_net_getIP_554200
	legacy.Nox_xxx_netOnPacketRecvCli_48EA70 = nox_xxx_netOnPacketRecvCli_48EA70
	legacy.Sub_43C6E0 = sub_43C6E0
	legacy.Sub_43CF40 = sub_43CF40
	legacy.Sub_43CF70 = sub_43CF70
	legacy.Sub_47A1F0 = guiCloseNPCDialog
	legacy.Nox_xxx_unitMonsterInit_4F0040 = objectMonsterInit
	legacy.Nox_xxx_setNPCColor_4E4A90 = nox_xxx_setNPCColor_4E4A90
	legacy.Nox_xxx_checkSummonedCreaturesLimit_500D70 = nox_xxx_checkSummonedCreaturesLimit_500D70
	legacy.Nox_xxx_unitDoSummonAt_5016C0 = nox_xxx_unitDoSummonAt_5016C0
	legacy.Sub_57AEE0 = sub_57AEE0
	legacy.Sub_4E71F0 = sub_4E71F0
	legacy.Nox_bomberDead_54A150 = nox_bomberDead_54A150
	legacy.Nox_xxx_dieGlyph_54DF30 = nox_xxx_dieGlyph_54DF30
	legacy.Nox_xxx_collideGlyph_4E9A00 = nox_xxx_collideGlyph_4E9A00
	legacy.Sub_51A100 = func() {
		noxServer.mapSend.sub_51A100()
	}
	legacy.Sub_41CC00 = sub_41CC00
	legacy.Nox_xxx_playerSendMOTD_4DD140 = nox_xxx_playerSendMOTD_4DD140
	legacy.Sub_526CA0 = sub_526CA0
	legacy.Nox_xxx_monsterCreateFn_54C480 = nox_xxx_monsterCreateFn_54C480
	legacy.Nox_xxx_monsterClearActionStack_50A3A0 = nox_xxx_monsterClearActionStack_50A3A0
	legacy.Sub_4A1A40 = sub_4A1A40
}
