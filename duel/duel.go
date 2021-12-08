package duel

import (
	"github.com/RyoJerryYu/GocgCore/card"
	"github.com/RyoJerryYu/GocgCore/effect"
	"github.com/RyoJerryYu/GocgCore/field"
	"github.com/RyoJerryYu/GocgCore/group"
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
	// TODO: d.Lua = interpreter.NewInterpreter()
	d.GameField = field.NewField(d)
	// TODO: game_field->temp_card = new_card(0)
	return d
}

func (d *Duel) Clear() {
	d.Cards = nil
	d.Groups = nil
	d.Effects = nil
	d.GameField = nil
	d.GameField = field.NewField(d)
	// TODO: game_field-> temp_card = new_card(0)
}

func (d *Duel) NewCard(code uint32) interfaces.Card {
	var pCard interfaces.Card = card.NewCard(d)
	d.Cards[pCard] = struct{}{}
	if code > 0 {
		// TODO: d.readCard(pCard, code)
	} else {
		// TODO: pCard.Data.Clear()
	}
	return nil
}

func (d *Duel) registerGroup(pgroup interfaces.Group) interfaces.Group {
	d.Groups[pgroup] = struct{}{}
	if d.Lua.CallDepth != 0 {
		d.SGroups[pgroup] = struct{}{}
	}
	// TODO: d.Lua.RegisterGroup(pgroup)
	return pgroup
}

func (d *Duel) NewGroup() interfaces.Group {
	pGroup := group.NewGroup(d)
	return d.registerGroup(pGroup)
}

func (d *Duel) NewGroupByCard(card interfaces.Card) interfaces.Group {
	pGroup := group.NewGroupWithCard(d, card)
	return d.registerGroup(pGroup)
}

func (d *Duel) NewGroupByCardSet(cset interfaces.CardSet) interfaces.Group {
	pGroup := group.NewGroupWithCardSet(d, cset)
	return d.registerGroup(pGroup)
}

func (d *Duel) NewEffect() interfaces.Effect {
	pEffect := effect.NewEffect(d)
	d.Effects[pEffect] = struct{}{}
	// TODO: d.Lua.RegisterEffect(pEffect)
	return pEffect
}

func (d *Duel) DeleteCard(card interfaces.Card) {
	delete(d.Cards, card)
}

func (d *Duel) DeleteGroup(group interfaces.Group) {
	// TODO: d.Lua.UnregisterGroup(group)
	delete(d.Groups, group)
	delete(d.SGroups, group)
}

func (d *Duel) DeleteEffect(effect interfaces.Effect) {
	// TODO: d.Lua.UnregisterEffect(effect)
	delete(d.Effects, effect)
}

func (d *Duel) ReleaseScriptGroup() {
	for group := range d.SGroups {
		if true { // TODO:group.IsReadonly()
			d.DeleteGroup(group)
		}
	}
}

func (d *Duel) RestoreAssumes() {
	// for card := range d.Assumes {
	// TODO: card.assumeType = 0
	// }
	d.Assumes = nil
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
	// TODO: game_field->core.duel_options
	// TODO: reduce type change
	return int32(d.Random.GetRandomInteger(int(lo), int(hi)))
}
