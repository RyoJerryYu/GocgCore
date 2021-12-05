package interfaces

type Duel interface {
	NewCard(code uint32) Card
	NewGroup() Group
	NewGroupByCard(card Card) Group
	NewGroupByCardSet(cards CardSet) Group
	NewEffect() Effect
	DeleteCard(card Card)
	DeleteGroup(group Group)
	DeleteEffect(effect Effect)
	ReleaseScriptGroup()
	RestoreAssumes()
	ReadBuffer(buf *byte) int32 // TODO: maybe slice
	WriteBuffer32(value uint32)
	WriteBuffer16(value uint16)
	WriteBuffer8(value uint8)
	ClearBuffer()
	SetResponsei(resp uint32)
	SetResponseb(resp *byte)
	GetNextInteger(lo int32, hi int32) int32
}
