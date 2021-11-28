package duel

import (
	"github.com/RyoJerryYu/GocgCore/field"
	"github.com/RyoJerryYu/GocgCore/group"
	"github.com/RyoJerryYu/GocgCore/interfaces"
	"github.com/RyoJerryYu/GocgCore/interpreter"
)

// type cardSet map[interfaces.Card]struct{} // TODO: comparator

type Duel struct {
	StrBuffer string
	Buffer    [0x1000]byte
	BufferLen uint32
	BufferP   *byte
	Lua       interpreter.Interpreter
	GameField *field.Field
	Random    Randomer
	Cards     map[interfaces.Card]struct{}
	Assumes   map[interfaces.Card]struct{}
	Groups    map[*group.Group]struct{}
	SGroups   map[*group.Group]struct{}
	Effects   map[interfaces.Effect]struct{}
	UnCopy    map[interfaces.Effect]struct{}
}
