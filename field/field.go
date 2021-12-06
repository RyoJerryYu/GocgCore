package field

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/interfaces"
)

// type (
// effectContainer map[uint32][]*Effect
// cardSet map[interfaces.Card]struct{}
// effectVector []*Effect
// cardList []interfaces.Card
// eventList []*TEvent
// instantFList map[*Effect]Chain
// chainArray []*Chain
// processorList []*ProcessorUnit
// )

type Field struct {
	PDuel    interfaces.Duel
	Player   [2]PlayerInfo
	TempCard interfaces.Card
	Infos    fieldInfo
	// Cost     [2]LpCost
	Effects  FieldEffect
	Core     Processor
	Returns  ReturnValue
	NilEvent interfaces.TEvent
}

var _ interfaces.Field = (*Field)(nil)

var FieldUsedCount [32]int32

func NewField(duel interfaces.Duel) *Field {
	return &Field{
		PDuel: duel,
	}
}

func (f *Field) ReloadFieldInfo() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AddCard(playerid uint8, pcard interfaces.Card, location uint8, sequence uint8, pzone uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RemoveCard(pcard interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) MoveCard(playerid uint8, pcard interfaces.Card, location uint8, sequence uint8, pzone uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SwapCardWithSequence(pcard1 interfaces.Card, pcard2 interfaces.Card, newSequence1 uint8, newSequence2 uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SwapCard(pcard1 interfaces.Card, pcard2 interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SetControl(pcard interfaces.Card, playerid uint8, resetPhase uint16, resetCount uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetFieldCard(playerid uint32, location uint32, sequence uint32) interfaces.Card {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsLocationUseable(playerid uint32, location uint32, sequence uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetUseableCount(pcard interfaces.Card, playerid uint8, location uint8, uplayer uint8, reason uint32, zone uint32, list *uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetUseableCountFromex(pcard interfaces.Card, playerid uint8, uplayer uint8, zone uint32, list *uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetSpsummonableCount(pcard interfaces.Card, playerid uint8, zone uint32, list *uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetSpsummonableCountFromex(pcard interfaces.Card, playerid uint8, uplayer uint8, zone uint32, list *uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetUseableCountOther(pcard interfaces.Card, playerid uint8, location uint8, uplayer uint8, reason uint32, zone uint32, list *uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetTofieldCount(pcard interfaces.Card, playerid uint8, location uint8, uplayer uint32, reason uint32, zone uint32, list *uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetUseableCountFromexRule4(pcard interfaces.Card, playerid uint8, uplayer uint8, zone uint32, list *uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetSpsummonableCountFromexRule4(pcard interfaces.Card, playerid uint8, uplayer uint8, zone uint32, list *uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetMzoneLimit(playerid uint8, uplayer uint8, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetSzoneLimit(playerid uint8, uplayer uint8, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetLinkedZone(playerid int32) uint32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetRuleZoneFromex(playerid int32, pcard interfaces.Card) uint32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) FilterMustUseMzone(playerid uint8, uplayer uint8, reason uint32, pcard interfaces.Card, flag *uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetLinkedCards(self uint8, s uint8, o uint8, cset interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckExtraLink(playerid int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckExtraLinkWithSequence(playerid int32, pcard interfaces.Card, sequence int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetCardsInZone(cset interfaces.CardSet, zone uint32, playerid int32, location int32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Shuffle(playerid uint8, location uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ResetSequence(playerid uint8, location uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SwapDeckAndGrave(playerid uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ReverseDeck(playerid uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) TagSwap(playerid uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AddEffect(peffect interfaces.Effect, ownerPlayer uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RemoveEffect(peffect interfaces.Effect) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RemoveOathEffect(reasonEffect interfaces.Effect) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ReleaseOathRelation(reasonEffect interfaces.Effect) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ResetPhase(phase uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ResetChain() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AddEffectCode(code uint32, playerid uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetEffectCode(code uint32, playerid uint32) uint32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DecEffectCode(code uint32, playerid uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) FilterFieldEffect(code uint32, eset interfaces.EffectSet, sort uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) FilterAffectedCards(peffect interfaces.Effect, cset interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) FilterInrangeCards(peffect interfaces.Effect, cset interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) FilterPlayerEffect(playerid uint8, code uint32, eset interfaces.EffectSet, sort uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) FilterMatchingCard(findex int32, self uint8, location1 uint32, location2 uint32, pgroup interfaces.Group, pexception interfaces.Card, pexgroup interfaces.Group, extraargs uint32, pret *interfaces.Card, fcount int32, isTarget int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) FilterFieldCard(self uint8, location uint32, location2 uint32, pgroup interfaces.Group) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerAffectedByEffect(playerid uint8, code uint32) interfaces.Effect {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetReleaseList(playerid uint8, releaseList interfaces.CardSet, exList interfaces.CardSet, exListOneof interfaces.CardSet, useCon int32, useHand int32, fun int32, exarg int32, exc interfaces.Card, exg interfaces.Group) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckReleaseList(playerid uint8, count int32, useCon int32, useHand int32, fun int32, exarg int32, exc interfaces.Card, exg interfaces.Group) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetSummonReleaseList(target interfaces.Card, releaseList interfaces.CardSet, exList interfaces.CardSet, exListOneof interfaces.CardSet, mg interfaces.Group, ex uint32, releasable uint32, pos uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetSummonCountLimit(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetDrawCount(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetRitualMaterial(playerid uint8, peffect interfaces.Effect, material interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetFusionMaterial(playerid uint8, materialAll interfaces.CardSet, materialBase interfaces.CardSet, location uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RitualRelease(material interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetXyzMaterial(scard interfaces.Card, findex int32, lv uint32, maxc int32, mg interfaces.Group) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetOverlayGroup(self uint8, s uint8, o uint8, pset interfaces.CardSet) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetOverlayCount(self uint8, s uint8, o uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) UpdateDisableCheckList(peffect interfaces.Effect) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AddToDisableCheckList(pcard interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AdjustDisableCheckList() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AdjustSelfDestroySet() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) EraseGrantEffect(peffect interfaces.Effect) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AdjustGrantEffect() int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AddUniqueCard(pcard interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RemoveUniqueCard(pcard interfaces.Card) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckUniqueOnfield(pcard interfaces.Card, controler uint8, location uint8, icard interfaces.Card) interfaces.Effect {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckSpsummonOnce(pcard interfaces.Card, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckCardCounterFromCard(pcard interfaces.Card, counterType int32, playerid int32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckCardCounterFromGroup(pgroup interfaces.Group, counterType int32, playerid int32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckChainCounter(peffect interfaces.Effect, playerid int32, chainid int32, cancel bool) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SetSpsummonCounter(playerid uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckSpsummonCounter(playerid uint8, ct uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckLpCost(playerid uint8, cost uint32, mustPay uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SaveLpCost() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RestoreLpCost() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) PayLpCost(step uint32, playerid uint8, cost uint32, mustPay uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetFieldCounter(self uint8, s uint8, o uint8, countertype uint16) uint32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) EffectReplaceCheck(code uint32, e interfaces.TEvent) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetAttackTarget(pcard interfaces.Card, v interfaces.CardVector, chainAttack uint8, selectTarget bool) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ConfirmAttackTarget() bool {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AttackAllTargetCheck() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckSynchroMaterial(pcard interfaces.Card, findex1 int32, findex2 int32, min int32, max int32, smat interfaces.Card, mg interfaces.Group) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckTunerMaterial(pcard interfaces.Card, tuner interfaces.Card, findex1 int32, findex2 int32, min int32, max int32, smat interfaces.Card, mg interfaces.Group) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckOtherSynchroMaterial(nsyn interfaces.CardVector, lv int32, min int32, max int32, mcount int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckTribute(pcard interfaces.Card, min int32, max int32, mg interfaces.Group, toplayer uint8, zone uint32, releasable uint32, pos uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckWithSumLimit(mats interfaces.CardVector, acc int32, index int32, count int32, min int32, max int32, opmin int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckWithSumLimitM(mats interfaces.CardVector, acc int32, index int32, min int32, max int32, opmin int32, mustCount int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckWithSumGreaterLimit(mats interfaces.CardVector, acc int32, index int32, opmin int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckWithSumGreaterLimitM(mats interfaces.CardVector, acc int32, index int32, opmin int32, mustCount int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckXyzMaterial(pcard interfaces.Card, findex int32, lv int32, min int32, max int32, mg interfaces.Group) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanDraw(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanDiscardDeck(playerid uint8, count int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanDiscardDeckAsCost(playerid uint8, count int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanDiscardHand(playerid uint8, pcard interfaces.Card, reasonEffect interfaces.Effect, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanAction(playerid uint8, actionlimit uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSummon(sumtype uint32, playerid uint8, pcard interfaces.Card, toplayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanMset(sumtype uint32, playerid uint8, pcard interfaces.Card, toplayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSset(playerid uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSpsummon(playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSpsummonThat(reasonEffect interfaces.Effect, sumtype uint32, sumpos uint8, playerid uint8, toplayer uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanFlipsummon(playerid uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSpsummonMonster(playerid uint8, toplayer uint8, sumpos uint8, sumtype uint32, pdata interfaces.CardData) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSpsummonCount(playerid uint8, count uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanRelease(playerid uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanPlaceCounter(playerid uint8, pcard interfaces.Card, countertype uint16, count uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanRemoveCounter(playerid uint8, pcard interfaces.Card, s uint8, o uint8, countertype uint16, count uint16, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanRemoveOverlayCard(playerid uint8, pcard interfaces.Card, s uint8, o uint8, count uint16, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSendToGrave(playerid uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSendToHand(playerid uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanSendToDeck(playerid uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsPlayerCanRemove(playerid uint8, pcard interfaces.Card, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsChainNegatable(chaincount uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsChainDisablable(chaincount uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsChainDisabled(chaincount uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckChainTarget(chaincount uint8, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetChain(chaincount uint32) interfaces.Chain {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetCteffect(peffect interfaces.Effect, playerid int32, store int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetCteffectEvt(feffect interfaces.Effect, playerid int32, e interfaces.TEvent, store int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckCteffectHint(peffect interfaces.Effect, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckNonpublicTrigger(ch interfaces.Chain) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckTriggerEffect(ch interfaces.Chain) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckSpselfFromHandTrigger(ch interfaces.Chain) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) IsAbleToEnterBp() int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AddProcess(t uint16, step uint16, peffect interfaces.Effect, target interfaces.Group, arg1 common.Ptr, arg2 common.Ptr, arg3 common.Ptr, arg4 common.Ptr, ptr1 int32, ptr2 int32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Process() int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ExecuteCost(step uint16, peffect interfaces.Effect, triggeringPlayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ExecuteOperation(step uint16, peffect interfaces.Effect, triggeringPlayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ExecuteTarget(step uint16, peffect interfaces.Effect, triggeringPlayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RaiseEventFromCard(eventCard interfaces.Card, eventCode uint32, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8, eventPlayer uint8, eventValue uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RaiseEventFromCardSet(eventCards interfaces.CardSet, eventCode uint32, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8, eventPlayer uint8, eventValue uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RaiseSingleEvent(triggerCard interfaces.Card, eventCards interfaces.CardSet, eventCode uint32, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8, eventPlayer uint8, eventValue uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckEvent(code uint32, pe interfaces.TEvent) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckEventC(peffect interfaces.Effect, playerid uint8, neglectCon int32, neglectCost int32, copyInfo int32, pe interfaces.TEvent) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CheckHintTiming(peffect interfaces.Effect) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessPhaseEvent(step int16, phaseEvent int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessPointEvent(step int16, skipTrigger int32, skipFreechain int32, skipNew int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessQuickEffect(step int16, skipFreechain int32, priority uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessInstantEvent() int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessSingleEvent() int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessSingleEventThat(peffect interfaces.Effect, e interfaces.TEvent, tp interfaces.ChainList, ntp interfaces.ChainList) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessIdleCommand(step uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessBattleCommand(step uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessDamageStep(step uint16, newAttack uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) CalculateBattleDamage(pdamchange *interfaces.Effect, preasonCard *interfaces.Card, battleDestroyed *uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ProcessTurn(step uint16, turnPlayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AddChain(step uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SolveContinuousOfEffect(playerid uint8, peffect interfaces.Effect, e interfaces.TEvent) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SolveContinuousOfStep(step uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SolveChain(step uint16, chainendArg1 uint32, chainendArg2 uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) BreakEffect() int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AdjustInstant() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AdjustAll() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RefreshLocationInfoInstant() {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RefreshLocationInfo(step uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AdjustStep(step uint16) int32 {
	panic("not implemented") // TODO: Implement
}

//operations
func (f *Field) NegateChain(chaincount uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DisableChain(chaincount uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ChangeChainEffect(chaincount uint8, replaceOp int32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ChangeTarget(chaincount uint8, targets interfaces.Group) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ChangeTargetPlayer(chaincount uint8, playerid uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ChangeTargetParam(chaincount uint8, param int32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RemoveCounter(reason uint32, pcard interfaces.Card, rplayer uint32, s uint32, o uint32, countertype uint32, count uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RemoveOverlayCard(reason uint32, pcard interfaces.Card, rplayer uint32, s uint32, o uint32, min uint16, max uint16) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetControlOfCardSet(targets interfaces.CardSet, reasonEffect interfaces.Effect, reasonPlayer uint32, playerid uint32, resetPhase uint32, resetCount uint32, zone uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetControlOfCard(target interfaces.Card, reasonEffect interfaces.Effect, reasonPlayer uint32, playerid uint32, resetPhase uint32, resetCount uint32, zone uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SwapControlOfCardSet(reasonEffect interfaces.Effect, reasonPlayer uint32, targets1 interfaces.CardSet, targets2 interfaces.CardSet, resetPhase uint32, resetCount uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SwapControlOfCard(reasonEffect interfaces.Effect, reasonPlayer uint32, pcard1 interfaces.Card, pcard2 interfaces.Card, resetPhase uint32, resetCount uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Equip(equipPlayer uint32, equipCard interfaces.Card, target interfaces.Card, up uint32, isStep uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Draw(reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32, playerid uint32, count uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Damage(reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32, reasonCard interfaces.Card, playerid uint32, amount uint32, isStep uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Recover(reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32, playerid uint32, amount uint32, isStep uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Summon(sumplayer uint32, target interfaces.Card, proc interfaces.Effect, ignoreCount uint32, minTribute uint32, zone uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Mset(setplayer uint32, target interfaces.Card, proc interfaces.Effect, ignoreCount uint32, minTribute uint32, zone uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SpecialSummonRule(sumplayer uint32, target interfaces.Card, summonType uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SpecialSummon(target interfaces.CardSet, sumtype uint32, sumplayer uint32, playerid uint32, nocheck uint32, nolimit uint32, positions uint32, zone uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SpecialSummonStep(target interfaces.Card, sumtype uint32, sumplayer uint32, playerid uint32, nocheck uint32, nolimit uint32, positions uint32, zone uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SpecialSummonComplete(reasonEffect interfaces.Effect, reasonPlayer uint8) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DestroyCardSet(targets interfaces.CardSet, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32, playerid uint32, destination uint32, sequence uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DestroyCard(target interfaces.Card, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32, playerid uint32, destination uint32, sequence uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ReleaseCardSet(targets interfaces.CardSet, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ReleaseCard(target interfaces.Card, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SendToForCardSet(targets interfaces.CardSet, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32, playerid uint32, destination uint32, sequence uint32, position uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SendToForCard(target interfaces.Card, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint32, playerid uint32, destination uint32, sequence uint32, position uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) MoveToField(target interfaces.Card, movePlayer uint32, playerid uint32, destination uint32, positions uint32, enable uint32, ret uint32, pzone uint32, zone uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ChangePositionForCardSet(targets interfaces.CardSet, reasonEffect interfaces.Effect, reasonPlayer uint32, au uint32, ad uint32, du uint32, dd uint32, flag uint32, enable uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ChangePositionForCard(target interfaces.Card, reasonEffect interfaces.Effect, reasonPlayer uint32, npos uint32, flag uint32, enable uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) OperationReplace(t int32, step int32, targets interfaces.Group) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectTributeCards(target interfaces.Card, playerid uint8, cancelable uint8, min int32, max int32, toplayer uint8, zone uint32) {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RemoveCounterForStep(step uint16, reason uint32, pcard interfaces.Card, rplayer uint8, s uint8, o uint8, countertype uint16, count uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RemoveOverlayCardForStep(step uint16, reason uint32, pcard interfaces.Card, rplayer uint8, s uint8, o uint8, min uint16, max uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) GetControlOfGroup(step uint16, reasonEffect interfaces.Effect, reasonPlayer uint8, targets interfaces.Group, playerid uint8, resetPhase uint16, resetCount uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SwapControlOfGroup(step uint16, reasonEffect interfaces.Effect, reasonPlayer uint8, targets1 interfaces.Group, targets2 interfaces.Group, resetPhase uint16, resetCount uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelfDestroy(step uint16, ucard interfaces.Card, p int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) TrapMonsterAdjust(step uint16) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) EquipForStep(step uint16, equipPlayer uint8, equipCard interfaces.Card, target interfaces.Card, up uint32, isStep uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DrawForStep(step uint16, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8, playerid uint8, count uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DamageForStep(step uint16, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8, reasonCard interfaces.Card, playerid uint8, amount uint32, isStep uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RecoverForStep(step uint16, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8, playerid uint8, amount uint32, isStep uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SummonForStep(step uint16, sumplayer uint8, target interfaces.Card, proc interfaces.Effect, ignoreCount uint8, minTribute uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) FlipSummon(step uint16, sumplayer uint8, target interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) MsetForStep(step uint16, setplayer uint8, ptarget interfaces.Card, proc interfaces.Effect, ignoreCount uint8, minTribute uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) Sset(step uint16, setplayer uint8, toplayer uint8, ptarget interfaces.Card, reasonEffect interfaces.Effect) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SsetG(step uint16, setplayer uint8, toplayer uint8, ptarget interfaces.Group, confirm uint8, reasonEffect interfaces.Effect) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SpecialSummonRuleForStep(step uint16, sumplayer uint8, target interfaces.Card, summonType uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SpecialSummonStepForStep(step uint16, targets interfaces.Group, target interfaces.Card, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SpecialSummonForStep(step uint16, reasonEffect interfaces.Effect, reasonPlayer uint8, targets interfaces.Group, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DestroyReplace(step uint16, targets interfaces.Group, target interfaces.Card, battle uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DestroyGroup(step uint16, targets interfaces.Group, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ReleaseReplace(step uint16, targets interfaces.Group, target interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ReleaseGroup(step uint16, targets interfaces.Group, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SendReplace(step uint16, targets interfaces.Group, target interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SendToForGroup(step uint16, targets interfaces.Group, reasonEffect interfaces.Effect, reason uint32, reasonPlayer uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) DiscardDeck(step uint16, playerid uint8, count uint8, reason uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) MoveToFieldForGroup(step uint16, target interfaces.Card, enable uint32, ret uint32, pzone uint32, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ChangePositionForStep(step uint16, targets interfaces.Group, reasonEffect interfaces.Effect, reasonPlayer uint8, enable uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) OperationReplaceThat(step uint16, replaceEffect interfaces.Effect, targets interfaces.Group, target interfaces.Card, isDestroy int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) ActivateEffect(step uint16, peffect interfaces.Effect) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectSynchroMaterial(step int16, playerid uint8, pcard interfaces.Card, min int32, max int32, smat interfaces.Card, mg interfaces.Group) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectXyzMaterial(step int16, playerid uint8, lv uint32, pcard interfaces.Card, min int32, max int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectReleaseCards(step int16, playerid uint8, cancelable uint8, min int32, max int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectTributeCardsForStep(step int16, target interfaces.Card, playerid uint8, cancelable uint8, min int32, max int32, toplayer uint8, zone uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) TossCoin(step uint16, reasonEffect interfaces.Effect, reasonPlayer uint8, playerid uint8, count uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) TossDice(step uint16, reasonEffect interfaces.Effect, reasonPlayer uint8, playerid uint8, count1 uint8, count2 uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) RockPaperScissors(step uint16, repeat uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectBattleCommand(step uint16, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectIdleCommand(step uint16, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectEffectYesNo(step uint16, playerid uint8, description uint32, pcard interfaces.Card) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectYesNo(step uint16, playerid uint8, description uint32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectOption(step uint16, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectCard(step uint16, playerid uint8, cancelable uint8, min uint8, max uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectUnselectCard(step uint16, playerid uint8, cancelable uint8, min uint8, max uint8, finishable uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectChain(step uint16, playerid uint8, speCount uint8, forced uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectPlace(step uint16, playerid uint8, flag uint32, count uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectPosition(step uint16, playerid uint8, code uint32, positions uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectTribute(step uint16, playerid uint8, cancelable uint8, min uint8, max uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectCounter(step uint16, playerid uint8, countertype uint16, count uint16, s uint8, o uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SelectWithSumLimit(step int16, playerid uint8, acc int32, min int32, max int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) SortCard(step int16, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AnnounceRace(step int16, playerid uint8, count int32, available int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AnnounceAttribute(step int16, playerid uint8, count int32, available int32) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AnnounceCard(step int16, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}

func (f *Field) AnnounceNumber(step int16, playerid uint8) int32 {
	panic("not implemented") // TODO: Implement
}
