package field

// type (
// effectContainer map[uint32][]*Effect
// cardSet map[interfaces.Card]struct{}
// effectVector []*Effect
// cardList []interfaces.Card
// eventList []*TEvent
// chainList []*Chain
// instantFList map[*Effect]Chain
// chainArray []*Chain
// processorList []*ProcessorUnit
// )

type Field struct {
	PDuel    Duel
	Player   [2]PlayerInfo
	TempCard Card
	Infos    fieldInfo
	// Cost     [2]LpCost
	Effects  FieldEffect
	Core     Processor
	Returns  ReturnValue
	NilEvent TEvent
}

var FieldUsedCount [32]int32

func NewField(duel Duel) *Field {
	return &Field{
		PDuel: duel,
	}
}
