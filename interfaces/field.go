package interfaces

import "github.com/RyoJerryYu/GocgCore/common"

type (
	ChainList []*Chain
)

type Field interface {
	ReloadFieldInfo()

	AddCard(playerid uint8, pcard Card, location uint8, sequence uint8, pzone uint8)
	RemoveCard(pcard Card)
	MoveCard(playerid uint8, pcard Card, location uint8, sequence uint8, pzone uint8)
	SwapCardWithSequence(pcard1 Card, pcard2 Card, newSequence1 uint8, newSequence2 uint8)
	SwapCard(pcard1 Card, pcard2 Card)
	SetControl(pcard Card, playerid uint8, resetPhase uint16, resetCount uint8)
	GetFieldCard(playerid uint32, location uint32, sequence uint32) Card
	IsLocationUseable(playerid uint32, location uint32, sequence uint32) int32
	GetUseableCount(pcard Card, playerid uint8, location uint8, uplayer uint8, reason uint32, zone uint32, list *uint32) int32
	GetUseableCountFromex(pcard Card, playerid uint8, uplayer uint8, zone uint32, list *uint32) int32
	GetSpsummonableCount(pcard Card, playerid uint8, zone uint32, list *uint32) int32
	GetSpsummonableCountFromex(pcard Card, playerid uint8, uplayer uint8, zone uint32, list *uint32) int32
	GetUseableCountOther(pcard Card, playerid uint8, location uint8, uplayer uint8, reason uint32, zone uint32, list *uint32) int32
	GetTofieldCount(pcard Card, playerid uint8, location uint8, uplayer uint32, reason uint32, zone uint32, list *uint32) int32
	GetUseableCountFromexRule4(pcard Card, playerid uint8, uplayer uint8, zone uint32, list *uint32) int32
	GetSpsummonableCountFromexRule4(pcard Card, playerid uint8, uplayer uint8, zone uint32, list *uint32) int32
	GetMzoneLimit(playerid uint8, uplayer uint8, reason uint32) int32
	GetSzoneLimit(playerid uint8, uplayer uint8, reason uint32) int32
	GetLinkedZone(playerid int32) uint32
	GetRuleZoneFromex(playerid int32, pcard Card) uint32
	FilterMustUseMzone(playerid uint8, uplayer uint8, reason uint32, pcard Card, flag *uint32)
	GetLinkedCards(self uint8, s uint8, o uint8, cset CardSet)
	CheckExtraLink(playerid int32) int32
	CheckExtraLinkWithSequence(playerid int32, pcard Card, sequence int32) int32
	GetCardsInZone(cset CardSet, zone uint32, playerid int32, location int32)
	Shuffle(playerid uint8, location uint8)
	ResetSequence(playerid uint8, location uint8)
	SwapDeckAndGrave(playerid uint8)
	ReverseDeck(playerid uint8)
	TagSwap(playerid uint8)

	AddEffect(peffect Effect, ownerPlayer uint8)
	RemoveEffect(peffect Effect)
	RemoveOathEffect(reasonEffect Effect)
	ReleaseOathRelation(reasonEffect Effect)
	ResetPhase(phase uint32)
	ResetChain()
	AddEffectCode(code uint32, playerid uint32)
	GetEffectCode(code uint32, playerid uint32) uint32
	DecEffectCode(code uint32, playerid uint32)

	FilterFieldEffect(code uint32, eset EffectSet, sort uint8)
	FilterAffectedCards(peffect Effect, cset CardSet)
	FilterInrangeCards(peffect Effect, cset CardSet)
	FilterPlayerEffect(playerid uint8, code uint32, eset EffectSet, sort uint8)
	FilterMatchingCard(findex int32, self uint8, location1 uint32, location2 uint32, pgroup Group, pexception Card, pexgroup Group, extraargs uint32, pret *Card, fcount int32, isTarget int32) int32
	FilterFieldCard(self uint8, location uint32, location2 uint32, pgroup Group) int32
	IsPlayerAffectedByEffect(playerid uint8, code uint32) Effect

	GetReleaseList(playerid uint8, releaseList CardSet, exList CardSet, exListOneof CardSet, useCon int32, useHand int32, fun int32, exarg int32, exc Card, exg Group) int32
	CheckReleaseList(playerid uint8, count int32, useCon int32, useHand int32, fun int32, exarg int32, exc Card, exg Group) int32
	GetSummonReleaseList(target Card, releaseList CardSet, exList CardSet, exListOneof CardSet, mg Group, ex uint32, releasable uint32, pos uint32) int32
	GetSummonCountLimit(playerid uint8) int32
	GetDrawCount(playerid uint8) int32
	GetRitualMaterial(playerid uint8, peffect Effect, material CardSet)
	GetFusionMaterial(playerid uint8, materialAll CardSet, materialBase CardSet, location uint32)
	RitualRelease(material CardSet)
	GetXyzMaterial(scard Card, findex int32, lv uint32, maxc int32, mg Group)
	GetOverlayGroup(self uint8, s uint8, o uint8, pset CardSet)
	GetOverlayCount(self uint8, s uint8, o uint8) int32
	UpdateDisableCheckList(peffect Effect)
	AddToDisableCheckList(pcard Card)
	AdjustDisableCheckList()
	AdjustSelfDestroySet()
	EraseGrantEffect(peffect Effect)
	AdjustGrantEffect() int32
	AddUniqueCard(pcard Card)
	RemoveUniqueCard(pcard Card)
	CheckUniqueOnfield(pcard Card, controler uint8, location uint8, icard Card) Effect
	CheckSpsummonOnce(pcard Card, playerid uint8) int32
	CheckCardCounterFromCard(pcard Card, counterType int32, playerid int32)
	CheckCardCounterFromGroup(pgroup Group, counterType int32, playerid int32)
	CheckChainCounter(peffect Effect, playerid int32, chainid int32, cancel bool)
	SetSpsummonCounter(playerid uint8)
	CheckSpsummonCounter(playerid uint8, ct uint8) int32

	CheckLpCost(playerid uint8, cost uint32, mustPay uint32) int32
	SaveLpCost()
	RestoreLpCost()
	PayLpCost(step uint32, playerid uint8, cost uint32, mustPay uint32) int32

	GetFieldCounter(self uint8, s uint8, o uint8, countertype uint16) uint32
	EffectReplaceCheck(code uint32, e TEvent) int32
	GetAttackTarget(pcard Card, v CardVector, chainAttack uint8, selectTarget bool) int32
	ConfirmAttackTarget() bool
	AttackAllTargetCheck()
	CheckSynchroMaterial(pcard Card, findex1 int32, findex2 int32, min int32, max int32, smat Card, mg Group) int32
	CheckTunerMaterial(pcard Card, tuner Card, findex1 int32, findex2 int32, min int32, max int32, smat Card, mg Group) int32
	CheckOtherSynchroMaterial(nsyn CardVector, lv int32, min int32, max int32, mcount int32) int32
	CheckTribute(pcard Card, min int32, max int32, mg Group, toplayer uint8, zone uint32, releasable uint32, pos uint32) int32
	CheckWithSumLimit(mats CardVector, acc int32, index int32, count int32, min int32, max int32, opmin int32) int32
	CheckWithSumLimitM(mats CardVector, acc int32, index int32, min int32, max int32, opmin int32, mustCount int32) int32
	CheckWithSumGreaterLimit(mats CardVector, acc int32, index int32, opmin int32) int32
	CheckWithSumGreaterLimitM(mats CardVector, acc int32, index int32, opmin int32, mustCount int32) int32
	CheckXyzMaterial(pcard Card, findex int32, lv int32, min int32, max int32, mg Group) int32

	IsPlayerCanDraw(playerid uint8) int32
	IsPlayerCanDiscardDeck(playerid uint8, count int32) int32
	IsPlayerCanDiscardDeckAsCost(playerid uint8, count int32) int32
	IsPlayerCanDiscardHand(playerid uint8, pcard Card, reasonEffect Effect, reason uint32) int32
	IsPlayerCanAction(playerid uint8, actionlimit uint32) int32
	IsPlayerCanSummon(sumtype uint32, playerid uint8, pcard Card, toplayer uint8) int32
	IsPlayerCanMset(sumtype uint32, playerid uint8, pcard Card, toplayer uint8) int32
	IsPlayerCanSset(playerid uint8, pcard Card) int32
	IsPlayerCanSpsummon(playerid uint8) int32
	IsPlayerCanSpsummonThat(reasonEffect Effect, sumtype uint32, sumpos uint8, playerid uint8, toplayer uint8, pcard Card) int32
	IsPlayerCanFlipsummon(playerid uint8, pcard Card) int32
	IsPlayerCanSpsummonMonster(playerid uint8, toplayer uint8, sumpos uint8, sumtype uint32, pdata CardData) int32
	IsPlayerCanSpsummonCount(playerid uint8, count uint32) int32
	IsPlayerCanRelease(playerid uint8, pcard Card) int32
	IsPlayerCanPlaceCounter(playerid uint8, pcard Card, countertype uint16, count uint16) int32
	IsPlayerCanRemoveCounter(playerid uint8, pcard Card, s uint8, o uint8, countertype uint16, count uint16, reason uint32) int32
	IsPlayerCanRemoveOverlayCard(playerid uint8, pcard Card, s uint8, o uint8, count uint16, reason uint32) int32
	IsPlayerCanSendToGrave(playerid uint8, pcard Card) int32
	IsPlayerCanSendToHand(playerid uint8, pcard Card) int32
	IsPlayerCanSendToDeck(playerid uint8, pcard Card) int32
	IsPlayerCanRemove(playerid uint8, pcard Card, reason uint32) int32
	IsChainNegatable(chaincount uint8) int32
	IsChainDisablable(chaincount uint8) int32
	IsChainDisabled(chaincount uint8) int32
	CheckChainTarget(chaincount uint8, pcard Card) int32
	GetChain(chaincount uint32) Chain
	GetCteffect(peffect Effect, playerid int32, store int32) int32
	GetCteffectEvt(feffect Effect, playerid int32, e TEvent, store int32) int32
	CheckCteffectHint(peffect Effect, playerid uint8) int32
	CheckNonpublicTrigger(ch Chain) int32
	CheckTriggerEffect(ch Chain) int32
	CheckSpselfFromHandTrigger(ch Chain) int32
	IsAbleToEnterBp() int32

	AddProcess(t uint16, step uint16, peffect Effect, target Group, arg1 common.Ptr, arg2 common.Ptr, arg3 common.Ptr, arg4 common.Ptr, ptr1 int32, ptr2 int32) // TODO: how to handle void*?
	Process() int32
	ExecuteCost(step uint16, peffect Effect, triggeringPlayer uint8) int32
	ExecuteOperation(step uint16, peffect Effect, triggeringPlayer uint8) int32
	ExecuteTarget(step uint16, peffect Effect, triggeringPlayer uint8) int32
	RaiseEventFromCard(eventCard Card, eventCode uint32, reasonEffect Effect, reason uint32, reasonPlayer uint8, eventPlayer uint8, eventValue uint32)
	RaiseEventFromCardSet(eventCards CardSet, eventCode uint32, reasonEffect Effect, reason uint32, reasonPlayer uint8, eventPlayer uint8, eventValue uint32)
	RaiseSingleEvent(triggerCard Card, eventCards CardSet, eventCode uint32, reasonEffect Effect, reason uint32, reasonPlayer uint8, eventPlayer uint8, eventValue uint32)
	CheckEvent(code uint32, pe TEvent) int32
	CheckEventC(peffect Effect, playerid uint8, neglectCon int32, neglectCost int32, copyInfo int32, pe TEvent) int32
	CheckHintTiming(peffect Effect) int32
	ProcessPhaseEvent(step int16, phaseEvent int32) int32
	ProcessPointEvent(step int16, skipTrigger int32, skipFreechain int32, skipNew int32) int32
	ProcessQuickEffect(step int16, skipFreechain int32, priority uint8) int32
	ProcessInstantEvent() int32
	ProcessSingleEvent() int32
	ProcessSingleEventThat(peffect Effect, e TEvent, tp ChainList, ntp ChainList) int32
	ProcessIdleCommand(step uint16) int32
	ProcessBattleCommand(step uint16) int32
	ProcessDamageStep(step uint16, newAttack uint32) int32
	CalculateBattleDamage(pdamchange *Effect, preasonCard *Card, battleDestroyed *uint8)
	ProcessTurn(step uint16, turnPlayer uint8) int32

	AddChain(step uint16) int32
	SolveContinuousOfEffect(playerid uint8, peffect Effect, e TEvent)
	SolveContinuousOfStep(step uint16) int32
	SolveChain(step uint16, chainendArg1 uint32, chainendArg2 uint32) int32
	BreakEffect() int32
	AdjustInstant()
	AdjustAll()
	RefreshLocationInfoInstant()
	RefreshLocationInfo(step uint16) int32
	AdjustStep(step uint16) int32

	//operations
	NegateChain(chaincount uint8) int32
	DisableChain(chaincount uint8) int32
	ChangeChainEffect(chaincount uint8, replaceOp int32)
	ChangeTarget(chaincount uint8, targets Group)
	ChangeTargetPlayer(chaincount uint8, playerid uint8)
	ChangeTargetParam(chaincount uint8, param int32)
	RemoveCounter(reason uint32, pcard Card, rplayer uint32, s uint32, o uint32, countertype uint32, count uint32)
	RemoveOverlayCard(reason uint32, pcard Card, rplayer uint32, s uint32, o uint32, min uint16, max uint16)
	GetControlOfCardSet(targets CardSet, reasonEffect Effect, reasonPlayer uint32, playerid uint32, resetPhase uint32, resetCount uint32, zone uint32)
	GetControlOfCard(target Card, reasonEffect Effect, reasonPlayer uint32, playerid uint32, resetPhase uint32, resetCount uint32, zone uint32)
	SwapControlOfCardSet(reasonEffect Effect, reasonPlayer uint32, targets1 CardSet, targets2 CardSet, resetPhase uint32, resetCount uint32)
	SwapControlOfCard(reasonEffect Effect, reasonPlayer uint32, pcard1 Card, pcard2 Card, resetPhase uint32, resetCount uint32)
	Equip(equipPlayer uint32, equipCard Card, target Card, up uint32, isStep uint32)
	Draw(reasonEffect Effect, reason uint32, reasonPlayer uint32, playerid uint32, count uint32)
	Damage(reasonEffect Effect, reason uint32, reasonPlayer uint32, reasonCard Card, playerid uint32, amount uint32, isStep uint32)
	Recover(reasonEffect Effect, reason uint32, reasonPlayer uint32, playerid uint32, amount uint32, isStep uint32)
	Summon(sumplayer uint32, target Card, proc Effect, ignoreCount uint32, minTribute uint32, zone uint32)
	Mset(setplayer uint32, target Card, proc Effect, ignoreCount uint32, minTribute uint32, zone uint32)
	SpecialSummonRule(sumplayer uint32, target Card, summonType uint32)
	SpecialSummon(target CardSet, sumtype uint32, sumplayer uint32, playerid uint32, nocheck uint32, nolimit uint32, positions uint32, zone uint32)
	SpecialSummonStep(target Card, sumtype uint32, sumplayer uint32, playerid uint32, nocheck uint32, nolimit uint32, positions uint32, zone uint32)
	SpecialSummonComplete(reasonEffect Effect, reasonPlayer uint8)
	DestroyCardSet(targets CardSet, reasonEffect Effect, reason uint32, reasonPlayer uint32, playerid uint32, destination uint32, sequence uint32)
	DestroyCard(target Card, reasonEffect Effect, reason uint32, reasonPlayer uint32, playerid uint32, destination uint32, sequence uint32)
	ReleaseCardSet(targets CardSet, reasonEffect Effect, reason uint32, reasonPlayer uint32)
	ReleaseCard(target Card, reasonEffect Effect, reason uint32, reasonPlayer uint32)
	SendToForCardSet(targets CardSet, reasonEffect Effect, reason uint32, reasonPlayer uint32, playerid uint32, destination uint32, sequence uint32, position uint32)
	SendToForCard(target Card, reasonEffect Effect, reason uint32, reasonPlayer uint32, playerid uint32, destination uint32, sequence uint32, position uint32)
	MoveToField(target Card, movePlayer uint32, playerid uint32, destination uint32, positions uint32, enable uint32, ret uint32, pzone uint32, zone uint32)
	ChangePositionForCardSet(targets CardSet, reasonEffect Effect, reasonPlayer uint32, au uint32, ad uint32, du uint32, dd uint32, flag uint32, enable uint32)
	ChangePositionForCard(target Card, reasonEffect Effect, reasonPlayer uint32, npos uint32, flag uint32, enable uint32)
	OperationReplace(t int32, step int32, targets Group)
	SelectTributeCards(target Card, playerid uint8, cancelable uint8, min int32, max int32, toplayer uint8, zone uint32)

	RemoveCounterForStep(step uint16, reason uint32, pcard Card, rplayer uint8, s uint8, o uint8, countertype uint16, count uint16) int32
	RemoveOverlayCardForStep(step uint16, reason uint32, pcard Card, rplayer uint8, s uint8, o uint8, min uint16, max uint16) int32
	GetControlOfGroup(step uint16, reasonEffect Effect, reasonPlayer uint8, targets Group, playerid uint8, resetPhase uint16, resetCount uint8, zone uint32) int32
	SwapControlOfGroup(step uint16, reasonEffect Effect, reasonPlayer uint8, targets1 Group, targets2 Group, resetPhase uint16, resetCount uint8) int32
	SelfDestroy(step uint16, ucard Card, p int32) int32
	TrapMonsterAdjust(step uint16) int32
	EquipForStep(step uint16, equipPlayer uint8, equipCard Card, target Card, up uint32, isStep uint32) int32
	DrawForStep(step uint16, reasonEffect Effect, reason uint32, reasonPlayer uint8, playerid uint8, count uint32) int32
	DamageForStep(step uint16, reasonEffect Effect, reason uint32, reasonPlayer uint8, reasonCard Card, playerid uint8, amount uint32, isStep uint32) int32
	RecoverForStep(step uint16, reasonEffect Effect, reason uint32, reasonPlayer uint8, playerid uint8, amount uint32, isStep uint32) int32
	SummonForStep(step uint16, sumplayer uint8, target Card, proc Effect, ignoreCount uint8, minTribute uint8, zone uint32) int32
	FlipSummon(step uint16, sumplayer uint8, target Card) int32
	MsetForStep(step uint16, setplayer uint8, ptarget Card, proc Effect, ignoreCount uint8, minTribute uint8, zone uint32) int32
	Sset(step uint16, setplayer uint8, toplayer uint8, ptarget Card, reasonEffect Effect) int32
	SsetG(step uint16, setplayer uint8, toplayer uint8, ptarget Group, confirm uint8, reasonEffect Effect) int32
	SpecialSummonRuleForStep(step uint16, sumplayer uint8, target Card, summonType uint32) int32
	SpecialSummonStepForStep(step uint16, targets Group, target Card, zone uint32) int32
	SpecialSummonForStep(step uint16, reasonEffect Effect, reasonPlayer uint8, targets Group, zone uint32) int32
	DestroyReplace(step uint16, targets Group, target Card, battle uint8) int32
	DestroyGroup(step uint16, targets Group, reasonEffect Effect, reason uint32, reasonPlayer uint8) int32
	ReleaseReplace(step uint16, targets Group, target Card) int32
	ReleaseGroup(step uint16, targets Group, reasonEffect Effect, reason uint32, reasonPlayer uint8) int32
	SendReplace(step uint16, targets Group, target Card) int32
	SendToForGroup(step uint16, targets Group, reasonEffect Effect, reason uint32, reasonPlayer uint8) int32
	DiscardDeck(step uint16, playerid uint8, count uint8, reason uint32) int32
	MoveToFieldForGroup(step uint16, target Card, enable uint32, ret uint32, pzone uint32, zone uint32) int32
	ChangePositionForStep(step uint16, targets Group, reasonEffect Effect, reasonPlayer uint8, enable uint32) int32
	OperationReplaceThat(step uint16, replaceEffect Effect, targets Group, target Card, isDestroy int32) int32
	ActivateEffect(step uint16, peffect Effect) int32
	SelectSynchroMaterial(step int16, playerid uint8, pcard Card, min int32, max int32, smat Card, mg Group) int32
	SelectXyzMaterial(step int16, playerid uint8, lv uint32, pcard Card, min int32, max int32) int32
	SelectReleaseCards(step int16, playerid uint8, cancelable uint8, min int32, max int32) int32
	SelectTributeCardsForStep(step int16, target Card, playerid uint8, cancelable uint8, min int32, max int32, toplayer uint8, zone uint32) int32
	TossCoin(step uint16, reasonEffect Effect, reasonPlayer uint8, playerid uint8, count uint8) int32
	TossDice(step uint16, reasonEffect Effect, reasonPlayer uint8, playerid uint8, count1 uint8, count2 uint8) int32
	RockPaperScissors(step uint16, repeat uint8) int32

	SelectBattleCommand(step uint16, playerid uint8) int32
	SelectIdleCommand(step uint16, playerid uint8) int32
	SelectEffectYesNo(step uint16, playerid uint8, description uint32, pcard Card) int32
	SelectYesNo(step uint16, playerid uint8, description uint32) int32
	SelectOption(step uint16, playerid uint8) int32
	SelectCard(step uint16, playerid uint8, cancelable uint8, min uint8, max uint8) int32
	SelectUnselectCard(step uint16, playerid uint8, cancelable uint8, min uint8, max uint8, finishable uint8) int32
	SelectChain(step uint16, playerid uint8, speCount uint8, forced uint8) int32
	SelectPlace(step uint16, playerid uint8, flag uint32, count uint8) int32
	SelectPosition(step uint16, playerid uint8, code uint32, positions uint8) int32
	SelectTribute(step uint16, playerid uint8, cancelable uint8, min uint8, max uint8) int32
	SelectCounter(step uint16, playerid uint8, countertype uint16, count uint16, s uint8, o uint8) int32
	SelectWithSumLimit(step int16, playerid uint8, acc int32, min int32, max int32) int32
	SortCard(step int16, playerid uint8) int32
	AnnounceRace(step int16, playerid uint8, count int32, available int32) int32
	AnnounceAttribute(step int16, playerid uint8, count int32, available int32) int32
	AnnounceCard(step int16, playerid uint8) int32
	AnnounceNumber(step int16, playerid uint8) int32
}
