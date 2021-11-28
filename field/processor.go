package field

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/effectset"
	"github.com/RyoJerryYu/GocgCore/group"
	"github.com/RyoJerryYu/GocgCore/interfaces"
)

type (
	effectVector []interfaces.Effect
	// TODO: check redeclear and in-class type define
	// cardVector   []interfaces.Card
	optionVector []uint32
	// cardList      []interfaces.Card //  unused
	eventList     []*TEvent // TODO: use list
	chainList     []*Chain
	instantFList  map[interfaces.Effect]Chain
	chainArray    []Chain
	processorList []*ProcessorUnit
	cardSet       map[interfaces.Card]struct{}

	delayedEffectCollection map[struct {
		interfaces.Effect
		TEvent
	}]struct{}
	chainLimitT struct {
		Function int32
		Player   int32
	}
	chainLimitList []chainLimitT
	pairUint32     struct {
		first  uint32
		second uint32
	}
)

type Processor struct {
	Units                  processorList
	SubUnits               processorList
	Reserved               ProcessorUnit
	SelectCards            cardVector
	UnSelectCards          cardVector
	SummonableCards        cardVector
	SpSummonableCards      cardVector
	RepositionableCards    cardVector
	MSetableCards          cardVector
	SSetableCards          cardVector
	AttackableCards        cardVector
	SelectEffects          effectVector
	SelectOptions          optionVector
	MustSelectCards        cardVector
	PointEvent             eventList
	InstantEvent           eventList
	QueueEvent             eventList
	DelayedActivateEvent   eventList
	FullEvent              eventList
	UsedEvent              eventList
	SingleEvent            eventList
	SolvingEvent           eventList
	SubSolvingEvent        eventList
	SelectChains           chainArray
	CurrentChain           chainArray
	IgnitionPriorityChains chainArray
	ContinuousChain        chainList
	SolvingContinuous      chainList
	SubSolvingContinuous   chainList
	DelayedContinuousTp    chainList
	DelayedContinuousNtp   chainList
	DesrepChain            chainList
	NewFChain              chainList
	NewFChainS             chainList
	NewOChain              chainList
	NewOChainS             chainList
	NewFChainB             chainList
	NewOChainB             chainList
	NewOChainH             chainList
	NewChains              chainList
	DelayedQuickTemp       delayedEffectCollection
	DelayedQuick           delayedEffectCollection
	QuickFChain            instantFList
	LeaveConfirmed         cardSet
	SpecialSummoning       cardSet
	SSToGraveSet           cardSet
	EquipingCards          cardSet
	ControlAdjustSet       [2]cardSet
	UniqueDestroySet       cardSet
	SelfDestroySet         cardSet
	SelfToGraveSet         cardSet
	TrapMonsterAdjustSet   [2]cardSet
	ReleaseCards           cardSet
	ReleaseCardsEx         cardSet
	ReleaseCardsExOneOf    cardSet
	BattleDestroyRep       cardSet
	FusionMaterials        cardSet
	SynchroMaterials       cardSet
	OperatedSet            cardSet
	DiscardedSet           cardSet
	DestroyCanceled        cardSet
	DelayedEnableSet       cardSet
	SetGroupPreSet         cardSet
	SetGroupSet            cardSet
	DisfieldEffects        effectset.EffectSet // TODO: use effect set v
	ExtraMzoneEffects      effectset.EffectSet
	ExtraSzoneEffects      effectset.EffectSet
	ResetedEffects         map[interfaces.Effect]struct{}
	ReadjustMap            map[interfaces.Card]uint32
	UniqueCards            [2]map[interfaces.Card]struct{}
	EffectCountCode        map[uint32]uint32
	EffecctCountCodeDuel   map[uint32]uint32
	SpSummonOnceMap        [2]map[uint32]uint32
	XmaterialLst           map[int32][]interfaces.Card // TODO: use comarator

	TempVar                    [4]common.Ptr
	GlobalFlag                 uint32
	PreField                   [2]uint16
	OppMzone                   map[uint16]struct{}
	ChainLimit                 chainLimitList
	ChainLimitP                chainLimitList
	ChainSolving               uint8
	ContiSolving               uint8
	WinPlayer                  uint8
	WinReason                  uint8
	ReAdjust                   uint8
	ReasonEffect               interfaces.Effect
	ReasonPlayer               uint8
	SummoningCard              interfaces.Card
	SummonDepth                uint8
	SummonCancelable           uint8
	Attacker                   interfaces.Card
	AttackTarget               interfaces.Card
	LimitExtraSummonZone       uint32
	LimitExtraSummonReleasable uint32
	LimitTuner                 interfaces.Card
	LimitSyn                   *group.Group
	LimitSynMinc               int32
	LimitSynMaxc               int32
	LimitXyz                   *group.Group
	LimitXyzMinc               int32
	LimitXyzMaxc               int32
	LimitLink                  *group.Group
	LimitLinkCard              interfaces.Card
	LimitLinkMinc              int32
	LimitLinkMaxc              int32
	NotMaterial                uint8
	AttackCancelable           uint8
	AttackRollback             uint8
	EffectDamageStep           uint8
	BattleDamage               [2]int32
	SummonCount                [2]int32
	ExtraSummon                [2]uint8
	SpeEffect                  [2]int32
	DuelOptions                int32
	DuelRule                   int32 //current rule: 5, Master Rule 2020 int32
	CopyReset                  uint32
	CopyResetCount             uint8
	LastControlChangedId       uint32
	SetGroupUsedZones          uint32
	SetGroupSeq                [7]uint8
	DiceResult                 [5]uint8
	CoinResult                 [5]uint8
	ToBp                       uint8
	ToM2                       uint8
	ToEp                       uint8
	SkipM2                     uint8
	ChainAttack                uint8
	ChainAttackerId            uint32
	ChainAttackTarget          interfaces.Card
	AttackPlayer               uint8
	SelfdesDisabled            uint8
	Overdraw                   [2]uint8
	CheckLevel                 int32
	ShuffleCheckDisabled       uint8
	ShuffleHandCheck           [2]uint8
	ShuffleDeckCheck           [2]uint8
	DeckReversed               uint8
	RemoveBrainwashing         uint8
	FlipDelayed                uint8
	DamageCalculated           uint8
	HandAdjusted               uint8
	SummonStateCount           [2]uint8
	NormalsummonStateCount     [2]uint8
	FlipSummonStateCount       [2]uint8
	SpSummonStateCount         [2]uint8
	AttackStateCount           [2]uint8
	BattlePhaseCount           [2]uint8
	BattledCount               [2]uint8
	PhaseAction                uint8
	HintTiming                 [2]uint32
	CurrentPlayer              uint8
	ContiPlayer                uint8
	SummonCounter              map[uint32]pairUint32
	NormalSummonCounter        map[uint32]pairUint32
	SpSummonCounter            map[uint32]pairUint32
	FlipSummonCounter          map[uint32]pairUint32
	AttackCounter              map[uint32]pairUint32
	ChainCounter               map[uint32]pairUint32
	RecoverDamageReserve       processorList
	DecCountReserve            effectVector
}
