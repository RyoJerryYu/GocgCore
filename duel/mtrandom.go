package duel

import (
	"math"
	"math/rand"
	"time"
)

type randomMachine interface {
	GetRandMax() uint32
	Reset(seed uint32)
	Rand() uint32
}

type buildInU32RandomMachine struct {
}

func (r buildInU32RandomMachine) GetRandMax() uint32 {
	return math.MaxUint32
}

func (r buildInU32RandomMachine) Reset(seed uint32) {
	rand.Seed(int64(seed))
}

func (r buildInU32RandomMachine) Rand() uint32 {
	return rand.Uint32()
}

type Randomer struct {
	randomMachine
}

func (r Randomer) GetRandomInteger(lo, hi int) int {
	rng := uint32(hi - lo + 1) // range
	secureMax := uint32(r.GetRandMax() - r.GetRandMax()%rng)
	x := r.Rand()
	for x >= secureMax {
		x = r.Rand()
	}
	return int(x%rng) + lo
}

func (r Randomer) GetRandomIntegerOld(lo, hi int) int {
	result := int(float64(r.Rand())/
		float64(r.GetRandMax())*
		float64(hi-lo+1) +
		float64(lo))
	if result > hi {
		return hi
	}
	return result
}

func (r Randomer) ShuffleVector(v []interface{}) {
	n := len(v)
	for i := 0; i < n; i++ {
		j := r.GetRandomInteger(i, n-1)
		v[i], v[j] = v[j], v[i]
	}
}
func (r Randomer) ShuffleVectorOld(v []interface{}) {
	n := len(v)
	for i := 0; i < n; i++ {
		j := r.GetRandomIntegerOld(i, n-1)
		v[i], v[j] = v[j], v[i]
	}
}

func NewRandomer() *Randomer {
	rand.Seed(int64(time.Now().UnixNano()))
	return &Randomer{buildInU32RandomMachine{}}
}

func NewRandomerWithSeed(seed uint32) *Randomer {
	rand.Seed(int64(seed))
	return &Randomer{buildInU32RandomMachine{}}
}
