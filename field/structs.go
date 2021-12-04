package field

import (
	"github.com/RyoJerryYu/GocgCore/common"
)

type TEvent struct {
	TriggerCard  Card
	EventCards   Group
	ReasonEffect Effect
	EventCode    uint32
	EventValue   uint32
	Reason       uint32
	EventPlayer  uint8
	ReasonPlayer uint8
}

func NewTEvent() *TEvent {
	// TODO: Change to default zero value
	return &TEvent{
		TriggerCard:  nil,
		EventCards:   nil,
		ReasonEffect: nil,
		EventCode:    0,
		EventValue:   0,
		Reason:       0,
		EventPlayer:  common.PLAYER_NONE,
		ReasonPlayer: common.PLAYER_NONE,
	}
}

func (te *TEvent) LessThan(v *TEvent) bool {
	return false
}

type OpTarget struct {
	OpCards  Group
	OpCount  uint8
	OpPlayer uint8
	OpParam  uint32
}

func NewOpTarget() *OpTarget {
	// TODO: Change to default zero value
	return &OpTarget{
		OpCards:  nil,
		OpCount:  0,
		OpPlayer: common.PLAYER_NONE,
		OpParam:  0,
	}
}

type opMap map[uint32]*OpTarget

type Chain struct {
	ChainId             uint16
	ChainCount          uint8
	TriggeringPlayer    uint8
	TriggeringControler uint8
	TriggeringLocation  uint16
	TriggeringSequence  uint8
	TriggeringPosition  uint8
	TriggeringState     CardState
	TriggeringEffect    Effect
	TargetCards         Group
	ReplaceOp           int32
	TargetPlayer        uint8
	TargetParam         int32
	DisableReason       Effect
	DisablePlayer       uint8
	Evt                 TEvent
	OpInfos             opMap
	Flag                uint32
}

func NewChain() *Chain {
	return &Chain{
		ChainId:             0,
		ChainCount:          0,
		TriggeringPlayer:    common.PLAYER_NONE,
		TriggeringControler: common.PLAYER_NONE,
		TriggeringLocation:  0,
		TriggeringSequence:  0,
		TriggeringPosition:  0,
		TriggeringState:     *NewCardState(), // TODO: use default value
		TriggeringEffect:    nil,
		TargetCards:         nil,
		ReplaceOp:           0,
		TargetPlayer:        common.PLAYER_NONE,
		TargetParam:         0,
		DisableReason:       nil,
		DisablePlayer:       common.PLAYER_NONE,
		Evt:                 *NewTEvent(), // TODO: USE DEFAULT VALUE
		Flag:                0,
	}
}

type cardVector []Card

type PlayerInfo struct {
	Lp               int32
	StartCount       int32
	DrawCount        int32
	UsedLocation     uint32
	DisabledLocation uint32
	ExtraPCount      uint32
	TagExtraPCount   uint32
	ListMZone        cardVector
	ListSZone        cardVector
	ListMain         cardVector
	ListGrave        cardVector
	ListHand         cardVector
	ListRemove       cardVector
	ListExtra        cardVector
	TagListMain      cardVector
	TagListHand      cardVector
	TagListExtra     cardVector
}

func NewPlayerInfo() *PlayerInfo {
	return &PlayerInfo{
		Lp:               0,
		StartCount:       0,
		DrawCount:        0,
		UsedLocation:     0,
		DisabledLocation: 0,
		ExtraPCount:      0,
		TagExtraPCount:   0,
	}
}

type (
	effectContainer      map[uint32][]Effect // multimap
	effectIndexer        map[Effect]uint32   // TODO: How to sync iterator?
	oathEffects          map[Effect]Effect
	effectCollection     map[Effect]struct{}
	gainEffects          map[Card]Effect
	grantEffectContainer map[Effect]gainEffects
)

type FieldEffect struct {
	AuraEffect       effectContainer
	IgnitionEffect   effectContainer
	ActivateEffect   effectContainer
	TriggerOEffect   effectContainer
	TriggerFEffect   effectContainer
	QuickOEffect     effectContainer
	QuickFEffect     effectContainer
	ContinuousEffect effectContainer
	Indexer          effectIndexer
	Oath             oathEffects
	PhEff            effectCollection
	ChEff            effectCollection
	Rechargeable     effectCollection
	SpSummonCountEff effectCollection

	DisableCheckList []Card
	DisableCheckSet  map[Card]struct{}

	GrantEffect grantEffectContainer
}

type fieldInfo struct {
	FieldId        int32
	CopyId         int16
	TurnId         int16
	TurnIdByPlayer [2]int16
	CardId         int16
	Phase          uint16
	TurnPlayer     uint8
	Priorities     [2]uint8
	CanShuffle     uint8
}

func NewFieldInfo() *fieldInfo {
	return &fieldInfo{
		TurnIdByPlayer: [2]int16{0, 0},
		Priorities:     [2]uint8{0, 0},
		CanShuffle:     common.TRUE,
	}
}

type LpCost struct {
	Count   int32
	Amount  int32
	LpStack [8]int32
}

type ProcessorUnit struct {
	Type    uint16
	Step    uint16
	PEffect Effect
	PTarget Group
	Arg1    common.Ptr
	Arg2    common.Ptr
	Arg3    common.Ptr
	Arg4    common.Ptr
	Ptr1    interface{}
	Ptr2    interface{}
}

// TODO: sync union, check
type ReturnValue [64]byte

func (rv *ReturnValue) AsBValue() [64]int8 {
	var bv [64]int8
	for i := 0; i < 64; i++ {
		bv[i] = int8(rv[i])
	}
	return bv
}

func (rv *ReturnValue) AsSValue() [32]int16 {
	var sv [32]int16
	for i := 0; i < 32; i++ {
		sv[i] = int16(rv[i*2]) | int16(rv[i*2+1])<<8
	}
	return sv
}

func (rv *ReturnValue) AsIValue() [16]int32 {
	var iv [16]int32
	for i := 0; i < 16; i++ {
		iv[i] = int32(rv[i*4]) |
			int32(rv[i*4+1])<<8 |
			int32(rv[i*4+2])<<16 |
			int32(rv[i*4+3])<<24
	}
	return iv
}

func (rv *ReturnValue) AsLValue() [8]int64 {
	var lv [8]int64
	for i := 0; i < 8; i++ {
		lv[i] = int64(rv[i*8]) |
			int64(rv[i*8+1])<<8 |
			int64(rv[i*8+2])<<16 |
			int64(rv[i*8+3])<<24 |
			int64(rv[i*8+4])<<32 |
			int64(rv[i*8+5])<<40 |
			int64(rv[i*8+6])<<48 |
			int64(rv[i*8+7])<<56
	}
	return lv
}
