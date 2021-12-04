package interpreter

import (
	"github.com/RyoJerryYu/GocgCore/field"
	lua "github.com/yuin/gopher-lua"
)

type coroutineMap map[int32]lua.LState
type paramList []struct { // TODO: pair
	First  interface{}
	Second uint32
}

type Interpreter struct {
	PDuel        field.Duel
	MsgBuf       [64]byte // TODO: char[64]
	LuaState     *lua.LState
	CurrentState *lua.LState
	Params       paramList
	Resumes      paramList
	Coroutines   coroutineMap
	NoAction     int32
	CallDepth    int32
}
