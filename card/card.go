package card

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/interfaces"
)

type EffectRelationHash struct {
}

// TODO: What's the meaning?
/*
struct effect_relation_hash {
	inline std::size_t operator()(const std::pair<effect*, uint16>& v) const {
		return std::hash<uint16>()(v.second);
	}
};
*/

type (
	cardVector      []*Card
	effectContainer map[uint32][]interfaces.Effect // symulate multimap
	cardSet         map[*Card]struct{}             // TODO: symulate set, ignoring comparator
	effectIndexer   map[interfaces.Effect]uint     // TODO: iterator
	effectRelation  map[uint16]interfaces.Effect   // TODO: comparator
	relationMap     map[*Card]uint32
	counterMap      map[uint16][2]uint16
	effectCount     map[uint32]int32
)

// TODO: symulate map of pair
type AttackerMap map[uint16]struct {
	first  *Card
	second uint32
}

func (am *AttackerMap) AddCard(card *Card) {

}

func (am *AttackerMap) FindCard(card *Card) uint32 {
	return 0
}

type sendToParamT struct {
	PlayerId uint8
	Position uint8
	Location uint8
	Sequence uint8
}

// TODO: ignore default seq
func (stp *sendToParamT) Set(p uint8, pos uint8, loc uint8, seq uint8) {
	stp.PlayerId = p
	stp.Position = pos
	stp.Location = loc
	stp.Sequence = seq
}

func (stp *sendToParamT) Clear() {
	stp.PlayerId = 0
	stp.Position = 0
	stp.Location = 0
	stp.Sequence = 0
}

type Card struct {
	RefHandle             int32
	PDuel                 interfaces.Duel
	Data                  interfaces.CardData
	Previous              interfaces.CardState
	Temp                  interfaces.CardState
	Current               interfaces.CardState
	QCache                interfaces.CardState
	Owner                 uint8
	SummonPlayer          uint8
	SummonInfo            uint32
	Status                uint32
	SendToParam           sendToParamT
	ReleaseParam          uint32
	SumParam              uint32
	PositionParam         uint32
	SPSummonParam         uint32
	ToFieldParam          uint32
	AttackAnnounceCount   uint8
	DirectAttackable      uint8
	AnnounceCount         uint8
	AttackAllTarget       uint8
	AttackControler       uint8
	CardId                uint16
	FieldId               uint32
	FieldIdR              uint32
	TurnId                uint16
	TurnCounter           uint16
	UniquePos             [2]uint8
	UniqueFieldId         uint32
	UniqueCode            uint32
	UniqueLocation        uint32
	UniqueFunction        uint32
	UniqueEffect          interfaces.Effect
	SPSummonCode          uint32
	SPSummonCounter       [2]uint16
	AssumeType            uint8
	AssumeValue           uint32
	EquipingTarget        *Card
	PreEquipTarget        *Card
	OverlayTarget         *Card
	Relations             relationMap
	Counters              counterMap
	IndestructableEffects effectCount
	AnnouncedCards        AttackerMap
	AttackedCards         AttackerMap
	BattledCards          AttackerMap
	EquipingCards         cardSet
	MaterialCards         cardSet
	EffectTargetOwner     cardSet
	EffectTargetCards     cardSet
	XyzMaterials          cardVector
	SingleEffect          effectContainer
	FieldEffect           effectContainer
	EquipEffect           effectContainer
	TargetEffect          effectContainer
	XMaterialEffect       effectContainer
	Indexer               effectIndexer
	RelateEffect          effectRelation
	ImmuneEffect          interfaces.EffectSet // TODO: EffectSetV
}

var _ interfaces.Card = (*Card)(nil)

func (c *Card) CardOperationSort(c1, c2 *Card) bool {
	// pDuel := c.PDuel
	var (
		cp1 uint8 // TODO: why int32 here?
		cp2 uint8
	)
	if c1.OverlayTarget != nil {
		cp1 = c1.OverlayTarget.Current.Controler // Controler Is a Peaple?
	} else {
		cp1 = c1.Current.Controler
	}
	if c2.OverlayTarget != nil {
		cp2 = c2.OverlayTarget.Current.Controler
	} else {
		cp2 = c2.Current.Controler
	}

	if cp1 == cp2 {
		if cp1 == common.PLAYER_NONE || cp2 == common.PLAYER_NONE {
			return cp1 < cp2
		}
		// TODO: Finish This After Duel Defined
	}
	return false
}

func NewCard(pd interfaces.Duel) *Card {
	c := &Card{} //TODO:
	return c
}

func (c *Card) GetId() uint16 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsExtraDeckMonster() bool {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetInfos(buf *byte, queryFlag int32, useCache int32) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetInfoLocation() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) SecondCode(code uint32) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetCode() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetAnotherCode() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsSetCard(setCode uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsOriginSetCard(setCode uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsPreSetCard(setCode uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsFusionSetCard(setCode uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsLinkSetCard(setCode uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetType() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetFusionType() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetSynchroType() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetXyzType() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLinkType() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetBaseAttack() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetAttack() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetBaseDefense() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetDefense() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetBattleAttack() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetBattleDefense() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLevel() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetRank() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLink() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetSynchroLevel(pcard interfaces.Card) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetRitualLevel(pcard interfaces.Card) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CheckXyzLevel(pcard interfaces.Card, lv uint32) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetAttribute() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetFusionAttribute(playerid uint8) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLinkAttribute(playerid uint8) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetGraveAttribute(playerid uint8) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetRace() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLinkRace(playerid uint8) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetGraveRace(playerid uint8) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLscale() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetRscale() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLinkMarker() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsLinkMarker(dir uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLinkedZone() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetLinkedCards(cset interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetMutualLinkedZone() uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetMutualLinkedCards(cset interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsLinkState() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsExtraLinkState() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsPosition(pos int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) SetStatus(status uint32, enabled int32) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetStatus(status uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsStatus(status uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetColumnZone(location int32) uint32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetColumnCards(cset interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsAllColumn() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) Equip(target interfaces.Card, sendMsg uint32) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) Unequip() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetUnionCount() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetOldUnionCount() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) XyzOverlay(materials interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) XyzAdd(mat interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) XyzRemove(mat interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) ApplyFieldEffect() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CancelFieldEffect() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) EnableFieldEffect(enabled bool) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) AddEffect(peffect interfaces.Effect) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) RemoveEffect(peffect interfaces.Effect) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) RemoveEffectInContainer(peffect interfaces.Effect, iterator interfaces.EffectContainer) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CopyEffect(code uint32, reset uint32, count uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) ReplaceEffect(code uint32, reset uint32, count uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) Reset(id uint32, resetType uint32) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) ResetEffectCount() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) RefreshDisableStatus() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) RefreshControlStatus() interfaces.Effect {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CountTurn(ct uint16) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CreateRelationWithCard(target interfaces.Card, reset uint32) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsHasRelationWithCard(target interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) ReleaseRelationWithCard(target interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CreateRelationWithChain(ch interfaces.Chain) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsHasRelationWithChain(ch interfaces.Chain) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) ReleaseRelationWithChain(ch interfaces.Chain) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) ClearRelateEffect() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CreateRelationWithEffect(peffect interfaces.Effect) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsHasRelationWithEffect(peffect interfaces.Effect) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) ReleaseRelationWithEffect(peffect interfaces.Effect) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) LeaveFieldRedirect(reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) DestinationRedirect(destination uint8, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) AddCounter(playerid uint8, countertype uint16, count uint16, singly uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) RemoveCounter(countertype uint16, count uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanAddCounter(playerid uint8, countertype uint16, count uint16, singly uint8, loc uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanHaveCounter(countertype uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetCounter(countertype uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) SetMaterial(materials interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) AddCardTarget(pcard interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CancelCardTarget(pcard interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) ClearCardTarget() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterEffect(code int32, eset *interfaces.EffectSet, sort uint8) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterSingleContinuousEffect(code int32, eset *interfaces.EffectSet, sort uint8) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterSelfEffect(code int32, eset *interfaces.EffectSet, sort uint8) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterImmuneEffect() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterDisableRelatedCards() {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterSummonProcedure(playerid uint8, eset *interfaces.EffectSet, ignoreCount uint8, minTribute uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CheckSummonProcedure(proc interfaces.Effect, playerid uint8, ignoreCount uint8, minTribute uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterSetProcedure(playerid uint8, eset *interfaces.EffectSet, ignoreCount uint8, minTribute uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CheckSetProcedure(proc interfaces.Effect, playerid uint8, ignoreCount uint8, minTribute uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterSpsummonProcedure(playerid uint8, eset *interfaces.EffectSet, summonType uint32, info interfaces.MaterialInfo) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FilterSpsummonProcedureG(playerid uint8, eset *interfaces.EffectSet) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsAffectedByEffect(code int32) interfaces.Effect {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsAffectedByEffectOfCard(code int32, target interfaces.Card) interfaces.Effect {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FusionCheck(fusionM interfaces.Group, cg interfaces.Card, chkf uint32, notMaterial uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) FusionSelect(playerid uint8, fusionM interfaces.Group, cg interfaces.Card, chkf uint32, notMaterial uint8) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CheckFusionSubstitute(fcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsNotTuner(scard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CheckUniqueCode(pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetUniqueTarget(cset interfaces.CardSet, controler int32, icard interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CheckCostCondition(ecode int32, playerid int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CheckCostConditionWithSumtype(ecode int32, playerid int32, sumtype int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsSummonableCard() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsFusionSummonableCard(summonType uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsSpsummonable(proc interfaces.Effect, info interfaces.MaterialInfo) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsSummonable(proc interfaces.Effect, minTribute uint8, zone uint32, releasable uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanBeSummoned(playerid uint8, ingoreCount uint8, peffect interfaces.Effect, minTribute uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetSummonTributeCount() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) GetSetTributeCount() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanBeFlipSummoned(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsSpecialSummonable(playerid uint8, summonType uint32, info interfaces.MaterialInfo) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanBeSpecialSummoned(reasonEffect interfaces.Effect, sumtype uint32, sumpos uint8, sumplayer uint8, toplayer uint8, nocheck uint8, nolimit uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsSetableMzone(playerid uint8, ignoreCount uint8, peffect interfaces.Effect, minTribute uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsSetableSzone(playerid uint8, ignoreFd uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsAffectByEffect(reasonEffect interfaces.Effect) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsDestructable() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsDestructableByBattle(pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) CheckIndestructableByEffect(reasonEffect interfaces.Effect, playerid uint8) interfaces.Effect {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsDestructableByEffect(reasonEffect interfaces.Effect, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsRemoveable(playerid uint8, pos uint8, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsRemoveableAsCost(playerid uint8, pos uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsReleasableBySummon(playerid uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsReleasableByNonsummon(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsReleasableByEffect(playerid uint8, reasonEffect interfaces.Effect) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableSendToGrave(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableSendToHand(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableSendToDeck(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableSendToExtra(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableCostToGrave(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableCostToHand(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableCostToDeck(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableCostToExtra(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableAttack() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableAttackAnnounce(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableChangePosition(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableChangePositionByEffect(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableTurnSet(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableChangeControl() int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsControlCanBeChanged(ignoreMzone int32, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableBeBattleTarget(pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableBeEffectTarget(reasonEffect interfaces.Effect, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCapableOverlay(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanBeFusionMaterial(fcard interfaces.Card, summonType uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanBeSynchroMaterial(scard interfaces.Card, tuner interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanBeRitualMaterial(scard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanBeXyzMaterial(scard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (c *Card) IsCanBeLinkMaterial(scard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}
