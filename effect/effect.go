package effect

import "github.com/RyoJerryYu/GocgCore/common"

type (
	Card   interface{}
	Duel   interface{}
	Group  interface{}
	TEvent interface{}
)

// TODO: Maybe should type EffectSet interface{} ?

type EffectFlag uint32
type EffectFlag2 uint32

type Effect struct {
	RefHandle      int32
	PDual          Duel
	Owner          Card
	Handler        Card
	EffectOwner    uint8
	Description    uint32
	Code           uint32
	Flag           [2]uint32 // TODO: Flag could be a struct{EffectFlag, EffectFlag2}
	Id             uint32
	Type           uint16
	CopyId         uint16
	Range          uint16
	SRange         uint16
	ORange         uint16
	CountLimit     uint8
	CountLimitMax  uint8
	ResetCount     uint16
	ResetFlag      uint32
	CountCode      uint32
	Category       uint32
	HintTiming     [2]uint32
	CardType       uint32
	ActiveType     uint32
	ActiveLocation uint16
	ActiveSequence uint16
	ActiveHandler  Card
	Status         uint16
	Label          []uint32
	LabelObject    int32
	Condition      int32
	Cost           int32
	Target         int32
	Value          int32
	Operation      int32
}

// In Golang, there is no need to use keyword explicit.
// In Golang, there is no need to define a destructor.
func NewEffect(pd Duel) *Effect {
	// TODO: use default zero value instead
	return &Effect{
		RefHandle:      0,
		PDual:          pd,
		Owner:          nil,
		Handler:        nil,
		Description:    0,
		EffectOwner:    common.PLAYER_NONE,
		CardType:       0,
		ActiveType:     0,
		ActiveLocation: 0,
		ActiveSequence: 0,
		ActiveHandler:  nil,
		Id:             0,
		Code:           0,
		Type:           0,
		Flag:           [2]uint32{0, 0},
		CopyId:         0,
		Range:          0,
		SRange:         0,
		ORange:         0,
		CountLimit:     0,
		CountLimitMax:  0,
		ResetCount:     0,
		ResetFlag:      0,
		CountCode:      0,
		Category:       0,
		Label:          make([]uint32, 0, 4),
		LabelObject:    0,
		HintTiming:     [2]uint32{0, 0},
		Status:         0,
		Condition:      0,
		Cost:           0,
		Target:         0,
		Value:          0,
		Operation:      0,
	}
}

func EffectSortId(e1, e2 *Effect) bool {
	return e1.Id < e2.Id
}

// TODO: change the type of e.Flags, and use bit operation
func (e *Effect) IsFlag(flag EffectFlag) bool {
	return EffectFlag(e.Flag[0])&flag > 0
}

// rename for conflict
func (e *Effect) IsFlag2(flag EffectFlag2) bool {
	return EffectFlag2(e.Flag[1])&flag > 0
}
