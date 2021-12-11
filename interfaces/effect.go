package interfaces

type Effect interface {
	IsDisableRelated() bool
	IsSelfDestroyRelated() bool
	IsCanBeForbidden() bool
	IsAvailable() bool
	IsSingleReady() bool
	CheckCountLimit(playerid uint8) bool
	IsActivateable(playerid uint8, e TEvent, neglectCond int32, neglectCost int32, neglectTarget int32, neglectLoc int32, neglectFaceup int32) bool
	IsActionCheck(playerid uint8) int32
	IsActivateReadyWithEffect(reasonEffect Effect, playerid uint8, e TEvent, neglectCond int32, neglectCost int32, neglectTarget int32) int32
	IsActivateReady(playerid uint8, e TEvent, neglectCond int32, neglectCost int32, neglectTarget int32) int32
	IsConditionCheck(playerid uint8, e TEvent) int32
	IsActivateCheck(playerid uint8, e TEvent, neglectCond int32, neglectCost int32, neglectTarget int32) int32
	IsTarget(pcard Card) int32
	IsFitTargetFunction(pcard Card) int32
	IsTargetPlayer(playerid uint8) int32
	IsPlayerEffectTarget(pcard Card) int32
	IsImmuned(pcard Card) int32
	IsChainable(tp uint8) int32
	IsHandTrigger() int32
	Reset(resetLevel uint32, resetType uint32) int32
	DecCount(playerid uint32)
	Recharge()
	GetValue(extraargs uint32) int32
	GetValueFromCard(pcard Card, extraargs uint32) int32
	GetValueFromEffect(peffect Effect, extraargs uint32) int32
	GetValueToRes(extraargs uint32, result *[]int32)
	GetValueFromCardToRes(pcard Card, extraargs uint32, result *[]int32)
	GetValueFromEffectToRes(peffect Effect, extraargs uint32, result *[]int32)
	CheckValueCondition(extraargs uint32) int32
	GetLabelObject() *int32 // TODO: *void, how to handle it?
	GetSpeed() int32
	Clone() Effect
	GetOwner() Card
	GetOwnerPlayer() uint8
	GetHandler() Card
	GetHandlerPlayer() uint8
	InRangeFromCard(pcard Card) bool
	InRangeFromChain(ch Chain) bool
	SetActivateLocation()
	SetActiveType()
	GetActiveType() uint32
	GetCodeType() int32
}
