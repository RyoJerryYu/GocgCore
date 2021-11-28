package card

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/group"
	"github.com/RyoJerryYu/GocgCore/interfaces"
)

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
	ReasonCard   *Card
	ReasonPlayer uint8
	ReasonEffect interfaces.Effect
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

type QueryCache struct {
	Code        uint32
	Alias       uint32
	Type        uint32
	Level       uint32
	Rank        uint32
	Link        uint32
	Attribute   uint32
	Race        uint32
	Attack      int32
	Defense     int32
	BaseAttack  int32
	BaseDefense int32
	Reason      uint32
	Status      int32
	LScale      uint32
	RScale      uint32
	LinkMarker  uint32
}

type MaterialInfo struct {
	// Synchron
	LimitTuner   *Card
	LimitSyn     *group.Group
	LimitSynMinc int32
	LimitSynMaxc int32
	// Xyz
	LimitXyz     *group.Group
	LimitXyzMinc int32
	LimitXyzMaxc int32
	// Link
	LimitLink     *group.Group
	LimitLinkCard *Card
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
