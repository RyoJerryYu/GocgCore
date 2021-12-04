package effect

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/field"
)

// TODO: Maybe should type EffectSet interface{} ?

type EffectFlag uint32
type EffectFlag2 uint32

type Effect struct {
	RefHandle      int32
	PDual          field.Duel
	Owner          field.Card
	Handler        field.Card
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
	ActiveHandler  field.Card
	Status         uint16
	Label          []uint32
	LabelObject    int32
	Condition      int32
	Cost           int32
	Target         int32
	Value          int32
	Operation      int32
}

// var _ field.Effect = (*Effect)(nil)

// In Golang, there is no need to use keyword explicit.
// In Golang, there is no need to define a destructor.
func NewEffect(pd field.Duel) *Effect {
	// TODO: use default zero value instead
	return &Effect{
		PDual:       pd,
		EffectOwner: common.PLAYER_NONE,
		Label:       make([]uint32, 0, 4),
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

// TODO: Why not return boolean?
// TODO: use bit operation
func (e *Effect) IsDisableRelated() int32 {
	switch e.Code {
	case EFFECT_IMMUNE_EFFECT, EFFECT_DISABLE, EFFECT_CANNOT_DISABLE, EFFECT_FORBIDDEN:
		return common.TRUE
	default:
		return common.FALSE
	}
}

func (e *Effect) IsSelfDestroyRelated() int32 {
	switch e.Code {
	case EFFECT_UNIQUE_CHECK, EFFECT_SELF_DESTROY, EFFECT_SELF_TOGRAVE:
		return common.TRUE
	default:
		return common.FALSE
	}
}

func (e *Effect) IsCanBeForbidden() int32 {
	if e.IsFlag(EFFECT_FLAG_CANNOT_DISABLE) && !e.IsFlag(EFFECT_FLAG_CANNOT_NEGATE) {
		return common.FALSE
	}
	if e.Code == EFFECT_CHANGE_CODE {
		return common.FALSE
	}
	return common.TRUE
}

// IsAvailable check if a single/field/equip effect is available
// check properties: range, EFFECT_FLAG_OWNER_RELATE, STATUS_BATTLE_DESTROYED, STATUS_EFFECT_ENABLED, disabled/forbidden
// check fucntions: condition
func (e *Effect) IsAvailable() int32 {
	if (e.Type & EFFECT_TYPE_ACTIONS) > 0 {
		return common.FALSE
	}
	// if (e.Type&(EFFECT_TYPE_SINGLE|EFFECT_TYPE_XMATERIAL) > 0) &&
	// 	!(e.Type&EFFECT_TYPE_FIELD > 0) {
	// pHandler := e.GetHandler()
	// pOwner := e.GetOwner()
	// TODO: Card.current.controler
	// if e.IsFlag(EFFECT_FLAG_SINGLE_RANGE) &&
	// TODO: Effect.InRage()
	// TODO: Card.current.location
	// TODO: Card.IsPosition()
	// TODO: Card.IsStatus()
	// TODO: Card.GetStatus()
	// TODO: Duel.Lua.AddParam()
	// TODO: Duel.GameField.Infos.FieldID++, Field, FieldInfo
	// }
	return common.TRUE
}

// check if a effect is EFFECT_TYPE_SINGLE and is ready
// check: range, enabled, condition
func (e *Effect) IsSingleReady() int32 {
	if e.Type&EFFECT_TYPE_ACTIONS > 0 {
		return common.FALSE
	}
	// if (e.Type&(EFFECT_TYPE_SINGLE|EFFECT_TYPE_XMATERIAL) > 0) &&
	// 	!(e.Type&EFFECT_TYPE_FIELD > 0) {
	// 	pHandler := e.GetHandler()
	// 	pOwner := e.GetOwner()
	// TODO: Card.current.controler
	// TODO: Effect.InRage()
	// TODO: Card.GetStatus()
	// TODO: Card.current.location
	// TODO: Card.IsPosition()
	// TODO: Duel.Lua.AddParam()
	// TODO: Duel.Lua.CheckCondition()
	// }
	return common.TRUE
}

// reset_count: count of effect reset
// count_limit: left count of activation
// count_limit_max: max count of activation
func (e *Effect) CheckCountLimit(playerId uint8) int32 {
	if e.IsFlag(EFFECT_FLAG_COUNT_LIMIT) {
		if e.CountLimit == 0 {
			return common.FALSE
		}
		if e.CountCode > 0 {
			code := e.CountCode & 0xfffffff // reduce the first 4 bit
			// count := e.CountLimitMax
			if code == 1 {
				// TODO: Duel.GameField.GetEffectCode()
				// TODO: Card.FieldID
				return common.FALSE
			}
		}
	}
	return common.TRUE
}

// check if an EFFECT_TYPE_ACTIONS effect can be activated
// for triggering effects, it checks EFFECT_FLAG_DAMAGE_STEP, EFFECT_FLAG_SET_AVAILABLE
func (e *Effect) IsActivatable(
	playerId uint8,
	te field.TEvent,
	neglectCost int32,
	neglectTarget int32,
	neglectLoc int32,
	neglectFaceup int32,
) int32 {
	if !(e.Type&EFFECT_TYPE_ACTIONS > 0) {
		return common.FALSE
	}
	if e.CheckCountLimit(playerId) == common.FALSE {
		return common.FALSE
	}
	if e.IsFlag(EFFECT_FLAG_FIELD_ONLY) {
		// TODO: Card.current.controler
		// TODO: Duel.GameField.CheckUniqueOnfield()
		return common.FALSE
	}
	return common.TRUE
}

func (e *Effect) GetOwner() field.Card {
	if e.ActiveHandler != nil {
		return e.ActiveHandler
	}
	if e.Type&EFFECT_TYPE_XMATERIAL > 0 {
		// TODO
		// return e.Handler.OverlayTarget()
		return nil
	}
	return e.Owner
}

func (e *Effect) GetHandler() field.Card {
	if e.ActiveHandler != nil {
		return e.ActiveHandler
	}
	if e.Type&EFFECT_TYPE_XMATERIAL > 0 {
		// TODO
		// return e.Handler.OverlayTarget()
		return nil
	}
	return e.Handler
}
