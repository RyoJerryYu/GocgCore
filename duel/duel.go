package duel

import (
	"github.com/RyoJerryYu/GocgCore/card"
	"github.com/RyoJerryYu/GocgCore/effect"
	"github.com/RyoJerryYu/GocgCore/field"
	"github.com/RyoJerryYu/GocgCore/group"
	"github.com/RyoJerryYu/GocgCore/interpreter"
)

// type cardSet map[*card.Card]struct{} // TODO: comparator

type Duel struct {
	StrBuffer string
	Buffer    [0x1000]byte
	BufferLen uint32
	BufferP   *byte
	Lua       interpreter.Interpreter
	GameField *field.Field
	Random    Randomer
	Cards     map[*card.Card]struct{}
	Assumes   map[*card.Card]struct{}
	Groups    map[*group.Group]struct{}
	SGroups   map[*group.Group]struct{}
	Effects   map[*effect.Effect]struct{}
	UnCopy    map[*effect.Effect]struct{}
}
