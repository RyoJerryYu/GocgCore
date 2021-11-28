package effect

const (
	EFFECT_STATUS_AVAILABLE = 0x0001
	// EFFECT_STATUS_ACTIVATED = 0x0002
	EFFECT_STATUS_SPSELF = 0x0004
)

const (
	EFFECT_COUNT_CODE_OATH = 0x10000000
	EFFECT_COUNT_CODE_DUEL = 0x20000000
)

//========== Reset ==========
const (
	RESET_SELF_TURN = 0x10000000
	RESET_OPPO_TURN = 0x20000000
	RESET_PHASE     = 0x40000000
	RESET_CHAIN     = 0x80000000
	RESET_EVENT     = 0x1000
	RESET_CARD      = 0x2000
	RESET_CODE      = 0x4000
	RESET_COPY      = 0x8000

	RESET_DISABLE     = 0x00010000
	RESET_TURN_SET    = 0x00020000
	RESET_TOGRAVE     = 0x00040000
	RESET_REMOVE      = 0x00080000
	RESET_TEMP_REMOVE = 0x00100000
	RESET_TOHAND      = 0x00200000
	RESET_TODECK      = 0x00400000
	RESET_LEAVE       = 0x00800000
	RESET_TOFIELD     = 0x01000000
	RESET_CONTROL     = 0x02000000
	RESET_OVERLAY     = 0x04000000
	RESET_MSCHANGE    = 0x08000000
)

//========== Types ==========
const (
	EFFECT_TYPE_SINGLE     = 0x0001
	EFFECT_TYPE_FIELD      = 0x0002
	EFFECT_TYPE_EQUIP      = 0x0004
	EFFECT_TYPE_ACTIONS    = 0x0008
	EFFECT_TYPE_ACTIVATE   = 0x0010
	EFFECT_TYPE_FLIP       = 0x0020
	EFFECT_TYPE_IGNITION   = 0x0040
	EFFECT_TYPE_TRIGGER_O  = 0x0080
	EFFECT_TYPE_QUICK_O    = 0x0100
	EFFECT_TYPE_TRIGGER_F  = 0x0200
	EFFECT_TYPE_QUICK_F    = 0x0400
	EFFECT_TYPE_CONTINUOUS = 0x0800
	EFFECT_TYPE_XMATERIAL  = 0x1000
	EFFECT_TYPE_GRANT      = 0x2000
	EFFECT_TYPE_TARGET     = 0x4000
)

// TODO: Change to Golang build-in 1<<iota
//========== Flags ==========
const (
	EFFECT_FLAG_INITIAL           EffectFlag = 0x0001
	EFFECT_FLAG_FUNC_VALUE        EffectFlag = 0x0002
	EFFECT_FLAG_COUNT_LIMIT       EffectFlag = 0x0004
	EFFECT_FLAG_FIELD_ONLY        EffectFlag = 0x0008
	EFFECT_FLAG_CARD_TARGET       EffectFlag = 0x0010
	EFFECT_FLAG_IGNORE_RANGE      EffectFlag = 0x0020
	EFFECT_FLAG_ABSOLUTE_TARGET   EffectFlag = 0x0040
	EFFECT_FLAG_IGNORE_IMMUNE     EffectFlag = 0x0080
	EFFECT_FLAG_SET_AVAILABLE     EffectFlag = 0x0100
	EFFECT_FLAG_CANNOT_NEGATE     EffectFlag = 0x0200
	EFFECT_FLAG_CANNOT_DISABLE    EffectFlag = 0x0400
	EFFECT_FLAG_PLAYER_TARGET     EffectFlag = 0x0800
	EFFECT_FLAG_BOTH_SIDE         EffectFlag = 0x1000
	EFFECT_FLAG_COPY_INHERIT      EffectFlag = 0x2000
	EFFECT_FLAG_DAMAGE_STEP       EffectFlag = 0x4000
	EFFECT_FLAG_DAMAGE_CAL        EffectFlag = 0x8000
	EFFECT_FLAG_DELAY             EffectFlag = 0x10000
	EFFECT_FLAG_SINGLE_RANGE      EffectFlag = 0x20000
	EFFECT_FLAG_UNCOPYABLE        EffectFlag = 0x40000
	EFFECT_FLAG_OATH              EffectFlag = 0x80000
	EFFECT_FLAG_SPSUM_PARAM       EffectFlag = 0x100000
	EFFECT_FLAG_REPEAT            EffectFlag = 0x200000
	EFFECT_FLAG_NO_TURN_RESET     EffectFlag = 0x400000
	EFFECT_FLAG_EVENT_PLAYER      EffectFlag = 0x800000
	EFFECT_FLAG_OWNER_RELATE      EffectFlag = 0x1000000
	EFFECT_FLAG_CANNOT_INACTIVATE EffectFlag = 0x2000000
	EFFECT_FLAG_CLIENT_HINT       EffectFlag = 0x4000000
	EFFECT_FLAG_CONTINUOUS_TARGET EffectFlag = 0x8000000
	EFFECT_FLAG_LIMIT_ZONE        EffectFlag = 0x10000000
	// EFFECT_FLAG_COF               EffectFlag = 0x20000000
	// EFFECT_FLAG_CVAL_CHECK        EffectFlag = 0x40000000
	EFFECT_FLAG_IMMEDIATELY_APPLY EffectFlag = 0x80000000
)

const (
	EFFECT_FLAG2_MILLENNIUM_RESTRICT EffectFlag2 = 0x0001
	EFFECT_FLAG2_COF                 EffectFlag2 = 0x0002
)

// Golang have build-in type interface and need not to define operattor|

//========== Codes ==========
const (
	EFFECT_IMMUNE_EFFECT              = 1
	EFFECT_DISABLE                    = 2
	EFFECT_CANNOT_DISABLE             = 3
	EFFECT_SET_CONTROL                = 4
	EFFECT_CANNOT_CHANGE_CONTROL      = 5
	EFFECT_CANNOT_ACTIVATE            = 6
	EFFECT_CANNOT_TRIGGER             = 7
	EFFECT_DISABLE_EFFECT             = 8
	EFFECT_DISABLE_CHAIN              = 9
	EFFECT_DISABLE_TRAPMONSTER        = 10
	EFFECT_CANNOT_INACTIVATE          = 12
	EFFECT_CANNOT_DISEFFECT           = 13
	EFFECT_CANNOT_CHANGE_POSITION     = 14
	EFFECT_TRAP_ACT_IN_HAND           = 15
	EFFECT_TRAP_ACT_IN_SET_TURN       = 16
	EFFECT_REMAIN_FIELD               = 17
	EFFECT_MONSTER_SSET               = 18
	EFFECT_CANNOT_SUMMON              = 20
	EFFECT_CANNOT_FLIP_SUMMON         = 21
	EFFECT_CANNOT_SPECIAL_SUMMON      = 22
	EFFECT_CANNOT_MSET                = 23
	EFFECT_CANNOT_SSET                = 24
	EFFECT_CANNOT_DRAW                = 25
	EFFECT_CANNOT_DISABLE_SUMMON      = 26
	EFFECT_CANNOT_DISABLE_SPSUMMON    = 27
	EFFECT_SET_SUMMON_COUNT_LIMIT     = 28
	EFFECT_EXTRA_SUMMON_COUNT         = 29
	EFFECT_SPSUMMON_CONDITION         = 30
	EFFECT_REVIVE_LIMIT               = 31
	EFFECT_SUMMON_PROC                = 32
	EFFECT_LIMIT_SUMMON_PROC          = 33
	EFFECT_SPSUMMON_PROC              = 34
	EFFECT_EXTRA_SET_COUNT            = 35
	EFFECT_SET_PROC                   = 36
	EFFECT_LIMIT_SET_PROC             = 37
	EFFECT_DIVINE_LIGHT               = 38
	EFFECT_CANNOT_DISABLE_FLIP_SUMMON = 39
	EFFECT_INDESTRUCTABLE             = 40
	EFFECT_INDESTRUCTABLE_EFFECT      = 41
	EFFECT_INDESTRUCTABLE_BATTLE      = 42
	EFFECT_UNRELEASABLE_SUM           = 43
	EFFECT_UNRELEASABLE_NONSUM        = 44
	EFFECT_DESTROY_SUBSTITUTE         = 45
	EFFECT_CANNOT_RELEASE             = 46
	EFFECT_INDESTRUCTABLE_COUNT       = 47
	EFFECT_UNRELEASABLE_EFFECT        = 48
	EFFECT_DESTROY_REPLACE            = 50
	EFFECT_RELEASE_REPLACE            = 51
	EFFECT_SEND_REPLACE               = 52
	EFFECT_CANNOT_DISCARD_HAND        = 55
	EFFECT_CANNOT_DISCARD_DECK        = 56
	EFFECT_CANNOT_USE_AS_COST         = 57
	EFFECT_CANNOT_PLACE_COUNTER       = 58
	EFFECT_CANNOT_TO_GRAVE_AS_COST    = 59
	EFFECT_LEAVE_FIELD_REDIRECT       = 60
	EFFECT_TO_HAND_REDIRECT           = 61
	EFFECT_TO_DECK_REDIRECT           = 62
	EFFECT_TO_GRAVE_REDIRECT          = 63
	EFFECT_REMOVE_REDIRECT            = 64
	EFFECT_CANNOT_TO_HAND             = 65
	EFFECT_CANNOT_TO_DECK             = 66
	EFFECT_CANNOT_REMOVE              = 67
	EFFECT_CANNOT_TO_GRAVE            = 68
	EFFECT_CANNOT_TURN_SET            = 69
	EFFECT_CANNOT_BE_BATTLE_TARGET    = 70
	EFFECT_CANNOT_BE_EFFECT_TARGET    = 71
	EFFECT_IGNORE_BATTLE_TARGET       = 72
	EFFECT_CANNOT_DIRECT_ATTACK       = 73
	EFFECT_DIRECT_ATTACK              = 74
	EFFECT_DUAL_STATUS                = 75
	EFFECT_EQUIP_LIMIT                = 76
	EFFECT_DUAL_SUMMONABLE            = 77
	EFFECT_UNION_LIMIT                = 78
	EFFECT_REVERSE_DAMAGE             = 80
	EFFECT_REVERSE_RECOVER            = 81
	EFFECT_CHANGE_DAMAGE              = 82
	EFFECT_REFLECT_DAMAGE             = 83
	EFFECT_CANNOT_ATTACK              = 85
	EFFECT_CANNOT_ATTACK_ANNOUNCE     = 86
	EFFECT_CANNOT_CHANGE_POS_E        = 87
	EFFECT_ACTIVATE_COST              = 90
	EFFECT_SUMMON_COST                = 91
	EFFECT_SPSUMMON_COST              = 92
	EFFECT_FLIPSUMMON_COST            = 93
	EFFECT_MSET_COST                  = 94
	EFFECT_SSET_COST                  = 95
	EFFECT_ATTACK_COST                = 96

	EFFECT_UPDATE_ATTACK     = 100
	EFFECT_SET_ATTACK        = 101
	EFFECT_SET_ATTACK_FINAL  = 102
	EFFECT_SET_BASE_ATTACK   = 103
	EFFECT_UPDATE_DEFENSE    = 104
	EFFECT_SET_DEFENSE       = 105
	EFFECT_SET_DEFENSE_FINAL = 106
	EFFECT_SET_BASE_DEFENSE  = 107
	EFFECT_REVERSE_UPDATE    = 108
	EFFECT_SWAP_AD           = 109
	EFFECT_SWAP_BASE_AD      = 110
	// EFFECT_SWAP_ATTACK_FINAL    = 111
	// EFFECT_SWAP_DEFENSE_FINAL   = 112
	EFFECT_ADD_CODE             = 113
	EFFECT_CHANGE_CODE          = 114
	EFFECT_ADD_TYPE             = 115
	EFFECT_REMOVE_TYPE          = 116
	EFFECT_CHANGE_TYPE          = 117
	EFFECT_ADD_RACE             = 120
	EFFECT_REMOVE_RACE          = 121
	EFFECT_CHANGE_RACE          = 122
	EFFECT_ADD_ATTRIBUTE        = 125
	EFFECT_REMOVE_ATTRIBUTE     = 126
	EFFECT_CHANGE_ATTRIBUTE     = 127
	EFFECT_UPDATE_LEVEL         = 130
	EFFECT_CHANGE_LEVEL         = 131
	EFFECT_UPDATE_RANK          = 132
	EFFECT_CHANGE_RANK          = 133
	EFFECT_UPDATE_LSCALE        = 134
	EFFECT_CHANGE_LSCALE        = 135
	EFFECT_UPDATE_RSCALE        = 136
	EFFECT_CHANGE_RSCALE        = 137
	EFFECT_SET_POSITION         = 140
	EFFECT_SELF_DESTROY         = 141
	EFFECT_SELF_TOGRAVE         = 142
	EFFECT_DOUBLE_TRIBUTE       = 150
	EFFECT_DECREASE_TRIBUTE     = 151
	EFFECT_DECREASE_TRIBUTE_SET = 152
	EFFECT_EXTRA_RELEASE        = 153
	EFFECT_TRIBUTE_LIMIT        = 154
	EFFECT_EXTRA_RELEASE_SUM    = 155
	// EFFECT_TRIPLE_TRIBUTE       = 156
	EFFECT_ADD_EXTRA_TRIBUTE    = 157
	EFFECT_EXTRA_RELEASE_NONSUM = 158
	EFFECT_PUBLIC               = 160
	EFFECT_COUNTER_PERMIT       = 0x10000
	EFFECT_COUNTER_LIMIT        = 0x20000
	EFFECT_RCOUNTER_REPLACE     = 0x30000
	EFFECT_LPCOST_CHANGE        = 170
	EFFECT_LPCOST_REPLACE       = 171
	EFFECT_SKIP_DP              = 180
	EFFECT_SKIP_SP              = 181
	EFFECT_SKIP_M1              = 182
	EFFECT_SKIP_BP              = 183
	EFFECT_SKIP_M2              = 184
	EFFECT_CANNOT_BP            = 185
	EFFECT_CANNOT_M2            = 186
	EFFECT_CANNOT_EP            = 187
	EFFECT_SKIP_TURN            = 188
	EFFECT_SKIP_EP              = 189
	EFFECT_DEFENSE_ATTACK       = 190
	EFFECT_MUST_ATTACK          = 191
	EFFECT_FIRST_ATTACK         = 192
	EFFECT_ATTACK_ALL           = 193
	EFFECT_EXTRA_ATTACK         = 194
	// EFFECT_MUST_BE_ATTACKED               = 195
	EFFECT_ONLY_BE_ATTACKED               = 196
	EFFECT_ATTACK_DISABLED                = 197
	EFFECT_NO_BATTLE_DAMAGE               = 200
	EFFECT_AVOID_BATTLE_DAMAGE            = 201
	EFFECT_REFLECT_BATTLE_DAMAGE          = 202
	EFFECT_PIERCE                         = 203
	EFFECT_BATTLE_DESTROY_REDIRECT        = 204
	EFFECT_BATTLE_DAMAGE_TO_EFFECT        = 205
	EFFECT_BOTH_BATTLE_DAMAGE             = 206
	EFFECT_ALSO_BATTLE_DAMAGE             = 207
	EFFECT_CHANGE_BATTLE_DAMAGE           = 208
	EFFECT_TOSS_COIN_REPLACE              = 220
	EFFECT_TOSS_DICE_REPLACE              = 221
	EFFECT_FUSION_MATERIAL                = 230
	EFFECT_CHAIN_MATERIAL                 = 231
	EFFECT_SYNCHRO_MATERIAL               = 232
	EFFECT_XYZ_MATERIAL                   = 233
	EFFECT_FUSION_SUBSTITUTE              = 234
	EFFECT_CANNOT_BE_FUSION_MATERIAL      = 235
	EFFECT_CANNOT_BE_SYNCHRO_MATERIAL     = 236
	EFFECT_SYNCHRO_MATERIAL_CUSTOM        = 237
	EFFECT_CANNOT_BE_XYZ_MATERIAL         = 238
	EFFECT_CANNOT_BE_LINK_MATERIAL        = 239
	EFFECT_SYNCHRO_LEVEL                  = 240
	EFFECT_RITUAL_LEVEL                   = 241
	EFFECT_XYZ_LEVEL                      = 242
	EFFECT_EXTRA_RITUAL_MATERIAL          = 243
	EFFECT_NONTUNER                       = 244
	EFFECT_OVERLAY_REMOVE_REPLACE         = 245
	EFFECT_SCRAP_CHIMERA                  = 246
	EFFECT_TUNE_MAGICIAN_X                = 247
	EFFECT_PRE_MONSTER                    = 250
	EFFECT_MATERIAL_CHECK                 = 251
	EFFECT_DISABLE_FIELD                  = 260
	EFFECT_USE_EXTRA_MZONE                = 261
	EFFECT_USE_EXTRA_SZONE                = 262
	EFFECT_MAX_MZONE                      = 263
	EFFECT_MAX_SZONE                      = 264
	EFFECT_MUST_USE_MZONE                 = 265
	EFFECT_HAND_LIMIT                     = 270
	EFFECT_DRAW_COUNT                     = 271
	EFFECT_SPIRIT_DONOT_RETURN            = 280
	EFFECT_SPIRIT_MAYNOT_RETURN           = 281
	EFFECT_CHANGE_ENVIRONMENT             = 290
	EFFECT_NECRO_VALLEY                   = 291
	EFFECT_FORBIDDEN                      = 292
	EFFECT_NECRO_VALLEY_IM                = 293
	EFFECT_REVERSE_DECK                   = 294
	EFFECT_REMOVE_BRAINWASHING            = 295
	EFFECT_BP_TWICE                       = 296
	EFFECT_UNIQUE_CHECK                   = 297
	EFFECT_MATCH_KILL                     = 300
	EFFECT_SYNCHRO_CHECK                  = 310
	EFFECT_QP_ACT_IN_NTPHAND              = 311
	EFFECT_MUST_BE_SMATERIAL              = 312
	EFFECT_TO_GRAVE_REDIRECT_CB           = 313
	EFFECT_CHANGE_INVOLVING_BATTLE_DAMAGE = 314
	// EFFECT_CHANGE_RANK_FINAL              = 315
	EFFECT_MUST_BE_FMATERIAL           = 316
	EFFECT_MUST_BE_XMATERIAL           = 317
	EFFECT_MUST_BE_LMATERIAL           = 318
	EFFECT_SPSUMMON_PROC_G             = 320
	EFFECT_SPSUMMON_COUNT_LIMIT        = 330
	EFFECT_LEFT_SPSUMMON_COUNT         = 331
	EFFECT_CANNOT_SELECT_BATTLE_TARGET = 332
	EFFECT_CANNOT_SELECT_EFFECT_TARGET = 333
	EFFECT_ADD_SETCODE                 = 334
	EFFECT_NO_EFFECT_DAMAGE            = 335
	// EFFECT_UNSUMMONABLE_CARD           = 336
	EFFECT_DISCARD_COST_CHANGE   = 338
	EFFECT_HAND_SYNCHRO          = 339
	EFFECT_ADD_FUSION_CODE       = 340
	EFFECT_ADD_FUSION_SETCODE    = 341
	EFFECT_ONLY_ATTACK_MONSTER   = 343
	EFFECT_MUST_ATTACK_MONSTER   = 344
	EFFECT_PATRICIAN_OF_DARKNESS = 345
	EFFECT_EXTRA_ATTACK_MONSTER  = 346
	EFFECT_UNION_STATUS          = 347
	EFFECT_OLDUNION_STATUS       = 348
	// EFFECT_ADD_FUSION_ATTRIBUTE    = 349
	// EFFECT_REMOVE_FUSION_ATTRIBUTE = 350
	EFFECT_CHANGE_FUSION_ATTRIBUTE = 351
	EFFECT_EXTRA_FUSION_MATERIAL   = 352
	EFFECT_TUNER_MATERIAL_LIMIT    = 353
	EFFECT_ADD_LINK_CODE           = 354
	// EFFECT_ADD_LINK_SETCODE        = 355
	EFFECT_ADD_LINK_ATTRIBUTE      = 356
	EFFECT_ADD_LINK_RACE           = 357
	EFFECT_EXTRA_LINK_MATERIAL     = 358
	EFFECT_QP_ACT_IN_SET_TURN      = 359
	EFFECT_EXTRA_PENDULUM_SUMMON   = 360
	EFFECT_MATERIAL_LIMIT          = 361
	EFFECT_SET_BATTLE_ATTACK       = 362
	EFFECT_SET_BATTLE_DEFENSE      = 363
	EFFECT_OVERLAY_RITUAL_MATERIAL = 364
	EFFECT_CHANGE_GRAVE_ATTRIBUTE  = 365
	EFFECT_CHANGE_GRAVE_RACE       = 366
)

const (
	EVENT_STARTUP          = 1000
	EVENT_FLIP             = 1001
	EVENT_FREE_CHAIN       = 1002
	EVENT_DESTROY          = 1010
	EVENT_REMOVE           = 1011
	EVENT_TO_HAND          = 1012
	EVENT_TO_DECK          = 1013
	EVENT_TO_GRAVE         = 1014
	EVENT_LEAVE_FIELD      = 1015
	EVENT_CHANGE_POS       = 1016
	EVENT_RELEASE          = 1017
	EVENT_DISCARD          = 1018
	EVENT_LEAVE_FIELD_P    = 1019
	EVENT_CHAIN_SOLVING    = 1020
	EVENT_CHAIN_ACTIVATING = 1021
	EVENT_CHAIN_SOLVED     = 1022
	// EVENT_CHAIN_ACTIVATED	= 1023
	EVENT_CHAIN_NEGATED        = 1024
	EVENT_CHAIN_DISABLED       = 1025
	EVENT_CHAIN_END            = 1026
	EVENT_CHAINING             = 1027
	EVENT_BECOME_TARGET        = 1028
	EVENT_DESTROYED            = 1029
	EVENT_MOVE                 = 1030
	EVENT_LEAVE_GRAVE          = 1031
	EVENT_ADJUST               = 1040
	EVENT_BREAK_EFFECT         = 1050
	EVENT_SUMMON_SUCCESS       = 1100
	EVENT_FLIP_SUMMON_SUCCESS  = 1101
	EVENT_SPSUMMON_SUCCESS     = 1102
	EVENT_SUMMON               = 1103
	EVENT_FLIP_SUMMON          = 1104
	EVENT_SPSUMMON             = 1105
	EVENT_MSET                 = 1106
	EVENT_SSET                 = 1107
	EVENT_BE_MATERIAL          = 1108
	EVENT_BE_PRE_MATERIAL      = 1109
	EVENT_DRAW                 = 1110
	EVENT_DAMAGE               = 1111
	EVENT_RECOVER              = 1112
	EVENT_PREDRAW              = 1113
	EVENT_SUMMON_NEGATED       = 1114
	EVENT_FLIP_SUMMON_NEGATED  = 1115
	EVENT_SPSUMMON_NEGATED     = 1116
	EVENT_CONTROL_CHANGED      = 1120
	EVENT_EQUIP                = 1121
	EVENT_ATTACK_ANNOUNCE      = 1130
	EVENT_BE_BATTLE_TARGET     = 1131
	EVENT_BATTLE_START         = 1132
	EVENT_BATTLE_CONFIRM       = 1133
	EVENT_PRE_DAMAGE_CALCULATE = 1134
	// EVENT_DAMAGE_CALCULATING	= 1135
	EVENT_PRE_BATTLE_DAMAGE = 1136
	// EVENT_BATTLE_END			= 1137
	EVENT_BATTLED           = 1138
	EVENT_BATTLE_DESTROYING = 1139
	EVENT_BATTLE_DESTROYED  = 1140
	EVENT_DAMAGE_STEP_END   = 1141
	EVENT_ATTACK_DISABLED   = 1142
	EVENT_BATTLE_DAMAGE     = 1143
	EVENT_TOSS_DICE         = 1150
	EVENT_TOSS_COIN         = 1151
	EVENT_TOSS_COIN_NEGATE  = 1152
	EVENT_TOSS_DICE_NEGATE  = 1153
	EVENT_LEVEL_UP          = 1200
	EVENT_PAY_LPCOST        = 1201
	EVENT_DETACH_MATERIAL   = 1202
	EVENT_TURN_END          = 1210
	EVENT_PHASE             = 0x1000
	EVENT_PHASE_START       = 0x2000
	EVENT_ADD_COUNTER       = 0x10000
	EVENT_REMOVE_COUNTER    = 0x20000
	EVENT_CUSTOM            = 0x10000000
)

const (
	DOUBLE_DAMAGE = 0x80000000
	HALF_DAMAGE   = 0x80000001
)

// The type of bit field in code
const (
	CODE_CUSTOM  = 1 // header + id (28 bits)
	CODE_COUNTER = 2 // header + counter_id (16 bits)
	CODE_PHASE   = 3 // header + phase_id (12 bits)
	CODE_VALUE   = 4 // numeric value, max = 4095
)

// To Symulate const unordered_set
func ContinuousEvent() map[uint32]struct{} {
	return map[uint32]struct{}{
		EVENT_ADJUST:       {},
		EVENT_BREAK_EFFECT: {},
		EVENT_TURN_END:     {},
	}
}

// TODO: Maybe we should define something like bit mask for readability
func IsContinuousEvent(code uint32) bool {
	if code&EVENT_CUSTOM > 0 {
		// codes start with 1 after 28 bits all belong to EVENT_CUSTOM
		return false
	} else if code&0xf0000 > 0 {
		// code is a counter
		return false
	} else if code&0xf000 > 0 {
		// code is a phase
		// !! means to boolean in C++. It's no need to do this in Golang
		return code&EVENT_PHASE_START > 0
	} else {
		// use set in golang
		_, ok := ContinuousEvent()[code]
		return ok
	}
}