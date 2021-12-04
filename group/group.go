package group

import (
	"github.com/RyoJerryYu/GocgCore/common"
	"github.com/RyoJerryYu/GocgCore/field"
)

type cardSet map[field.Card]struct{}

type Group struct {
	RefHandle  int32
	PDuel      field.Duel
	Container  cardSet
	It         uint32 // TODO: iterator of cardSet
	IsReadOnly uint32
}

func (g *Group) HasCard(c field.Card) bool {
	_, ok := g.Container[c]
	return ok
}

func NewGroup(pd field.Duel) *Group {
	return &Group{
		RefHandle:  0,
		PDuel:      pd,
		IsReadOnly: common.FALSE,
	}
}

func NewGroupWithCard(pd field.Duel, c field.Card) *Group {
	return &Group{
		RefHandle: 0,
		PDuel:     pd,
		Container: cardSet{
			c: {},
		},
		IsReadOnly: common.FALSE,
	}
}

func NewGroupWithCardSet(pd field.Duel, cs cardSet) *Group {
	return &Group{
		RefHandle:  0,
		PDuel:      pd,
		Container:  cs,
		IsReadOnly: common.FALSE,
	}
}
