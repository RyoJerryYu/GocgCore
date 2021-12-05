package field

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/interfaces"
)

type PlayerInfo struct {
	Lp               int32
	StartCount       int32
	DrawCount        int32
	UsedLocation     uint32
	DisabledLocation uint32
	ExtraPCount      uint32
	TagExtraPCount   uint32
	ListMZone        interfaces.CardVector
	ListSZone        interfaces.CardVector
	ListMain         interfaces.CardVector
	ListGrave        interfaces.CardVector
	ListHand         interfaces.CardVector
	ListRemove       interfaces.CardVector
	ListExtra        interfaces.CardVector
	TagListMain      interfaces.CardVector
	TagListHand      interfaces.CardVector
	TagListExtra     interfaces.CardVector
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
	effectContainer      map[uint32][]interfaces.Effect // multimap
	effectIndexer        map[interfaces.Effect]uint32   // TODO: How to sync iterator?
	oathEffects          map[interfaces.Effect]interfaces.Effect
	effectCollection     map[interfaces.Effect]struct{}
	gainEffects          map[interfaces.Card]interfaces.Effect
	grantEffectContainer map[interfaces.Effect]gainEffects
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

	DisableCheckList []interfaces.Card
	DisableCheckSet  map[interfaces.Card]struct{}

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
	PEffect interfaces.Effect
	PTarget interfaces.Group
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
