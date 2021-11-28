package card

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/effectset"
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
	effectContainer map[uint32][]Effect // symulate multimap
	cardSet         map[*Card]struct{}  // TODO: symulate set, ignoring comparator
	effectIndexer   map[Effect]uint     // TODO: iterator
	effectRelation  map[uint16]Effect   // TODO: comparator
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
	PDuel                 Duel
	Data                  CardData
	Previous              CardState
	Temp                  CardState
	Current               CardState
	QCache                CardState
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
	UniqueEffect          Effect
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
	ImmuneEffect          effectset.EffectSet // TODO: EffectSetV
}