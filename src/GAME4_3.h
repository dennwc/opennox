#ifndef NOX_PORT_GAME4_3
#define NOX_PORT_GAME4_3

#include "defs.h"


int nox_xxx_onFrameLightning_52F8A0(float a1);
void nox_xxx_lightningCanAttackCheck_52FF10(int a1, int a2);
_DWORD* nox_xxx_lightningSpellDuration_52FFD0(int a1, int a2, int a3);
void nox_xxx_lightningSpellTrapEffect_530020(int a1, int a2);
char sub_530100(_DWORD* a1);
int nox_xxx_spellTagCreature_530160(_DWORD* a1);
unsigned int sub_530250(int a1);
int sub_530270(int a1);
int nox_xxx_spellBlink2_530310(_DWORD* a1);
int nox_xxx_spellBlink1_530380(int* a1);
_DWORD* nox_xxx_spellTeleportCreateWake_530560(int a1, int* a2, _DWORD* a3);
int sub_5305D0(_DWORD* a1);
int sub_530650(int* a1);
int nox_xxx_castTele_530820(int a1);
int sub_530880(int* a1);
int nox_xxx_castTTT_530B70(int* a1);
int sub_530CA0(int a1);
int sub_530D30(int* a1);
int nox_xxx_manaBomb_530F90(_DWORD* a1);
int nox_xxx_manaBombBoom_5310C0(int* a1);
int sub_531290(int a1);
int nox_xxx_spellTurnUndeadCreate_531310(_DWORD* a1);
int nox_xxx_spellTurnUndeadUpdate_531410();
int nox_xxx_spellTurnUndeadDelete_531420(int a1);
int sub_531490(_DWORD* a1);
BOOL sub_5314F0(int a1);
int sub_531560(int a1);
int nox_xxx_plasmaSmth_531580(int a1);
int nox_xxx_plasmaShot_531600(int a1);
void sub_531920(int a1, int a2);
int sub_5319E0(int a1);
int nox_xxx_spellCreateMoonglow_531A00(_DWORD* a1);
int sub_531AF0(int a1);
int* nox_xxx_TODOsomeCallingMeleeAttack_531B40(int a1, int a2);
int* sub_531C10(int a1, int a2);
int* nox_xxx_monsterAction_531C60(int a1, int a2);
int* sub_531D50(int a1, int a2);
int nox_xxx_mobActionFightStart_531E20(float a1);
int sub_531E90(int a1);
char nox_xxx_mobActionFight_531EC0(int a1);
int sub_532040(int a1, int a2);
char nox_xxx_monsterShieldBlockStart_532070(int a1);
char nox_xxx_monsterShieldBlockStop_5320E0(int a1);
char sub_532100(int a1);
char sub_532110(int a1);
int* nox_xxx_mobActionMelee1_532130(int a1);
void sub_532390(int a1, int a2);
char nox_xxx_mobActionMeleeAtt_532440(int a1);
_DWORD* sub_532540(int a1);
char nox_xxx_mobActionMissileAtt_532610(int a1);
char nox_xxx_monsterPlayHurtSound_532800(int a1);
int sub_532880(int a1);
int nox_xxx_soundPlayerDamageSound_5328B0(int a1, int a2);
int sub_532930(int a1, unsigned __int16 a2, unsigned __int16 a3);
int nox_xxx_soundDefaultDamageSound_532E20(int a1, int a2);
_DWORD* sub_532EC0(int a1, unsigned __int16 a2);
int sub_532F70(int a1);
BOOL sub_532FB0(__int16 a1);
int sub_532FE0(unsigned __int16 a1, int a2);
int sub_533010(unsigned __int16 a1, int a2);
unsigned __int8* nox_xxx_unused_533030(int a1, int a2);
int nox_xxx_projAddVelocitySmth_533080(int a1, int a2, float a3, int a4);
BOOL nox_xxx_unitIsEnemyTo_5330C0(int a1, int a2);
BOOL sub_533360(int a1, int a2);
int nox_xxx_unused_5333F0(int a1);
void nox_xxx_fnPickEnemyAggro_533460(int a1, int a2);
int sub_533570(int a1);
int sub_5335D0(int a1, float a2);
BOOL nox_xxx_unused_533610(int a1, int a2);
char* sub_533660(int a1);
long double sub_5336D0(int a1);
void sub_533720(int a1, int a2);
int nox_xxx_mobActionToAnimation_533790(int a1);
void nox_xxx_orderUnit_533900(int owner, int creature, int orderType);
void nox_xxx_enactUnitOrder_5339A0(int source, int unit, int orderId);
BOOL sub_533C80(int a1, int a2);
void nox_xxx_mobCalcDir_533CC0(int a1, float* a2);
unsigned __int8* nox_xxx_unitNPCActionToAnim_533D00(int a1);
int nox_xxx_monsterTestBlockShield_533E70(int a1);
void sub_533EB0(int a1, int a2);
int sub_534020(int a1);
char nox_xxx_monsterMoveAudio_534030(int a1);
BOOL sub_534120(int a1, float2* a2);
int sub_534160(int a1);
void nox_xxx_debugPrintf_5341A0(char* a1, ...);
int sub_5341D0(int a1);
BOOL sub_5341F0(nox_object_t* a1p);
BOOL nox_xxx_monsterCanMelee_534220(int a1);
BOOL nox_xxx_monsterCanShoot_534280(int a1);
int nox_xxx_monsterHasShield_5342C0(int a1);
int nox_xxx_monsterCanCast_534300(int a1);
BOOL nox_xxx_monsterIsMoveing_534320(int a1);
BOOL sub_534340(int a1);
BOOL nox_xxx_monsterCanAttackAtWill_534390(int a1);
BOOL sub_5343C0(int a1);
BOOL sub_534400(int a1);
BOOL sub_534440(int a1);
double sub_534470(int a1);
char* nox_xxx_unused_534490(float a1);
double nox_xxx_unused_5344C0(const char* a1);
char* nox_xxx_unused_534530(int a1);
int nox_xxx_unused_534550(const char* a1);
char* sub_5345B0(int a1);
int nox_xxx_actionNByNameMB_5345F0(const char* a1);
char* sub_534650(int a1);
int nox_xxx_actionByName_534670(const char* a1);
int sub_5346D0(int a1);
int nox_xxx_monsterResetEnemy_5346F0(int a1);
BOOL sub_534710(int a1);
int sub_534750(int a1);
int sub_534780(int a1);
int sub_5347A0(int a1);
BOOL sub_5347C0(int a1);
BOOL nox_xxx_isNotPoisoned_5347F0(int a1);
BOOL nox_xxx_mobGetMoveAttemptTime_534810(int a1);
BOOL nox_xxx_unitIsMimic_534840(int a1);
int* nox_xxx_unused_534870(int a1);
int* nox_xxx_unused_5348A0(int a1);
char nox_xxx_mobActionMorphToChest_5348D0(int a1);
char nox_xxx_mobActionMorphBackToSelf_534910(int a1);
void nox_xxx_monsterMimicCheckMorph_534950(int a1);
BOOL nox_xxx_unitIsPlant_534A10(int a1);
BOOL nox_xxx_unitIsZombie_534A40(int a1);
char nox_xxx_mobActionGetUp_534A90(int a1);
unsigned int nox_xxx_mobRaiseZombie_534AB0(int a1);
BOOL nox_xxx_unitIsAFish_534B10(int a1);
BOOL nox_xxx_unitIsARat_534B60(int a1);
BOOL nox_xxx_unitIsAFrog_534B90(int a1);
int nox_xxx_damageToMap_534BC0(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_wallPreDestroy_534DA0(int* a1);
bool nox_xxx_mapDamageToWalls_534FC0(int4* a1, int a2, float a3, int a4, int a5, int a6);
int nox_xxx_mapTraceRay_535250(float4* a1, float2* a2, int2* a3, char a4);
int nox_xxx_parseXP_535970(int a1, int a2, char* a3);
int nox_xxx_parseExtent_535990(int a1, int a2, char* a3);
int nox_xxx_parseZSize_5359B0(int a1, int a2, char* a3);
int nox_xxx_parseWorth_535A00(int a1, int a2, char* a3);
int nox_xxx_parseSpeed_535A20(int a1, int a2, char* a3);
int nox_xxx_parseHealth_535A60(int a1, int a2, char* a3);
int nox_xxx_parseFlags_535AD0(int a1, int a2, const char* a3);
int nox_xxx_parseClass_535B00(int a1, int a2, const char* a3);
int nox_xxx_parseSubclass_535B30(int a1, int a2, const char* a3);
int nox_xxx_parseMass_535B60(int a1, int a2, char* a3);
int nox_xxx_parseCapacity_535B80(int a1, int a2, char* a3);
int nox_xxx_parseWeight_535BB0(int a1, int a2, char* a3);
int nox_xxx_parseMaterial_535BE0(int a1, int a2, const char* a3);
int nox_xxx_parseLightIntensity_535C20(int a1, int a2, char* a3);
int nox_xxx_parseMenuIcon_535C30(int a1, int a2, char* a3);
int nox_xxx_parsePrettyImage_0_535C80(int a1, int a2);
int nox_xxx_parseDraw_535CD0(int a1, _DWORD* a2, void* a3);
int nox_xxx_parseXFer_5360A0(int a1, int a2, char* a3);
char* sub_536130(char* a1, int* a2);
char* sub_536180(char* a1, int* a2);
char* sub_5361B0(char* a1, int a2);
char* sub_536260(char* a1, int a2);
char* sub_536390(char* a1, int* a2);
char* sub_5363C0(char* a1, int* a2);
int nox_xxx_parseUseFn_5363F0(_DWORD* a1, int a2, char* a3);
int sub_5364E0(char* a1, int a2);
int sub_536550(char* a1, _DWORD* a2);
int sub_536580(char* a1, int a2);
int sub_5365B0(char* a1, int a2);
int sub_536600(char* a1, int a2);
int nox_xxx_parseUpdate_536620(_DWORD* a1, int a2, char* a3);
int nox_xxx_parsePickup_536710(int a1, int a2, char* a3);
char* sub_5367B0(__int16* a1);
int nox_xxx_parseCreateProc_536830(int a1, int a2, char* a3);
int sub_5368C0(char* a1, int a2);
int sub_536910(int a1, int a2);
int nox_xxx_parseInitProc_536930(_DWORD* a1, int a2, char* a3);
int nox_xxx_parseDrop_536A20(int a1, int a2, char* a3);
char* sub_536AC0(__int16* a1);
int sub_536B40(char* a1, int a2);
int nox_xxx_parseDieProc_536B80(int a1, int a2, char* a3);
int nox_xxx_parseDamageFn_536C60(int a1, int a2, char* a3);
int nox_xxx_parseDamageSound_536CF0(int a1, int a2, char* a3);
int sub_536D80(char* a1, int a2);
BOOL sub_536DA0(char* a1, int* a2);
int sub_536DE0(char* a1, _BYTE* a2);
BOOL nox_xxx_collideDamageLoad_536E10(char* a1, int a2);
int sub_536E50(char* a1, _BYTE* a2);
int sub_536E80(char* a1, int* a2);
int nox_xxx_parseCollide_536EC0(_DWORD* a1, int a2, char* a3);
BOOL nox_xxx_unitCanSee_536FB0(int a1, int a2, char a3);
int nox_xxx_unitCanInteractWith_5370E0(int a1, int a2, char a3);
int nox_xxx_mapCheck_537110(int a1, int a2);
void nox_xxx_lineCollisionChk_537230(float* a1, int arg4);
int nox_xxx_traceRay_5374B0(float4* a1);
void sub_5374D0(_DWORD* a1);
_DWORD* nox_xxx_harpoonBreakForPlr_537520(_DWORD* a1);
void sub_537540(int a1);
int sub_537580(int a1);
BOOL sub_537590();
void sub_5375A0(int a1);
char nox_xxx_unitHasCollideOrUpdateFn_537610(int a1);
int sub_537700();
int sub_537740();
int sub_537750(int a1);
unsigned int sub_537760();
char sub_537770(int a1);
char nox_xxx_projectileTraceHit_537850(int a1, int* a2, float2* a3);
_DWORD* nox_xxx_sMakeScorch_537AF0(int* a1, int a2);
int nox_xxx_scorchInit_537BD0();
char nox_xxx_trapBAH_537C10(int a1, int a2);
void sub_537DD0(float* a1, int a2);
int sub_537E60(int a1, int a2, int a3, int a4);
void sub_537F00(float* a1, int a2);
int nox_xxx_castGlyph_537FA0(int a1, int a2, int a3, int a4, int a5);
int sub_538250(int a1);
char nox_xxx___mkgmtime_538280(int a1);
int nox_xxx_playerPreAttackEffects_538290(int a1, int a2, int a3, int a4);
int nox_xxx_playerTraceAttack_538330(int a1, int a2);
void sub_538510(int a1, int a2);
void sub_5386A0(int a3, int a2);
int nox_xxx_itemApplyAttackEffect_538840(int a1, int a2, int a3);
int nox_xxx_createHarpoonBolt_538890(int a1);
BOOL nox_xxx_playerAttack_538960(int a1);
__int16 nox_xxx_warcryStunMonsters_539B90(int a1, int a2);
int nox_xxx_shootBowCrossbow1_539BD0(int a1, int a2);
_DWORD* nox_xxx_shootBowCrossbow2_539D80(int a1, int a2, int a3, char* a4);
int nox_xxx_shootApplyEffects_539F40(int a1, int a2, int a3);
int sub_539FB0(_DWORD* a1);
int nox_xxx_playerTryReloadQuiver_539FF0(_DWORD* a1);
int nox_xxx_equipWeaponNPC_53A030(int a1, int a2);
void sub_53A0F0(int a1, int a2, int a3);
int nox_xxx_playerDequipWeapon_53A140(_DWORD* a1, nox_object_t* item, int a3, int a4);
int nox_xxx_NPCEquipWeapon_53A2C0(int a1, nox_object_t* item);
void sub_53A3D0(_DWORD* a1);
int nox_xxx_playerEquipWeapon_53A420(_DWORD* a1, nox_object_t* item, int a3, int a4);
int sub_53A680(int a1);
void sub_53A6C0(int a1, nox_object_t* item);
int sub_53A720(int a1, nox_object_t* item, int a3, int a4);
int nox_xxx_sendMsgOblivionPickup_53A9C0(int a1, nox_object_t* item, int a3, int a4);
void sub_53AAB0(int a1);
int nox_xxx_dropWeapon_53AB10(int a1, _DWORD* a2, int* a3);
void sub_53AB90(int a1, int a2);
void nox_xxx_updateProjectile_53AC10(int a1);
char nox_xxx_updateDoor_53AC50(int a1);
void nox_xxx_updateSpark_53ADC0(int a1);
char sub_53AE00(int a1);
float* nox_xxx_updateProjTrail_53AEC0(int a1);
void nox_xxx_updatePush_53B030(int a1);
char nox_xxx_updateToggle_53B060(_DWORD* a1);
char nox_xxx_updateTrigger_53B1B0(int a1);
char sub_53B300(int a1);
char nox_xxx_updateSwitch_53B320(_DWORD* a1);
char nox_xxx_updateElevatorShaft_53B380(int a1);
void nox_xxx_fnElevatorShaft_53B410(int a1, int a2);
_DWORD* nox_xxx_elevatorAud_53B490(int a1, int a2);
void nox_xxx_updateElevator_53B5D0(_DWORD* a1);
void nox_xxx_elevatorFn_53B750(int a1, int a2);
void nox_xxx_updatePhantomPlayer_53B860(int a1);
void nox_xxx_updateLifetime_53B8F0(int unit);
void nox_xxx_spellFlyUpdate_53B940(int a1);
void nox_xxx_updateAntiSpellProj_53BB00(int a1);
void sub_53BD10(int a1, int a2);
int nox_xxx_updateMagicMissile_53BDA0(int a1);
int nox_xxx_updateTeleportPentagram_53BEF0(int a1);
void nox_xxx_fnPentagramTeleport_53C060(float* a1, int a2);
int nox_xxx_updateInvisiblePentagram_53C0C0(int a1);
void sub_53C140(float* a1, int a2);
void nox_xxx_updateBlow_53C160(int a3);
void sub_53C240(float* a1, int arg4);
int nox_xxx_rechargeItem_53C520(int a1, int a2);
signed int nox_xxx_updateObelisk_53C580(int a1);
int nox_xxx_getRechargeRate_53C940(_DWORD* a1);
void nox_xxx_updateBlackPowderBarrel_53C9A0(float* a1);
void nox_xxx_updateOneSecondDie_53CB60(int a1);
void nox_xxx_updateWaterBarrel_53CB90(int a1);
void nox_xxx_waterBarrel_53CC30(float* a1, int a2);
void nox_xxx_updateSelfDestruct_53CC90(int a1);
void nox_xxx_updateBlackPowderBurn_53CCB0(int a1);
void nox_xxx_updatePixie_53CD20(int a1);
__int16 nox_xxx_pixieIdleAnimate_53CF90(int a1, float2* a2, __int16 a3);
__int16 sub_53D010(int a1, int a2);
void nox_xxx_updateDeathBall_53D080(int a1);
void sub_53D170(int a1, int a2);
void nox_xxx_updateDeathBallFragment_53D220(int a1);
void nox_xxx_updateMoonglow_53D270(int a1);
void nox_xxx_updateTelekinesis_53D330(int a1);
void nox_xxx_updateFist_53D400(int a1);
void nox_xxx_updateFlameCleanse_53D510(int a1);
void nox_xxx_updateMeteorShower_53D5A0(float* a2);
void nox_xxx_meteorExplode_53D6E0(int a6);
void nox_xxx_updateToxicCloud_53D850(int a1);
void sub_53D8C0(int a1, int a2);
void nox_xxx_updateSmallToxicCloud_53D960(int a1);
void nox_xxx_toxicCloudPoison_53D9D0(int a1, int a2);
void nox_xxx_updateArachnaphobia_53DA60(int* a1);
void nox_xxx_updateExpire_53DB00(int a1);
int* nox_xxx_updateBreak_53DB30(_DWORD* a1);
int* nox_xxx_updateOpen_53DBB0(_DWORD* a1);
void nox_xxx_updateBreakAndRemove_53DC30(_DWORD* a1);
void nox_xxx_updateChakramInMotion_53DCC0(int a1);
int nox_xxx_updateFlag_53DDF0(int a1);
int* nox_xxx_updateTrapDoor_53DE80(_DWORD* a1);
void nox_xxx_updateGameBall_53DF40(int a3);
void nox_xxx_updateUndeadKiller_53E190(int a1);
void nox_xxx_updateCrown_53E1D0(int a1);
char sub_53E2B0(int a1);
BOOL sub_53E2D0(int a1);
int nox_xxx_recalculateArmorVal_53E300(_DWORD* a1);
int sub_53E3A0(int a1, nox_object_t* object);
int sub_53E430(_DWORD* a1, nox_object_t* object, int a3, int a4);
int nox_xxx_NPCEquipArmor_53E520(int a1, _DWORD* a2);
void sub_53E600(_DWORD* a1);
int nox_xxx_playerEquipArmor_53E650(_DWORD* a1, nox_object_t* item, int a3, int a4);
_DWORD* nox_xxx_armorHaveSameSubclass_53E7B0(int a1, int a2);
int nox_xxx_pickupArmor_53E7F0(int a1, int a2, int a3, int a4);
void sub_53EAE0(int a1);
int nox_xxx_dropArmor_53EB70(int a1, _DWORD* a2, int* a3);
int nox_xxx_ItemIsDroppable_53EBF0(int a1);
CHAR* sub_53EC40();
BOOL sub_53EC80(int a1, int a2);
int nox_xxx_useMushroom_53ECE0(int a1, int a2);
int nox_xxx_useEnchant_53ED60(int a1, int a2);
int nox_xxx_useCast_53ED90(int a1, _DWORD* a2);
int nox_xxx_useConsume_53EE10(int a1, int a2);
int nox_xxx_useCiderConfuse_53EF00(int a1, int a2);
int nox_xxx_usePotion_53EF70(int a1, int a2);
int nox_xxx_useLesserFireballStaff_53F290(int a1, _DWORD* a2);
_DWORD* nox_xxx_wandShot_53F480(int a1, int a2, int* a3, _DWORD* a4);
int nox_xxx_useWandCastSpell_53F4F0(int a1, _DWORD* a2);
int nox_xxx_useFireWand_53F670(int a1, int a2);
int nox_xxx_useRead_53F7C0(int a1, int a2);
int sub_53F830(int a1, int a2);
int nox_xxx_useByNetCode_53F8E0(int a1, int a2);
int sub_53F930(int a1, int a2);
int nox_xxx_useSpellReward_53F9E0(int a1, int a2);
int nox_xxx_useAbilityReward_53FAE0(int a1, int a2);
_DWORD* nox_xxx_respawnPlayerImpl_53FBC0(float* a1, int a2);
void nox_xxx_createCorpse_53FCA0();
void nox_xxx_warriorBerserker_53FEB0(int a1);
void nox_xxx_warriorWarcry_53FF40(int a1);
int sub_53FFB0(_DWORD* a1, _DWORD* a2);
int nox_xxx_warriorHarpoon_540070(_DWORD* a1);
void nox_xxx_warriorTreadLightly_5400B0(_DWORD* a1, __int16 a2);
void nox_xxx_warriorInfravis_540110(int a1, __int16 a2);
int nox_xxx_castMissilesOM_540160(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_castPixies_540440(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_spellFlySearchTarget_540610(float2* a1, int a2, int a3, float a4, int a5, int a6);
void sub_540780(float* a1, int a2);
BOOL sub_5408A0(int a1);
int nox_xxx_mobCastInversion_5408D0(int a1);
int* nox_xxx_monsterCast_540A30(int a1, int a2, int a3);
int nox_xxx_unitIsMagicMissile_540B60(int a1, int a2);
int nox_xxx_monsterBuffSelf_540B90(int a1);
BOOL sub_540CE0(int a1, int a2);
BOOL sub_540D20(int a1);
int sub_540D40(int a1);
int nox_xxx_mobCastRelated2_540D90(int a1, int a2);
int nox_xxx_monsterCastOffensive_540F20(int a1, int a2);
int nox_xxx_mobCastRelated_541050(int a1);
int nox_xxx_mobHealSomeone_5411A0(int a3);
void nox_xxx_mobMayHealThis_5412A0(float* a1, int a2);
char nox_xxx_mobCast_541300(int a1, _DWORD* a2, int a3);
char nox_xxx_mobActionCastOnObj_541360(_DWORD* a1);
char nox_xxx_mobActionCast_5413B0(_DWORD* a1, int a2);
void nox_xxx_mobCastRandomRecoil_541490(int a1, float* a2, float2* a3);
char nox_xxx_mobActionCastOnPoint_541560(_DWORD* a1);
_DWORD* nox_xxx_mobActionCastStopMB_541590(int a1);
char nox_xxx_mobActionCastFinishMB_5415C0(int a1);
char nox_xxx_mobActionCastStart_5415F0(_DWORD* a1);
char sub_541630(int a1, int a2);
char* sub_542BF0(int a1, int a2, int a3);
char* sub_5435C0(int a1, int a2, int a3, int a4);
char* sub_543620(int a1, int a2);
int sub_543680(float* a1);
void nox_xxx_unused_543720(int a1, int a2, int a3, int a4);
int nox_xxx_unused_543770(_DWORD* a1, _DWORD* a2, _DWORD* a3, _DWORD* a4);
BOOL sub_5437E0(int* a1, int a2, int a3);
void sub_543BC0(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_tile_543C50(_DWORD* a1, int a2, int a3, int a4, int a5, int a6);
int sub_543E60(int a1, int a2);
int nox_xxx_mapGenEdge_543EB0(int a1, int a2);
int sub_543FB0(const char* a1);
int sub_544020(char* a1);
int nox_xxx_tileCheckByte3_544070(int a1);
int nox_xxx_tileCheckByte4_5440A0(int a1);
int sub_5440F0(int a1, signed int a2);
int sub_544240(int a1);
int sub_544270(const char* a1);
int sub_5442D0(float* a1);
int nox_xxx_tileSubtile_544310(float2* a1);
char nox_xxx_mobActionMoveTo_5443F0(int a1);
char nox_xxx_mobActionMoveToFar_5445C0(int* a1);
void nox_xxx_mobActionDodge_544640(int a1);
int sub_544740(int a1);
int sub_544750(int a1);
char nox_xxx_mobActionFlee_544760(int a1);
int nox_xxx_mobActionReturnToHome_544920(int a1);
int sub_544930(int a1);
int sub_544940(int a1);
char sub_544950(int a1);
char nox_xxx_mobActionWait_544960(int a1);
char nox_xxx_mobActionWaitRelative_544990(int a1);
int* nox_xxx_mobActionHunt_5449D0(int a1);
int nox_xxx_mobSearchEdible_544A00(int a1, float a2);
void nox_xxx_mobSearchEdible2_544A40(int a1, int a2);
int sub_544AE0(int a1, float a2);
void sub_544B20(int a1, int a2);
char nox_xxx_mobActionPickupObject_544B90(int a1);
int nox_xxx_mobGenericDeath_544C40(int a1);
void nox_xxx_zombieBurnDeleteCheck_544CA0(_DWORD* a1);
void nox_xxx_zombieBurnDelete_544CE0(_DWORD* a1);
char sub_544D60(int a1);
char nox_xxx_mobActionDead1_544D80(_DWORD* a1);
void nox_xxx_createReleasedSoul_544E60(int a1);
int sub_544F70(int a1);
void nox_xxx_mobActionDead2_544EC0(int a1);
unsigned __int8* nox_xxx_mobActionReportComplete_544FF0(int a1);


#endif // NOX_PORT_GAME4_3