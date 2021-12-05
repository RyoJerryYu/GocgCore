package interfaces

// TODO: Should not be a type but an actual function
// Maybe could be an method of Effect
type EffectSortId func(Effect, Effect) bool

type EffectSet struct {
	count     int
	container [64]Effect
}

func NewEffectSet() *EffectSet {
	return &EffectSet{
		count:     0,
		container: [64]Effect{nil},
	}
}

func (es *EffectSet) AddItem(item Effect) {
	if es.count >= len(es.container) {
		return
	}
	es.container[es.count] = item
	es.count++
}

func (es *EffectSet) RemoveItem(index int) {
	if index >= es.count {
		return
	}
	if index == es.count-1 {
		es.count--
		return
	}
	for i := index; i < es.count-1; i++ {
		es.container[i] = es.container[i+1]
	}
	es.count--
}

func (es *EffectSet) Clear() {
	es.count = 0
}

func (es *EffectSet) Size() int {
	return es.count
}

func (es *EffectSet) Sort() {
	if es.count < 2 {
		return
	}
	// TODO: EffectSortId should be used here
}

/**
In origin, there are two functions.
- One for const:
	```C++
	effect* const& get_last() const {
		return container[count - 1];
	}
	```
- One for non-const:
	```C++
	effect*& get_last() {
		return container[count - 1];
	}
	```
Here we egnore the difference and just define one.
Same as operator[](int)Effect and at(int)Effect.
*/
func (es *EffectSet) GetLast() Effect {
	return es.container[es.count-1]
}

func (es *EffectSet) At(index int) Effect {
	return es.container[index]
}

// TODO: Ignore effect_set_v, and improve EffectSet
// with golang build-in data struct later
