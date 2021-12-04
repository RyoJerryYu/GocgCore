package card

type CardData struct {
	Code       uint32
	Alias      uint32
	SetCode    uint64
	Type       uint32
	Level      uint32
	Attribute  uint32
	Race       uint32
	Attack     int32
	Defense    int32
	LScale     uint32
	RScale     uint32
	LinkMarker uint32
}

func (cd *CardData) Clear() {
	*cd = CardData{}
}

type QueryCache struct {
	Code        uint32
	Alias       uint32
	Type        uint32
	Level       uint32
	Rank        uint32
	Link        uint32
	Attribute   uint32
	Race        uint32
	Attack      int32
	Defense     int32
	BaseAttack  int32
	BaseDefense int32
	Reason      uint32
	Status      int32
	LScale      uint32
	RScale      uint32
	LinkMarker  uint32
}
