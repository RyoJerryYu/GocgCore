package group

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/interfaces"
)

type cardSet map[interfaces.Card]struct{}

type Group struct {
	RefHandle  int32
	PDuel      interfaces.Duel
	Container  cardSet
	It         uint32 // TODO: iterator of cardSet
	IsReadOnly uint32
}

func (g *Group) HasCard(c interfaces.Card) bool {
	_, ok := g.Container[c]
	return ok
}

func NewGroup(pd interfaces.Duel) *Group {
	return &Group{
		RefHandle:  0,
		PDuel:      pd,
		IsReadOnly: common.FALSE,
	}
}

func NewGroupWithCard(pd interfaces.Duel, c interfaces.Card) *Group {
	return &Group{
		RefHandle: 0,
		PDuel:     pd,
		Container: cardSet{
			c: {},
		},
		IsReadOnly: common.FALSE,
	}
}

func NewGroupWithCardSet(pd interfaces.Duel, cs cardSet) *Group {
	return &Group{
		RefHandle:  0,
		PDuel:      pd,
		Container:  cs,
		IsReadOnly: common.FALSE,
	}
}
