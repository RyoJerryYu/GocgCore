package field

import (
	"github.com/RyoJerryYu/GocgCore/card"
	"github.com/RyoJerryYu/GocgCore/duel"
)

// type (
// effectContainer map[uint32][]*Effect
// cardSet map[*card.Card]struct{}
// effectVector []*Effect
// cardList []*card.Card
// eventList []*TEvent
// chainList []*Chain
// instantFList map[*Effect]Chain
// chainArray []*Chain
// processorList []*ProcessorUnit
// )

type Field struct {
	PDuel    *duel.Duel
	Player   [2]PlayerInfo
	TempCard *card.Card
	Infos    fieldInfo
	// Cost     [2]LpCost
	Effects  FieldEffect
	Core     Processor
	Returns  ReturnValue
	NilEvent TEvent
}

var FieldUsedCount [32]int32

func NewField(duel *duel.Duel) *Field {
	return &Field{
		PDuel: duel,
	}
}
