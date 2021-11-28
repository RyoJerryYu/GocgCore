package field

//Location Use Reason
const (
	LOCATION_REASON_TOFIELD = 0x1
	LOCATION_REASON_CONTROL = 0x2
)

//Chain Info
const (
	CHAIN_DISABLE_ACTIVATE = 0x01
	CHAIN_DISABLE_EFFECT   = 0x02
	CHAIN_HAND_EFFECT      = 0x04
	CHAIN_CONTINUOUS_CARD  = 0x08
	CHAIN_ACTIVATING       = 0x10
	CHAIN_HAND_TRIGGER     = 0x20
	//#define CHAIN_DECK_EFFECT		0x40

	CHAININFO_CHAIN_COUNT          = 0x01
	CHAININFO_TRIGGERING_EFFECT    = 0x02
	CHAININFO_TRIGGERING_PLAYER    = 0x04
	CHAININFO_TRIGGERING_CONTROLER = 0x08
	CHAININFO_TRIGGERING_LOCATION  = 0x10
	CHAININFO_TRIGGERING_SEQUENCE  = 0x20
	CHAININFO_TARGET_CARDS         = 0x40
	CHAININFO_TARGET_PLAYER        = 0x80
	CHAININFO_TARGET_PARAM         = 0x100
	CHAININFO_DISABLE_REASON       = 0x200
	CHAININFO_DISABLE_PLAYER       = 0x400
	CHAININFO_CHAIN_ID             = 0x800
	CHAININFO_TYPE                 = 0x1000
	CHAININFO_EXTTYPE              = 0x2000
	CHAININFO_TRIGGERING_POSITION  = 0x4000
	CHAININFO_TRIGGERING_CODE      = 0x8000
	CHAININFO_TRIGGERING_CODE2     = 0x10000
	//#define CHAININFO_TRIGGERING_TYPE		0x20000
	CHAININFO_TRIGGERING_LEVEL     = 0x40000
	CHAININFO_TRIGGERING_RANK      = 0x80000
	CHAININFO_TRIGGERING_ATTRIBUTE = 0x100000
	CHAININFO_TRIGGERING_RACE      = 0x200000
	CHAININFO_TRIGGERING_ATTACK    = 0x400000
	CHAININFO_TRIGGERING_DEFENSE   = 0x800000
)

//Timing
const (
	TIMING_DRAW_PHASE      = 0x1
	TIMING_STANDBY_PHASE   = 0x2
	TIMING_MAIN_END        = 0x4
	TIMING_BATTLE_START    = 0x8
	TIMING_BATTLE_END      = 0x10
	TIMING_END_PHASE       = 0x20
	TIMING_SUMMON          = 0x40
	TIMING_SPSUMMON        = 0x80
	TIMING_FLIPSUMMON      = 0x100
	TIMING_MSET            = 0x200
	TIMING_SSET            = 0x400
	TIMING_POS_CHANGE      = 0x800
	TIMING_ATTACK          = 0x1000
	TIMING_DAMAGE_STEP     = 0x2000
	TIMING_DAMAGE_CAL      = 0x4000
	TIMING_CHAIN_END       = 0x8000
	TIMING_DRAW            = 0x10000
	TIMING_DAMAGE          = 0x20000
	TIMING_RECOVER         = 0x40000
	TIMING_DESTROY         = 0x80000
	TIMING_REMOVE          = 0x100000
	TIMING_TOHAND          = 0x200000
	TIMING_TODECK          = 0x400000
	TIMING_TOGRAVE         = 0x800000
	TIMING_BATTLE_PHASE    = 0x1000000
	TIMING_EQUIP           = 0x2000000
	TIMING_BATTLE_STEP_END = 0x4000000
	TIMING_BATTLED         = 0x8000000
)

const (
	GLOBALFLAG_DECK_REVERSE_CHECK = 0x1
	GLOBALFLAG_BRAINWASHING_CHECK = 0x2
	GLOBALFLAG_SCRAP_CHIMERA      = 0x4
	// GLOBALFLAG_DELAYED_QUICKEFFECT = 0x8
	GLOBALFLAG_DETACH_EVENT      = 0x10
	GLOBALFLAG_MUST_BE_SMATERIAL = 0x20
	GLOBALFLAG_SPSUMMON_COUNT    = 0x40
	GLOBALFLAG_XMAT_COUNT_LIMIT  = 0x80
	GLOBALFLAG_SELF_TOGRAVE      = 0x100
	GLOBALFLAG_SPSUMMON_ONCE     = 0x200
	GLOBALFLAG_TUNE_MAGICIAN     = 0x400
)

const (
	PROCESSOR_NONE    = 0
	PROCESSOR_WAITING = 0x10000
	PROCESSOR_END     = 0x20000
)

const (
	PROCESSOR_ADJUST               = 1
	PROCESSOR_HINT                 = 2
	PROCESSOR_TURN                 = 3
	PROCESSOR_WAIT                 = 4
	PROCESSOR_REFRESH_LOC          = 5
	PROCESSOR_SELECT_IDLECMD       = 10
	PROCESSOR_SELECT_EFFECTYN      = 11
	PROCESSOR_SELECT_BATTLECMD     = 12
	PROCESSOR_SELECT_YESNO         = 13
	PROCESSOR_SELECT_OPTION        = 14
	PROCESSOR_SELECT_CARD          = 15
	PROCESSOR_SELECT_CHAIN         = 16
	PROCESSOR_SELECT_UNSELECT_CARD = 17
	PROCESSOR_SELECT_PLACE         = 18
	PROCESSOR_SELECT_POSITION      = 19
	PROCESSOR_SELECT_TRIBUTE_P     = 20
	PROCESSOR_SELECT_COUNTER       = 22
	PROCESSOR_SELECT_SUM           = 23
	PROCESSOR_SELECT_DISFIELD      = 24
	PROCESSOR_SORT_CARD            = 25
	PROCESSOR_SELECT_RELEASE       = 26
	PROCESSOR_SELECT_TRIBUTE       = 27
	PROCESSOR_POINT_EVENT          = 30
	PROCESSOR_QUICK_EFFECT         = 31
	PROCESSOR_IDLE_COMMAND         = 32
	PROCESSOR_PHASE_EVENT          = 33
	PROCESSOR_BATTLE_COMMAND       = 34
	PROCESSOR_DAMAGE_STEP          = 35
	PROCESSOR_ADD_CHAIN            = 40
	PROCESSOR_SOLVE_CHAIN          = 42
	PROCESSOR_SOLVE_CONTINUOUS     = 43
	PROCESSOR_EXECUTE_COST         = 44
	PROCESSOR_EXECUTE_OPERATION    = 45
	PROCESSOR_EXECUTE_TARGET       = 46
	PROCESSOR_DESTROY              = 50
	PROCESSOR_RELEASE              = 51
	PROCESSOR_SENDTO               = 52
	PROCESSOR_MOVETOFIELD          = 53
	PROCESSOR_CHANGEPOS            = 54
	PROCESSOR_OPERATION_REPLACE    = 55
	PROCESSOR_DESTROY_REPLACE      = 56
	PROCESSOR_RELEASE_REPLACE      = 57
	PROCESSOR_SENDTO_REPLACE       = 58
	PROCESSOR_SUMMON_RULE          = 60
	PROCESSOR_SPSUMMON_RULE        = 61
	PROCESSOR_SPSUMMON             = 62
	PROCESSOR_FLIP_SUMMON          = 63
	PROCESSOR_MSET                 = 64
	PROCESSOR_SSET                 = 65
	PROCESSOR_SPSUMMON_STEP        = 66
	PROCESSOR_SSET_G               = 67
	PROCESSOR_DRAW                 = 70
	PROCESSOR_DAMAGE               = 71
	PROCESSOR_RECOVER              = 72
	PROCESSOR_EQUIP                = 73
	PROCESSOR_GET_CONTROL          = 74
	PROCESSOR_SWAP_CONTROL         = 75
	// PROCESSOR_CONTROL_ADJUST       = 76
	PROCESSOR_SELF_DESTROY        = 77
	PROCESSOR_TRAP_MONSTER_ADJUST = 78
	PROCESSOR_PAY_LPCOST          = 80
	PROCESSOR_REMOVE_COUNTER      = 81
	PROCESSOR_ATTACK_DISABLE      = 82
	PROCESSOR_ACTIVATE_EFFECT     = 83

	PROCESSOR_ANNOUNCE_RACE       = 110
	PROCESSOR_ANNOUNCE_ATTRIB     = 111
	PROCESSOR_ANNOUNCE_LEVEL      = 112
	PROCESSOR_ANNOUNCE_CARD       = 113
	PROCESSOR_ANNOUNCE_TYPE       = 114
	PROCESSOR_ANNOUNCE_NUMBER     = 115
	PROCESSOR_ANNOUNCE_COIN       = 116
	PROCESSOR_TOSS_DICE           = 117
	PROCESSOR_TOSS_COIN           = 118
	PROCESSOR_ROCK_PAPER_SCISSORS = 119

	PROCESSOR_SELECT_FUSION    = 131
	PROCESSOR_SELECT_SYNCHRO   = 132
	PROCESSOR_SELECT_XMATERIAL = 139
	PROCESSOR_DISCARD_HAND     = 150
	PROCESSOR_DISCARD_DECK     = 151
	PROCESSOR_SORT_DECK        = 152
	PROCESSOR_REMOVE_OVERLAY   = 160

//#define PROCESSOR_SORT_CHAIN		21
//#define PROCESSOR_DESTROY_S			100
//#define PROCESSOR_RELEASE_S			101
//#define PROCESSOR_SENDTO_S			102
//#define PROCESSOR_CHANGEPOS_S		103
//#define PROCESSOR_SELECT_YESNO_S	120
//#define PROCESSOR_SELECT_OPTION_S	121
//#define PROCESSOR_SELECT_CARD_S		122
//#define PROCESSOR_SELECT_EFFECTYN_S	123
//#define PROCESSOR_SELECT_UNSELECT_CARD_S	124
//#define PROCESSOR_SELECT_PLACE_S	125
//#define PROCESSOR_SELECT_POSITION_S	126
//#define PROCESSOR_SELECT_TRIBUTE_S	127
//#define PROCESSOR_SORT_CARDS_S		128
//#define PROCESSOR_SELECT_RELEASE_S	129
//#define PROCESSOR_SELECT_TARGET		130
//#define PROCESSOR_SELECT_SUM_S		133
//#define PROCESSOR_SELECT_DISFIELD_S	134
//#define PROCESSOR_SPSUMMON_S		135
//#define PROCESSOR_SPSUMMON_STEP_S	136
//#define PROCESSOR_SPSUMMON_COMP_S	137
//#define PROCESSOR_RANDOM_SELECT_S	138
//#define PROCESSOR_DRAW_S			140
//#define PROCESSOR_DAMAGE_S			141
//#define PROCESSOR_RECOVER_S			142
//#define PROCESSOR_EQUIP_S			143
//#define PROCESSOR_GET_CONTROL_S		144
//#define PROCESSOR_SWAP_CONTROL_S	145
//#define PROCESSOR_MOVETOFIELD_S		161
)
