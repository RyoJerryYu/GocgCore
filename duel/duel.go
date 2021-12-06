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
	GameField interfaces.Field
	Random    Randomer
	Cards     map[interfaces.Card]struct{}
	Assumes   map[interfaces.Card]struct{}
	Groups    map[interfaces.Group]struct{}
	SGroups   map[interfaces.Group]struct{}
	Effects   map[interfaces.Effect]struct{}
	UnCopy    map[interfaces.Effect]struct{}
}

var _ interfaces.Duel = (*Duel)(nil)

func NewDuel() *Duel {
	d := &Duel{}
	// d.Lua = interpreter.NewInterpreter()
	// d.GameField = field.NewField(d)
	// TODO: game_field->temp_card = new_card(0)
	return d
}

func (d *Duel) Clear() {
	d.Cards = nil
	d.Groups = nil
	d.Effects = nil
	d.GameField = nil
	// TODO: finish here after implementing Duel interfaces for Duel
	// d.GameField = field.NewField(d)
	// game_field-> temp_card = new_card(0)
}

func (d *Duel) NewCard(code uint32) interfaces.Card {
	return nil
}

func (d *Duel) registerGroup(pgroup interfaces.Group) interfaces.Group {
	return nil
}

func (d *Duel) NewGroup() interfaces.Group {
	return nil
}

func (d *Duel) NewGroupByCard(card interfaces.Card) interfaces.Group {
	return nil
}

func (d *Duel) NewGroupByCardSet(cset interfaces.CardSet) interfaces.Group {
	return nil
}

func (d *Duel) NewEffect() interfaces.Effect {
	return nil
}

func (d *Duel) DeleteCard(card interfaces.Card) {

}

func (d *Duel) DeleteGroup(group interfaces.Group) {

}

func (d *Duel) DeleteEffect(effect interfaces.Effect) {

}

func (d *Duel) ReleaseScriptGroup() {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) RestoreAssumes() {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) ReadBuffer(buf *byte) int32 {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) WriteBuffer32(value uint32) {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) WriteBuffer16(value uint16) {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) WriteBuffer8(value uint8) {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) ClearBuffer() {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) SetResponsei(resp uint32) {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) SetResponseb(resp *byte) {
	panic("not implemented") // TODO: Implement
}

func (d *Duel) GetNextInteger(lo int32, hi int32) int32 {
	panic("not implemented") // TODO: Implement
}
