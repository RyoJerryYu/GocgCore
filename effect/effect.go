package effect

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/interfaces"
)

// TODO: Maybe should type EffectSet interface{} ?

type EffectFlag uint32
type EffectFlag2 uint32

type Effect struct {
	RefHandle      int32
	PDual          interfaces.Duel
	Owner          interfaces.Card
	Handler        interfaces.Card
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
	ActiveHandler  interfaces.Card
	Status         uint16
	Label          []uint32
	LabelObject    int32
	Condition      int32
	Cost           int32
	Target         int32
	Value          int32
	Operation      int32
}

// var _ interfaces.Effect = (*Effect)(nil)

// In Golang, there is no need to use keyword explicit.
// In Golang, there is no need to define a destructor.
func NewEffect(pd interfaces.Duel) *Effect {
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
// TODO: use bool, check
func (e *Effect) IsDisableRelated() bool {
	switch e.Code {
	case EFFECT_IMMUNE_EFFECT, EFFECT_DISABLE, EFFECT_CANNOT_DISABLE, EFFECT_FORBIDDEN:
		return true
	default:
		return false
	}
}

// TODO: use bool, check
func (e *Effect) IsSelfDestroyRelated() bool {
	switch e.Code {
	case EFFECT_UNIQUE_CHECK, EFFECT_SELF_DESTROY, EFFECT_SELF_TOGRAVE:
		return true
	default:
		return false
	}
}

func (e *Effect) IsCanBeForbidden() bool {
	if e.IsFlag(EFFECT_FLAG_CANNOT_DISABLE) && !e.IsFlag(EFFECT_FLAG_CANNOT_NEGATE) {
		return false
	}
	if e.Code == EFFECT_CHANGE_CODE {
		return false
	}
	return true
}

// IsAvailable check if a single/field/equip effect is available
// check properties: range, EFFECT_FLAG_OWNER_RELATE, STATUS_BATTLE_DESTROYED, STATUS_EFFECT_ENABLED, disabled/forbidden
// check fucntions: condition
// TODO: use bool, check
func (e *Effect) IsAvailable() bool {
	if (e.Type & EFFECT_TYPE_ACTIONS) > 0 {
		return false
	}
	if (e.Type&(EFFECT_TYPE_SINGLE|EFFECT_TYPE_XMATERIAL) > 0) &&
		!(e.Type&EFFECT_TYPE_FIELD > 0) {
		pHandler := e.GetHandler()
		pOwner := e.GetOwner()
		// // TODO: Card.current.controler
		if e.IsFlag(EFFECT_FLAG_SINGLE_RANGE) && !e.InRangeFromCard(pHandler) {
			return false
		}
		if e.IsFlag(EFFECT_FLAG_SINGLE_RANGE) &&
			!(pHandler.GetStatus(common.STATUS_EFFECT_ENABLED) > 0) &&
			!e.IsFlag(EFFECT_FLAG_IMMEDIATELY_APPLY) {
			return false
		}
		// if e.IsFlag(EFFECT_FLAG_SINGLE_RANGE) &&
		// (pHandler.c)
		// TODO: Card.current.location
		if e.IsFlag(EFFECT_FLAG_SINGLE_RANGE) &&
			e.IsCanBeForbidden() &&
			(pOwner.IsStatus(common.STATUS_FORBIDDEN) > 0) {
			return false
		}
		if pOwner == pHandler &&
			e.IsCanBeForbidden() &&
			pOwner.IsStatus(common.STATUS_FORBIDDEN) > 0 {
			return false
		}
		if e.IsFlag(EFFECT_FLAG_OWNER_RELATE) &&
			!e.IsFlag(EFFECT_FLAG_CANNOT_DISABLE) &&
			pOwner.IsStatus(common.STATUS_BATTLE_DESTROYED) > 0 {
			return false
		}
		if pOwner == pHandler &&
			!e.IsFlag(EFFECT_FLAG_CANNOT_DISABLE) &&
			(pHandler.IsStatus(common.STATUS_DISABLED) > 0) {
			return false
		}
	}

	if e.Type&EFFECT_TYPE_EQUIP > 0 {
		// if e.Handler.current.controler == PLAYER_NONE
		// TODO: Card.current.controler
		if e.IsFlag(EFFECT_FLAG_OWNER_RELATE) &&
			e.IsCanBeForbidden() &&
			e.Owner.IsStatus(common.STATUS_FORBIDDEN) > 0 {
			return false
		}
		if e.Owner == e.Handler &&
			e.IsCanBeForbidden() &&
			e.Owner.IsStatus(common.STATUS_FORBIDDEN) > 0 {
			return false
		}
		if e.IsFlag(EFFECT_FLAG_OWNER_RELATE) &&
			!e.IsFlag(EFFECT_FLAG_CANNOT_DISABLE) &&
			e.Owner.IsStatus(common.STATUS_DISABLED) > 0 {
			return false
		}
		if e.Owner == e.Handler &&
			!e.IsFlag(EFFECT_FLAG_CANNOT_DISABLE) &&
			e.Handler.GetStatus(common.STATUS_DISABLED) > 0 {
			return false
		}
		if !e.IsFlag(EFFECT_FLAG_SET_AVAILABLE) {
			if !(e.Handler.GetStatus(common.STATUS_EFFECT_ENABLED) > 0) {
				return false
			}
			if !(e.Handler.IsPosition(common.POS_FACEUP) > 0) {
				return false
			}
		}
	}

	if e.Type&(EFFECT_TYPE_FIELD|EFFECT_TYPE_EQUIP) > 0 {
		pHandler := e.GetHandler()
		pOwner := e.GetOwner()
		if !e.IsFlag(EFFECT_FLAG_FIELD_ONLY) {
			// TODO: Card.current.controler
			// if Handler-> current.controler == Player_none

			if !e.InRangeFromCard(pHandler) {
				return false
			}
			if !(pHandler.GetStatus(common.STATUS_EFFECT_ENABLED) > 0) &&
				!e.IsFlag(EFFECT_FLAG_IMMEDIATELY_APPLY) {
				return false
			}
			// TODO: Card.current.location
			// if handler.current.location & LOCATION_ONFIELD

			if e.IsFlag(EFFECT_FLAG_OWNER_RELATE) &&
				e.IsCanBeForbidden() &&
				pOwner.IsStatus(common.STATUS_FORBIDDEN) > 0 {
				return false
			}
			if pOwner == pHandler &&
				e.IsCanBeForbidden() &&
				pHandler.GetStatus(common.STATUS_FORBIDDEN) > 0 {
				return false
			}
			if e.IsFlag(EFFECT_FLAG_OWNER_RELATE) &&
				!e.IsFlag(EFFECT_FLAG_CANNOT_DISABLE) &&
				pOwner.IsStatus(common.STATUS_DISABLED) > 0 {
				return false
			}
			if pOwner == pHandler &&
				!e.IsFlag(EFFECT_FLAG_CANNOT_DISABLE) &&
				pHandler.GetStatus(common.STATUS_DISABLED) > 0 {
				return false
			}
			if pHandler.IsStatus(common.STATUS_BATTLE_DESTROYED) > 0 {
				return false
			}
		}
	}
	if e.Condition == 0 {
		return true
	}

	// TODO: Duel.Lua.AddParam()
	// TODO: Duel.Gameinterfaces.Infos.FieldID++, Field, FieldInfo
	return true
}

// check if a effect is EFFECT_TYPE_SINGLE and is ready
// check: range, enabled, condition
// TODO: use bool, check
func (e *Effect) IsSingleReady() bool {
	if e.Type&EFFECT_TYPE_ACTIONS > 0 {
		return false
	}
	if (e.Type&(EFFECT_TYPE_SINGLE|EFFECT_TYPE_XMATERIAL) > 0) &&
		!(e.Type&EFFECT_TYPE_FIELD > 0) {
		return false
		// 	pHandler := e.GetHandler()
		// 	pOwner := e.GetOwner()
		// TODO: Card.current.controler
		// TODO: Effect.InRage()
		// TODO: Card.GetStatus()
		// TODO: Card.current.location
		// TODO: Card.IsPosition()
		// TODO: Duel.Lua.AddParam()
		// TODO: Duel.Lua.CheckCondition()
	}
	return true
}

// reset_count: count of effect reset
// count_limit: left count of activation
// count_limit_max: max count of activation
// TODO: change into bool, check if correct
// TODO: rename the function to isxxxx
func (e *Effect) CheckCountLimit(playerId uint8) bool {
	if e.IsFlag(EFFECT_FLAG_COUNT_LIMIT) {
		if e.CountLimit == 0 {
			return false
		}
		if e.CountCode > 0 {
			code := e.CountCode & 0xfffffff // reduce the first 4 bit
			// count := e.CountLimitMax
			if code == 1 {
				// TODO: Duel.Gameinterfaces.GetEffectCode()
				// TODO: Card.FieldID
				return false
			}
		}
	}
	return true
}

// check if an EFFECT_TYPE_ACTIONS effect can be activated
// for triggering effects, it checks EFFECT_FLAG_DAMAGE_STEP, EFFECT_FLAG_SET_AVAILABLE
func (e *Effect) IsActivatable(
	playerId uint8,
	te interfaces.TEvent,
	neglectCost int32,
	neglectTarget int32,
	neglectLoc int32,
	neglectFaceup int32,
) int32 {
	if !(e.Type&EFFECT_TYPE_ACTIONS > 0) {
		return common.FALSE
	}
	if !e.CheckCountLimit(playerId) {
		return common.FALSE
	}
	if e.IsFlag(EFFECT_FLAG_FIELD_ONLY) {
		// TODO: Card.current.controler
		// TODO: Duel.Gameinterfaces.CheckUniqueOnfield()
		return common.FALSE
	}
	return common.TRUE
}

func (e *Effect) GetOwner() interfaces.Card {
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

func (e *Effect) GetHandler() interfaces.Card {
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

// check if an EFFECT_TYPE_ACTIONS effect can be activated
// for triggering effects, it checks EFFECT_FLAG_DAMAGE_STEP, EFFECT_FLAG_SET_AVAILABLE
// TODO: changed into bool, check if correct
func (ef *Effect) IsActivateable(
	playerid uint8,
	e interfaces.TEvent,
	neglectCond int32,
	neglectCost int32,
	neglectTarget int32,
	neglectLoc int32,
	neglectFaceup int32,
) bool {
	if ef.Type&EFFECT_TYPE_ACTIONS == 0 {
		return false
	}
	if !ef.CheckCountLimit(playerid) {
		return false
	}
	if !ef.IsFlag(EFFECT_FLAG_FIELD_ONLY) {
		return false //TODO:
	}
	return true // TODO:
}

func (ef *Effect) IsActionCheck(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsActivateReadyWithEffect(reasonEffect interfaces.Effect, playerid uint8, e interfaces.TEvent, neglectCond int32, neglectCost int32, neglectTarget int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsActivateReady(playerid uint8, e interfaces.TEvent, neglectCond int32, neglectCost int32, neglectTarget int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsConditionCheck(playerid uint8, e interfaces.TEvent) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsActivateCheck(playerid uint8, e interfaces.TEvent, neglectCond int32, neglectCost int32, neglectTarget int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsTarget(pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsFitTargetFunction(pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsTargetPlayer(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsPlayerEffectTarget(pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsImmuned(pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsChainable(tp uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) IsHandTrigger() int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) Reset(resetLevel uint32, resetType uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) DecCount(playerid uint32) {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) Recharge() {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetValue(extraargs uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetValueFromCard(pcard interfaces.Card, extraargs uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetValueFromEffect(peffect interfaces.Effect, extraargs uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetValueToRes(extraargs uint32, result *[]int32) {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetValueFromCardToRes(pcard interfaces.Card, extraargs uint32, result *[]int32) {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetValueFromEffectToRes(peffect interfaces.Effect, extraargs uint32, result *[]int32) {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) CheckValueCondition(extraargs uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetLabelObject() *int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetSpeed() int32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) Clone() interfaces.Effect {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetOwnerPlayer() uint8 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetHandlerPlayer() uint8 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) InRangeFromCard(pcard interfaces.Card) bool {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) InRangeFromChain(ch interfaces.Chain) bool {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) SetActivateLocation() {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) SetActiveType() {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetActiveType() uint32 {
	panic("not implemented") // TODO: Implement
}

func (ef *Effect) GetCodeType() int32 {
	panic("not implemented") // TODO: Implement
}
