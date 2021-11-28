package interpreter

import (
	"github.com/RyoJerryYu/GocgCore/duel"
	lua "github.com/yuin/gopher-lua"
)

type coroutineMap map[int32]lua.LState
type paramList []struct { // TODO: pair
	First  interface{}
	Second uint32
}

type Interpreter struct {
	PDuel        *duel.Duel
	MsgBuf       [64]byte // TODO: char[64]
	LuaState     *lua.LState
	CurrentState *lua.LState
	Params       paramList
	Resumes      paramList
	Coroutines   coroutineMap
	NoAction     int32
	CallDepth    int32
}
