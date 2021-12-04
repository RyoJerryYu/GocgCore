package duel

import (
	"github.com/RyoJerryYu/GocgCore/field"
	"github.com/RyoJerryYu/GocgCore/interpreter"
)

// type cardSet map[field.Card]struct{} // TODO: comparator

type Duel struct {
	StrBuffer string
	Buffer    [0x1000]byte
	BufferLen uint32
	BufferP   *byte
	Lua       interpreter.Interpreter
	GameField *field.Field
	Random    Randomer
	Cards     map[field.Card]struct{}
	Assumes   map[field.Card]struct{}
	Groups    map[field.Group]struct{}
	SGroups   map[field.Group]struct{}
	Effects   map[field.Effect]struct{}
	UnCopy    map[field.Effect]struct{}
}
