package interfaces

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
