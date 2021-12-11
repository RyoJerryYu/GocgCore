package interfaces

// TODO: check redeclaration
type (
	EffectContainer map[uint32][]Effect
)

type Card interface {
	GetId() uint16 // defined as card_sort

	IsExtraDeckMonster() bool

	GetInfos(buf *byte, queryFlag int32, useCache int32) uint32
	GetInfoLocation() uint32
	SecondCode(code uint32) uint32
	GetCode() uint32
	GetAnotherCode() uint32
	IsSetCard(setCode uint32) int32
	IsOriginSetCard(setCode uint32) int32
	IsPreSetCard(setCode uint32) int32
	IsFusionSetCard(setCode uint32) int32
	IsLinkSetCard(setCode uint32) int32
	GetType() uint32
	GetFusionType() uint32
	GetSynchroType() uint32
	GetXyzType() uint32
	GetLinkType() uint32
	GetBaseAttack() int32
	GetAttack() int32
	GetBaseDefense() int32
	GetDefense() int32
	GetBattleAttack() int32
	GetBattleDefense() int32
	GetLevel() uint32
	GetRank() uint32
	GetLink() uint32
	GetSynchroLevel(pcard Card) uint32
	GetRitualLevel(pcard Card) uint32
	CheckXyzLevel(pcard Card, lv uint32) uint32
	GetAttribute() uint32
	GetFusionAttribute(playerid uint8) uint32
	GetLinkAttribute(playerid uint8) uint32
	GetGraveAttribute(playerid uint8) uint32
	GetRace() uint32
	GetLinkRace(playerid uint8) uint32
	GetGraveRace(playerid uint8) uint32
	GetLscale() uint32
	GetRscale() uint32
	GetLinkMarker() uint32
	IsLinkMarker(dir uint32) int32
	GetLinkedZone() uint32
	GetLinkedCards(cset CardSet)
	GetMutualLinkedZone() uint32
	GetMutualLinkedCards(cset CardSet)
	IsLinkState() int32
	IsExtraLinkState() int32
	IsPosition(pos int32) int32
	SetStatus(status uint32, enabled int32)
	GetStatus(status uint32) int32
	IsStatus(status uint32) int32
	GetColumnZone(location int32) uint32
	GetColumnCards(cset CardSet)
	IsAllColumn() int32

	Equip(target Card, sendMsg uint32)
	Unequip()
	GetUnionCount() int32
	GetOldUnionCount() int32
	XyzOverlay(materials CardSet)
	XyzAdd(mat Card)
	XyzRemove(mat Card)
	ApplyFieldEffect()
	CancelFieldEffect()
	EnableFieldEffect(enabled bool)
	AddEffect(peffect Effect) int32
	RemoveEffect(peffect Effect)
	RemoveEffectInContainer(peffect Effect, iterator EffectContainer) // TODO: How to define a iterator
	CopyEffect(code uint32, reset uint32, count uint32) int32
	ReplaceEffect(code uint32, reset uint32, count uint32) int32
	Reset(id uint32, resetType uint32)
	ResetEffectCount()
	RefreshDisableStatus()
	RefreshControlStatus() Effect

	CountTurn(ct uint16)
	CreateRelationWithCard(target Card, reset uint32)
	IsHasRelationWithCard(target Card) int32
	ReleaseRelationWithCard(target Card)
	CreateRelationWithChain(ch Chain)
	IsHasRelationWithChain(ch Chain) int32
	ReleaseRelationWithChain(ch Chain)
	ClearRelateEffect()
	CreateRelationWithEffect(peffect Effect)
	IsHasRelationWithEffect(peffect Effect) int32
	ReleaseRelationWithEffect(peffect Effect)
	LeaveFieldRedirect(reason uint32) int32
	DestinationRedirect(destination uint8, reason uint32) int32
	AddCounter(playerid uint8, countertype uint16, count uint16, singly uint8) int32
	RemoveCounter(countertype uint16, count uint16) int32
	IsCanAddCounter(playerid uint8, countertype uint16, count uint16, singly uint8, loc uint32) int32
	IsCanHaveCounter(countertype uint16) int32
	GetCounter(countertype uint16) int32
	SetMaterial(materials CardSet)
	AddCardTarget(pcard Card)
	CancelCardTarget(pcard Card)
	ClearCardTarget()

	FilterEffect(code int32, eset *EffectSet, sort uint8)
	FilterSingleContinuousEffect(code int32, eset *EffectSet, sort uint8)
	FilterSelfEffect(code int32, eset *EffectSet, sort uint8)
	FilterImmuneEffect()
	FilterDisableRelatedCards()
	FilterSummonProcedure(playerid uint8, eset *EffectSet, ignoreCount uint8, minTribute uint8, zone uint32) int32
	CheckSummonProcedure(proc Effect, playerid uint8, ignoreCount uint8, minTribute uint8, zone uint32) int32
	FilterSetProcedure(playerid uint8, eset *EffectSet, ignoreCount uint8, minTribute uint8, zone uint32) int32
	CheckSetProcedure(proc Effect, playerid uint8, ignoreCount uint8, minTribute uint8, zone uint32) int32
	FilterSpsummonProcedure(playerid uint8, eset *EffectSet, summonType uint32, info MaterialInfo)
	FilterSpsummonProcedureG(playerid uint8, eset *EffectSet)
	IsAffectedByEffect(code int32) Effect
	IsAffectedByEffectOfCard(code int32, target Card) Effect
	FusionCheck(fusionM Group, cg Card, chkf uint32, notMaterial uint8) int32
	FusionSelect(playerid uint8, fusionM Group, cg Card, chkf uint32, notMaterial uint8)
	CheckFusionSubstitute(fcard Card) int32
	IsNotTuner(scard Card) int32

	CheckUniqueCode(pcard Card) int32
	GetUniqueTarget(cset CardSet, controler int32, icard Card)
	CheckCostCondition(ecode int32, playerid int32) int32
	CheckCostConditionWithSumtype(ecode int32, playerid int32, sumtype int32) int32
	IsSummonableCard() int32
	IsFusionSummonableCard(summonType uint32) int32
	IsSpsummonable(proc Effect, info MaterialInfo) int32
	IsSummonable(proc Effect, minTribute uint8, zone uint32, releasable uint32) int32
	IsCanBeSummoned(playerid uint8, ingoreCount uint8, peffect Effect, minTribute uint8, zone uint32) int32
	GetSummonTributeCount() int32
	GetSetTributeCount() int32
	IsCanBeFlipSummoned(playerid uint8) int32
	IsSpecialSummonable(playerid uint8, summonType uint32, info MaterialInfo) int32
	IsCanBeSpecialSummoned(reasonEffect Effect, sumtype uint32, sumpos uint8, sumplayer uint8, toplayer uint8, nocheck uint8, nolimit uint8, zone uint32) int32
	IsSetableMzone(playerid uint8, ignoreCount uint8, peffect Effect, minTribute uint8, zone uint32) int32
	IsSetableSzone(playerid uint8, ignoreFd uint8) int32
	IsAffectByEffect(reasonEffect Effect) int32
	IsDestructable() int32
	IsDestructableByBattle(pcard Card) int32
	CheckIndestructableByEffect(reasonEffect Effect, playerid uint8) Effect
	IsDestructableByEffect(reasonEffect Effect, playerid uint8) int32
	IsRemoveable(playerid uint8, pos uint8, reason uint32) int32
	IsRemoveableAsCost(playerid uint8, pos uint8) int32
	IsReleasableBySummon(playerid uint8, pcard Card) int32
	IsReleasableByNonsummon(playerid uint8) int32
	IsReleasableByEffect(playerid uint8, reasonEffect Effect) int32
	IsCapableSendToGrave(playerid uint8) int32
	IsCapableSendToHand(playerid uint8) int32
	IsCapableSendToDeck(playerid uint8) int32
	IsCapableSendToExtra(playerid uint8) int32
	IsCapableCostToGrave(playerid uint8) int32
	IsCapableCostToHand(playerid uint8) int32
	IsCapableCostToDeck(playerid uint8) int32
	IsCapableCostToExtra(playerid uint8) int32
	IsCapableAttack() int32
	IsCapableAttackAnnounce(playerid uint8) int32
	IsCapableChangePosition(playerid uint8) int32
	IsCapableChangePositionByEffect(playerid uint8) int32
	IsCapableTurnSet(playerid uint8) int32
	IsCapableChangeControl() int32
	IsControlCanBeChanged(ignoreMzone int32, zone uint32) int32
	IsCapableBeBattleTarget(pcard Card) int32
	IsCapableBeEffectTarget(reasonEffect Effect, playerid uint8) int32
	IsCapableOverlay(playerid uint8) int32
	IsCanBeFusionMaterial(fcard Card, summonType uint32) int32
	IsCanBeSynchroMaterial(scard Card, tuner Card) int32
	IsCanBeRitualMaterial(scard Card) int32
	IsCanBeXyzMaterial(scard Card) int32
	IsCanBeLinkMaterial(scard Card) int32
}

type CardSet map[uint16]Card // TODO: Now Unsorted set. But actually sorted set in origin.

func NewCardSet(c Card) CardSet {
	return map[uint16]Card{c.GetId(): c}
}

func (s CardSet) Insert(c Card) {
	s[c.GetId()] = c
}

func (s CardSet) Remove(c Card) {
	delete(s, c.GetId())
}

func (s CardSet) HasCard(c Card) bool {
	_, ok := s[c.GetId()]
	return ok
}

type CardVector []Card
