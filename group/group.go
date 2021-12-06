package group

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/interfaces"
)

type Group struct {
	RefHandle int32
	PDuel     interfaces.Duel
	Container interfaces.CardSet
	// It         uint32 // TODO: iterator of cardSet
	IsReadOnly uint32
}

func (g *Group) HasCard(c interfaces.Card) bool {
	return g.Container.HasCard(c)
}

func NewGroup(pd interfaces.Duel) *Group {
	return &Group{
		PDuel:      pd,
		IsReadOnly: common.FALSE,
	}
}

func NewGroupWithCard(pd interfaces.Duel, c interfaces.Card) *Group {
	return &Group{
		PDuel:      pd,
		Container:  interfaces.NewCardSet(c),
		IsReadOnly: common.FALSE,
	}
}

func NewGroupWithCardSet(pd interfaces.Duel, cs interfaces.CardSet) *Group {
	return &Group{
		PDuel:      pd,
		Container:  cs,
		IsReadOnly: common.FALSE,
	}
}
