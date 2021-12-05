package field

import "github.com/RyoJerryYu/GocgCore/interfaces"

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

var FieldUsedCount [32]int32

func NewField(duel interfaces.Duel) *Field {
	return &Field{
		PDuel: duel,
	}
}
