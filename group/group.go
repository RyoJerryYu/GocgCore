package group

import (
	"github.com/RyoJerryYu/GocgCore/card"
	"github.com/RyoJerryYu/GocgCore/common"
)

type cardSet map[*card.Card]struct{}
type Duel struct{}

type Group struct {
	RefHandle  int32
	PDuel      *Duel
	Container  cardSet
	It         uint32 // TODO: iterator of cardSet
	IsReadOnly uint32
}

func (g *Group) HasCard(c *card.Card) bool {
	_, ok := g.Container[c]
	return ok
}

func NewGroup(pd *Duel) *Group {
	return &Group{
		RefHandle:  0,
		PDuel:      pd,
		IsReadOnly: common.FALSE,
	}
}

func NewGroupWithCard(pd *Duel, c *card.Card) *Group {
	return &Group{
		RefHandle: 0,
		PDuel:     pd,
		Container: cardSet{
			c: {},
		},
		IsReadOnly: common.FALSE,
	}
}

func NewGroupWithCardSet(pd *Duel, cs cardSet) *Group {
	return &Group{
		RefHandle:  0,
		PDuel:      pd,
		Container:  cs,
		IsReadOnly: common.FALSE,
	}
}
