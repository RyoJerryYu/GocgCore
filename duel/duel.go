package duel

import (
	"github.com/RyoJerryYu/GocgCore/interfaces"
	"github.com/RyoJerryYu/GocgCore/interpreter"
)

// type cardSet map[field.Card]struct{} // TODO: comparator

type Duel struct {
	StrBuffer string
	Buffer    [0x1000]byte
	BufferLen uint32
	BufferP   *byte
	Lua       interpreter.Interpreter
	GameField *interfaces.Field
	Random    Randomer
	Cards     map[interfaces.Card]struct{}
	Assumes   map[interfaces.Card]struct{}
	Groups    map[interfaces.Group]struct{}
	SGroups   map[interfaces.Group]struct{}
	Effects   map[interfaces.Effect]struct{}
	UnCopy    map[interfaces.Effect]struct{}
}
