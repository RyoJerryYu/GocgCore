package interfaces

import (
	"github.com/RyoJerryYu/GocgCore/common"
)

// field.h
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

// field.h
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

// field.h
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

// card.h
type MaterialInfo struct {
	// Synchron
	LimitTuner   Card
	LimitSyn     Group
	LimitSynMinc int32
	LimitSynMaxc int32
	// Xyz
	LimitXyz     Group
	LimitXyzMinc int32
	LimitXyzMaxc int32
	// Link
	LimitLink     Group
	LimitLinkCard Card
	LimitLinkMinc int32
	LimitLinkMaxc int32
}

func NewMaterialInfo() *MaterialInfo {
	// TODO: use default 0 value
	return &MaterialInfo{
		LimitTuner:    nil,
		LimitSyn:      nil,
		LimitSynMinc:  0,
		LimitSynMaxc:  0,
		LimitXyz:      nil,
		LimitXyzMinc:  0,
		LimitXyzMaxc:  0,
		LimitLink:     nil,
		LimitLinkCard: nil,
		LimitLinkMinc: 0,
		LimitLinkMaxc: 0,
	}
}

var NullInfo MaterialInfo

// card.h
type CardState struct {
	Code         uint32
	Code2        uint32
	SetCode      uint16
	Type         uint32
	Level        uint32
	Rank         uint32
	Link         uint32
	LScale       uint32
	RScale       uint32
	Attribute    uint32
	Race         uint32
	Attack       int32
	Defense      int32
	BaseAttack   int32
	BaseDefense  int32
	Controler    uint8
	Location     uint8
	Sequence     uint8
	Position     uint8
	Reason       uint32
	PZone        bool
	ReasonCard   Card
	ReasonPlayer uint8
	ReasonEffect Effect
}

func NewCardState() *CardState {
	// TODO: use default 0 value
	return &CardState{
		Code:         0,
		Code2:        0,
		SetCode:      0,
		Type:         0,
		Level:        0,
		Rank:         0,
		Link:         0,
		LScale:       0,
		RScale:       0,
		Attribute:    0,
		Race:         0,
		Attack:       0,
		Defense:      0,
		BaseAttack:   0,
		BaseDefense:  0,
		Controler:    common.PLAYER_NONE,
		Location:     0,
		Sequence:     0,
		Position:     0,
		Reason:       0,
		PZone:        false,
		ReasonCard:   nil,
		ReasonPlayer: common.PLAYER_NONE,
		ReasonEffect: nil,
	}
}

func (cs *CardState) IsLocation(loc int32) bool {
	if (loc&common.LOCATION_FZONE) > 0 &&
		cs.Location == common.LOCATION_SZONE &&
		cs.Sequence == 5 {
		return true
	}
	if (loc&common.LOCATION_PZONE) > 0 &&
		cs.Location == common.LOCATION_SZONE &&
		cs.PZone {
		return true
	}
	if (cs.Location & uint8(loc)) > 0 { // TODO: different Type ?
		return true
	}
	return false
}

// card.h
type CardData struct {
	Code       uint32
	Alias      uint32
	SetCode    uint64
	Type       uint32
	Level      uint32
	Attribute  uint32
	Race       uint32
	Attack     int32
	Defense    int32
	LScale     uint32
	RScale     uint32
	LinkMarker uint32
}

func (cd *CardData) Clear() {
	*cd = CardData{}
}
